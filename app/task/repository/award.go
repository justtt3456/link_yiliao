package repository

import (
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type Award struct {
	Times int
}

const prefixdayincome = "prefixdayincome"

// 收盘前业务 全部收益到可提现
func (this *Award) Income(orders []*model.OrderProduct) {
	now := time.Now().Unix()
	for _, v := range orders {
		member := model.Member{Id: v.UId}
		if !member.Get() {
			continue
		}
		//根据收益类型
		switch v.Product.Type {
		case 1, 5: //到期返本
			income := v.PayMoney.Mul(v.IncomeRate)
			//存入收益列表
			trade := model.Trade{
				UId:        v.UId,
				TradeType:  16,
				ItemId:     v.Id,
				Amount:     income,
				Before:     member.WithdrawBalance,
				After:      member.WithdrawBalance.Add(income),
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			err := trade.Insert()
			if err != nil {
				logrus.Errorf("存入账单失败  今日%v  用户Id %v err= %v", time.Now().Format("20060102"), v.UId, err)
			}
			//更改用户可提现余额
			member.WithdrawBalance = member.WithdrawBalance.Add(income)
			//是否到期返回本金
			if now >= v.EndTime {
				//存入收益列表
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     v.PayMoney,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(v.PayMoney),
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				member.WithdrawBalance = member.WithdrawBalance.Add(v.PayMoney)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 2: //延期返本
			if now < v.EndTime {
				//计算收益
				income := v.PayMoney.Mul(v.IncomeRate)
				//存入收益列表
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     income,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(income),
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				err := trade.Insert()
				if err != nil {
					logrus.Errorf("存入账单失败  今日%v  用户Id %v err= %v", time.Now().Format("20060102"), v.UId, err)
				}
				//更改用户可提现余额
				member.WithdrawBalance = member.WithdrawBalance.Add(income)
			}
			if now >= v.EndTime+int64(v.Product.DelayTime)*86400 {
				//可提现
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     v.PayMoney,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(v.PayMoney),
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				member.WithdrawBalance = member.WithdrawBalance.Add(v.PayMoney)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 3: //到期返本返息
			if now >= v.EndTime {
				income := v.PayMoney.Mul(v.Product.IncomeRate).Mul(decimal.NewFromInt(int64(v.Product.Interval))).Round(2)
				logrus.Infof("今日已经结算%v  用户Id %v 收益 %v", time.Now().Format("20060102"), v.UId, income)
				//利息
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     income,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(income),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//更改用户余额
				member.WithdrawBalance = member.WithdrawBalance.Add(income)
				//返本
				trade2 := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     v.PayMoney,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(v.PayMoney),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.WithdrawBalance = member.WithdrawBalance.Add(v.PayMoney)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		}
		member.Update("withdraw_balance")
	}
}

func (this *Award) Run() {
	defer func() {
		// 团队收益结算
		go this.TeamIncome()
	}()
	now := time.Now().Unix()
	//今日是否已结算
	today := common.GetTodayZero()
	s := global.REDIS.Get(prefixdayincome + fmt.Sprint(today))
	if s.Val() != "" {
		logrus.Errorf("今日已经结算%v", today)
		return
	}
	//订单列表
	o := model.OrderProduct{}
	productOrder := o.GetValidOrderList(today)
	if len(productOrder) == 0 {
		return
	}
	//获取基础配置表信息
	config := model.SetBase{}
	config.Get()
	this.Income(productOrder)
	global.REDIS.Set(prefixdayincome+fmt.Sprint(today), now, -1)
}

// 团队收益结算
func (this *Award) TeamIncome() {
	now := time.Now().Unix()
	today := common.GetTodayZero()
	//执行时间:每月1日执行
	if time.Unix(today, 0).Day() == 1 {
		year, month, _ := time.Now().Date()
		thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		//上个月开始与结束时间
		teamStartTime := thisMonth.AddDate(0, -1, 0).Unix()
		teamEndTime := thisMonth.Unix() - 1
		//获取有下级的用户
		parents := make([]int, 0)
		global.DB.Model(model.MemberParents{}).Select("parent_id").Where("parent_id > 0").Group("parent_id").Scan(&parents)
		for _, v := range parents {
			member := model.Member{Id: v}
			if !member.Get() {
				continue
			}
			//获取团队代理Id列表
			orderModel := model.OrderProduct{}
			userIds := orderModel.GetOrderUserIds(teamStartTime, teamEndTime)
			count := len(userIds)
			var rate decimal.Decimal
			if count >= 6 {
				rate = decimal.NewFromFloat(1.8)
			} else if count >= 4 {
				rate = decimal.NewFromFloat(1.25)
			} else if count >= 2 {
				rate = decimal.NewFromFloat(1)
			} else if count >= 1 {
				rate = decimal.NewFromFloat(0.8)
			} else {
				continue
			}
			var amount float64
			tx := global.DB.Model(model.OrderProduct{}).Select("COALESCE(sum(pay_money),0)").Where("create_time between ? and ?", teamStartTime, teamEndTime).Scan(&amount)
			if tx.Error != nil {
				logrus.Error(tx.Error)
				continue
			}
			income := rate.Mul(decimal.NewFromFloat(amount)).Div(decimal.NewFromInt(100)).Round(2)
			//存入收益列表
			trade := model.Trade{
				UId:        v,
				TradeType:  21,
				ItemId:     0,
				Amount:     income,
				Before:     member.WithdrawBalance,
				After:      member.WithdrawBalance.Add(income),
				Desc:       "团队收益",
				CreateTime: now,
				UpdateTime: now,
				IsFrontend: 1,
			}
			_ = trade.Insert()
			//更改账户余额
			member.WithdrawBalance = member.WithdrawBalance.Add(income)
			//更改账户余额
			member.TotalIncome = member.TotalIncome.Add(income)
			err := member.Update("total_income", "withdraw_balance")
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}
