package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	//"yiliao/model"
)

type MemberAddressList struct {
}

func (MemberAddressList) List(member model.Member) response.MemberAddressList {
	memberBank := model.MemberAddress{
		UId: member.Id,
	}
	list := memberBank.List()
	res := make([]response.MemberAddress, 0)
	for _, v := range list {
		i := response.MemberAddress{
			Id:      v.Id,
			Name:    v.Name,
			Phone:   v.Phone,
			Address: v.Address,
			Other:   v.Other,
		}
		res = append(res, i)
	}
	return response.MemberAddressList{
		List: res,
	}
}

type MemberAddressCreate struct {
	request.MemberAddressCreate
}

func (this MemberAddressCreate) Create(member model.Member) error {
	if this.Name == "" {
		return errors.New("收货人不能为空")
	}
	if this.Phone == "" {
		return errors.New("手机号码不能为空")
	}
	if this.Address == "" {
		return errors.New("地址不能为空")
	}
	//if this.Other == "" {
	//	return errors.New("详细地址不能为空")
	//}
	//允许N张银行卡
	memberBank := model.MemberAddress{
		UId:     member.Id,
		Name:    this.Name,
		Phone:   this.Phone,
		Address: this.Address,
		Other:   this.Other,
	}
	return memberBank.Insert()
}

type MemberAddressUpdate struct {
	request.MemberAddressUpdate
}

func (this MemberAddressUpdate) Update(member model.Member) error {
	if this.Id == 0 {
		return errors.New(lang.Lang("Parameter error"))
	}
	if this.Name == "" {
		return errors.New("收货人不能为空")
	}
	if this.Phone == "" {
		return errors.New("手机号码不能为空")
	}
	if this.Address == "" {
		return errors.New("地址不能为空")
	}
	//if this.Other == "" {
	//	return errors.New("详细地址不能为空")
	//}
	memberAddress := model.MemberAddress{
		Id:  this.Id,
		UId: member.Id,
	}
	if !memberAddress.Get() {
		return errors.New("收货地址不存在")
	}
	memberAddress.Name = this.Name
	memberAddress.Phone = this.Phone
	memberAddress.Address = this.Address
	memberAddress.Other = this.Other
	return memberAddress.Update("name", "phone", "address", "other")
}

type MemberAddressRemove struct {
	request.MemberAddressRemove
}

func (this MemberAddressRemove) Remove(member model.Member) error {
	memberAddress := model.MemberAddress{
		Id:  this.Id,
		UId: member.Id,
	}
	if !memberAddress.Get() {
		return errors.New("收货地址不存在")
	}
	//return errors.New("请联系客服修改")
	return memberAddress.Remove()
}
