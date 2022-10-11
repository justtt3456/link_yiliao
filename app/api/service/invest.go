package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/common"
	"finance/lang"
	"finance/model"
	"fmt"
	"time"
)

type InvestIndex struct {
	request.Pagination
}

func (this InvestIndex) Index(member model.Member) *response.InvestIndexData {
	//余额宝配置
	invest := model.Invest{}
	if !invest.Get() {
		return nil
	}
	if invest.Status != model.StatusOk {
		return nil
	}
	//累计收益 余额 昨日收益
	income := model.InvestLog{}
	where := "uid = ?"
	args := []interface{}{member.ID}
	//sum := income.Sum(where, args, "income")
	yesterday := income.YesterdayIncome(member.ID)
	m := response.InvestMember{
		Balance:         float64(member.InvestAmount+member.InvestFreeze) / model.UNITY,
		TotalIncome:     float64(member.InvestIncome) / model.UNITY,
		YesterdayIncome: float64(yesterday) / model.UNITY,
	}
	//收益记录
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	list, page := income.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InvestIncome, 0)
	for _, v := range list {
		i := response.InvestIncome{
			Income:     float64(v.Income) / model.UNITY,
			Balance:    float64(v.Balance) / model.UNITY,
			CreateTime: v.CreateTime,
		}
		res = append(res, i)
	}
	return &response.InvestIndexData{
		Info: response.InvestInfo{
			Name:           invest.Name,
			Ratio:          invest.Ratio,
			FreezeDay:      invest.FreezeDay,
			IncomeInterval: invest.IncomeInterval,
			Status:         invest.Status,
			Description:    invest.Description,
		},
		Income: response.InvestIncomeData{
			List: res,
			Page: FormatPage(page),
		},
		Member: m,
	}
}

type InvestOrder struct {
	request.InvestOrder
}

func (this InvestOrder) Insert(member model.Member) error {
	if this.Amount == 0 {
		return errors.New(lang.Lang("Wrong amount"))
	}
	if member.FundsStatus != model.StatusOk {
		return errors.New(lang.Lang("The current account has been frozen!"))
	}
	amount := int64(this.Amount * model.UNITY)
	now := time.Now().Unix()
	switch this.Type {
	case 1: //转入
		if member.Balance < amount {
			return errors.New(lang.Lang("Insufficient account balance"))
		}
		c := model.Invest{}
		if !c.Get() {
			return errors.New(lang.Lang("Parameter error"))
		}
		if c.Status != model.StatusOk {
			return errors.New(lang.Lang("Parameter error"))
		}
		if amount < c.MinAmount {
			return errors.New(fmt.Sprintf("最小转入金额%v", c.MinAmount))
		}
		today := common.GetTodayZero()
		m := model.InvestOrder{
			UID:            member.ID,
			Type:           this.Type,
			Amount:         amount,
			Rate:           c.Ratio,
			UnfreezeTime:   now + int64(c.FreezeDay*86400),
			IncomeTime:     today + int64(c.IncomeInterval*86400),
			Balance:        member.InvestAmount + member.InvestFreeze + amount,
			UnfreezeStatus: model.StatusClose,
		}
		if err := m.Insert(); err != nil {
			return err
		}
		//账变
		trade := model.Trade{
			UID:       member.ID,
			TradeType: 11,
			ItemID:    m.ID,
			Amount:    amount,
			Before:    member.Balance,
			After:     member.Balance - amount,
			Desc:      "余额宝转入",
		}
		trade.Insert()
		//扣除余额
		member.Balance -= amount
		member.InvestFreeze += amount
		return member.Update("balance", "invest_freeze")
	case 2: //转出
		//获取可转出余额 余额-不可转出
		if member.InvestAmount < amount {
			return errors.New(lang.Lang("Insufficient transferable balance"))
		}
		m := model.InvestOrder{
			UID:     member.ID,
			Type:    this.Type,
			Amount:  amount,
			Balance: member.InvestAmount + member.InvestFreeze - amount,
		}
		if err := m.Insert(); err != nil {
			return err
		}
		//账变
		trade := model.Trade{
			UID:       member.ID,
			TradeType: 12,
			ItemID:    m.ID,
			Amount:    amount,
			Before:    member.Balance,
			After:     member.Balance + amount,
			Desc:      "余额宝转出",
		}
		trade.Insert()
		//增加余额
		member.Balance += amount
		member.InvestAmount -= amount
		return member.Update("balance", "invest_amount")
	default:
		return errors.New(lang.Lang("Parameter error"))
	}
}

type InvestOrderList struct {
	request.Pagination
}

func (this InvestOrderList) PageList(member model.Member) response.InvestOrderData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.InvestOrder{}
	where := "uid = ?"
	args := []interface{}{member.ID}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InvestOrder, 0)
	for _, v := range list {
		item := response.InvestOrder{
			Type:         v.Type,
			Amount:       float64(v.Amount) / model.UNITY,
			CreateTime:   v.CreateTime,
			UnfreezeTime: v.UnfreezeTime,
			IncomeTime:   v.IncomeTime,
			Balance:      float64(v.Balance) / model.UNITY,
		}
		res = append(res, item)
	}
	return response.InvestOrderData{List: res, Page: FormatPage(page)}
}
