package repository

import (
	"finance/common"
	"finance/model"
	"github.com/sirupsen/logrus"
	"time"
)

type Guquan struct {
}

func (this *Guquan) Do() {

	now := time.Now().Unix()

	g := model.Guquan{}
	if !g.Get(true) {
		logrus.Errorf("未开启股权")
		return
	}

	//获取基础配置表信息
	config := model.SetBase{}
	config.Get()

	//收盘状态分析
	isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)

	if now == g.OpenTime {
		//发行
		g := model.OrderGuquan{}
		orders := g.List("", nil)
		if len(orders) > 0 {
			for i := range orders {
				//未中签的钱 + 返回的钱
				weiMoney := (orders[i].PayMoney * int64(int(model.UNITY)-orders[i].Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(orders[i].Guquan.ReturnRate)) / int64(model.UNITY)

				logrus.Infof("发行返回的钱  用户ID%v  收益%v", orders[i].UID, weiMoney)

				//获取用户余额
				m := model.Member{ID: orders[i].UID}
				m.Get()

				if isRetreatStatus == true {
					//可用余额转换比例分析, 默认为90%
					if config.IncomeBalanceRate == 0 {
						config.IncomeBalanceRate = 9000
					}

					//可用余额,可提现余额分析
					balanceAmount := int64(config.IncomeBalanceRate) / int64(model.UNITY) * weiMoney
					useBalanceAmount := weiMoney - balanceAmount

					//加入账变记录
					trade := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     balanceAmount,
						Before:     orders[i].Member.Balance,
						After:      orders[i].Member.Balance + balanceAmount,
						Desc:       "股权发行收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err := trade.Insert()
					if err != nil {
						logrus.Errorf("发行返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, weiMoney, err)
					}

					trade2 := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     useBalanceAmount,
						Before:     orders[i].Member.UseBalance,
						After:      orders[i].Member.UseBalance + balanceAmount,
						Desc:       "股权发行收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err = trade2.Insert()
					if err != nil {
						logrus.Errorf("发行返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, weiMoney, err)
					}

					//用户加钱
					m.Balance += balanceAmount
					m.UseBalance += useBalanceAmount
				} else {
					//加入账变记录
					trade := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     weiMoney,
						Before:     orders[i].Member.UseBalance,
						After:      orders[i].Member.UseBalance + weiMoney,
						Desc:       "股权发行收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err := trade.Insert()
					if err != nil {
						logrus.Errorf("发行返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, weiMoney, err)
					}

					//用户加钱
					m.Balance += 0
					m.UseBalance += weiMoney
				}

				//更改用户余额
				m.TotalBalance += weiMoney
				m.Income += weiMoney
				err := m.Update("total_balance", "balance", "use_balance", "income")
				if err != nil {
					logrus.Errorf("发行 修改余额失败  用户ID %v 收益 %v  err= &v", orders[i].UID, weiMoney, err)
				}
			}
		}
	}

	if now == g.ReturnTime {
		//回购
		g := model.OrderGuquan{}
		orders := g.List("", nil)
		if len(orders) > 0 {
			for i := range orders {
				//回购 + 返回的钱
				huiMoney := (orders[i].PayMoney * int64(orders[i].Rate) / int64(model.UNITY)) * int64(orders[i].Guquan.ReturnLuckyRate) / int64(model.UNITY)

				logrus.Infof("发行返回的钱  用户ID%v  收益%v", orders[i].UID, huiMoney)

				//获取用户余额
				m := model.Member{ID: orders[i].UID}
				m.Get()

				if isRetreatStatus == true {
					//可用余额转换比例分析, 默认为90%
					if config.IncomeBalanceRate == 0 {
						config.IncomeBalanceRate = 9000
					}

					//可用余额,可提现余额分析
					balanceAmount := int64(config.IncomeBalanceRate) / int64(model.UNITY) * huiMoney
					useBalanceAmount := huiMoney - balanceAmount

					//加入账变记录
					trade := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     balanceAmount,
						Before:     m.Balance,
						After:      m.Balance + balanceAmount,
						Desc:       "股权回购收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err := trade.Insert()
					if err != nil {
						logrus.Errorf("回购返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, balanceAmount, err)
					}

					trade2 := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     useBalanceAmount,
						Before:     m.UseBalance,
						After:      m.UseBalance + useBalanceAmount,
						Desc:       "股权回购收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err = trade2.Insert()
					if err != nil {
						logrus.Errorf("回购返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, useBalanceAmount, err)
					}

					//用户加钱
					m.Balance += balanceAmount
					m.UseBalance += useBalanceAmount
				} else {
					//加入账变记录
					trade := model.Trade{
						UID:        orders[i].UID,
						TradeType:  17,
						ItemID:     orders[i].ID,
						Amount:     huiMoney,
						Before:     m.UseBalance,
						After:      m.UseBalance + huiMoney,
						Desc:       "股权回购收益",
						CreateTime: time.Now().Unix(),
						UpdateTime: time.Now().Unix(),
						IsFrontend: 1,
					}
					err := trade.Insert()
					if err != nil {
						logrus.Errorf("回购返回的钱  用户ID%v  收益%v  err=%v", orders[i].UID, huiMoney, err)
					}

					//用户加钱
					m.Balance += 0
					m.UseBalance += huiMoney
				}

				m.TotalBalance += huiMoney
				m.Income += huiMoney
				err := m.Update("total_balance", "balance", "use_balance", "income")
				if err != nil {
					logrus.Errorf("回购 修改余额失败  用户ID %v 收益 %v  err= &v", orders[i].UID, huiMoney, err)
				}
			}
		}
	}

}
