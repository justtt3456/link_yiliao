package repository

import (
	"finance/common"
	"finance/global"
	"finance/model"
	"strconv"
	"time"
)

type Invest struct {
}

func (this Invest) Do() {
	tickerKey := "invest_ticker"
	zero := common.GetTodayZero()
	investTicker := global.REDIS.Get(tickerKey).Val()
	parseInt, _ := strconv.ParseInt(investTicker, 10, 64)
	if parseInt >= zero {
		return
	}
	//设置已发放
	global.REDIS.Set(tickerKey, zero, -1)
	redis := model.Redis{}
	redis.Lock("invest_do")
	defer redis.Unlock("invest_do")
	invest := model.Invest{}
	if !invest.Get() {
		return
	}
	//发放收益
	//可用余额+冻结中可发放
	//查询用户余额宝余额
	member := model.Member{}
	memberWhere := "invest_amount > ? or invest_freeze > ?"
	memberArgs := []interface{}{0, 0}
	memberList := member.List(memberWhere, memberArgs)
	for _, iv := range memberList {
		this.sendIncome(iv)
	}

}

//发放余额宝奖励
func (this Invest) sendIncome(member model.Member) {
	in := model.InvestOrder{}
	inWhere := "uid = ? and type = ? and unfreeze_status = ? and income_time <= ?"
	inArgs := []interface{}{member.ID, 1, 2, time.Now().Unix()}
	amount := in.Sum(inWhere, inArgs, "amount")
	invest := model.Invest{}
	invest.Get()
	income := (member.InvestAmount + amount) * int64(invest.Ratio) / 365 / int64(model.UNITY)
	member.InvestAmount += income
	member.InvestIncome += income
	member.TotalBalance += income
	member.Income += income
	member.Update("invest_amount", "invest_income", "total_balance", "income")
	investBalance := member.InvestAmount + member.InvestFreeze
	r := model.InvestLog{
		UID:     member.ID,
		Income:  income,
		Balance: investBalance,
	}
	r.Insert()
	trade := model.Trade{
		UID:       member.ID,
		TradeType: 13,
		ItemID:    r.ID,
		Amount:    income,
		Before:    investBalance - income,
		After:     investBalance,
		Desc:      "余额宝收益",
	}
	trade.Insert()
}
