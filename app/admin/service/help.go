package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type HelpList struct {
	request.HelpList
}

func (this HelpList) PageList() response.HelpData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Help{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Help, 0)
	for _, v := range list {
		i := response.Help{
			Id:         v.Id,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Sort:       v.Sort,
			Status:     v.Status,
			Lang:       v.Lang,
			Category:   v.Category,
		}
		res = append(res, i)
	}
	return response.HelpData{List: res, Page: FormatPage(page)}
}
func (this HelpList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Status > 0 {
		where["status"] = this.Status
	}
	if this.Category > 0 {
		where["category"] = this.Category
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

type HelpCreate struct {
	request.HelpCreate
}

func (this HelpCreate) Create() error {
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	if this.Category == 0 {
		return errors.New("分类不能为空")
	}
	m := model.Help{
		Title:    this.Title,
		Content:  this.Content,
		Lang:     this.Lang,
		Sort:     this.Sort,
		Status:   this.Status,
		Category: this.Category,
	}
	return m.Insert()
}

type HelpUpdate struct {
	request.HelpUpdate
}

func (this HelpUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Title == "" {
		return errors.New("标题不能为空")
	}
	if this.Content == "" {
		return errors.New("内容不能为空")
	}
	if this.Category == 0 {
		return errors.New("分类不能为空")
	}
	m := model.Help{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	m.Lang = this.Lang
	m.Title = this.Title
	m.Content = this.Content
	m.Sort = this.Sort
	m.Status = this.Status
	m.Category = this.Category
	return m.Update("lang", "title", "sort", "lang", "content", "category")
}

type HelpUpdateStatus struct {
	request.HelpUpdateStatus
}

func (this HelpUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Help{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type HelpRemove struct {
	request.HelpRemove
}

func (this HelpRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Help{
		Id: this.Id,
	}
	return m.Remove()
}
