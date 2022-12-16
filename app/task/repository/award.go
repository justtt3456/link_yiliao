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
	productOrder := o.GetAll(today)
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
		if productOrder[i].Product.Type == 1 {
			//时间到了就反
			desc = "到期返回本金"
			if overtime == now {
				capital = productOrder[i].PayMoney
			}
		} else {
			//延期返回本金
			desc = "延期返回本金"
			if overtime+int64(productOrder[i].Product.DelayTime*3600*24) == now {
				capital = productOrder[i].PayMoney
			}
		}

		memberModel := model.Member{ID: productOrder[i].UID}
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
			memberModel.TotalBalance += capital
			memberModel.UseBalance += capital
			memberModel.Income += capital
			memberModel.PIncome += capital
			err := memberModel.Update("total_balance", "use_balance", "income", "p_income")
			if err != nil {
				logrus.Errorf("修改余额失败  今日%v  用户ID %v 收益 %v err= &v", today, productOrder[i].UID, capital, err)
			}
		}

		//当前还没有到开始收益时间
		if starttime > now {
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
		err = memberModel.Update("total_balance", "use_balance", "income", "p_income")
		if err != nil {
			logrus.Errorf("修改余额失败  今日%v  用户ID %v 收益 %v err= &v", today, productOrder[i].UID, income2, err)
		}

		//团队注册人数分析(每月一日执行)
		if time.Unix(today, 0).Day() == 1 {
			teams := model.MemberRelation{}
			where := "puid = ? and Member.is_buy = 1"
			args := []interface{}{productOrder[i].UID}
			users, count := teams.GetByPuidAll(where, args)
			var totoalMoney int64
			var income3 int64
			var uids []int
			if count >= 100 {
				for i := range users {
					uids = append(uids, users[i].Member.ID)
				}
				//获取下级所有总价值
				o := model.OrderProduct{}
				where1 := "uid in (?) "
				args1 := []interface{}{uids}
				omoney := o.Sum(where1, args1, "pay_money")
				totoalMoney = omoney
			}

			if count >= 100 && count < 500 {
				income3 = totoalMoney * 750 / int64(model.UNITY)
			} else if count >= 500 && count < 1000 {
				income3 = totoalMoney * 990 / int64(model.UNITY)
			} else if count >= 1000 && count < 3000 {
				income3 = totoalMoney * 1370 / int64(model.UNITY)
			} else if count >= 3000 && count < 5000 {
				income3 = totoalMoney * 1690 / int64(model.UNITY)
			} else if count >= 5000 {
				income3 = totoalMoney * 2020 / int64(model.UNITY)
			}

			if income3 > 0 {
				//获取当前余额
				memberModel.Get()
				//存入收益列表
				trade := model.Trade{
					UID:        productOrder[i].UID,
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

				memberModel.TotalBalance += income3
				memberModel.UseBalance += income3
				memberModel.Income += income3
				err = memberModel.Update("total_balance", "use_balance", "income")
				if err != nil {
					logrus.Errorf("修改余额失败  今日%v  用户ID %v 团队收益 %v err= &v", today, productOrder[i].UID, income3, err)
				}
			}
		}
	}
}
