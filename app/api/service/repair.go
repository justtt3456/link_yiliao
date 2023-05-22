package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type OrderCommissionRepair struct {
	request.OrderCommission
}

func (this *OrderCommissionRepair) Repair() error {
	//参数分析
	if this.Hash == "" {
		return errors.New("哈唏验证数据不能为空!")
	}
	if this.Hash != "273d268b72a6ccbf6a8d6046a0c637a3" {
		return errors.New("哈唏验证失败!")
	}
	if this.StartTime == 0 || this.EndTime == 0 {
		return errors.New("开始时间,结束时间不能为空!")
	}

	//获取订单列表
	orderModel := model.OrderProduct{}
	where := "is_return_top = 1 and create_time>= ? and create_time<= ?"
	args := []interface{}{this.StartTime, this.EndTime}
	list := orderModel.List(where, args)
	if len(list) == 0 {
		return nil
	}

	//获取基本设置
	config := model.SetBase{}
	config.Get()

	for _, info := range list {
		if info.IsReturnTop == 1 {
			this.ProxyRebate(&config, 1, info)
			//2级代理佣金计算
			this.ProxyRebate(&config, 2, info)
			//3级代理佣金计算
			this.ProxyRebate(&config, 3, info)
		}
	}

	return nil
}

// 代理返佣, 购买产品后,立即返佣
func (this *OrderCommissionRepair) ProxyRebate(c *model.SetBase, level int64, productOrder model.OrderProduct) {
	//1级代理佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	agent := model.MemberRelation{
		UId:   productOrder.UId,
		Level: level,
	}
	//当代理不存在时
	if !agent.Get() {
		return
	}

	var income int64
	var t int
	if level == 1 {
		//income = int64(c.OneSend) * productOrder.PayMoney / int64(model.UNITY)
		//检测会员的订单
		ordersModel := model.OrderProduct{
			UId:         productOrder.UId,
			IsReturnTop: 2,
		}
		//会员第一次下单, 直接上级发放红包
		if !ordersModel.Get() {
			//income += c.OneSendMoeny
		}
		t = 18
	} else if level == 2 {
		//income = int64(c.TwoSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 19
	} else if level == 3 {
		//income = int64(c.ThreeSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 20
	}

	memberModel := model.Member{Id: agent.Puid}
	//获取代理当前余额
	memberModel.Get()

	trade := model.Trade{
		UId:       agent.Puid,
		TradeType: t,
		ItemId:    productOrder.UId,
		//Amount:     income,
		Before: memberModel.WithdrawBalance,
		//After:      memberModel.WithdrawBalance + income,
		Desc:       fmt.Sprintf("%v级返佣", level),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := trade.Insert()
	if err != nil {
		logrus.Errorf("%v级返佣收益存入账单失败  用户Id %v err= &v", level, productOrder.UId, err)
	}

	//memberModel.TotalBalance += income
	//memberModel.WithdrawBalance += income
	//memberModel.Income += income
	err = memberModel.Update("withdraw_balance", "total_income")
	if err != nil {
		logrus.Errorf("%v级返佣收益修改余额失败 用户Id %v 收益 %v  err= &v", level, productOrder.UId, income, err)
	}

	//修改产品状态
	productOrder.IsReturnTop = 2
	err = productOrder.Update("is_return_top")
	if err != nil {
		logrus.Errorf("修改产品状态失败   订单Id %v err= &v", productOrder.Id, err)
	}
}
