package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MedicineOrderListService struct {
	request.MedicineOrderListRequest
}

func (this MedicineOrderListService) PageList() *response.MedicineOrderPageListResponse {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.MedicineOrder{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	items := make([]response.MedicineOrder, 0)
	var totalAmount decimal.Decimal
	for _, v := range list {
		agent := model.Agent{Id: v.Member.AgentId}
		if v.Member.AgentId > 0 {
			agent.Get()
		}
		items = append(items, response.MedicineOrder{
			Username:          v.Member.Username,
			Uid:               v.Member.Id,
			Name:              v.Medicine.Name,
			CreateTime:        int(v.CreateTime),
			Amount:            v.PayMoney,
			RealName:          v.Member.RealName,
			AgentName:         agent.Account,
			Address:           v.Address.Name + v.Address.Phone + v.Address.Address + v.Address.Other,
			WithdrawThreshold: v.WithdrawThreshold,
			Interval:          v.Interval,
			Status:            v.Status,
			Current:           v.Current,
		})
		totalAmount = totalAmount.Add(v.PayMoney)
	}
	return &response.MedicineOrderPageListResponse{List: items, Page: FormatPage(page), TotalAmount: totalAmount}
}
func (this MedicineOrderListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.AgentName != "" {
		agent := model.Agent{Account: this.AgentName}
		if agent.Get() {
			where["Member.agent_id"] = agent.Id
		}
	}
	if this.MedicineName != "" {
		where["Medicine.name"] = this.MedicineName
	}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	if this.RealName != "" {
		where["Member.real_name"] = this.RealName
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
