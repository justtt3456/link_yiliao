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

type PayChannelListService struct {
	request.PayChannelListRequest
}
type PayChannelCreateService struct {
	request.PayChannelCreateRequest
}
type PayChannelUpdateService struct {
	request.PayChannelUpdateRequest
}
type PayChannelRemoveService struct {
	request.PayChannelRemoveRequest
}
type PayChannelUpdateStatusService struct {
	request.PayChannelUpdateStatusRequest
}

func (this PayChannelListService) PageList() *response.PayChannelData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.PayChannel{}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.PayChannel, 0)
	for _, v := range list {
		i := response.PayChannel{
			Id:          v.Id,
			Name:        v.Name,
			PaymentId:   v.PaymentId,
			PaymentName: v.Payment.PayName,
			Code:        v.Code,
			Min:         (v.Min),
			Max:         (v.Max),
			Status:      v.Status,
			MethodId:    v.MethodId,
			Sort:        v.Sort,
			Icon:        v.Icon,
			Fee:         v.Fee,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		res = append(res, i)
	}
	return &response.PayChannelData{
		List: res,
		Page: FormatPage(page),
	}
}

func (this PayChannelCreateService) Create() error {
	if this.Name == "" {
		return errors.New("通道名称不能为空")
	}
	if this.PaymentId == 0 {
		return errors.New("支付名称不能为空")
	}
	if this.MethodId == 0 {
		return errors.New("支付方式不能为空")
	}
	if this.Code == "" {
		return errors.New("通道编码不能为空")
	}
	if this.Min.LessThanOrEqual(decimal.Zero) {
		return errors.New("最小值不能为空")
	}
	if this.Max.LessThan(this.Min) {
		return errors.New("最大值错误")
	}
	m := model.PayChannel{
		Name:      this.Name,
		PaymentId: this.PaymentId,
		Code:      this.Code,
		Min:       this.Min,
		Max:       this.Max,
		Icon:      this.Icon,
		Fee:       this.Fee,
		MethodId:  this.MethodId,
	}
	return m.Insert()
}
func (this PayChannelUpdateService) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("通道名称不能为空")
	}
	if this.PaymentId == 0 {
		return errors.New("支付名称不能为空")
	}
	if this.MethodId == 0 {
		return errors.New("支付方式不能为空")
	}
	if this.Code == "" {
		return errors.New("通道编码不能为空")
	}
	if this.Min.LessThanOrEqual(decimal.Zero) {
		return errors.New("最小值不能为空")
	}
	if this.Max.LessThan(this.Min) {
		return errors.New("最大值错误")
	}
	m := model.PayChannel{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("通道不存在")
	}
	m.Name = this.Name
	m.PaymentId = this.PaymentId
	m.Code = this.Code
	m.Min = (this.Min)
	m.Max = (this.Max)
	m.Icon = this.Icon
	m.Fee = this.Fee
	m.MethodId = this.MethodId
	return m.Update("name", "payment_id", "code", "min", "max", "icon", "fee", "method_id")
}
func (this PayChannelRemoveService) Remove() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.PayChannel{
		Id: this.Id,
	}
	return m.Remove()
}
func (this PayChannelUpdateStatusService) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.PayChannel{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("通道不存在")
	}
	if m.Status == this.Status {
		return nil
	}
	m.Status = this.Status
	return m.Update("status")
}
func (this PayChannelListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{
		"1": 1,
	}
	if this.Name != "" {
		where["name"] = this.Name
	}
	if this.PaymentId != 0 {
		where["payment_id"] = this.PaymentId
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
