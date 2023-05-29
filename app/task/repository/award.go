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
func (this *Award) Before(orders []*model.OrderProduct) {
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

// 收盘业务 部分到可用余额 部分到可提现余额
func (this *Award) After(orders []*model.OrderProduct, config model.SetBase) {
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
			logrus.Infof("今日已经结算%v  用户Id %v 收益 %v", time.Now().Format("20060102"), v.UId, income)
			//根据收益比例计算可用可提现余额
			balanceAmount := config.IncomeBalanceRate.Mul(income).Div(decimal.NewFromInt(100)).Round(2)
			useBalanceAmount := income.Sub(balanceAmount)
			//可用余额
			trade := model.Trade{
				UId:        v.UId,
				TradeType:  16,
				ItemId:     v.Id,
				Amount:     balanceAmount,
				Before:     member.Balance,
				After:      member.Balance.Add(balanceAmount),
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			_ = trade.Insert()
			//可提现余额
			trade2 := model.Trade{
				UId:        v.UId,
				TradeType:  16,
				ItemId:     v.Id,
				Amount:     useBalanceAmount,
				Before:     member.WithdrawBalance,
				After:      member.WithdrawBalance.Add(useBalanceAmount),
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			_ = trade2.Insert()
			//更改用户余额
			member.Balance = member.Balance.Add(balanceAmount)
			member.WithdrawBalance = member.WithdrawBalance.Add(useBalanceAmount)
			//是否到期返回本金
			if now >= v.EndTime {
				//根据收益比例计算可用可提现余额
				balanceAmount := config.IncomeBalanceRate.Mul(v.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
				useBalanceAmount := v.PayMoney.Sub(balanceAmount)
				//可提现
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance.Add(balanceAmount),
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现
				trade2 := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     useBalanceAmount,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(useBalanceAmount),
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.Balance = member.Balance.Add(balanceAmount)
				member.WithdrawBalance = member.WithdrawBalance.Add(useBalanceAmount)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 2: //延期返本
			if now < v.EndTime {
				income := v.PayMoney.Mul(v.IncomeRate).Div(decimal.NewFromInt(100)).Round(2)
				logrus.Infof("今日已经结算%v  用户Id %v 收益 %v", time.Now().Format("20060102"), v.UId, income)
				//根据收益比例计算可用可提现余额
				balanceAmount := config.IncomeBalanceRate.Mul(income).Div(decimal.NewFromInt(100)).Round(2)
				useBalanceAmount := income.Sub(balanceAmount)
				//可用余额
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance.Add(balanceAmount),
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现余额
				trade2 := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     useBalanceAmount,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(useBalanceAmount),
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				//更改用户余额
				member.Balance = member.Balance.Add(balanceAmount)
				member.WithdrawBalance = member.WithdrawBalance.Add(useBalanceAmount)
			}
			if now >= v.EndTime+int64(v.Product.DelayTime)*86400 {
				//根据收益比例计算可用可提现余额
				balanceAmount := config.IncomeBalanceRate.Mul(v.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
				useBalanceAmount := v.PayMoney.Sub(balanceAmount)
				//可提现
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance.Add(balanceAmount),
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现
				trade2 := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     useBalanceAmount,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(useBalanceAmount),
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.Balance = member.Balance.Add(balanceAmount)
				member.WithdrawBalance = member.WithdrawBalance.Add(useBalanceAmount)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 3: //到期返本返息
			if now >= v.EndTime {
				income := v.PayMoney.Mul(v.Product.IncomeRate).Mul(decimal.NewFromInt(int64(v.Product.Interval))).Div(decimal.NewFromInt(100)).Round(2)
				logrus.Infof("今日已经结算%v  用户Id %v 收益 %v", time.Now().Format("20060102"), v.UId, income)
				//根据收益比例计算可用可提现余额
				balanceAmount := config.IncomeBalanceRate.Mul(income).Div(decimal.NewFromInt(100)).Round(2)
				useBalanceAmount := income.Sub(balanceAmount)
				//可用余额
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance.Add(balanceAmount),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现余额
				trade2 := model.Trade{
					UId:        v.UId,
					TradeType:  16,
					ItemId:     v.Id,
					Amount:     useBalanceAmount,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(useBalanceAmount),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				//更改用户余额
				member.Balance = member.Balance.Add(balanceAmount)
				member.WithdrawBalance = member.WithdrawBalance.Add(useBalanceAmount)
				//根据收益比例计算可用可提现余额
				srcBalanceAmount := config.IncomeBalanceRate.Mul(v.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
				srcUseBalanceAmount := v.PayMoney.Sub(srcBalanceAmount)
				//可提现
				trade3 := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     srcBalanceAmount,
					Before:     member.Balance,
					After:      member.Balance.Add(srcBalanceAmount),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade3.Insert()
				//可提现
				trade4 := model.Trade{
					UId:        v.UId,
					TradeType:  24,
					ItemId:     v.Id,
					Amount:     srcUseBalanceAmount,
					Before:     member.WithdrawBalance,
					After:      member.WithdrawBalance.Add(srcUseBalanceAmount),
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade4.Insert()
				member.Balance = member.Balance.Add(srcBalanceAmount)
				member.WithdrawBalance = member.WithdrawBalance.Add(srcUseBalanceAmount)
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		}
		member.Update("balance", "withdraw_balance")
	}
}
func (this *Award) Run() {
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
	//收盘状态 如果到达收盘时间 则执行释放业务 否则每日收益到可提现余额
	isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
	if isRetreatStatus {
		this.After(productOrder, config)
	} else {
		this.Before(productOrder)
	}
	global.REDIS.Set(prefixdayincome+fmt.Sprint(today), now, -1)
	// 团队收益结算
	this.TeamIncome()
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

		//获取团队代理Id列表
		orderModel := model.OrderProduct{}
		userIds := orderModel.GetOrderUserIds(teamStartTime, teamEndTime)
		if len(userIds) == 0 {
			return
		}

		teamModel := model.MemberParents{}
		proxyIds := teamModel.GetTeamLeaderIds(userIds)
		if len(proxyIds) == 0 {
			return
		}

		for _, proxyId := range proxyIds {
			//获取下线会员总人数
			teams := model.MemberParents{}
			where := "c_member_parents.parent_id = ? and c_member_parents.level > 0 and Member.is_buy = 1"
			args := []interface{}{proxyId}
			users, count := teams.GetByPuidAll(where, args)

			var totalMoney decimal.Decimal
			var income decimal.Decimal
			var uids []int
			//当下线有效会员总人数少于最低要求时
			if count < 100 {
				continue
			}

			for i := range users {
				if users[i].Level <= 3 {
					uids = append(uids, users[i].Member.Id)
				}
			}
			//当没有3级内下线会员时则跳过
			if len(uids) == 0 {
				continue
			}

			//获取下级所有总价值
			o := model.OrderProduct{}
			where1 := "uid in (?) and create_time >= ? and create_time <= ? and is_return_team = 0"
			args1 := []interface{}{uids, teamStartTime, teamEndTime}
			totalMoney = o.Sum(where1, args1, "pay_money")
			if totalMoney.LessThanOrEqual(decimal.Zero) {
				continue
			}

			//if count >= 100 && count < 500 {
			//	income = totoalMoney * 75 / int64(model.UNITY)
			//} else if count >= 500 && count < 1000 {
			//	income = totoalMoney * 99 / int64(model.UNITY)
			//} else if count >= 1000 && count < 3000 {
			//	income = totoalMoney * 137 / int64(model.UNITY)
			//} else if count >= 3000 && count < 5000 {
			//	income = totoalMoney * 169 / int64(model.UNITY)
			//} else if count >= 5000 {
			//	income = totoalMoney * 202 / int64(model.UNITY)
			//}

			if decimal.Zero.LessThan(income) {
				//获取基础配置表信息
				config := model.SetBase{}
				config.Get()

				memberModel := model.Member{Id: proxyId}
				//获取当前余额
				memberModel.Get()

				//收盘状态分析
				isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
				if isRetreatStatus == true {
					//可用余额转换比例分析, 默认为90%
					if config.IncomeBalanceRate.LessThanOrEqual(decimal.Zero) {
						config.IncomeBalanceRate = decimal.NewFromFloat(0.9)
					}

					//可用余额,可提现余额分析
					balanceAmount := config.IncomeBalanceRate.Mul(income)
					useBalanceAmount := income.Sub(balanceAmount)

					//存入收益列表
					trade := model.Trade{
						UId:        proxyId,
						TradeType:  21,
						ItemId:     int(count),
						Amount:     balanceAmount,
						Before:     memberModel.Balance,
						After:      memberModel.Balance.Add(balanceAmount),
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade.Insert()

					trade2 := model.Trade{
						UId:        proxyId,
						TradeType:  21,
						ItemId:     int(count),
						Amount:     useBalanceAmount,
						Before:     memberModel.WithdrawBalance,
						After:      memberModel.WithdrawBalance.Add(useBalanceAmount),
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade2.Insert()

					//更改账户余额
					memberModel.Balance = memberModel.Balance.Add(balanceAmount)
					memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(useBalanceAmount)
				} else {
					//存入收益列表
					trade := model.Trade{
						UId:        proxyId,
						TradeType:  21,
						ItemId:     int(count),
						Amount:     income,
						Before:     memberModel.WithdrawBalance,
						After:      memberModel.WithdrawBalance.Add(income),
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade.Insert()
					//更改账户余额
					memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(income)
				}
				//更改账户余额
				memberModel.TotalIncome = memberModel.TotalIncome.Add(income)
				err := memberModel.Update("total_income", "balance", "withdraw_balance")
				if err != nil {
					logrus.Errorf("修改余额失败  今日%v  用户Id %v 团队收益 %v err= &v", today, proxyId, income, err)
				}
			}
		}

		//更改订单团队结算状态
		orderModel2 := model.OrderProduct{IsReturnTeam: 1}
		where2 := "create_time >= ? and create_time <= ? and is_return_team = 0"
		args2 := []interface{}{teamStartTime, teamEndTime}
		_ = orderModel2.UpdateTeamSettleStatus(where2, args2)
	}
}
