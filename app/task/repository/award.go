package repository

import (
	"finance/common"
	"finance/global"
	"finance/model"
	"fmt"
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
		member := model.Member{ID: v.UID}
		if !member.Get() {
			continue
		}
		//根据收益类型
		switch v.Product.Type {
		case 1: //到期返本

			income := int64(float64(v.PayMoney*int64(v.IncomeRate)) / model.UNITY)
			logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, float64(income)/model.UNITY)
			//存入收益列表
			trade := model.Trade{
				UID:        v.UID,
				TradeType:  16,
				ItemID:     v.ID,
				Amount:     income,
				Before:     member.UseBalance,
				After:      member.UseBalance + income,
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			err := trade.Insert()
			if err != nil {
				logrus.Errorf("存入账单失败  今日%v  用户ID %v err= %v", time.Now().Format("20060102"), v.UID, err)
			}
			//更改用户可提现余额
			member.UseBalance += income
			//是否到期返回本金
			if now >= v.EndTime {
				//存入收益列表
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     v.PayMoney,
					Before:     member.UseBalance,
					After:      member.UseBalance + v.PayMoney,
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				member.UseBalance += v.PayMoney
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 2: //延期返本
			if now < v.EndTime {
				//计算收益
				income := int64(float64(v.PayMoney*int64(v.IncomeRate)) / model.UNITY)
				logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, float64(income)/model.UNITY)
				//存入收益列表
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     income,
					Before:     member.UseBalance,
					After:      member.UseBalance + income,
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				err := trade.Insert()
				if err != nil {
					logrus.Errorf("存入账单失败  今日%v  用户ID %v err= %v", time.Now().Format("20060102"), v.UID, err)
				}
				//更改用户可提现余额
				member.UseBalance += income
			}
			if now >= v.EndTime+int64(v.Product.DelayTime)*86400 {
				//可提现
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     v.PayMoney,
					Before:     member.UseBalance,
					After:      member.UseBalance + v.PayMoney,
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				member.UseBalance += v.PayMoney
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 3: //到期返本返息
			if now >= v.EndTime {
				income := int64(float64(v.PayMoney*int64(v.Product.Dayincome*v.Product.TimeLimit)) / model.UNITY)
				logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, income)
				//利息
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     income,
					Before:     member.UseBalance,
					After:      member.UseBalance + income,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//更改用户余额
				member.UseBalance += income
				//返本
				trade2 := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     v.PayMoney,
					Before:     member.UseBalance,
					After:      member.UseBalance + v.PayMoney,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.UseBalance += v.PayMoney
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		}
		member.Update("use_balance")
	}
}

