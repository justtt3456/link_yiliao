package logic

import (
	"china-russia/common"
	"china-russia/global"
	"china-russia/model"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type productBuyLogic struct {
	tx     *gorm.DB
	config *model.SetBase
}

func NewProductBuyLogic() *productBuyLogic {
	//基础配置表
	config := model.SetBase{}
	config.Get()
	return &productBuyLogic{
		tx:     global.DB.Begin(),
		config: &config,
	}
}
func (this productBuyLogic) ProductBuy(pt int, member *model.Member, product model.Product, amount decimal.Decimal, coupon model.Coupon, mc model.MemberCoupon) error {
	if this.tx == nil {
		this.tx = global.DB
	}
	if this.config == nil {
		config := model.SetBase{}
		config.Get()
		this.config = &config
	}
	var err error
	defer func() {
		if err != nil {
			this.tx.Rollback()
		} else {
			this.tx.Commit()
		}
	}()
	switch pt {
	case 1:
		//订单入库
		order, err := this.createOrder(member, product, amount)
		if err != nil {
			return err
		}
		//减去可投余额
		product.Current = product.Current.Add(amount)
		//验证是否已购满
		if product.Total.LessThanOrEqual(product.Current) || product.Total.Sub(product.Current).LessThan(product.Price) {
			product.IsFinished = model.StatusOk
		}
		err = this.tx.Select("current", "is_finished").Updates(product).Error
		if err != nil {
			logrus.Errorf("购买产品减去可投余额失败%v", err)
		}
		//账变记录
		err = this.createTrade(member, order, amount)
		if err != nil {
			return err
		}
		//扣减可用余额 记录已购状态
		member.Balance = member.Balance.Sub(amount.Sub(amount))
		member.IsBuy = 1
		//优惠券使用
		err = this.useCoupon(member, coupon, mc)
		if err != nil {
			return err
		}
		//首次购买赠送
		err = this.firstBuy(member, order.Id, amount)
		if err != nil {
			return err
		}
		//记录用户待收益利息及本金
		member.PreIncome = member.PreIncome.Add(product.IncomeRate.Mul(decimal.NewFromInt(int64(product.Interval))))
		member.PreCapital = member.PreCapital.Add(product.Price)
		err = this.tx.Select("balance", "withdraw_balance", "total_income").Updates(member).Error
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
		}
		//检查是否有满送活动
		//if product.IsCouponGift == model.StatusOk {
		//	full := model.FullDelivery{}
		//	if full.Find(amount) {
		//		//满送活动加入账变记录
		//		trade3 := model.Trade{
		//			UId:        member.Id,
		//			TradeType:  9,
		//			ItemId:     int(full.Coupon.Id),
		//			Amount:     full.Coupon.Price,
		//			Before:     member.Balance,
		//			After:      member.Balance,
		//			Desc:       fmt.Sprintf("赠送%v优惠券", full.Coupon.Price),
		//			CreateTime: time.Now().Unix(),
		//			UpdateTime: time.Now().Unix(),
		//			IsFrontend: 1,
		//		}
		//		err = trade3.Insert()
		//		if err != nil {
		//			logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		//		}
		//		couponIns := model.MemberCoupon{
		//			Uid:      int64(member.Id),
		//			CouponId: full.Coupon.Id,
		//			IsUse:    1,
		//		}
		//		err = couponIns.Insert()
		//		if err != nil {
		//			logrus.Errorf("赠赠送优惠券记录失败%v %v", err, member.Id)
		//		}
		//		//更改会员当前余额
		//		//memberModel.Balance += full.Coupon.Price
		//		//err = memberModel.Update("balance")
		//		//if err != nil {
		//		//	logrus.Errorf("更改会员余额信息失败%v", err)
		//		//}
		//	}
		//}
		//上级返佣
		if order.IsReturnTop == 1 {
			//1级代理佣金计算
			this.proxyRebate(1, order)
			//2级代理佣金计算
			this.proxyRebate(2, order)
			//3级代理佣金计算
			this.proxyRebate(3, order)
		}
	}
	return nil
}
func (this productBuyLogic) createOrder(member *model.Member, product model.Product, amount decimal.Decimal) (*model.OrderProduct, error) {
	//购买
	inc := &model.OrderProduct{
		UId:          member.Id,
		Pid:          product.Id,
		PayMoney:     amount.Sub(amount),
		IsReturnTop:  1,
		AfterBalance: member.Balance.Sub(amount.Sub(amount)),
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
		IncomeRate:   product.IncomeRate,
		EndTime:      time.Now().Unix() + int64(product.Interval*86400),
	}
	err := this.tx.Create(inc).Error
	if err != nil {
		return nil, err
	}
	return inc, nil
}
func (this productBuyLogic) createTrade(member *model.Member, order *model.OrderProduct, amount decimal.Decimal) error {
	//加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  1,
		ItemId:     order.Id,
		Amount:     amount,
		Before:     member.Balance,
		After:      member.Balance.Sub(amount),
		Desc:       "购买产品",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(trade).Error
	if err != nil {
		logrus.Errorf("购买产品加入账变记录失败%v", err)
		return err
	}
	return nil
}
func (this productBuyLogic) useCoupon(member *model.Member, coupon model.Coupon, mc model.MemberCoupon) error {
	//优惠券使用记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  10,
		ItemId:     coupon.Id,
		Amount:     coupon.Price,
		Before:     member.Balance,
		After:      member.Balance,
		Desc:       "使用优惠券",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(trade).Error
	if err != nil {
		logrus.Errorf("使用优惠券记录 加入账变记录失败%v", err)
		return err
	}
	//更改优惠券状态
	mc.IsUse = 2
	err = this.tx.Select("is_use").Updates(mc).Error
	if err != nil {
		logrus.Errorf("修改用户优惠券失败%v", err)
		return err
	}
	return nil
}
func (this productBuyLogic) firstBuy(member *model.Member, orderId int, amount decimal.Decimal) error {
	//赠送礼金 加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  7,
		ItemId:     orderId,
		Amount:     amount,
		Before:     member.WithdrawBalance,
		After:      member.WithdrawBalance.Add(amount),
		Desc:       "第一次购买赠送礼金",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(trade).Error
	if err != nil {
		logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		return err
	}
	//赠送礼金
	member.WithdrawBalance = member.WithdrawBalance.Add(amount)
	return this.tx.Select("withdraw_balance").Updates(member).Error
}
func (this *productBuyLogic) proxyRebate(level int, productOrder *model.OrderProduct) {
	//佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	parent := model.MemberParents{
		Uid:   productOrder.UId,
		Level: level,
	}
	if !parent.Get() {
		return
	}
	var income decimal.Decimal
	var t int
	if level == 1 {
		income = this.config.OneSend.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		//检测会员的订单
		ordersModel := model.OrderProduct{
			UId:         productOrder.UId,
			IsReturnTop: 2,
		}
		//会员第一次下单, 直接上级发放红包
		if !ordersModel.Get() {
			income = income.Add(this.config.OneSendMoney)
		}
		t = 18
	} else if level == 2 {
		income = this.config.TwoSend.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		t = 19
	} else if level == 3 {
		income = this.config.ThreeSend.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		t = 20
	}
	if parent.ParentId <= 0 {
		return
	}
	memberModel := model.Member{Id: parent.ParentId}
	//获取代理当前余额
	if !memberModel.Get() {
		return
	}
	//收盘状态分析
	isRetreatStatus := common.ParseRetreatStatus(this.config.RetreatStartDate)
	if isRetreatStatus == true && level != 1 {
		trade := model.Trade{
			UId:        parent.ParentId,
			TradeType:  t,
			ItemId:     productOrder.UId,
			Amount:     income,
			Before:     memberModel.Balance,
			After:      memberModel.Balance.Add(income),
			Desc:       fmt.Sprintf("%v级返佣", level),
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err := trade.Insert()
		if err != nil {
			logrus.Errorf("%v级返佣收益存入账单失败  用户Id %v err= &v", level, productOrder.UId, err)
		}
		memberModel.Balance = memberModel.Balance.Add(income)
		memberModel.TotalRebate = memberModel.TotalRebate.Add(income)
		err = memberModel.Update("balance", "total_income")
		if err != nil {
			logrus.Errorf("%v级返佣收益修改余额失败 用户Id %v 收益 %v  err= &v", level, productOrder.UId, income, err)
		}

		//修改产品状态
		productOrder.IsReturnTop = 2
		err = productOrder.Update("is_return_top")
		if err != nil {
			logrus.Errorf("修改产品状态失败   订单Id %v err= &v", productOrder.Id, err)
		}
	} else {
		trade := model.Trade{
			UId:        parent.ParentId,
			TradeType:  t,
			ItemId:     productOrder.UId,
			Amount:     income,
			Before:     memberModel.WithdrawBalance,
			After:      memberModel.WithdrawBalance.Add(income),
			Desc:       fmt.Sprintf("%v级返佣", level),
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err := trade.Insert()
		if err != nil {
			logrus.Errorf("%v级返佣收益存入账单失败  用户Id %v err= &v", level, productOrder.UId, err)
		}

		//memberModel.TotalBalance = memberModel.TotalBalance.Add(income)
		memberModel.Balance = memberModel.Balance.Add(income)
		memberModel.TotalRebate = memberModel.TotalRebate.Add(income)
		err = memberModel.Update("balance", "total_income")
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

	//释放一级,二级,三级代理的可用余额
	if isRetreatStatus == true && decimal.Zero.LessThan(memberModel.Balance) {
		var freeAmount decimal.Decimal
		var t2 int
		switch level {
		case 1:
			if this.config.OneReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.OneReleaseRate = 1000
			}
			t2 = 25
			freeAmount = this.config.OneReleaseRate.Mul(productOrder.PayMoney)
		case 2:
			if this.config.TwoReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.TwoReleaseRate = 500
			}
			t2 = 26
			freeAmount = this.config.TwoReleaseRate.Mul(productOrder.PayMoney)
		case 3:
			if this.config.ThreeReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.ThreeReleaseRate = 200
			}
			t2 = 27
			freeAmount = this.config.ThreeReleaseRate.Mul(productOrder.PayMoney)
		}
		//可用余额分析
		if memberModel.Balance.LessThan(freeAmount) {
			freeAmount = memberModel.Balance
		}

		trade2 := model.Trade{
			UId:        parent.ParentId,
			TradeType:  t2,
			ItemId:     productOrder.UId,
			Amount:     freeAmount,
			Before:     memberModel.WithdrawBalance,
			After:      memberModel.WithdrawBalance.Add(freeAmount),
			Desc:       fmt.Sprintf("%v级释放可用余额", level),
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err := trade2.Insert()
		if err != nil {
			logrus.Errorf("%v级释放可用余额失败  代理UId %v err= &v", level, parent.ParentId, err)
		}

		memberModel.Balance = memberModel.Balance.Sub(freeAmount)
		memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(freeAmount)
		err = memberModel.Update("balance", "withdraw_balance")
		if err != nil {
			logrus.Errorf("%v级释放可用余额失败 代理UId %v 收益 %v  err= &v", level, parent.ParentId, freeAmount, err)
		}
	}
}
