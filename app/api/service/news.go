package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/model"
)

type News struct {
	request.NewsPageListRequest
}

func (this News) PageList() response.NewsData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.News{}
	where := ""
	args := []interface{}{}
	if this.Category > 0 {
		where = "category = ?"
		args = []interface{}{this.Category}
	}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.News, 0)
	for _, v := range list {
		item := response.News{
			Id:         v.Id,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Sort:       v.Sort,
			Intro:      v.Intro,
			Cover:      v.Cover,
			Category:   v.Category,
			Source:     v.Source,
			DateTime:   v.DateTime,
		}
		res = append(res, item)
	}
	return response.NewsData{List: res, Page: FormatPage(page)}
}
