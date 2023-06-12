package service

import (
	"china-russia/app/agent/swag/request"
	"china-russia/app/agent/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type OrderListService struct {
	request.OrderListRequest
}

func (this OrderListService) PageList(agent *model.Agent) *response.BuyListResp {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.OrderProduct{}
	where, args := this.getWhere(agent)
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.BuyList, 0)
	var totalAmount decimal.Decimal
	for _, v := range list {
		//订单状态
		orderStatus := 1
		if v.IsReturnCapital == 1 {
			orderStatus = 2
		}
		items = append(items, response.BuyList{
			Username: v.Member.Username,
			Uid:      v.Member.Id,
			Name:     v.Product.Name,
			BuyTime:  int(v.CreateTime),
			Amount:   v.PayMoney,
			Status:   orderStatus,
		})
		totalAmount = totalAmount.Add(v.PayMoney)
	}
	return &response.BuyListResp{List: items, Page: FormatPage(page), TotalAmount: totalAmount}
}
func (this OrderListService) getWhere(agent *model.Agent) (string, []interface{}) {
	where := map[string]interface{}{
		"Member.agent_id": agent.Id,
	}
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

func (this OrderListService) GuQuanPageList(agent *model.Agent) *response.BuyGuquanResp {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.OrderEquity{}
	where, args := this.getWhere1(agent)
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.BuyGuquan, 0)
	for _, v := range list {
		items = append(items, response.BuyGuquan{
			Id:         v.Id,
			Rate:       v.Rate,
			Username:   v.Member.Username,
			Uid:        v.Member.Id,
			Num:        v.PayMoney.Div(v.Equity.Price).IntPart(),
			Price:      v.Equity.Price,
			CreateTime: v.CreateTime,
			TotalPrice: v.PayMoney,
		})
	}
	return &response.BuyGuquanResp{List: items, Page: FormatPage(page)}
}
func (this OrderListService) getWhere1(agent *model.Agent) (string, []interface{}) {
	where := map[string]interface{}{
		"Member.agent_id": agent.Id,
	}
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
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Rate.LessThanOrEqual(decimal.Zero) || decimal.NewFromInt(100).LessThan(this.Rate) {
		return errors.New("参数错误")
	}
	order := model.OrderEquity{
		Id: this.Id,
	}

	if !order.Get() {
		return errors.New("订单不存在")
	}
	order.Rate = this.Rate
	return order.Update("rate")
}
