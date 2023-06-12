package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type BankList struct {
	request.BankList
}

func (this BankList) PageList() response.BankData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Bank{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.BankInfo, 0)
	for _, v := range list {
		i := response.BankInfo{
			Id:         v.Id,
			BankName:   v.BankName,
			Sort:       v.Sort,
			Status:     v.Status,
			Lang:       v.Lang,
			Code:       v.Code,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		res = append(res, i)
	}
	return response.BankData{List: res, Page: FormatPage(page)}
}
func (this BankList) getWhere() (string, []interface{}) {
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

type BankCreate struct {
	request.BankCreate
}

func (this BankCreate) Create() error {
	if this.BankName == "" {
		return errors.New("银行名称不能为空")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Code == "" {
		return errors.New("银行编码不能为空")
	}
	m := model.Bank{
		BankName: this.BankName,
		Sort:     this.Sort,
		Status:   this.Status,
		Lang:     this.Lang,
		Code:     this.Code,
	}
	return m.Insert()
}

type BankUpdate struct {
	request.BankUpdate
}

func (this BankUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.BankName == "" {
		return errors.New("银行名称不能为空")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Code == "" {
		return errors.New("银行编码不能为空")
	}
	m := model.Bank{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("银行不存在")
	}
	m.BankName = this.BankName
	m.Sort = this.Sort
	m.Status = this.Status
	m.Lang = this.Lang
	m.Code = this.Code
	return m.Update("bank_name", "status", "sort", "lang", "code")
}

type BankUpdateStatus struct {
	request.BankUpdateStatus
}

func (this BankUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Bank{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("银行不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type BankRemove struct {
	request.BankRemove
}

func (this BankRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Bank{
		Id: this.Id,
	}
	return m.Remove()
}
