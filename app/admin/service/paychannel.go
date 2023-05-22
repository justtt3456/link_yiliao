package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
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
			//Min:         float64(v.Min) ,
			//Max:         float64(v.Max) ,
			Status:     v.Status,
			Category:   v.Category,
			Sort:       v.Sort,
			Icon:       v.Icon,
			Fee:        v.Fee,
			Lang:       v.Lang,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
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
	if this.Code == "" {
		return errors.New("通道编码不能为空")
	}
	//if this.Min == 0 {
	//	return errors.New("最小值不能为空")
	//}
	//if this.Max == 0 {
	//	return errors.New("最大值不能为空")
	//}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	m := model.PayChannel{
		Name:      this.Name,
		PaymentId: this.PaymentId,
		Code:      this.Code,
		//Min:       int64(this.Min),
		//Max:       int64(this.Max),
		Icon: this.Icon,
		Fee:  this.Fee,
		Lang: this.Lang,
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
	if this.Code == "" {
		return errors.New("通道编码不能为空")
	}
	//if this.Min == 0 {
	//	return errors.New("最小值不能为空")
	//}
	//if this.Max == 0 {
	//	return errors.New("最大值不能为空")
	//}
	if this.Lang == "" {
		return errors.New("语言不能为空")
	}
	if this.Lang == "" {
		return errors.New("语言不能为空")
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
	//m.Min = int64(this.Min)
	//m.Max = int64(this.Max)
	m.Icon = this.Icon
	m.Fee = this.Fee
	m.Lang = this.Lang
	return m.Update("name", "payment_id", "code", "min", "max", "icon", "fee", "lang")
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
