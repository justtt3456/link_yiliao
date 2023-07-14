package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/model"
	"errors"
	//"yiliao/model"
)

type MemberBankList struct {
	request.MemberBankList
}

func (this MemberBankList) List() response.MemberBankList {
	memberBank := model.MemberBank{
		UId: this.UId,
	}
	list := memberBank.List()
	res := make([]response.MemberBank, 0)
	for _, v := range list {
		i := response.MemberBank{
			Id:         v.Id,
			BankName:   v.BankName,
			CardNumber: v.CardNumber,
			BranchBank: v.BranchBank,
			RealName:   v.RealName,
		}
		res = append(res, i)
	}
	return response.MemberBankList{
		List: res,
	}
}

type MemberBankUpdate struct {
	request.MemberBankUpdate
}

func (this MemberBankUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}

	if this.BankName == "" {
		return errors.New("银行不能为空")
	}
	if this.CardNumber == "" {
		return errors.New("银行卡号不能为空")
	}
	if this.RealName == "" {
		return errors.New("真实姓名不能为空")
	}
	if this.BranchBank == "" {
		return errors.New("支行不能为空")
	}
	memberBank := model.MemberBank{
		Id: this.Id,
	}
	if !memberBank.Get() {
		return errors.New("银行卡不存在")
	}
	memberBank.BankName = this.BankName
	memberBank.RealName = this.RealName
	memberBank.CardNumber = this.CardNumber
	memberBank.BranchBank = this.BranchBank
	return memberBank.Update("real_name", "card_number", "branch_bank", "bank_name")
}

type MemberBankRemove struct {
	request.MemberBankRemove
}

func (this MemberBankRemove) Remove() error {

	memberBank := model.MemberBank{
		Id: this.Id,
	}
	if !memberBank.Get() {
		return errors.New("银行卡不存在")
	}
	return memberBank.Remove()
}
