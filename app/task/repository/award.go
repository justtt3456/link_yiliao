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
	global.REDIS.Set(prefixdayincome+fmt.Sprint(today), "111", -1)
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
		overtime := common.GetTimeByYMD(productOrder[i].CreateTime) + int64(productOrder[i].Product.TimeLimit+1)*3600*24+2*3600
		starttime := common.GetTimeByYMD(productOrder[i].CreateTime) + 26*3600
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
		if capital > 0 {

			//存入收益列表
			trade := model.Trade{
				UID:        productOrder[i].UID,
				TradeType:  16,
				ItemID:     productOrder[i].ID,
				Amount:     capital,
				Before:     productOrder[i].Member.UseBalance + capital,
				After:      productOrder[i].Member.UseBalance + capital,
				Desc:       desc,
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			_ = trade.Insert()
		}
		if starttime > now {
			continue
		}
		if overtime < now {
			continue
		}
		//计算收益
		income := float64(productOrder[i].PayMoney*int64(productOrder[i].Product.Dayincome)) / model.UNITY / model.UNITY
		logrus.Infof("今日已经结算%v  用户ID %v 收益 &v", today, productOrder[i].UID, income)
		income2 := int64(income * model.UNITY)

		//存入收益列表
		trade := model.Trade{
			UID:        productOrder[i].UID,
			TradeType:  16,
			ItemID:     productOrder[i].ID,
			Amount:     income2,
			Before:     productOrder[i].Member.UseBalance,
			After:      productOrder[i].Member.UseBalance + income2,
			Desc:       "产品每日收益",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err := trade.Insert()

		if err != nil {
			logrus.Errorf("存入账单失败  今日%v  用户ID %v err= &v", today, productOrder[i].UID, err)
		}

		//是否应该返还上3级代理  佣金
		if productOrder[i].IsReturnTop == 1 {

			//1级代理佣金计算

			dealTop(c, 1, productOrder[i], today)

			//2级代理佣金计算

			dealTop(c, 2, productOrder[i], today)

			//3级代理佣金计算
			dealTop(c, 3, productOrder[i], today)
		}

		//查看团队代理人数  并且发放奖励
		//var team = map[int]int{
		//	100:  75,
		//	500:  99,
		//	1000: 137,
		//	3000: 169,
		//	5000: 202,
		//}
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
			//g := model.OrderGuquan{}
			//gmoney := o.Sum2(where1, args1, "pay_money")

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
			//存入收益列表
			trade := model.Trade{
				UID:        productOrder[i].UID,
				TradeType:  21,
				ItemID:     int(count),
				Amount:     income3,
				Before:     productOrder[i].Member.UseBalance + income2,
				After:      productOrder[i].Member.UseBalance + income2 + income3,
				Desc:       "团队收益",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			_ = trade.Insert()
		}

		//修改钱包余额
		m := model.Member{ID: productOrder[i].UID}
		m.Get()
		m.TotalBalance += income2 + income3 + capital
		m.UseBalance += income2 + income3 + capital
		m.Income += income2 + income3 + capital
		m.PIncome += income2 + capital
		err = m.Update("total_balance", "use_balance", "income", "p_income")
		if err != nil {
			logrus.Errorf("修改余额失败  今日%v  用户ID %v 收益 %v 团队收益 %v err= &v", today, productOrder[i].UID, income2, income3, err)
		}
	}
}

func dealTop(c model.SetBase, level int64, productOrder *model.OrderProduct, today int64) {
	//1级代理佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	agent := model.MemberRelation{
		UID:   productOrder.UID,
		Level: level,
	}
	if !agent.Get() {
		return
	}

	var income int64
	var t int
	if level == 1 {
		income = c.OneSendMoeny + int64(c.OneSend)*productOrder.PayMoney/int64(model.UNITY)
		t = 18
	} else if level == 2 {
		income = int64(c.TwoSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 19
	} else if level == 3 {
		income = int64(c.ThreeSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 20
	}
	logrus.Infof("今日已经结算%v  用户ID %v %v级返佣收益 &v", today, agent.Puid, level, income)
	trade := model.Trade{
		UID:        agent.Puid,
		TradeType:  t,
		ItemID:     productOrder.UID,
		Amount:     income,
		Before:     agent.Member.UseBalance,
		After:      agent.Member.UseBalance + income,
		Desc:       fmt.Sprintf("%v级返佣", level),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := trade.Insert()

	if err != nil {
		logrus.Errorf("%v级返佣收益存入账单失败  今日%v  用户ID %v err= &v", level, today, productOrder.UID, err)
	}

	//修改钱包余额
	m := model.Member{ID: agent.Puid}
	m.Get()
	m.TotalBalance += income
	m.UseBalance += income
	m.Income += income
	err = m.Update("total_balance", "use_balance", "income")
	if err != nil {
		logrus.Errorf("%v级返佣收益修改余额失败  今日%v  用户ID %v 收益 %v  err= &v", level, today, productOrder.UID, income, err)
	}
	//修改产品状态
	productOrder.IsReturnTop = 2
	err = productOrder.Update("is_return_top")
	if err != nil {
		logrus.Errorf("修改产品状态失败   订单ID %v err= &v", productOrder.ID, err)
	}
}
