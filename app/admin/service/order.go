package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type OrderListService struct {
	request.OrderListRequest
}

func (this OrderListService) PageList() *response.BuyListResp {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.OrderProduct{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.BuyList, 0)
	for i := range list {
		items = append(items, response.BuyList{
			Username: list[i].Member.Username,
			Uid:      list[i].Member.ID,
			Name:     list[i].Product.Name,
			BuyTime:  int(list[i].CreateTime),
			Amount:   float64(list[i].PayMoney) / model.UNITY,
			Status:   1,
		})
	}
	return &response.BuyListResp{List: items, Page: FormatPage(page)}
}
func (this OrderListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.ProductName != "" {
		where["Product.name"] = this.ProductName
	}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.Uid != 0 {
		where["Member.id"] = this.Uid
	}
	if this.StartTime != "" {
		where["create_time >="] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where["create_time <"] = common.DateToUnix(this.EndTime)
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

func (this OrderListService) GuQuanPageList() *response.BuyGuquanResp {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.OrderGuquan{}
	where, args := this.getWhere1()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.BuyGuquan, 0)
	for i := range list {
		items = append(items, response.BuyGuquan{
			Username:   list[i].Member.Username,
			Uid:        list[i].Member.ID,
			Num:        list[i].PayMoney / int64(model.UNITY) / list[i].Guquan.Price / int64(model.UNITY),
			Price:      float64(list[i].Guquan.Price) / model.UNITY,
			CreateTime: int64(list[i].CreateTime),
			TotalPrice: float64(list[i].PayMoney) / model.UNITY,
		})
	}
	return &response.BuyGuquanResp{List: items, Page: FormatPage(page)}
}
func (this OrderListService) getWhere1() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.ProductName != "" {
		where["Product.name"] = this.ProductName
	}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.Uid != 0 {
		where["Member.id"] = this.Uid
	}
	if this.StartTime != "" {
		where["create_time >="] = common.DateToUnix(this.StartTime)
	}
	if this.EndTime != "" {
		where["create_time <"] = common.DateToUnix(this.EndTime)
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type OrderUpdate struct {
	request.OrderUpdate
}

func (this OrderUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.CtlType == 0 {
		return errors.New("参数错误")
	}
	order := model.OrderProduct{
		ID: this.ID,
	}

	if !order.Get() {
		return errors.New("订单不存在")
	}

	return order.Update("ctl_type", "draw_result")
}
