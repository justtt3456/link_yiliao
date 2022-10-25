package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type WithdrawListService struct {
	request.WithdrawListRequest
}
type WithdrawUpdateService struct {
	request.WithdrawUpdateRequest
}

func (this WithdrawListService) PageList() response.WithdrawData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.Withdraw{}
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	res := make([]response.WithdrawInfo, 0)
	for _, v := range list {
		i := response.WithdrawInfo{
			ID:           v.ID,
			UID:          v.UID,
			WithdrawType: v.WithdrawType,
			BankName:     v.BankName,
			BranchBank:   v.BranchBank,
			RealName:     v.RealName,
			CardNumber:   v.CardNumber,
			BankPhone:    v.BankPhone,
			Amount:       float64(v.Amount) / model.UNITY,
			Fee:          float64(v.Fee) / model.UNITY,
			TotalAmount:  float64(v.TotalAmount) / model.UNITY,
			UsdtAmount:   float64(v.UsdtAmount) / model.UNITY,
			Description:  v.Description,
			Operator:     v.Operator1,
			ViewStatus:   v.ViewStatus,
			Status:       v.Status,
			SuccessTime:  v.SuccessTime,
			OrderSn:      v.OrderSn,
			PaymentID:    v.PaymentID,
			PaymentName:  v.Payment.PayName,
			TradeSn:      v.TradeSn,
			CreateTime:   v.CreateTime,
			UpdateTime:   v.UpdateTime,
			Username:     v.Member.Username,
			MethodName:   v.WithdrawMethod.Name,
		}
		res = append(res, i)
	}
	return response.WithdrawData{
		List: res,
		Page: FormatPage(page),
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
			ID: id,
		}
		if !m.Get() {
			return errors.New("记录不存在")
		}
		if m.Status != model.StatusReview {
			return errors.New("当前记录不允许操作")
		}
		user := model.Member{ID: m.UID}
		if !user.Get() {
			return errors.New("用户不存在")
		}
		switch this.Status {
		case model.StatusRollback:
			//生成账单
			trade := model.Trade{
				UID:       m.UID,
				TradeType: 4,
				ItemID:    m.ID,
				Amount:    m.TotalAmount,
				Before:    user.Balance,
				After:     user.Balance + m.TotalAmount,
				Desc:      "提现驳回",
			}
			if err := trade.Insert(); err != nil {
				return err
			}
			//回滚余额
			user.UseBalance += m.TotalAmount
			user.TotalBalance += m.TotalAmount
			if err := user.Update("balance", "total_balance"); err != nil {
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
			return  err
		}
	}

	return nil
}

func (this WithdrawListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.UID > 0 {
		where[model.Withdraw{}.TableName()+".uid"] = this.UID
	}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}

	if this.StartTime != "" {
		where[model.Withdraw{}.TableName()+".create_time >"] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where[model.Withdraw{}.TableName()+".create_time <"] = common.DateToUnix(this.EndTime)
	}
	if this.Status > 0 {
		where[model.Withdraw{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type WithdrawUpdateInfoService struct {
	request.WithdrawUpdateInfoRequest
}

func (this WithdrawUpdateInfoService) UpdateInfo() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Withdraw{
		ID: this.ID,
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
