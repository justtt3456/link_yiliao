package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type RechargeService struct {
	request.RechargeListRequest
}

func (this RechargeService) PageList(agent *model.Agent) *response.RechargeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	offset := this.PageSize * (this.Page - 1)
	m := model.Recharge{}
	db := global.DB.Model(m).Joins("Member").Joins("MemberVerified").Where("Member.agent_id = ?", agent.Id)
	db = this.getWhere(db)
	var total int64
	db2 := db.Session(&gorm.Session{})
	err := db2.Count(&total).Error
	if err != nil {
		return nil
	}
	page := common.Page{
		Page: this.Page,
	}
	page.SetPage(this.PageSize, total)
	list := make([]model.Recharge, 0)
	err = db.Order(m.TableName() + ".id desc").Limit(this.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return nil
	}
	sli := make([]response.RechargeInfo, 0)
	var totalAmount decimal.Decimal
	for _, v := range list {
		agent := model.Agent{}
		if v.Member.AgentId > 0 {
			agent.Id = v.Member.AgentId
			agent.Get()
		}
		item := response.RechargeInfo{
			Id:          v.Id,
			OrderSn:     v.OrderSn,
			UId:         v.UId,
			Type:        v.Type,
			Amount:      v.Amount,
			RealAmount:  v.RealAmount,
			From:        v.From,
			To:          v.To,
			Voucher:     v.Voucher,
			PaymentId:   v.PaymentId,
			Status:      v.Status,
			UsdtAmount:  v.UsdtAmount,
			Operator:    v.Operator,
			Description: v.Description,
			UpdateTime:  v.UpdateTime,
			CreateTime:  v.CreateTime,
			Username:    v.Member.Username,
			MethodName:  v.RechargeMethod.Name,
			PaymentName: v.Payment.PayName,
			SuccessTime: v.SuccessTime,
			TradeSn:     v.TradeSn,
			RealName:    v.MemberVerified.RealName,
			AgentName:   agent.Account,
		}
		sli = append(sli, item)
		totalAmount = totalAmount.Add(v.Amount)
	}
	return &response.RechargeData{
		List:        sli,
		Page:        FormatPage(page),
		TotalAmount: totalAmount,
	}
}

type RechargeUpdate struct {
	request.RechargeUpdateRequest
}

func (this RechargeUpdate) Update(admin model.Admin) error {
	if this.Ids == "" {
		return errors.New("参数错误")
	}
	ids := strings.Split(this.Ids, ",")
	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		m := model.Recharge{
			Id: id,
		}
		if !m.Get() {
			return errors.New("记录不存在")
		}
		if m.Status != model.StatusReview {
			return errors.New("当前状态无法修改")
		}
		member := model.Member{Id: m.UId}
		if !member.Get() {
			return errors.New("用户不存在")
		}
		if this.Status == model.StatusAccept {
			r := RechargeHandle{}
			r.Recharge(member, m.Id, m.Amount, 1, 3, 1)
		}
		m.Status = this.Status
		m.Description = this.Description
		m.Operator = admin.Id
		//更新状态 说明 操作者
		m.Update("status", "description", "operator")
	}
	return nil
}

func (this RechargeService) getWhere(db *gorm.DB) *gorm.DB {

	if this.UId > 0 {
		db.Where(model.Recharge{}.TableName()+".uid = ?", this.UId)
	}
	if this.OrderSn != "" {
		db.Where(model.Recharge{}.TableName()+".order_sn = ?", this.OrderSn)
	}
	if this.Username != "" {
		db.Where(".Member.username = ?", this.Username)
	}
	if this.StartTime != "" {
		db.Where(model.Recharge{}.TableName()+".create_time >= ?", common.DateToUnix(this.StartTime))
	}
	if this.EndTime != "" {
		db.Where(model.Recharge{}.TableName()+".create_time < ?", common.DateToUnix(this.EndTime))
	}
	if this.Status > 0 {
		db.Where(model.Recharge{}.TableName()+".status = ?", this.Status)
	}
	if this.AgentName != "" {
		agent := model.Agent{Account: this.AgentName}
		if agent.Get() {
			db.Where("Member.agent_id = ?", agent.Id)
		} else {
			db.Where("Member.agent_id = -1")
		}
	}
	return db
}

type RechargeHandle struct {
}

func (RechargeHandle) Recharge(member model.Member, item int, amount decimal.Decimal, way int, tradeType int, isfront int) error {
	//账单
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  tradeType,
		ItemId:     item,
		Amount:     amount,
		Before:     member.Balance,
		After:      member.Balance.Add(amount),
		IsFrontend: isfront,
	}
	switch way {
	case 1: //审核
		trade.Desc = "充值审核通过"
	case 2: //系统
		//trade.Desc = "福利派送"
		trade.Desc = "后台上分可用余额"
	}
	err := trade.Insert()
	if err != nil {
		logrus.Error(err)
		return err
	}
	//上分
	member.Balance = member.Balance.Add(amount)
	member.TotalRecharge = member.TotalRecharge.Add(amount)
	return member.Update("balance", "total_recharge")
}

func (RechargeHandle) TopupUseBalance(member model.Member, item int, amount decimal.Decimal, tradeType int, isfront int) error {
	//账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: tradeType,
		ItemId:    item,
		//Amount:     amount,
		Before: member.WithdrawBalance,
		//After:      member.WithdrawBalance + amount,
		Desc:       "提现冲正回调",
		IsFrontend: isfront,
	}

	err := trade.Insert()
	if err != nil {
		logrus.Error(err)
		return err
	}
	//上分
	//member.WithdrawBalance += amount
	//member.TotalBalance += amount
	return member.Update("withdraw_balance")
}
