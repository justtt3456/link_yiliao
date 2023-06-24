package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/global"
	"china-russia/model"
)

type Message struct {
	request.Pagination
}

func (this Message) PageList(member model.Member) response.MessageData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.MemberMessage{}
	where := m.TableName() + ".uid = ?"
	args := []interface{}{member.Id}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Message, 0)
	for _, v := range list {
		item := response.Message{
			Id:         v.Id,
			Title:      v.Message.Title,
			Content:    v.Message.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, item)
	}
	return response.MessageData{List: res, Page: FormatPage(page)}
}

type MessageRead struct {
	request.Msg
}

func (this MessageRead) Read(member *model.Member) error {
	return global.DB.Model(model.MemberMessage{}).Where("uid = ? and is_read = ?", member.Id, model.StatusClose).Update("is_read", model.StatusOk).Error
}
