package logic

import (
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
func (this productBuyLogic) ProductBuy(pt int, member *model.Member, product model.Product, amount decimal.Decimal, quantity int, mc model.MemberCoupon, isFirst bool, ybMoney decimal.Decimal, ybGive decimal.Decimal) error {
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
		order, err := this.createOrder(member, product, amount, quantity, ybMoney)
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
		err = this.createTrade(member, order, amount, mc, ybMoney)
		if err != nil {
			return err
		}
		//医保余额修改
		member.YiBaoBalance = member.YiBaoBalance.Add(ybGive)
		//扣减可用余额 记录已购状态
		member.Balance = member.Balance.Sub(amount.Sub(mc.Coupon.Price).Sub(ybMoney))
		member.TotalBuy = member.TotalBuy.Add(amount)
		member.IsBuy = 1
		if mc.Id > 0 {
			//优惠券使用
			err = this.useCoupon(member, mc)
			if err != nil {
				return err
			}
		}
		//首次购买赠送
		if isFirst {
			err = this.firstBuy(member, order.Id)
			if err != nil {
				return err
			}
		}
		//记录用户待收益利息及本金
		member.PreIncome = member.PreIncome.Add(amount.Mul(product.IncomeRate).Mul(decimal.NewFromInt(int64(product.Interval))).Div(decimal.NewFromInt(100).Round(2)))
		member.PreCapital = member.PreCapital.Add(amount)
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
		}
		//检查是否有满送活动
		if product.IsCouponGift == model.StatusOk {
			activity := model.CouponActivity{}
			if activity.Find(amount) {
				//满送活动加入账变记录
				trade := model.Trade{
					UId:        member.Id,
					TradeType:  9,
					ItemId:     activity.Coupon.Id,
					Amount:     activity.Coupon.Price,
					Before:     member.Balance,
					After:      member.Balance,
					Desc:       fmt.Sprintf("赠送%v优惠券", activity.Coupon.Price),
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = this.tx.Create(&trade).Error
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
					return err
				}
				couponIns := model.MemberCoupon{
					Uid:      member.Id,
					CouponId: activity.Coupon.Id,
					IsUse:    1,
				}
				err = this.tx.Create(&couponIns).Error
				if err != nil {
					logrus.Errorf("赠赠送优惠券记录失败%v %v", err, member.Id)
					return err
				}
			}
		}
		//上级返佣
		if order.IsReturnTop == 1 {
			//1级代理佣金计算
			err = this.proxyRebate(1, order)
			if err != nil {
				return err
			}
			//2级代理佣金计算
			err = this.proxyRebate(2, order)
			if err != nil {
				return err
			}
			//3级代理佣金计算
			err = this.proxyRebate(3, order)
			if err != nil {
				return err
			}
		}
		//赠品分析
		if product.GiftId > 0 {
			err = this.gift(member, product)
			if err != nil {
				return err
			}
		}
		//用户可提额度增加
		member.WithdrawThreshold = member.WithdrawThreshold.Add(product.WithdrawThresholdRate.Mul(amount).Div(decimal.NewFromInt(100)).Round(2))
		//扣除医保卡余额
		member.YiBaoBalance = member.YiBaoBalance.Sub(ybMoney)
		err = this.tx.Select("balance", "withdraw_balance", "is_buy", "pre_income", "pre_capital", "withdraw_threshold", "total_buy", "yibao_balance").Updates(member).Error
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
			return err
		}
	case 2: //股权

	}
	return nil
}
func (this productBuyLogic) createOrder(member *model.Member, product model.Product, amount decimal.Decimal, quantity int, ybMoney decimal.Decimal) (*model.OrderProduct, error) {
	//购买
	inc := &model.OrderProduct{
		UId:          member.Id,
		Pid:          product.Id,
		PayMoney:     amount,
		IsReturnTop:  1,
		AfterBalance: member.Balance.Sub(amount.Sub(ybMoney)),
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
		IncomeRate:   product.IncomeRate,
		EndTime:      time.Now().Unix() + int64(product.Interval*86400),
		Quantity:     quantity,
		YbAmount:     ybMoney,
	}
	err := this.tx.Create(&inc).Error
	if err != nil {
		return nil, err
	}
	return inc, nil
}
func (this productBuyLogic) createTrade(member *model.Member, order *model.OrderProduct, amount decimal.Decimal, mc model.MemberCoupon, ybMoney decimal.Decimal) error {
	//加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  1,
		ItemId:     order.Id,
		Amount:     amount.Sub(mc.Coupon.Price).Sub(ybMoney),
		Before:     member.Balance,
		After:      member.Balance.Sub(amount.Sub(mc.Coupon.Price).Sub(ybMoney)),
		Desc:       "购买产品",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(&trade).Error
	if err != nil {
		logrus.Errorf("购买产品加入账变记录失败%v", err)
		return err
	}
	return nil
}
func (this productBuyLogic) useCoupon(member *model.Member, mc model.MemberCoupon) error {
	//优惠券使用记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  10,
		ItemId:     mc.Coupon.Id,
		Amount:     mc.Coupon.Price,
		Before:     member.Balance,
		After:      member.Balance,
		Desc:       "使用优惠券",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(&trade).Error
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
func (this productBuyLogic) firstBuy(member *model.Member, orderId int) error {
	//赠送礼金 加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  7,
		ItemId:     orderId,
		Amount:     this.config.RegisterSend,
		Before:     member.WithdrawBalance,
		After:      member.WithdrawBalance.Add(this.config.RegisterSend),
		Desc:       "第一次购买赠送礼金",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(&trade).Error
	if err != nil {
		logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		return err
	}
	//赠送礼金
	member.WithdrawBalance = member.WithdrawBalance.Add(this.config.RegisterSend)
	return nil
}
func (this *productBuyLogic) proxyRebate(level int, productOrder *model.OrderProduct) error {
	var err error
	//佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	member := model.MemberParents{
		Uid:   productOrder.UId,
		Level: level,
	}
	if !member.Get() {
		return nil
	}
	var income decimal.Decimal
	var t int
	var threshold decimal.Decimal
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
		//提现额度
		if time.Now().Unix() < this.config.EquityStartDate {
			threshold = this.config.OneReleaseRate.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		}

	} else if level == 2 {
		income = this.config.TwoSend.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		t = 19
		//提现额度
		if time.Now().Unix() < this.config.EquityStartDate {
			threshold = this.config.TwoReleaseRate.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		}
	} else if level == 3 {
		income = this.config.ThreeSend.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		t = 20
		//提现额度
		if time.Now().Unix() < this.config.EquityStartDate {
			threshold = this.config.ThreeReleaseRate.Mul(productOrder.PayMoney).Div(decimal.NewFromInt(100)).Round(2)
		}
	}
	if member.ParentId <= 0 {
		return nil
	}
	parent := model.Member{Id: member.ParentId}
	//获取代理当前余额
	if !parent.Get() {
		return nil
	}

	trade := model.Trade{
		UId:        parent.Id,
		TradeType:  t,
		ItemId:     productOrder.UId,
		Amount:     income,
		Before:     parent.WithdrawBalance,
		After:      parent.WithdrawBalance.Add(income),
		Desc:       fmt.Sprintf("%v级返佣", level),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err = this.tx.Create(&trade).Error
	if err != nil {
		logrus.Errorf("%v级返佣收益存入账单失败  用户Id %v err= &v", level, productOrder.UId, err)
		return err
	}
	parent.WithdrawBalance = parent.WithdrawBalance.Add(income)

	//增加用户返佣金额
	parent.TotalRebate = parent.TotalRebate.Add(income)
	//增加用户可提现额度
	parent.WithdrawThreshold = parent.WithdrawThreshold.Add(threshold)

	err = this.tx.Select("withdraw_threshold", "total_rebate", "withdraw_balance").Updates(parent).Error
	if err != nil {
		logrus.Errorf("%v级返佣收益修改余额失败 用户Id %v 收益 %v  err= &v", level, productOrder.UId, income, err)
		return err
	}
	//修改产品状态
	productOrder.IsReturnTop = 2
	err = this.tx.Select("is_return_top").Updates(productOrder).Error
	if err != nil {
		logrus.Errorf("修改产品状态失败   订单Id %v err= &v", productOrder.Id, err)
		return err
	}
	return nil
}
func (this productBuyLogic) gift(member *model.Member, product model.Product) error {
	giftModel := model.Product{
		Id:     product.GiftId,
		Status: 1,
		Type:   5,
	}
	if !giftModel.Get() {
		return nil
	}
	//赠品订单
	orderModel := &model.OrderProduct{
		UId:          member.Id,
		Pid:          giftModel.Id,
		PayMoney:     giftModel.Price,
		AfterBalance: member.Balance,
		IsReturnTop:  1,
		IncomeRate:   giftModel.IncomeRate,
		EndTime:      time.Now().Unix() + int64(giftModel.Interval*86400),
	}
	err := this.tx.Create(&orderModel).Error
	if err != nil {
		return err
	}
	//加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  1,
		ItemId:     orderModel.Id,
		Amount:     decimal.Zero,
		Before:     member.Balance,
		After:      member.Balance,
		Desc:       fmt.Sprintf("赠送:%v", giftModel.Name),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err = this.tx.Create(&trade).Error
	if err != nil {
		logrus.Errorf("赠送赠品加入账变记录失败%v", err)
	}
	//更改用户收益
	member.PreIncome = member.PreIncome.Add(giftModel.Price.Mul(giftModel.IncomeRate).Mul(decimal.NewFromInt(int64(giftModel.Interval))).Div(decimal.NewFromInt(100).Round(2)))
	member.WithdrawThreshold = member.WithdrawThreshold.Add(giftModel.WithdrawThresholdRate.Mul(giftModel.Price).Div(decimal.NewFromInt(100)).Round(2))
	return nil
}
