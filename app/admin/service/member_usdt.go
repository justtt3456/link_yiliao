package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/model"
	"errors"
	//"yiliao/model"
)

type MemberUsdtList struct {
	request.MemberUsdtList
}

func (this MemberUsdtList) List() response.MemberUsdtList {
	MemberUsdt := model.MemberUsdt{
		UId: this.UId,
	}
	list := MemberUsdt.List()
	res := make([]response.MemberUsdt, 0)
	for _, v := range list {
		i := response.MemberUsdt{
			Id:       v.Id,
			Address:  v.Address,
			Protocol: v.Protocol,
		}
		res = append(res, i)
	}
	return response.MemberUsdtList{
		List: res,
	}
}

type MemberUsdtUpdate struct {
	request.MemberUsdtUpdate
}

func (this MemberUsdtUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}

	if this.Address == "" {
		return errors.New("地址不能为空")
	}
	if this.Protocol == "" {
		return errors.New("网络不能为空")
	}
	MemberUsdt := model.MemberUsdt{
		Id: this.Id,
	}
	if !MemberUsdt.Get() {
		return errors.New("记录不存在")
	}
	MemberUsdt.Address = this.Address
	MemberUsdt.Protocol = this.Protocol
	return MemberUsdt.Update("address", "protocol")
}

type MemberUsdtRemove struct {
	request.MemberUsdtRemove
}

func (this MemberUsdtRemove) Remove() error {

	MemberUsdt := model.MemberUsdt{
		Id: this.Id,
	}
	if !MemberUsdt.Get() {
		return errors.New("记录不存在")
	}
	return MemberUsdt.Remove()
}
