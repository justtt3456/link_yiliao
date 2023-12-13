package repository

import (
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type EquityScore struct {
}

func (this *EquityScore) Do() {
	//今日是否已结算
	today := common.GetTodayZero()
	s := global.REDIS.Get("medicine:" + fmt.Sprint(today))
	if s.Val() != "" {
		logrus.Errorf("今日已经结算%v", today)
		return
	}
	now := time.Now().Unix()
	defer global.REDIS.Set("medicine:"+fmt.Sprint(today), now, -1)
	config := model.SetBase{}
	if !config.Get() {
		logrus.Errorf("未开启股权分")
		return
	}
	if now >= config.EquityStartDate {
		//d := time.Now().Weekday()
		//if d == time.Saturday || d == time.Sunday {
		//	fmt.Println("周末不发放")
		//	return
		//}
		order := model.MedicineOrder{}
		orders := order.List("status = ?", []interface{}{model.StatusOk})
		log.Println("订单数量：", len(orders))
		for _, v := range orders {
			v.Current += 1
			if v.Current >= v.Interval {
				v.Status = model.StatusClose
			}
			err := v.Update("status", "current")
			if err != nil {
				logrus.Errorf("修改状态失败,err=%v", err)
			}
		}
	}

}
func (this *EquityScore) EquityScore() {
	now := time.Now().Unix()
	config := model.SetBase{}
	if !config.Get() {
		logrus.Errorf("未开启股权分")
		return
	}
	if now >= config.EquityStartDate {
		order := model.EquityScoreOrder{}
		orders := order.List(order.TableName()+".status = ? and create_time < ?", []interface{}{model.StatusOk, common.GetTodayZero()})
		log.Println("订单数量：", len(orders))
		for _, v := range orders {
			m := model.Member{Id: v.UId}
			if !m.Get() {
				continue
			}
			income := v.PayMoney.Mul(v.Rate).Div(decimal.NewFromInt(100)).Round(2)
			//加入账变记录
			trade := model.Trade{
				UId:        v.UId,
				TradeType:  17,
				ItemId:     v.Id,
				Amount:     income,
				Before:     v.Member.WithdrawBalance,
				After:      v.Member.WithdrawBalance.Add(income),
				Desc:       "股权分每日收益",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err := trade.Insert()
			if err != nil {
				logrus.Errorf("股权分每日收益：  用户Id%v  收益%v  err=%v", v.UId, income, err)
			}
			m.WithdrawBalance = m.WithdrawBalance.Add(income)
			m.PreIncome = m.PreIncome.Sub(income)
			//返回本金
			if now >= v.EndTime {
				//加入账变记录
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  18,
					ItemId:     v.Id,
					Amount:     income,
					Before:     v.Member.WithdrawBalance,
					After:      v.Member.WithdrawBalance.Add(v.PayMoney),
					Desc:       "股权分返回本金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err := trade.Insert()
				if err != nil {
					logrus.Errorf("股权分返回本金：  用户Id%v  收益%v  err=%v", v.UId, v.PayMoney, err)
				}
				m.WithdrawBalance = m.WithdrawBalance.Add(v.PayMoney)
				m.PreCapital = m.PreCapital.Sub(v.PayMoney)
				m.EquityScore -= int(v.PayMoney.IntPart())
				v.Status = model.StatusClose
				err = v.Update("status")
				if err != nil {
					logrus.Errorf("修改状态失败,err=%v", err)
				}
			}
			m.Update("withdraw_balance", "pre_income", "pre_capital", "equity_score")
		}
	}

}
