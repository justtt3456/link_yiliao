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
			ID:          v.ID,
			OrderSn:     v.OrderSn,
			UID:         v.UID,
			Type:        v.Type,
			Amount:      float64(v.Amount) / model.UNITY,
			RealAmount:  float64(v.RealAmount) / model.UNITY,
			From:        v.From,
			To:          v.To,
			Voucher:     v.Voucher,
			PaymentID:   v.PaymentID,
			Status:      v.Status,
			UsdtAmount:  float64(v.UsdtAmount) / model.UNITY,
			Operator:    v.Operator,
			Description: v.Description,
			UpdateTime:  v.UpdateTime,
			CreateTime:  v.CreateTime,
			Username:    v.Member.Username,
			MethodName:  v.RechargeMethod.Name,
			PaymentName: v.Payment.PayName,
			SuccessTime: v.SuccessTime,
			TradeSn:     v.TradeSn,
			ImageUrl:    v.ImageUrl,
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
			ID: id,
		}
		if !m.Get() {
			return errors.New("记录不存在")
		}
		if m.Status != model.StatusReview {
			return errors.New("当前状态无法修改")
		}
		member := model.Member{ID: m.UID}
		if !member.Get() {
			return errors.New("用户不存在")
		}
		if this.Status == model.StatusAccept {
			r := RechargeHandle{}
			r.Recharge(member, m.ID, m.Amount, 1, 3, 1)
		}
		m.Status = this.Status
		m.Description = this.Description
		m.Operator = admin.ID
		//更新状态 说明 操作者
		m.Update("status", "description", "operator")
	}
	return nil
}

func (this RechargeService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.UID > 0 {
		where[model.Recharge{}.TableName()+".uid"] = this.UID
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

func (RechargeHandle) Recharge(member model.Member, item int, amount int64, way int, tradeType int, isfront int) error {
	//账单
	trade := model.Trade{
		UID:        member.ID,
		TradeType:  tradeType,
		ItemID:     item,
		Amount:     amount,
		Before:     member.Balance,
		After:      member.Balance + amount,
		IsFrontend: isfront,
	}
	switch way {
	case 1: //审核
		trade.Desc = "充值审核通过"
	case 2: //系统
		trade.Desc = "福利派送"
	}
	err := trade.Insert()
	if err != nil {
		logrus.Error(err)
		return err
	}
	//上分
	member.Balance += amount
	member.TotalBalance += amount
	return member.Update("balance", "total_balance")

}
