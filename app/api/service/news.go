package service

import (
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/model"
)

type News struct {
	request.Pagination
}

func (this News) PageList() response.NewsData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.News{}
	where := "lang = ?"
	args := []interface{}{global.Language}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.News, 0)
	for _, v := range list {
		item := response.News{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Sort:       v.Sort,
			Intro:      v.Intro,
			Cover:      v.Cover,
		}
		res = append(res, item)
	}
	return response.NewsData{List: res, Page: FormatPage(page)}
}
