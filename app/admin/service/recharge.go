package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type RechargeService struct {
	request.RechargeRequest
}

func (this RechargeService) PageList() response.RechargeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.Recharge{}
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	sli := make([]response.RechargeInfo, 0)
	for _, v := range list {
		item := response.RechargeInfo{
			Id:      v.Id,
			OrderSn: v.OrderSn,
			UId:     v.UId,
			Type:    v.Type,
			//Amount:      float64(v.Amount),
			//RealAmount:  float64(v.RealAmount),
			From:      v.From,
			To:        v.To,
			Voucher:   v.Voucher,
			PaymentId: v.PaymentId,
			Status:    v.Status,
			//UsdtAmount:  float64(v.UsdtAmount),
			Operator:    v.Operator,
			Description: v.Description,
			UpdateTime:  v.UpdateTime,
			CreateTime:  v.CreateTime,
			Username:    v.Member.Username,
			MethodName:  v.RechargeMethod.Name,
			//PaymentName: v.Payment.PayName,
			SuccessTime: v.SuccessTime,
			TradeSn:     v.TradeSn,
			ImageUrl:    v.ImageUrl,
			RealName:    v.MemberVerified.RealName,
		}
		sli = append(sli, item)
	}
	return response.RechargeData{
		List: sli,
		Page: FormatPage(page),
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

func (this RechargeService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.UId > 0 {
		where[model.Recharge{}.TableName()+".uid"] = this.UId
	}
	if this.OrderSn != "" {
		where[model.Recharge{}.TableName()+".order_sn"] = this.OrderSn
	}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.StartTime != "" {
		where[model.Recharge{}.TableName()+".create_time >="] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where[model.Recharge{}.TableName()+".create_time <"] = common.DateToUnix(this.EndTime)
	}
	if this.Status > 0 {
		where[model.Recharge{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
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
		trade.Desc = "他人向您转账"
	}
	err := trade.Insert()
	if err != nil {
		logrus.Error(err)
		return err
	}
	//上分
	member.Balance = member.Balance.Add(amount)
	//member.TotalBalance = member.TotalBalance.Add(amount)
	return member.Update("balance")
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
