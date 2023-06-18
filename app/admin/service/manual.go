package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Manual struct {
	request.ManualRequest
}

//func (this Manual) Recharge() error {
//	s := strings.Split(this.UIds, ",")
//	for _, v := range s {
//		id, _ := strconv.Atoi(v)
//		if id == 0 {
//			return errors.New("用户错误")
//		}
//		member := model.Member{Id: id}
//		if !member.Get() {
//			return errors.New("用户不存在")
//		}
//
//		h := RechargeHandle{}
//		err := h.Recharge(member, 0, this.Amount, 2, 14, isfront)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

// 余额操作
func (this Manual) Balance() error {
	s := strings.Split(this.UIds, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		member := model.Member{Id: id}
		if !member.Get() {
			return errors.New("用户不存在")
		}
		var tradeType int
		var balance decimal.Decimal
		var desc string
		switch this.Handle {
		case 1:
			tradeType = 14
			balance = member.Balance
			desc = "他人向您转账"
			member.Balance = member.Balance.Add(this.Amount)
		case 2:
			tradeType = 15
			balance = member.Balance
			desc = "可用余额减少"
			member.Balance = member.Balance.Sub(this.Amount)
		case 3:
			tradeType = 16
			balance = member.WithdrawBalance
			desc = "充值"
			member.WithdrawBalance = member.WithdrawBalance.Add(this.Amount)
		case 4:
			tradeType = 17
			balance = member.WithdrawBalance
			desc = "提现"
			member.WithdrawBalance = member.WithdrawBalance.Sub(this.Amount)
		}
		trade := model.Trade{
			UId:        member.Id,
			TradeType:  tradeType,
			Amount:     this.Amount,
			Before:     balance,
			After:      member.Balance,
			IsFrontend: this.IsFrontend,
			Desc:       desc,
		}
		err := trade.Insert()
		if err != nil {
			logrus.Error(err)
			return err
		}
		if this.IsRecharge == 1 {
			member.TotalRecharge = member.TotalRecharge.Add(this.Amount)
		}
		err = member.Update("balance", "withdraw_balance", "total_recharge")
		if err != nil {
			return err
		}

	}
	return nil
}

// 股权操作
func (this Manual) Equity() error {
	s := strings.Split(this.UIds, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		member := model.Member{Id: id}
		if !member.Get() {
			return errors.New("用户不存在")
		}
		var tradeType int
		var equity int
		var desc string
		switch this.Handle {
		case 5:
			tradeType = 17
			equity = member.EquityScore
			desc = "增加股权"
			member.EquityScore = int(decimal.NewFromInt(int64(member.EquityScore)).Add(this.Amount).IntPart())
		case 6:
			tradeType = 18
			equity = member.EquityScore * -1
			desc = "减少股权"
			member.EquityScore = int(decimal.NewFromInt(int64(member.EquityScore)).Sub(this.Amount).IntPart())
		}
		trade := model.Trade{
			UId:        member.Id,
			TradeType:  tradeType,
			Amount:     this.Amount,
			Before:     decimal.NewFromInt(int64(equity)),
			After:      decimal.NewFromInt(int64(equity)).Add(this.Amount),
			IsFrontend: this.IsFrontend,
			Desc:       desc,
		}
		err := trade.Insert()
		if err != nil {
			logrus.Error(err)
			return err
		}
		err = member.Update("equity_score")
		if err != nil {
			return err
		}

	}
	return nil
}

//func (this Manual) Withdraw(admin model.Admin, isfront int) error {
//	s := strings.Split(this.UIds, ",")
//	for _, v := range s {
//		id, _ := strconv.Atoi(v)
//		if id == 0 {
//			return errors.New("用户错误")
//		}
//		user := model.Member{Id: id}
//		if !user.Get() {
//			return errors.New("用户不存在")
//		}
//		//if this.Handle == 2 && user.Balance < int64(this.Amount decimal.Decimal) {
//		//	return errors.New("用户可用余额不足")
//		//}
//		//if this.Handle == 3 && user.WithdrawBalance < int64(this.Amount decimal.Decimal) {
//		//	return errors.New("用户可提余额不足")
//		//}
//
//		//账单
//		trade := model.Trade{
//			UId:       user.Id,
//			TradeType: 15,
//			ItemId:    0,
//			//Amount:    int64(this.Amount),
//			IsFrontend: isfront,
//		}
//		if this.Handle == 2 {
//			trade.Before = user.Balance
//			//trade.After = user.Balance - int64(this.Amount decimal.Decimal)
//			trade.Desc = "系统回调"
//			//user.Balance -= int64(this.Amount)
//		} else {
//			trade.Before = user.WithdrawBalance
//			//trade.After = user.WithdrawBalance - int64(this.Amount)
//			trade.Desc = "自动回调可提现余额"
//			//user.WithdrawBalance -= int64(this.Amount)
//		}
//		//user.TotalBalance -= int64(this.Amount)
//		user.Update("balance", "withdraw_balance")
//		trade.Insert()
//	}
//
//	return nil
//}

func (this Manual) TopupUseBalance(admin model.Admin, t int, isfront int) error {
	s := strings.Split(this.UIds, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		//if this.Amount == 0 {
		//	return errors.New("金额错误")
		//}
		member := model.Member{Id: id}
		if !member.Get() {
			return errors.New("用户不存在")
		}

		h := RechargeHandle{}
		err := h.TopupUseBalance(member, 0, this.Amount, 14, isfront)
		if err != nil {
			return err
		}
	}
	return nil
}

type ManualList struct {
	request.ManualListRequest
}

func (this ManualList) PageList() response.ManualData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	where, args := this.getWhere()
	m := model.Manual{}
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	sli := make([]response.ManualInfo, 0)
	for _, v := range list {
		item := response.ManualInfo{
			Id:       v.Id,
			UserId:   v.UId,
			Username: v.Username,
			Type:     v.Type,
			//Amount:     float64(v.Amount / 100),
			AdminName:  v.Admin.Username,
			AgentName:  v.Agent.Account,
			CreateTime: v.CreateTime,
		}
		sli = append(sli, item)
	}
	return response.ManualData{
		List: sli,
		Page: FormatPage(page),
	}
}

func (this ManualList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Type != 0 {
		where[model.Manual{}.TableName()+".type"] = this.Type
	}
	if this.Username != "" {
		where[model.Manual{}.TableName()+".username"] = this.Username
	}
	if this.StartTime > 0 {
		where[model.Manual{}.TableName()+".create_time >="] = this.StartTime
	}
	if this.EndTime > 0 {
		where[model.Manual{}.TableName()+".create_time <"] = this.EndTime
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
