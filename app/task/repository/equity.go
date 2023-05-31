package repository

import (
	"china-russia/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type Equity struct {
}

func (this *Equity) Do() {

	now := time.Now().Unix()

	g := model.Equity{}
	if !g.Get(true) {
		logrus.Errorf("未开启股权")
		return
	}

	//获取基础配置表信息
	config := model.SetBase{}
	config.Get()

	//收盘状态分析
	//isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)

	if now == g.OpenTime {
		//发行
		g := model.OrderEquity{}
		orders := g.List("", nil)
		if len(orders) > 0 {
			for _, v := range orders {
				//未中签的钱 + 返回的钱
				//weiMoney := (v.PayMoney * int64(int(model.UNITY)-v.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(v.Equity.ReturnRate)) / int64(model.UNITY)
				weiMoney := v.PayMoney.Mul(decimal.NewFromInt(1).Sub(v.Rate)).Div(decimal.NewFromInt(100)).Mul(v.Equity.MissRate).Div(decimal.NewFromInt(100)).Round(2)
				logrus.Infof("发行返回的钱  用户Id%v  收益%v", v.UId, weiMoney)

				//获取用户余额
				m := model.Member{Id: v.UId}
				m.Get()

				//if isRetreatStatus == true {
				//	//可用余额转换比例分析, 默认为90%
				//	if config.IncomeBalanceRate.LessThanOrEqual(decimal.Zero) {
				//		config.IncomeBalanceRate = decimal.NewFromFloat(0.9)
				//	}
				//	//可用余额,可提现余额分析
				//	balanceAmount := config.IncomeBalanceRate.Mul(weiMoney).Div(decimal.NewFromInt(100)).Round(2)
				//	useBalanceAmount := weiMoney.Sub(balanceAmount)
				//
				//	//加入账变记录
				//	trade := model.Trade{
				//		UId:        v.UId,
				//		TradeType:  17,
				//		ItemId:     v.Id,
				//		Amount:     balanceAmount,
				//		Before:     v.Member.Balance,
				//		After:      v.Member.Balance.Add(balanceAmount),
				//		Desc:       "股权发行收益",
				//		CreateTime: time.Now().Unix(),
				//		UpdateTime: time.Now().Unix(),
				//		IsFrontend: 1,
				//	}
				//	err := trade.Insert()
				//	if err != nil {
				//		logrus.Errorf("发行返回的钱  用户Id%v  收益%v  err=%v", v.UId, weiMoney, err)
				//	}
				//
				//	trade2 := model.Trade{
				//		UId:        v.UId,
				//		TradeType:  17,
				//		ItemId:     v.Id,
				//		Amount:     useBalanceAmount,
				//		Before:     v.Member.WithdrawBalance,
				//		After:      v.Member.WithdrawBalance.Add(balanceAmount),
				//		Desc:       "股权发行收益",
				//		CreateTime: time.Now().Unix(),
				//		UpdateTime: time.Now().Unix(),
				//		IsFrontend: 1,
				//	}
				//	err = trade2.Insert()
				//	if err != nil {
				//		logrus.Errorf("发行返回的钱  用户Id%v  收益%v  err=%v", v.UId, weiMoney, err)
				//	}
				//
				//	//用户加钱
				//	m.Balance = m.Balance.Add(balanceAmount)
				//	m.WithdrawBalance = m.WithdrawBalance.Add(useBalanceAmount)
				//} else {
				//
				//}
				//加入账变记录
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  17,
					ItemId:     v.Id,
					Amount:     weiMoney,
					Before:     v.Member.WithdrawBalance,
					After:      v.Member.WithdrawBalance.Add(weiMoney),
					Desc:       "股权发行收益",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err := trade.Insert()
				if err != nil {
					logrus.Errorf("发行返回的钱  用户Id%v  收益%v  err=%v", v.UId, weiMoney, err)
				}

				//用户加钱
				m.WithdrawBalance = m.WithdrawBalance.Add(weiMoney)
				//更改用户余额
				m.TotalIncome = m.TotalIncome.Add(weiMoney)
				err = m.Update("balance", "withdraw_balance", "total_income")
				if err != nil {
					logrus.Errorf("发行 修改余额失败  用户Id %v 收益 %v  err= &v", v.UId, weiMoney, err)
				}
			}
		}
	}

	//if now == g.RecoverTime {
	//	//回购
	//	g := model.OrderEquity{}
	//	orders := g.List("", nil)
	//	if len(orders) > 0 {
	//		for i := range orders {
	//			//回购 + 返回的钱
	//			huiMoney := (v.PayMoney * int64(v.Rate) / int64(model.UNITY)) * int64(v.Equity.ReturnLuckyRate) / int64(model.UNITY)
	//
	//			logrus.Infof("发行返回的钱  用户Id%v  收益%v", v.UId, huiMoney)
	//
	//			//获取用户余额
	//			m := model.Member{Id: v.UId}
	//			m.Get()
	//
	//			if isRetreatStatus == true {
	//				//可用余额转换比例分析, 默认为90%
	//				if config.IncomeBalanceRate == 0 {
	//					config.IncomeBalanceRate = 9000
	//				}
	//
	//				//可用余额,可提现余额分析
	//				balanceAmount := int64(config.IncomeBalanceRate) / int64(model.UNITY) * huiMoney
	//				useBalanceAmount := huiMoney - balanceAmount
	//
	//				//加入账变记录
	//				trade := model.Trade{
	//					UId:        v.UId,
	//					TradeType:  17,
	//					ItemId:     v.Id,
	//					Amount:     balanceAmount,
	//					Before:     m.Balance,
	//					After:      m.Balance + balanceAmount,
	//					Desc:       "股权回购收益",
	//					CreateTime: time.Now().Unix(),
	//					UpdateTime: time.Now().Unix(),
	//					IsFrontend: 1,
	//				}
	//				err := trade.Insert()
	//				if err != nil {
	//					logrus.Errorf("回购返回的钱  用户Id%v  收益%v  err=%v", v.UId, balanceAmount, err)
	//				}
	//
	//				trade2 := model.Trade{
	//					UId:        v.UId,
	//					TradeType:  17,
	//					ItemId:     v.Id,
	//					Amount:     useBalanceAmount,
	//					Before:     m.WithdrawBalance,
	//					After:      m.WithdrawBalance + useBalanceAmount,
	//					Desc:       "股权回购收益",
	//					CreateTime: time.Now().Unix(),
	//					UpdateTime: time.Now().Unix(),
	//					IsFrontend: 1,
	//				}
	//				err = trade2.Insert()
	//				if err != nil {
	//					logrus.Errorf("回购返回的钱  用户Id%v  收益%v  err=%v", v.UId, useBalanceAmount, err)
	//				}
	//
	//				//用户加钱
	//				m.Balance += balanceAmount
	//				m.WithdrawBalance += useBalanceAmount
	//			} else {
	//				//加入账变记录
	//				trade := model.Trade{
	//					UId:        v.UId,
	//					TradeType:  17,
	//					ItemId:     v.Id,
	//					Amount:     huiMoney,
	//					Before:     m.WithdrawBalance,
	//					After:      m.WithdrawBalance + huiMoney,
	//					Desc:       "股权回购收益",
	//					CreateTime: time.Now().Unix(),
	//					UpdateTime: time.Now().Unix(),
	//					IsFrontend: 1,
	//				}
	//				err := trade.Insert()
	//				if err != nil {
	//					logrus.Errorf("回购返回的钱  用户Id%v  收益%v  err=%v", v.UId, huiMoney, err)
	//				}
	//
	//				//用户加钱
	//				m.Balance += 0
	//				m.WithdrawBalance += huiMoney
	//			}
	//
	//			m.TotalBalance += huiMoney
	//			m.Income += huiMoney
	//			err := m.Update("total_balance", "balance", "withdraw_balance", "income")
	//			if err != nil {
	//				logrus.Errorf("回购 修改余额失败  用户Id %v 收益 %v  err= &v", v.UId, huiMoney, err)
	//			}
	//		}
	//	}
	//}

}
