package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
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
			Id:         v.Id,
			UId:        v.UId,
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
	if this.UId == "" {
		return errors.New("接收用户错误")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	split := strings.Split(this.UId, ",")
	if len(split) == 0 {
		return errors.New("接收用户错误")
	}
	m := model.Message{
		UId:     this.UId,
		Title:   this.Title,
		Content: this.Content,
		Status:  model.StatusClose,
	}
	return m.Insert()
}

type MessageUpdate struct {
	request.MessageUpdate
}

func (this MessageUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	m := model.Message{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("站内信不存在")
	}
	//不能改接收用户
	m.Title = this.Title
	m.Content = this.Content
	//m.Status = this.Status
	return m.Update("title", "content")
}

type MessageUpdateStatus struct {
	request.MessageUpdateStatus
}

func (this MessageUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Message{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("站内信不存在")
	}
	if m.Status == model.StatusOk {
		return errors.New("请勿重复发放")
	}
	if this.Status == model.StatusOk {
		if m.UId == "-1" { //全部用户
			member := model.Member{}
			members := member.List("status = ?", []interface{}{model.StatusOk})
			for _, v := range members {
				msg := model.MemberMessage{
					Uid:       v.Id,
					MessageId: m.Id,
					IsRead:    model.StatusClose,
				}
				err := msg.Insert()
				if err != nil {
					return err
				}
			}
		} else {
			for _, v := range strings.Split(m.UId, ",") {
				uid, _ := strconv.Atoi(v)
				if uid > 0 {
					msg := model.MemberMessage{
						Uid:       uid,
						MessageId: m.Id,
						IsRead:    model.StatusClose,
					}
					err := msg.Insert()
					if err != nil {
						return err
					}
				}
			}
		}
	}
	m.Status = this.Status
	return m.Update("status")
}

type MessageRemove struct {
	request.MessageRemove
}

func (this MessageRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Message{
		Id: this.Id,
	}
	return m.Remove()
}
