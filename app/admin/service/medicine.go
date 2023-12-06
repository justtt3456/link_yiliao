package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MedicineList struct {
	request.MedicineList
}

func (this MedicineList) PageList() response.MedicineData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Medicine{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Medicine, 0)
	for _, v := range list {
		i := response.Medicine{
			Id:                v.Id,
			Name:              v.Name,
			Price:             v.Price,
			Img:               v.Img,
			Desc:              v.Desc,
			WithdrawThreshold: v.WithdrawThreshold,
			Interval:          v.Interval,
			Sort:              v.Sort,
			Status:            v.Status,
			CreateTime:        v.CreateTime,
		}
		res = append(res, i)
	}
	return response.MedicineData{List: res, Page: FormatPage(page)}
}
func (this MedicineList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where[model.Medicine{}.TableName()+".name"] = this.Name
	}
	if this.Status > 0 {
		where[model.Medicine{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type MedicineCreate struct {
	request.MedicineCreate
}

func (this MedicineCreate) Create() error {
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Img == "" {
		return errors.New("图片不能为空")
	}
	if this.Price.LessThanOrEqual(decimal.Zero) {
		return errors.New("价格不能为空")
	}
	if this.WithdrawThreshold.LessThanOrEqual(decimal.Zero) {
		return errors.New("提现额度不能为空")
	}
	if this.Interval <= 0 {
		return errors.New("周期不能为空")
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	m := model.Medicine{
		Name:              this.Name,
		Price:             this.Price,
		Img:               this.Img,
		Desc:              this.Desc,
		WithdrawThreshold: this.WithdrawThreshold,
		Sort:              this.Sort,
		Status:            this.Status,
		Interval:          this.Interval,
	}
	return m.Insert()
}

type MedicineUpdate struct {
	request.MedicineUpdate
}

func (this MedicineUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Img == "" {
		return errors.New("图片不能为空")
	}
	if this.Price.LessThanOrEqual(decimal.Zero) {
		return errors.New("价格不能为空")
	}
	if this.WithdrawThreshold.LessThanOrEqual(decimal.Zero) {
		return errors.New("提现额度不能为空")
	}
	if this.Interval <= 0 {
		return errors.New("周期不能为空")
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	m := model.Medicine{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}
	m.Name = this.Name
	m.Price = this.Price
	m.Img = this.Img
	m.Desc = this.Desc
	m.WithdrawThreshold = this.WithdrawThreshold
	m.Interval = this.Interval
	m.Sort = this.Sort
	m.Status = this.Status
	return m.Update("name", "price", "img", "desc", "withdraw_threshold", "interval", "sort", "status")
}

type MedicineUpdateStatus struct {
	request.MedicineUpdateStatus
}

func (this MedicineUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Medicine{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type MedicineRemove struct {
	request.MedicineRemove
}

func (this MedicineRemove) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Medicine{
		Id: this.Id,
	}
	return m.Remove()
}
