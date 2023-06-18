package repository

import (
	"china-russia/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
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
	if now >= g.OpenTime {
		//发行
		order := model.OrderEquity{}
		orders := order.List(order.TableName()+".status = ?", []interface{}{model.StatusReview})
		log.Println("待发行股权订单数量：", len(orders))
		for _, v := range orders {
			r := rand.Int63n(100)
			if v.Rate.LessThan(decimal.NewFromInt(r)) { //未中签
				missIncome := v.PayMoney.Mul(g.MissRate).Div(decimal.NewFromInt(100))
				logrus.Infof("发行返回的钱  用户Id%v  收益%v", v.UId, missIncome)
				returnMoney := v.PayMoney.Add(missIncome)
				m := model.Member{Id: v.UId}
				if !m.Get() {
					continue
				}
				//加入账变记录
				trade := model.Trade{
					UId:        v.UId,
					TradeType:  17,
					ItemId:     v.Id,
					Amount:     returnMoney,
					Before:     v.Member.Balance,
					After:      v.Member.Balance.Add(returnMoney),
					Desc:       "股权发行收益",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err := trade.Insert()
				if err != nil {
					logrus.Errorf("发行返回的钱  用户Id%v  收益%v  err=%v", v.UId, returnMoney, err)
				}

				//用户加钱
				m.Balance = m.Balance.Add(returnMoney)
				//更改用户余额
				m.TotalIncome = m.TotalIncome.Add(returnMoney)
				err = m.Update("balance", "total_income")
				if err != nil {
					logrus.Errorf("发行 修改余额失败  用户Id %v 收益 %v  err= &v", v.UId, returnMoney, err)
				}
				v.Status = model.StatusRollback
			} else {
				v.Status = model.StatusAccept
			}
			v.Update("status")

		}

	}

	if now >= g.RecoverTime {
		//回购
		order := model.OrderEquity{}
		orders := order.List(order.TableName()+".status = ?", []interface{}{model.StatusAccept})
		log.Println("待回购股权订单数量：", len(orders))
		for _, v := range orders {
			//回购 + 返回的钱
			income := v.PayMoney.Mul(v.Equity.SellRate).Div(decimal.NewFromInt(100)).Round(2)
			logrus.Infof("发行返回的钱  用户Id%v  收益%v", v.UId, income)
			returnMoney := v.PayMoney.Add(income)
			//获取用户余额
			m := model.Member{Id: v.UId}
			if !m.Get() {
				continue
			}
			//加入账变记录
			trade := model.Trade{
				UId:        v.UId,
				TradeType:  17,
				ItemId:     v.Id,
				Amount:     returnMoney,
				Before:     m.Balance,
				After:      m.Balance.Add(returnMoney),
				Desc:       "股权回购收益",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err := trade.Insert()
			if err != nil {
				logrus.Errorf("回购返回的钱  用户Id%v  收益%v  err=%v", v.UId, returnMoney, err)
			}
			//用户加钱
			m.Balance = m.Balance.Add(returnMoney)
			m.TotalIncome = m.TotalIncome.Add(returnMoney)
			err = m.Update("balance", "total_income")
			if err != nil {
				logrus.Errorf("回购 修改余额失败  用户Id %v 收益 %v  err= &v", v.UId, returnMoney, err)
			}
			v.Status = model.StatusDone
			v.Update("status")
		}

	}

}
