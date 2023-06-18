package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type AgentList struct {
	request.AgentList
}

func (this AgentList) List() response.AgentData {
	m := model.Agent{}
	where := "parent_id = ? and status = ?"
	args := []interface{}{0, model.StatusOk}
	list := m.List(where, args)
	res := make([]response.AgentInfo, 0)
	for _, v := range list {
		//parent := model.Agent{
		//	Id: v.ParentId,
		//}
		//if v.ParentId != 0 {
		//	parent.Get()
		//}
		i := response.AgentInfo{
			Id:   v.Id,
			Name: v.Account,
			//ParentId:   v.ParentId,
			//ParentName: parent.Name,
			//GroupName:  v.GroupName,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.AgentData{List: res}
}
func (this AgentList) PageList() response.AgentData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Agent{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.AgentInfo, 0)
	for _, v := range list {
		invite := model.InviteCode{
			AgentId: v.Id,
		}
		invite.Get()
		i := response.AgentInfo{
			Id:         v.Id,
			Name:       v.Account,
			InviteCode: invite.Code,
			//ParentId:   v.ParentId,
			//ParentName: parent.Name,
			//GroupName:  v.GroupName,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.AgentData{List: res, Page: FormatPage(page)}
}
func (this AgentList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where["name"] = this.Name
	}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type AgentCreate struct {
	request.AgentCreate
}

func (this AgentCreate) Create() error {
	if this.Account == "" {
		return errors.New("代理账号不能为空")
	}
	if this.Password == "" {
		return errors.New("代理密码不能为空")
	}
	//groupName := this.GroupName
	//if this.ParentId > 0 {
	//	parent := model.Agent{Id: this.ParentId}
	//	if !parent.Get() {
	//		return errors.New("上级代理不存在")
	//	}
	//	//groupName = parent.GroupName
	//}
	salt := common.RandStringRunes(6)
	m := model.Agent{
		Account:  this.Account,
		Password: common.Md5String(this.Password + salt),
		Salt:     salt,
		//ParentId:  this.ParentId,
		//GroupName: groupName,
		Status: this.Status,
	}
	err := m.Insert()
	if err != nil {
		return err
	}
	//创建邀请码
	invite := model.InviteCode{}
	code := invite.InviteCode()
	invite.AgentId = m.Id
	invite.AgentName = m.Account
	invite.Code = code
	return invite.Insert()
}

type AgentUpdate struct {
	request.AgentUpdate
}

func (this AgentUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Agent{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("代理不存在")
	}
	//groupName := this.GroupName
	//if this.ParentId > 0 {
	//	parent := model.Agent{Id: this.ParentId}
	//	if !parent.Get() {
	//		return errors.New("上级代理不存在")
	//	}
	//	groupName = parent.GroupName
	//}
	//m.ParentId = this.ParentId
	//m.GroupName = groupName
	m.Status = this.Status
	if this.Password != "" {
		salt := common.RandStringRunes(6)
		m.Salt = salt
		m.Password = common.Md5String(this.Password + salt)
		return m.Update("status", "salt", "password")
	} else {
		return m.Update("status")
	}
}

type AgentUpdateStatus struct {
	request.AgentUpdateStatus
}

func (this AgentUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Agent{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("代理不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type AgentRemove struct {
	request.AgentRemove
}

func (this AgentRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Agent{
		Id: this.Id,
	}
	return m.Remove()
}
