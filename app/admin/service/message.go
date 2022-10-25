package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type MessageList struct {
	request.MessageList
}

func (this MessageList) PageList() response.MessageData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Message{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.MessageInfo, 0)
	for _, v := range list {
		i := response.MessageInfo{
			ID:         v.ID,
			UID:        v.UID,
			Title:      v.Title,
			Content:    v.Content,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.MessageData{List: res, Page: FormatPage(page)}
}
func (this MessageList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type MessageCreate struct {
	request.MessageCreate
}

func (this MessageCreate) Create() error {
	if this.UID == 0 {
		return errors.New("接收用户错误")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	m := model.Message{
		UID:     this.UID,
		Title:   this.Title,
		Content: this.Content,
		Status:  this.Status,
		IsRead:  1,
	}
	return m.Insert()
}

type MessageUpdate struct {
	request.MessageUpdate
}

func (this MessageUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.UID == 0 {
		return errors.New("接收用户错误")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	m := model.Message{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("站内信不存在")
	}
	//不能改接收用户
	m.Title = this.Title
	m.Content = this.Content
	m.Status = this.Status
	return m.Update("title", "content", "status")
}

type MessageUpdateStatus struct {
	request.MessageUpdateStatus
}

func (this MessageUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Message{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("站内信不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type MessageRemove struct {
	request.MessageRemove
}

func (this MessageRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Message{
		ID: this.ID,
	}
	return m.Remove()
}
