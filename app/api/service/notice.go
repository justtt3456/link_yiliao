package service

import (
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/global"
	"finance/model"
)

type Notice struct {
	request.Pagination
}

func (this Notice) PageList() response.NoticeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Notice{}
	where := "lang = ?"
	args := []interface{}{global.Language}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Notice, 0)
	for _, v := range list {
		item := response.Notice{
			ID:         v.ID,
			Title:      v.Title,
			Intro:      v.Intro,
			Content:    v.Content,
			Lang:       v.Lang,
			CreateTime: v.CreateTime,
		}
		res = append(res, item)
	}
	return response.NoticeData{List: res, Page: FormatPage(page)}
}

type NoticeList struct {
	request.Pagination
}

func (this NoticeList) PageList() response.NoticeListResponse {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Notice{}
	where := "type = ? and status = ? and lang = ? "
	args := []interface{}{1, model.StatusOk, global.Language}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.NoticeInfo, 0)
	for _, v := range list {
		item := response.NoticeInfo{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, item)
	}
	return response.NoticeListResponse{List: res, Page: FormatPage(page)}
}
