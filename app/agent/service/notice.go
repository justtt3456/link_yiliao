package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type NoticeList struct {
	request.NoticeList
}

func (this NoticeList) PageList() response.NoticeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Notice{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Notice, 0)
	for _, v := range list {
		i := response.Notice{
			Id:         v.Id,
			Title:      v.Title,
			Intro:      v.Intro,
			Content:    v.Content,
			Type:       v.Type,
			Lang:       v.Lang,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.NoticeData{List: res, Page: FormatPage(page)}
}
func (this NoticeList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	if this.Lang != "" {
		where["lang"] = this.Lang
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type NoticeCreate struct {
	request.NoticeCreate
}

func (this NoticeCreate) Create() error {
	if this.Type != 1 && this.Type != 2 {
		return errors.New("公告类型错误")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Intro == "" {
		return errors.New("简介不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	m := model.Notice{
		Title:   this.Title,
		Intro:   this.Intro,
		Content: this.Content,
		Type:    this.Type,
		Lang:    this.Lang,
		Status:  this.Status,
	}
	return m.Insert()
}

type NoticeUpdate struct {
	request.NoticeUpdate
}

func (this NoticeUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Type != 1 && this.Type != 2 {
		return errors.New("公告类型错误")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Intro == "" {
		return errors.New("简介不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	m := model.Notice{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("公告不存在")
	}
	m.Lang = this.Lang
	m.Title = this.Title
	m.Intro = this.Intro
	m.Type = this.Type
	m.Content = this.Content
	m.Status = this.Status
	return m.Update("lang", "title", "lang", "intro", "content", "type")
}

type NoticeUpdateStatus struct {
	request.NoticeUpdateStatus
}

func (this NoticeUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Notice{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("公告不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type NoticeRemove struct {
	request.NoticeRemove
}

func (this NoticeRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Notice{
		Id: this.Id,
	}
	return m.Remove()
}
