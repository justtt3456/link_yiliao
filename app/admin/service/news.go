package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type NewsList struct {
	request.NewsList
}

func (this NewsList) PageList() response.NewsData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.News{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.News, 0)
	for _, v := range list {
		i := response.News{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Sort:       v.Sort,
			Intro:      v.Intro,
			Cover:      v.Cover,
			Status:     v.Status,
			Lang:       v.Lang,
		}
		res = append(res, i)
	}
	return response.NewsData{List: res, Page: FormatPage(page)}
}
func (this NewsList) getWhere() (string, []interface{}) {
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

type NewsCreate struct {
	request.NewsCreate
}

func (this NewsCreate) Create() error {
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
	if this.Cover == "" {
		return errors.New("封面不能为空")
	}
	m := model.News{
		Title:   this.Title,
		Content: this.Content,
		Status:  this.Status,
		Sort:    this.Sort,
		Intro:   this.Intro,
		Cover:   this.Cover,
		Lang:    this.Lang,
	}
	return m.Insert()
}

type NewsUpdate struct {
	request.NewsUpdate
}

func (this NewsUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
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
	if this.Cover == "" {
		return errors.New("封面不能为空")
	}
	m := model.News{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("资讯不存在")
	}
	m.Lang = this.Lang
	m.Title = this.Title
	m.Intro = this.Intro
	m.Cover = this.Cover
	m.Content = this.Content
	m.Sort = this.Sort
	m.Status = this.Status
	return m.Update("lang", "title", "sort", "lang", "intro", "content", "cover")
}

type NewsUpdateStatus struct {
	request.NewsUpdateStatus
}

func (this NewsUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.News{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("资讯不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type NewsRemove struct {
	request.NewsRemove
}

func (this NewsRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.News{
		ID: this.ID,
	}
	return m.Remove()
}
