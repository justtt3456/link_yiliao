package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type PaymentListService struct {
	request.PaymentListRequest
}
type PaymentAddService struct {
	request.PaymentAddRequest
}
type PaymentUpdateService struct {
	request.PaymentUpdateRequest
}
type PaymentRemoveService struct {
	request.PaymentRemoveRequest
}

func (this PaymentListService) PageList() *response.PaymentData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.Payment{}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Payment, 0)
	for _, v := range list {
		i := response.Payment{
			ID:             v.ID,
			PayName:        v.PayName,
			RechargeURL:    v.RechargeURL,
			WithdrawURL:    v.WithdrawURL,
			NotifyURL:      v.NotifyURL,
			MerchantNo:     v.MerchantNo,
			Secret:         v.Secret,
			PriKey:         v.PriKey,
			PubKey:         v.PubKey,
			ClassName:      v.ClassName,
			WithdrawStatus: v.WithdrawStatus,
			CreateTime:     v.CreateTime,
			UpdateTime:     v.UpdateTime,
			Type:           v.Type,
		}
		res = append(res, i)
	}
	return &response.PaymentData{
		List: res,
		Page: FormatPage(page),
	}
}
func (this PaymentAddService) Add() error {
	if this.PayName == "" {
		return errors.New("支付名称不能为空")
	}
	if this.RechargeURL == "" {
		return errors.New("充值地址不能为空")
	}
	if this.WithdrawURL == "" {
		return errors.New("提现地址不能为空")
	}
	if this.NotifyURL == "" {
		return errors.New("回调地址不能为空")
	}
	if this.MerchantNo == "" {
		return errors.New("商户号不能为空")
	}
	if this.ClassName == "" {
		return errors.New("类名不能为空")
	}
	if this.Type == 0 {
		return errors.New("支付类型不能为空")
	}
	m := model.Payment{
		PayName:        this.PayName,
		RechargeURL:    this.RechargeURL,
		WithdrawURL:    this.WithdrawURL,
		NotifyURL:      this.NotifyURL,
		MerchantNo:     this.MerchantNo,
		Secret:         this.Secret,
		PriKey:         this.PriKey,
		PubKey:         this.PubKey,
		ClassName:      this.ClassName,
		WithdrawStatus: this.WithdrawStatus,
		Type:           this.Type,
	}
	return m.Insert()
}
func (this PaymentUpdateService) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.PayName == "" {
		return errors.New("支付名称不能为空")
	}
	if this.RechargeURL == "" {
		return errors.New("充值地址不能为空")
	}
	if this.WithdrawURL == "" {
		return errors.New("提现地址不能为空")
	}
	if this.NotifyURL == "" {
		return errors.New("回调地址不能为空")
	}
	if this.MerchantNo == "" {
		return errors.New("商户号不能为空")
	}
	if this.ClassName == "" {
		return errors.New("类名不能为空")
	}
	m := model.Payment{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("支付不存在")
	}
	m.PayName = this.PayName
	m.RechargeURL = this.RechargeURL
	m.WithdrawURL = this.WithdrawURL
	m.NotifyURL = this.NotifyURL
	m.MerchantNo = this.MerchantNo
	m.Secret = this.Secret
	m.PriKey = this.PriKey
	m.PubKey = this.PubKey
	m.ClassName = this.ClassName
	m.Type = this.Type
	return m.Update("pay_name", "type", "recharge_url", "withdraw_url", "notify_url", "merchant_no", "secret", "pri_key", "pub_key", "class_name", "withdraw_status")
}
func (this PaymentRemoveService) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Payment{
		ID: this.ID,
	}
	return m.Remove()
}

func (this PaymentListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.PayName != "" {
		where["pay_name"] = this.PayName
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