// 收盘业务 部分到可用余额 部分到可提现余额
func (this *Award) After(orders []*model.OrderProduct, config model.SetBase) {
	now := time.Now().Unix()
	for _, v := range orders {
		member := model.Member{ID: v.UID}
		if !member.Get() {
			continue
		}
		//根据收益类型
		switch v.Product.Type {
		case 1: //到期返本

			income := int64(float64(v.PayMoney*int64(v.IncomeRate)) / model.UNITY)
			logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, income)

			//根据收益比例计算可用可提现余额
			balanceAmount := int64(config.IncomeBalanceRate) * income / int64(model.UNITY)
			useBalanceAmount := income - balanceAmount
			//可用余额
			trade := model.Trade{
				UID:        v.UID,
				TradeType:  16,
				ItemID:     v.ID,
				Amount:     balanceAmount,
				Before:     member.Balance,
				After:      member.Balance + balanceAmount,
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			_ = trade.Insert()
			//可提现余额
			trade2 := model.Trade{
				UID:        v.UID,
				TradeType:  16,
				ItemID:     v.ID,
				Amount:     useBalanceAmount,
				Before:     member.UseBalance,
				After:      member.UseBalance + useBalanceAmount,
				Desc:       "产品每日收益",
				IsFrontend: 1,
			}
			_ = trade2.Insert()
			//更改用户余额
			member.Balance += balanceAmount
			member.UseBalance += useBalanceAmount
			//是否到期返回本金
			if now >= v.EndTime {
				//根据收益比例计算可用可提现余额
				balanceAmount := int64(config.IncomeBalanceRate) * v.PayMoney / int64(model.UNITY)
				useBalanceAmount := v.PayMoney - balanceAmount
				//可提现
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance + balanceAmount,
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现
				trade2 := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     useBalanceAmount,
					Before:     member.UseBalance,
					After:      member.UseBalance + useBalanceAmount,
					Desc:       "到期返回本金",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.Balance += balanceAmount
				member.UseBalance += useBalanceAmount
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 2: //延期返本
			if now < v.EndTime {
				income := int64(float64(v.PayMoney*int64(v.IncomeRate)) / model.UNITY)
				logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, income)
				//根据收益比例计算可用可提现余额
				balanceAmount := int64(config.IncomeBalanceRate) * income / int64(model.UNITY)
				useBalanceAmount := income - balanceAmount
				//可用余额
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance + balanceAmount,
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现余额
				trade2 := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     useBalanceAmount,
					Before:     member.UseBalance,
					After:      member.UseBalance + useBalanceAmount,
					Desc:       "产品每日收益",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				//更改用户余额
				member.Balance += balanceAmount
				member.UseBalance += useBalanceAmount
			}
			if now >= v.EndTime+int64(v.Product.DelayTime)*86400 {
				//根据收益比例计算可用可提现余额
				balanceAmount := int64(config.IncomeBalanceRate) * v.PayMoney / int64(model.UNITY)
				useBalanceAmount := v.PayMoney - balanceAmount
				//可提现
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance + balanceAmount,
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现
				trade2 := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     useBalanceAmount,
					Before:     member.UseBalance,
					After:      member.UseBalance + useBalanceAmount,
					Desc:       "延期返本",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				member.Balance += balanceAmount
				member.UseBalance += useBalanceAmount
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		case 3: //到期返本返息
			if now >= v.EndTime {
				income := int64(float64(v.PayMoney*int64(v.Product.Dayincome*v.Product.TimeLimit)) / model.UNITY)
				logrus.Infof("今日已经结算%v  用户ID %v 收益 %v", time.Now().Format("20060102"), v.UID, income)
				//根据收益比例计算可用可提现余额
				balanceAmount := int64(config.IncomeBalanceRate) * income / int64(model.UNITY)
				useBalanceAmount := income - balanceAmount
				//可用余额
				trade := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     balanceAmount,
					Before:     member.Balance,
					After:      member.Balance + balanceAmount,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade.Insert()
				//可提现余额
				trade2 := model.Trade{
					UID:        v.UID,
					TradeType:  16,
					ItemID:     v.ID,
					Amount:     useBalanceAmount,
					Before:     member.UseBalance,
					After:      member.UseBalance + useBalanceAmount,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade2.Insert()
				//更改用户余额
				member.Balance += balanceAmount
				member.UseBalance += useBalanceAmount
				//根据收益比例计算可用可提现余额
				srcBalanceAmount := int64(config.IncomeBalanceRate) * v.PayMoney / int64(model.UNITY)
				srcUseBalanceAmount := v.PayMoney - srcBalanceAmount
				//可提现
				trade3 := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     srcBalanceAmount,
					Before:     member.Balance,
					After:      member.Balance + srcBalanceAmount,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade3.Insert()
				//可提现
				trade4 := model.Trade{
					UID:        v.UID,
					TradeType:  24,
					ItemID:     v.ID,
					Amount:     srcUseBalanceAmount,
					Before:     member.UseBalance,
					After:      member.UseBalance + srcUseBalanceAmount,
					Desc:       "到期返本返息",
					IsFrontend: 1,
				}
				_ = trade4.Insert()
				member.Balance += srcBalanceAmount
				member.UseBalance += srcUseBalanceAmount
				v.IsReturnCapital = 1
				v.Update("is_return_capital")
			}
		}
		member.Update("balance", "use_balance")
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

		//获取团队代理ID列表
		orderModel := model.OrderProduct{}
		userIds := orderModel.GetOrderUserIds(teamStartTime, teamEndTime)
		if len(userIds) == 0 {
			return
		}

		teamModel := model.MemberRelation{}
		proxyIds := teamModel.GetTeamLeaderIds(userIds)
		if len(proxyIds) == 0 {
			return
		}

		for _, proxyId := range proxyIds {
			//获取下线会员总人数
			teams := model.MemberRelation{}
			where := "c_member_relation.puid = ? and c_member_relation.level > 0 and Member.is_buy = 1"
			args := []interface{}{proxyId}
			users, count := teams.GetByPuidAll(where, args)

			var totoalMoney int64
			var income3 int64
			var uids []int
			//当下线有效会员总人数少于最低要求时
			if count < 100 {
				continue
			}

			for i := range users {
				if users[i].Level <= 3 {
					uids = append(uids, users[i].Member.ID)
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
			totoalMoney = o.Sum(where1, args1, "pay_money")
			if totoalMoney == 0 {
				continue
			}

			if count >= 100 && count < 500 {
				income3 = totoalMoney * 75 / int64(model.UNITY)
			} else if count >= 500 && count < 1000 {
				income3 = totoalMoney * 99 / int64(model.UNITY)
			} else if count >= 1000 && count < 3000 {
				income3 = totoalMoney * 137 / int64(model.UNITY)
			} else if count >= 3000 && count < 5000 {
				income3 = totoalMoney * 169 / int64(model.UNITY)
			} else if count >= 5000 {
				income3 = totoalMoney * 202 / int64(model.UNITY)
			}

			if income3 > 0 {
				//获取基础配置表信息
				config := model.SetBase{}
				config.Get()

				memberModel := model.Member{ID: proxyId}
				//获取当前余额
				memberModel.Get()

				//收盘状态分析
				isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
				if isRetreatStatus == true {
					//可用余额转换比例分析, 默认为90%
					if config.IncomeBalanceRate == 0 {
						config.IncomeBalanceRate = 9000
					}

					//可用余额,可提现余额分析
					balanceAmount := int64(config.IncomeBalanceRate) * income3 / int64(model.UNITY)
					useBalanceAmount := income3 - balanceAmount

					//存入收益列表
					trade := model.Trade{
						UID:        proxyId,
						TradeType:  21,
						ItemID:     int(count),
						Amount:     balanceAmount,
						Before:     memberModel.Balance,
						After:      memberModel.Balance + balanceAmount,
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade.Insert()

					trade2 := model.Trade{
						UID:        proxyId,
						TradeType:  21,
						ItemID:     int(count),
						Amount:     useBalanceAmount,
						Before:     memberModel.UseBalance,
						After:      memberModel.UseBalance + useBalanceAmount,
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade2.Insert()

					//更改账户余额
					memberModel.Balance += balanceAmount
					memberModel.UseBalance += useBalanceAmount
				} else {
					//存入收益列表
					trade := model.Trade{
						UID:        proxyId,
						TradeType:  21,
						ItemID:     int(count),
						Amount:     income3,
						Before:     memberModel.UseBalance,
						After:      memberModel.UseBalance + income3,
						Desc:       "团队收益",
						CreateTime: now,
						UpdateTime: now,
						IsFrontend: 1,
					}
					_ = trade.Insert()

					//更改账户余额
					memberModel.Balance += 0
					memberModel.UseBalance += income3
				}
				//更改账户余额
				memberModel.TotalBalance += income3
				memberModel.Income += income3
				err := memberModel.Update("total_balance", "balance", "use_balance", "income")
				if err != nil {
					logrus.Errorf("修改余额失败  今日%v  用户ID %v 团队收益 %v err= &v", today, proxyId, income3, err)
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
