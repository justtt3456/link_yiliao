package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/model"
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
	m := model.Message{}
	where := "uid = ? or uid = ? and status = ?"
	args := []interface{}{member.ID, -1,model.StatusOk}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Message, 0)
	for _, v := range list {
		item := response.Message{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, item)
	}
	return response.MessageData{List: res, Page: FormatPage(page)}
}
type MessageRead struct {
	request.Msg
}
func (this MessageRead)Read()(error)  {
	if this.Id == 0 {
		return errors.New("ID不能为空")
	}
	m := model.Message{ID:this.Id}
	if !m.Get() {
		return errors.New("信息不存在")
	}
	m.IsRead = 2
	return m.Update("is_read")
}