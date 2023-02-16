package repository

import (
	"finance/common"
	"finance/global"
	"finance/model"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type Award struct {
	Times int
}

const prefixdayincome = "prefixdayincome"

func (this *Award) Run() {
	now := time.Now().Unix()

	//加一层保险看看今天是否已经结算
	today := common.GetTodayZero()
	s := global.REDIS.Get(prefixdayincome + fmt.Sprint(today))
	if s.Val() != "" {
		logrus.Errorf("今日已经结算%v", today)
		return
	}
	global.REDIS.Set(prefixdayincome+fmt.Sprint(today), now, -1)

	//获取代理返佣配置
	c := model.SetBase{}
	if !c.Get() {
		logrus.Errorf("基础配置表  未配置")
		return
	}

	//产品收益
	o := model.OrderProduct{}
	productOrder := o.GetValidOrderList(today)
	if len(productOrder) == 0 {
		return
	}

	for i := range productOrder {
		//判断收益是否结束
		createDayTime := common.GetTimeByYMD(productOrder[i].CreateTime)
		overtime := createDayTime + int64(productOrder[i].Product.TimeLimit+1)*3600*24 + 2*3600
		starttime := createDayTime + 26*3600

		var capital int64
		var desc string
		//是否需要返回本金
		//if productOrder[i].Product.Type == 1 {
		//	//时间到了就反
		//	desc = "到期返回本金"
		//	if overtime == now {
		//		capital = productOrder[i].PayMoney
		//	}
		//} else {
		//	//延期返回本金
		//	desc = "延期返回本金"
		//	if overtime+int64(productOrder[i].Product.DelayTime*3600*24) == now {
		//		capital = productOrder[i].PayMoney
		//	}
		//}

		//返还本金
		switch productOrder[i].Product.Type {
		case 1: //到期返本
			desc = "到期返回本金"
			if overtime == now {
				capital = productOrder[i].PayMoney
			}
		case 2: //延期返本
			desc = "延期返回本金"
			if overtime+int64(productOrder[i].Product.DelayTime*3600*24) == now {
				capital = productOrder[i].PayMoney
			}
		case 3: //到期返本返息
			desc = "到期返本返息本金"
			if overtime == now {
				capital = productOrder[i].PayMoney
			}
		case 4: //每日返本返息
			desc = "每日返本返息本金"
			if overtime >= now {
				//计算每日返回本金金额
				if productOrder[i].Product.TimeLimit > 0 {
					totalAmount := decimal.NewFromInt(productOrder[i].PayMoney)
					totalDays := decimal.NewFromInt(int64(productOrder[i].Product.TimeLimit))
					capital = totalAmount.Div(totalDays).IntPart()
				}
			}
		}

		memberModel := model.Member{ID: productOrder[i].UID}
		orderModel := model.OrderProduct{ID: productOrder[i].ID}
		if capital > 0 {
			//获取当前余额
			memberModel.Get()
			//存入收益列表
			trade := model.Trade{
				UID:        productOrder[i].UID,
				TradeType:  16,
				ItemID:     productOrder[i].ID,
				Amount:     capital,
				Before:     memberModel.UseBalance,
				After:      memberModel.UseBalance + capital,
				Desc:       desc,
				CreateTime: now,
				UpdateTime: now,
				IsFrontend: 1,
			}
			_ = trade.Insert()

			//更改用户余额
			memberModel.UseBalance += capital
			memberModel.Income += capital
			memberModel.PIncome += capital
			err := memberModel.Update("use_balance", "income", "p_income")
			if err != nil {
				logrus.Errorf("修改余额失败  今日%v  用户ID %v 收益 %v err= &v", today, productOrder[i].UID, capital, err)
			}

			//更改订单状态:是否已返还投资本金
			orderModel.IsReturnCapital = 1
			err = orderModel.Update("is_return_capital")
			if err != nil {
				logrus.Errorf("修改产品订单返还本金状态失败  今日%v  订单ID %v err= &v", today, productOrder[i].ID, err)
			}
		}

		//当前还没有到开始收益时间
		if starttime >= now {
			continue
		}
		//当前已经过了收益的结束时间
		if overtime < now {
			continue
		}

		//计算收益
		income := float64(productOrder[i].PayMoney*int64(productOrder[i].Product.Dayincome)) / model.UNITY / model.UNITY
		logrus.Infof("今日已经结算%v  用户ID %v 收益 &v", today, productOrder[i].UID, income)
		income2 := int64(income * model.UNITY)

		//订单为返本返息结算类型时
		if productOrder[i].Product.Type == 3 {
			if overtime == now {
				income = income * float64(productOrder[i].Product.TimeLimit)
				income2 = income2 * int64(productOrder[i].Product.TimeLimit)
			} else {
				continue
			}
		}

		//获取当前余额
		memberModel.Get()
		//存入收益列表
		trade := model.Trade{
			UID:        productOrder[i].UID,
			TradeType:  16,
			ItemID:     productOrder[i].ID,
			Amount:     income2,
			Before:     memberModel.UseBalance,
			After:      memberModel.UseBalance + income2,
			Desc:       "产品每日收益",
			CreateTime: now,
			UpdateTime: now,
			IsFrontend: 1,
		}
		err := trade.Insert()
		if err != nil {
			logrus.Errorf("存入账单失败  今日%v  用户ID %v err= &v", today, productOrder[i].UID, err)
		}

		//更改用户余额
		memberModel.TotalBalance += income2
		memberModel.UseBalance += income2
		memberModel.Income += income2
		memberModel.PIncome += income2
		//待收益扣减
		if memberModel.WillIncome-income2 >= 0 {
			memberModel.WillIncome -= income2
		} else {
			memberModel.WillIncome = 0
		}
		err = memberModel.Update("total_balance", "use_balance", "income", "p_income", "wll_income")
		if err != nil {
			logrus.Errorf("修改余额失败  今日%v  用户ID %v 收益 %v err= &v", today, productOrder[i].UID, income, err)
		}
	}

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
				memberModel := model.Member{ID: proxyId}
				//获取当前余额
				memberModel.Get()
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
				memberModel.TotalBalance += income3
				memberModel.UseBalance += income3
				memberModel.Income += income3
				err := memberModel.Update("total_balance", "use_balance", "income")
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
