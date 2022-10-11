package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
)

type Invest struct {
}

func (this Invest) Get() (*response.InvestInfo, error) {
	invest := model.Invest{}
	if !invest.Get() {
		return nil, errors.New("参数错误")
	}
	return &response.InvestInfo{
		ID:             invest.ID,
		Name:           invest.Name,
		Ratio:          float64(invest.Ratio) / model.UNITY,
		FreezeDay:      invest.FreezeDay,
		IncomeInterval: invest.IncomeInterval,
		Status:         invest.Status,
		Description:    invest.Description,
		MinAmount:      float64(invest.MinAmount) / model.UNITY,
		CreateTime:     invest.CreateTime,
		UpdateTime:     invest.UpdateTime,
	}, nil
}

type InvestUpdate struct {
	request.InvestUpdate
}

func (this InvestUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("理财名称不能为空")
	}
	if this.Ratio == 0 {
		return errors.New("收益率不能为空")
	}

	if this.IncomeInterval == 0 {
		return errors.New("收益间隔不能为空")
	}
	m := model.Invest{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("记录不存在")
	}
	m.Name = this.Name
	m.Ratio = int(this.Ratio * model.UNITY)
	m.FreezeDay = this.FreezeDay
	m.IncomeInterval = this.IncomeInterval
	m.Status = this.Status
	m.Description = this.Description
	m.MinAmount = int64(this.MinAmount * model.UNITY)
	return m.Update("name", "ratio", "freeze_day", "income_interval", "status", "description", "min_amount")
}

type InvestOrderList struct {
	request.InvestOrder
}

func (this InvestOrderList) PageList() response.InvestOrderData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.InvestOrder{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InvestOrder, 0)
	for _, v := range list {
		i := response.InvestOrder{
			ID:           v.ID,
			UID:          v.UID,
			Username:     v.Member.Username,
			Type:         v.Type,
			Amount:       float64(v.Amount) / model.UNITY,
			CreateTime:   v.CreateTime,
			UnfreezeTime: v.UnfreezeTime,
			IncomeTime:   v.IncomeTime,
			Balance:      float64(v.Balance) / model.UNITY,
		}
		res = append(res, i)
	}
	return response.InvestOrderData{List: res, Page: FormatPage(page)}
}
func (this InvestOrderList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type InvestIncomeList struct {
	request.InvestIncome
}

func (this InvestIncomeList) PageList() response.InvestIncomeData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.InvestLog{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InvestIncome, 0)
	for _, v := range list {
		i := response.InvestIncome{
			ID:         v.ID,
			UID:        v.UID,
			Username:   v.Member.Username,
			Income:     float64(v.Income) / model.UNITY,
			Balance:    float64(v.Balance) / model.UNITY,
			CreateTime: v.CreateTime,
		}
		res = append(res, i)
	}
	return response.InvestIncomeData{List: res, Page: FormatPage(page)}
}
func (this InvestIncomeList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Username != "" {
		where["Member.username"] = this.Username
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
