package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/model"
)

type MemberLevel struct {
}

func (MemberLevel) List() response.MemberLevelData {
	m := model.MemberLevel{}
	list := m.List()
	res := make([]response.MemberLevelInfo, 0)
	for _, v := range list {
		item := response.MemberLevelInfo{
			ID:   v.ID,
			Name: v.Name,
			Img:  v.Img,
		}
		res = append(res, item)
	}
	return response.MemberLevelData{List: res}
}

type MemberLevelUpdate struct {
	request.MemberLevelUpdate
}

func (this MemberLevelUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("等级名称不能为空")
	}
	if this.Img == "" {
		return errors.New("等级图标不能为空")
	}
	m := model.MemberLevel{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("用户等级不存在")
	}
	m.Name = this.Name
	m.Img = this.Img
	return m.Update("name", "img")
}
