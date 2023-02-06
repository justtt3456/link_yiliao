package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/lang"
	"finance/model"
	//"finance/model"
)

type MemberBankList struct {
}

func (MemberBankList) List(member model.Member) response.MemberBankList {
	memberBank := model.MemberBank{
		UID: member.ID,
	}
	list := memberBank.List()
	res := make([]response.MemberBank, 0)
	for _, v := range list {
		i := response.MemberBank{
			ID:         v.ID,
			BankName:   v.BankName,
			CardNumber: v.CardNumber,
			Province:   v.Province,
			City:       v.City,
			BranchBank: v.BranchBank,
			RealName:   v.RealName,
			IDNumber:   v.IDNumber,
			BankPhone:  v.BankPhone,
			IsDefault:  v.IsDefault,
		}
		res = append(res, i)
	}
	return response.MemberBankList{
		List: res,
	}
}

type MemberBankCreate struct {
	request.MemberBankCreate
}

func (this MemberBankCreate) Create(member model.Member) error {
	if this.BankName == "" {
		return errors.New(lang.Lang("Bank cannot be empty"))
	}
	if this.CardNumber == "" {
		return errors.New(lang.Lang("Bank card number cannot be empty"))
	}
	if this.RealName == "" {
		return errors.New(lang.Lang("Cardholder cannot be empty"))
	}

	//允许N张银行卡
	memberBank := model.MemberBank{
		UID:        member.ID,
		BankName:   this.BankName,
		CardNumber: this.CardNumber,
		BranchBank: this.BranchBank,
		RealName:   this.RealName,
	}
	return memberBank.Insert()
}

type MemberBankUpdate struct {
	request.MemberBankUpdate
}

func (this MemberBankUpdate) Update(member model.Member) error {
	if this.ID == 0 {
		return errors.New(lang.Lang("Parameter error"))
	}
	memberBank := model.MemberBank{
		ID:  this.ID,
		UID: member.ID,
	}
	if !memberBank.Get() {
		return errors.New(lang.Lang("Bank card does not exist"))
	}
	if this.RealName != "" {
		memberBank.RealName = this.RealName
	}
	if this.CardNumber != "" {
		memberBank.CardNumber = this.CardNumber
	}
	if this.BankName != "" {
		memberBank.BankName = this.BankName

	}
	if this.BranchBank != "" {
		memberBank.BranchBank = this.BranchBank
	}

	return memberBank.Update("real_name", "bank_name", "card_number", "branch_bank")
}

type MemberBankRemove struct {
	request.MemberBankRemove
}

func (this MemberBankRemove) Remove(member model.Member) error {
	memberBank := model.MemberBank{
		ID:  this.ID,
		UID: member.ID,
	}
	if !memberBank.Get() {
		return errors.New(lang.Lang("Bank card does not exist"))
	}
	return memberBank.Remove()
}

type MemberUsdtCreate struct {
	request.MemberUsdtCreate
}

func (this MemberUsdtCreate) Create(member model.Member) error {
	if this.Protocol == "" {
		return errors.New(lang.Lang("Protocol number cannot be empty"))
	}
	if this.Address == "" {
		return errors.New(lang.Lang("Address cannot be empty"))
	}

	memberBank := model.MemberUsdt{
		UID:      member.ID,
		Protocol: this.Protocol,
		Address:  this.Address,
	}
	return memberBank.Insert()
}

type MemberUsdtList struct {
}

func (MemberUsdtList) List(member model.Member) response.MemberUsdtList {
	memberUsdt := model.MemberUsdt{
		UID: member.ID,
	}
	list := memberUsdt.List()
	res := make([]response.MemberUsdt, 0)
	for _, v := range list {
		i := response.MemberUsdt{
			ID:       v.ID,
			Protocol: v.Protocol,
			Address:  v.Address,
		}
		res = append(res, i)
	}
	return response.MemberUsdtList{
		List: res,
	}
}
