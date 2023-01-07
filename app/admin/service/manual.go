package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Manual struct {
	request.ManualRequest
}

func (this Manual) Recharge(admin model.Admin, t int, isfront int) error {
	s := strings.Split(this.UIDs, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		if this.Amount == 0 {
			return errors.New("金额错误")
		}
		member := model.Member{ID: id}
		if !member.Get() {
			return errors.New("用户不存在")
		}

		h := RechargeHandle{}
		err := h.Recharge(member, 0, int64(this.Amount*model.UNITY), 2, 14, isfront)
		if err != nil {
			return err
		}
	}

	return nil

}
func (this Manual) Withdraw(admin model.Admin, isfront int) error {
	s := strings.Split(this.UIDs, ",")
	for _, v := range s {
		id, _ := strconv.Atoi(v)
		if id == 0 {
			return errors.New("用户错误")
		}
		if this.Amount == 0 {
			return errors.New("金额错误")
		}
		user := model.Member{ID: id}
		if !user.Get() {
			return errors.New("用户不存在")
		}
		if this.Handle == 2 && user.Balance < int64(this.Amount*model.UNITY) {
			return errors.New("用户可用余额不足")
		}
		if this.Handle == 3 && user.UseBalance < int64(this.Amount*model.UNITY) {
			return errors.New("用户可提余额不足")
		}

		//账单
		trade := model.Trade{
			UID:        user.ID,
			TradeType:  15,
			ItemID:     0,
			Amount:     int64(this.Amount * model.UNITY),
			IsFrontend: isfront,
		}
		if this.Handle == 2 {
			trade.Before = user.Balance
			trade.After = user.Balance - int64(this.Amount*model.UNITY)
			trade.Desc = "系统回调"
			user.Balance -= int64(this.Amount * model.UNITY)
		} else {
			trade.Before = user.UseBalance
			trade.After = user.UseBalance - int64(this.Amount*model.UNITY)
			trade.Desc = "人工扣款可提现余额"
			user.UseBalance -= int64(this.Amount * model.UNITY)
		}
		user.TotalBalance -= int64(this.Amount * model.UNITY)
		user.Update("balance", "use_balance", "total_balance")
		trade.Insert()
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
			ID:         v.ID,
			UserID:     v.UID,
			Username:   v.Username,
			Type:       v.Type,
			Amount:     float64(v.Amount / 100),
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
