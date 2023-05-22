package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Manual struct {
	request.ManualRequest
}

func (this Manual) Recharge(admin model.Admin, t int, isfront int) error {
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
		err := h.Recharge(member, 0, this.Amount, 2, 14, isfront)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this Manual) Withdraw(admin model.Admin, isfront int) error {
	s := strings.Split(this.UIds, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		//if this.Amount == 0 {
		//	return errors.New("金额错误")
		//}
		user := model.Member{Id: id}
		if !user.Get() {
			return errors.New("用户不存在")
		}
		//if this.Handle == 2 && user.Balance < int64(this.Amount decimal.Decimal) {
		//	return errors.New("用户可用余额不足")
		//}
		//if this.Handle == 3 && user.WithdrawBalance < int64(this.Amount decimal.Decimal) {
		//	return errors.New("用户可提余额不足")
		//}

		//账单
		trade := model.Trade{
			UId:       user.Id,
			TradeType: 15,
			ItemId:    0,
			//Amount:    int64(this.Amount),
			IsFrontend: isfront,
		}
		if this.Handle == 2 {
			trade.Before = user.Balance
			//trade.After = user.Balance - int64(this.Amount decimal.Decimal)
			trade.Desc = "系统回调"
			//user.Balance -= int64(this.Amount)
		} else {
			trade.Before = user.WithdrawBalance
			//trade.After = user.WithdrawBalance - int64(this.Amount)
			trade.Desc = "自动回调可提现余额"
			//user.WithdrawBalance -= int64(this.Amount)
		}
		//user.TotalBalance -= int64(this.Amount)
		user.Update("balance", "withdraw_balance")
		trade.Insert()
	}

	return nil
}

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
			AgentName:  v.Agent.Name,
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
