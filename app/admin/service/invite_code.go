package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
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
			Id:        v.Id,
			UId:       v.UId,
			Username:  v.Username,
			AgentId:   v.AgentId,
			AgentName: v.AgentName,
			Code:      v.Code,
			//RegCount:   v.RegCount,
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
	if this.AgentId == 0 {
		return errors.New("代理不能为空")
	}
	if this.Code == "" {
		return errors.New("邀请码不能为空")
	}
	agent := model.Agent{Id: this.AgentId}
	if !agent.Get() {
		return errors.New("代理不存在")
	}
	m := model.InviteCode{
		AgentId:   this.AgentId,
		AgentName: agent.Name,
		Code:      this.Code,
	}
	return m.Insert()
}

type InviteCodeUpdate struct {
	request.InviteCodeUpdate
}

func (this InviteCodeUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Code == "" {
		return errors.New("邀请码不能为空")
	}
	m := model.InviteCode{Id: this.Id}
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
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.InviteCode{
		Id: this.Id,
	}
	return m.Remove()
}
