package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
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
	args := []interface{}{member.Id}
	yesterday := income.YesterdayIncome(member.Id)
	m := response.InvestMember{
		Balance:         member.InvestAmount.Add(member.InvestFreeze),
		TotalIncome:     member.InvestIncome,
		YesterdayIncome: yesterday,
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
			Income:     v.Income,
			Balance:    v.Balance,
			CreateTime: v.CreateTime,
		}
		res = append(res, i)
	}
	//万元日收益金额
	profitAmount := invest.Ratio.Div(decimal.NewFromInt(100)).Round(2)
	return &response.InvestIndexData{
		Info: response.InvestInfo{
			Name:           invest.Name,
			Ratio:          profitAmount, //万元每日收益
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
	if this.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New(lang.Lang("Wrong amount"))
	}
	if member.FundsStatus != model.StatusOk {
		return errors.New(lang.Lang("The current account has been frozen!"))
	}
	now := time.Now().Unix()
	switch this.Type {
	case 1: //转入
		if member.Balance.LessThanOrEqual(this.Amount) {
			return errors.New(lang.Lang("Insufficient account balance"))
		}
		c := model.Invest{}
		if !c.Get() {
			return errors.New(lang.Lang("Parameter error"))
		}
		if c.Status != model.StatusOk {
			return errors.New(lang.Lang("Parameter error"))
		}
		if this.Amount.LessThan(c.MinAmount) {
			return errors.New(fmt.Sprintf("最小转入金额%v", c.MinAmount))
		}
		today := common.GetTodayZero()
		m := model.InvestOrder{
			UId:            member.Id,
			Type:           this.Type,
			Amount:         this.Amount,
			Rate:           c.Ratio,
			UnfreezeTime:   now + int64(c.FreezeDay*86400),
			IncomeTime:     today + int64(c.IncomeInterval*86400),
			Balance:        member.InvestAmount.Add(member.InvestFreeze).Add(this.Amount),
			UnfreezeStatus: model.StatusClose,
		}
		if err := m.Insert(); err != nil {
			return err
		}
		//账变
		trade := model.Trade{
			UId:       member.Id,
			TradeType: 11,
			ItemId:    m.Id,
			Amount:    this.Amount,
			Before:    member.Balance,
			After:     member.Balance.Sub(this.Amount),
			Desc:      "余额宝转入",
		}
		trade.Insert()
		//扣除余额
		member.Balance = member.Balance.Sub(this.Amount)
		member.InvestFreeze = member.InvestFreeze.Add(this.Amount)
		return member.Update("balance", "invest_freeze")
	case 2: //转出
		//获取可转出余额 余额-不可转出
		if member.InvestAmount.LessThan(this.Amount) {
			return errors.New(lang.Lang("Insufficient transferable balance"))
		}
		m := model.InvestOrder{
			UId:     member.Id,
			Type:    this.Type,
			Amount:  this.Amount,
			Balance: member.InvestAmount.Add(member.InvestFreeze).Sub(this.Amount),
		}
		if err := m.Insert(); err != nil {
			return err
		}
		//账变
		trade := model.Trade{
			UId:       member.Id,
			TradeType: 12,
			ItemId:    m.Id,
			Amount:    this.Amount,
			Before:    member.Balance,
			After:     member.Balance.Add(this.Amount),
			Desc:      "余额宝转出",
		}
		trade.Insert()
		//增加余额
		member.Balance = member.Balance.Add(this.Amount)
		member.InvestAmount = member.InvestAmount.Sub(this.Amount)
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
	args := []interface{}{member.Id}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.InvestOrder, 0)
	for _, v := range list {
		item := response.InvestOrder{
			Type:         v.Type,
			Amount:       v.Amount,
			CreateTime:   v.CreateTime,
			UnfreezeTime: v.UnfreezeTime,
			IncomeTime:   v.IncomeTime,
			Balance:      v.Balance,
		}
		res = append(res, item)
	}
	return response.InvestOrderData{List: res, Page: FormatPage(page)}
}
