package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type InviteCodeList struct {
	request.InviteCodeList
}

func (this InviteCodeList) PageList() response.InviteCodeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.InviteCode{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InviteCodeInfo, 0)
	for _, v := range list {
		i := response.InviteCodeInfo{
			ID:         v.ID,
			UID:        v.UID,
			Username:   v.Username,
			AgentID:    v.AgentID,
			AgentName:  v.AgentName,
			Code:       v.Code,
			RegCount:   v.RegCount,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.InviteCodeData{List: res, Page: FormatPage(page)}
}
func (this InviteCodeList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Code != "" {
		where["code"] = this.Code
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type InviteCodeCreate struct {
	request.InviteCodeCreate
}

func (this InviteCodeCreate) Create() error {
	if this.AgentID == 0 {
		return errors.New("代理不能为空")
	}
	if this.Code == "" {
		return errors.New("邀请码不能为空")
	}
	agent := model.Agent{ID: this.AgentID}
	if !agent.Get() {
		return errors.New("代理不存在")
	}
	m := model.InviteCode{
		AgentID:   this.AgentID,
		AgentName: agent.Name,
		Code:      this.Code,
	}
	return m.Insert()
}

type InviteCodeUpdate struct {
	request.InviteCodeUpdate
}

func (this InviteCodeUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Code == "" {
		return errors.New("邀请码不能为空")
	}
	m := model.InviteCode{ID: this.ID}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	m.Code = this.Code
	return m.Update("code")

}

type InviteCodeRemove struct {
	request.InviteCodeRemove
}

func (this InviteCodeRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.InviteCode{
		ID: this.ID,
	}
	return m.Remove()
}
