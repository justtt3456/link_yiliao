package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type WithdrawListService struct {
	request.WithdrawListRequest
}
type WithdrawUpdateService struct {
	request.WithdrawUpdateRequest
}

func (this WithdrawListService) PageList(agent *model.Agent) *response.WithdrawData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	offset := this.PageSize * (this.Page - 1)
	db := global.DB.Model(model.Withdraw{}).Joins("Member").Where("Member.agent_id = ?", agent.Id)
	db = this.getWhere(db)
	db2 := db.Session(&gorm.Session{})
	var total int64
	err := db2.Count(&total).Error
	if err != nil {
		return nil
	}
	page := common.Page{
		Page: this.Page,
	}
	page.SetPage(this.PageSize, total)
	list := make([]model.Withdraw, 0)
	err = db.Order(model.Withdraw{}.TableName() + ".id desc").Limit(this.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return nil
	}
	res := make([]response.WithdrawInfo, 0)
	var totalAmount decimal.Decimal
	for _, v := range list {
		agent := model.Agent{}
		if v.Member.AgentId > 0 {
			agent.Id = v.Member.AgentId
			agent.Get()
		}
		i := response.WithdrawInfo{
			Id:               v.Id,
			UId:              v.UId,
			WithdrawType:     v.WithdrawType,
			BankName:         v.BankName,
			BranchBank:       v.BranchBank,
			RealName:         v.RealName,
			CardNumber:       v.CardNumber,
			BankPhone:        v.BankPhone,
			Amount:           v.Amount,
			Fee:              v.Fee,
			TotalAmount:      v.TotalAmount,
			UsdtAmount:       v.UsdtAmount,
			Description:      v.Description,
			Operator:         v.Operator1,
			ViewStatus:       v.ViewStatus,
			Status:           v.Status,
			SuccessTime:      v.SuccessTime,
			OrderSn:          v.OrderSn,
			PaymentId:        v.PaymentId,
			PaymentName:      v.Payment.PayName,
			TradeSn:          v.TradeSn,
			CreateTime:       v.CreateTime,
			UpdateTime:       v.UpdateTime,
			Username:         v.Member.Username,
			MethodName:       v.WithdrawMethod.Name,
			RegisterRealName: v.Member.RealName,
			AgentName:        agent.Account,
		}
		res = append(res, i)
		totalAmount = totalAmount.Add(v.TotalAmount)
	}
	return &response.WithdrawData{
		List:        res,
		Page:        FormatPage(page),
		TotalAmount: totalAmount,
	}
}

func (this WithdrawUpdateService) Update() error {
	if this.Ids == "" {
		return errors.New("参数错误")
	}
	ids := strings.Split(this.Ids, ",")
	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		m := model.Withdraw{
			Id: id,
		}
		if !m.Get() {
			return errors.New("记录不存在")
		}
		if m.Status != model.StatusReview {
			return errors.New("当前记录不允许操作")
		}
		user := model.Member{Id: m.UId}
		if !user.Get() {
			return errors.New("用户不存在")
		}
		switch this.Status {
		case model.StatusRollback:
			//生成账单
			trade := model.Trade{
				UId:       m.UId,
				TradeType: 4,
				ItemId:    m.Id,
				Amount:    m.TotalAmount,
				Before:    user.Balance,
				After:     user.WithdrawBalance.Add(m.TotalAmount),
				Desc:      "提现驳回",
			}
			if err := trade.Insert(); err != nil {
				return err
			}
			//回滚余额
			user.WithdrawBalance = user.WithdrawBalance.Add(m.TotalAmount)
			if err := user.Update("withdraw_balance"); err != nil {
				return err
			}
		case model.StatusAccept:
		}
		m.Status = this.Status
		m.Operator1 = this.Operator
		m.Description = this.Description
		//更新状态 说明 操作者
		err := m.Update("status", "description", "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func (this WithdrawListService) getWhere(db *gorm.DB) *gorm.DB {
	if this.UId > 0 {
		db.Where(model.Withdraw{}.TableName()+".uid = ?", this.UId)
	}
	if this.Username != "" {
		db.Where("Member.username = ?", this.Username)
	}

	if this.StartTime != "" {
		db.Where(model.Withdraw{}.TableName()+".create_time > ?", common.DateToUnix(this.StartTime))
	}
	if this.EndTime != "" {
		db.Where(model.Withdraw{}.TableName()+".create_time < ?", common.DateToUnix(this.EndTime))
	}
	if this.Status > 0 {
		db.Where(model.Withdraw{}.TableName()+".status = ?", this.Status)
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

type WithdrawUpdateInfoService struct {
	request.WithdrawUpdateInfoRequest
}

func (this WithdrawUpdateInfoService) UpdateInfo() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Withdraw{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	if this.BankName != "" {
		m.BankName = this.BankName
	}
	m.BranchBank = this.BranchBank
	if this.RealName != "" {
		m.RealName = this.RealName
	}
	if this.CardNumber != "" {
		m.CardNumber = this.CardNumber
	}
	//更新状态 说明 操作者
	return m.Update("bank_name", "branch_bank", "real_name", "card_number")
}
