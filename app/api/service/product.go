package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type ProductList struct {
	request.ProductList
}

func (this ProductList) PageList() response.ProductListData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Product{}
	where, args, _ := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Product, 0)
	act := make([]response.ManSongActive, 0)
	acts := model.CouponActivity{}
	FullDelivery := acts.List()
	for i := range FullDelivery {
		act = append(act, response.ManSongActive{
			Amount: FullDelivery[i].Amount,
			Price:  FullDelivery[i].Coupon.Price,
			Id:     FullDelivery[i].Coupon.Id,
		})
	}
	for _, v := range list {
		//获取赠品产品名称
		//giftName := ""
		//if v.GiftId > 0 {
		//	giftModel := model.Product{
		//		Id: v.GiftId,
		//	}
		//	if giftModel.Get() {
		//		giftName = giftModel.Name
		//	}
		//}

		i := response.Product{
			Id:           v.Id,
			Name:         v.Name,
			Category:     v.Category,
			Type:         v.Type,
			Price:        v.Price,
			Img:          v.Img,
			Interval:     v.Interval,
			IncomeRate:   v.IncomeRate,
			LimitBuy:     v.LimitBuy,
			Total:        v.Total,
			Current:      v.Current,
			Desc:         v.Desc,
			DelayTime:    v.DelayTime,
			IsHot:        v.IsHot,
			IsFinished:   v.IsFinished,
			IsCouponGift: v.IsCouponGift,
			Status:       v.Status,
		}
		res = append(res, i)
	}
	return response.ProductListData{List: res, Page: FormatPage(page)}
}

type RecommendList struct {
	request.Request
}

func (this RecommendList) PageList() response.ProductListData {

	m := model.Product{}
	list, page := m.PageList("is_recommend = ? and c_product.status = ?", []interface{}{1, 1}, 1, 10)
	res := make([]response.Product, 0)
	act := make([]response.ManSongActive, 0)
	acts := model.CouponActivity{}
	FullDelivery := acts.List()
	for i := range FullDelivery {
		act = append(act, response.ManSongActive{
			Amount: FullDelivery[i].Amount,
			Price:  FullDelivery[i].Coupon.Price,
			Id:     FullDelivery[i].Coupon.Id,
		})
	}
	for _, v := range list {

		i := response.Product{
			Id:         v.Id,
			Name:       v.Name,
			Category:   v.Category,
			Type:       v.Type,
			Price:      v.Price,
			Img:        v.Img,
			Interval:   v.Interval,
			IncomeRate: v.IncomeRate,
			LimitBuy:   v.LimitBuy,
			Total:      v.Total,
			Current:    v.Current,
			Desc:       v.Desc,
			DelayTime:  v.DelayTime,
			//GiftId:                v.GiftId,
			//WithdrawThresholdRate: decimal.Decimal{},
			IsHot:  v.IsHot,
			Status: v.Status,
		}
		res = append(res, i)
	}
	return response.ProductListData{List: res, Page: FormatPage(page)}
}

type GetProduct struct {
	request.GetProduct
}

//func (this GetProduct) GetOne() response.Product {
//
//	m := model.Product{
//		Id:     this.Id,
//		Status: 1,
//	}
//	m.Get()
//
//	act := make([]response.ManSongActive, 0)
//	acts := model.FullDelivery{}
//	FullDelivery := acts.List()
//	for i := range FullDelivery {
//		act = append(act, response.ManSongActive{
//			Amount: FullDelivery[i].Amount,
//			Price:  FullDelivery[i].Coupon.Price,
//			Id:     FullDelivery[i].Coupon.Id,
//		})
//	}
//
//	//获取当前项目进度
//	startProgress := m.Progress
//	var progress decimal.Decimal
//	if m.OtherPrice.LessThanOrEqual(m.TotalPrice) {
//		usedAmount := m.TotalPrice.Sub(m.OtherPrice)
//		trueProgress := usedAmount.Div(m.TotalPrice).Round(2)
//		if trueProgress.LessThan(startProgress) {
//			progress = startProgress
//		} else {
//			progress = trueProgress
//		}
//	} else {
//		progress = decimal.NewFromFloat(0)
//	}
//
//	//获取赠品产品名称
//	giftName := ""
//	if m.GiftId > 0 {
//		giftModel := model.Product{
//			Id:     m.GiftId,
//			Status: 1,
//		}
//		if giftModel.Get() {
//			giftName = giftModel.Name
//		}
//	}
//
//	res := response.Product{
//		Id:           m.Id,
//		Name:         m.Name,
//		Category:     m.Category,
//		CategoryName: m.ProductCategory.Name,
//		Status:       m.Status,
//		Tag:          m.Tag,
//		TimeLimit:    m.TimeLimit,
//		IsRecommend:  m.IsRecommend,
//		Dayincome:    m.DayIncome,
//		Price:        m.Price,
//		TotalPrice:   m.TotalPrice,
//		OtherPrice:   m.OtherPrice,
//		MoreBuy:      m.MoreBuy,
//		Desc:         m.Desc,
//		CreateTime:   m.CreateTime,
//		IsFinish:     m.IsFinish,
//		IsManjian:    m.IsManjian,
//		BuyTimeLimit: m.BuyTimeLimit,
//		Progress:     progress,
//		Type:         m.Type,
//		GiftName:     giftName,
//	}
//	if res.IsManjian == 1 {
//		res.ManSongActive = act
//	}
//	return res
//}

func (this ProductList) getWhere() (string, []interface{}, error) {
	where := map[string]interface{}{
		model.Product{}.TableName() + ".status": model.StatusOk,
		"ProductCategory.status":                model.StatusOk,
	}
	if this.Category > 0 {
		where[model.Product{}.TableName()+".category"] = this.Category
	}
	if this.Name != "" {
		where[model.Product{}.TableName()+".name"] = this.Name
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals, nil
}

type ProductOptional struct {
	request.ProductOptional
}

func (this ProductOptional) Optional(member model.Member) (error, bool) {
	if this.Id == 0 {
		return errors.New(lang.Lang("Parameter error")), false
	}
	mo := model.MemberOptional{
		PId: this.Id,
		UId: member.Id,
	}
	if mo.Get() {
		return nil, true
	}
	return nil, false
}

type ProductCategory struct {
}

func (this ProductCategory) Category() response.ProductCategoryData {
	pc := model.ProductCategory{}
	where := "status = ?"
	args := []interface{}{model.StatusOk}
	list := pc.List(where, args)
	res := make([]response.ProductCategoryItem, 0)
	for _, v := range list {
		i := response.ProductCategoryItem{
			Id:   v.Id,
			Name: v.Name,
		}
		res = append(res, i)
	}
	return response.ProductCategoryData{List: res}
}

type GuQuanList struct {
	request.Request
}

func (this GuQuanList) List() *response.EquityListResp {

	m := model.Equity{}
	if !m.Get(true) {
		return nil
	}
	return &response.EquityListResp{
		Id:           m.Id,
		Total:        m.Total,
		Current:      m.Current,
		ReleaseRate:  m.ReleaseRate,
		Price:        m.Price,
		MinBuy:       m.MinBuy,
		HitRate:      m.HitRate,
		MissRate:     m.MissRate,
		SellRate:     m.SellRate,
		PreStartTime: m.PreStartTime,
		PreEndTime:   m.PreEndTime,
		OpenTime:     m.OpenTime,
		RecoverTime:  m.RecoverTime,
		Status:       m.Status,
	}

}

type ProductBuy struct {
	request.BuyReq
}

func (this *ProductBuy) Buy(member *model.Member) error {
	//实名认证
	if member.IsReal != 2 {
		return errors.New("请实名认证！")
	}
	//产品Id
	if this.Id <= 0 {
		return errors.New("产品Id格式不正确！")
	}
	//添加Redis乐观锁
	lockKey := fmt.Sprintf("product_buy:%v:%v", member.Id, this.Id)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	if !redisLock.Lock(lockKey) {
		return errors.New(lang.Lang("During data processing, Please try again later"))
	}
	defer redisLock.Unlock(lockKey)
	if this.Quantity <= 0 {
		return errors.New("数量错误！")
	}
	product := model.Product{Id: this.Id}
	if !product.Get() {
		return errors.New("产品不存在！")
	}
	if product.Type == 5 {
		return errors.New("产品为赠品！不能购买")
	}
	if product.IsFinished == model.StatusOk {
		return errors.New("项目已投满！")
	}
	amount := product.Price.Mul(decimal.NewFromInt(int64(this.Quantity)))
	if product.Total.LessThan(product.Current.Add(amount)) {
		return errors.New("可投额度不足！")
	}
	//余额检查
	if member.Balance.LessThanOrEqual(amount) {
		return errors.New("余额不足,请先充值！")
	}
	//交易密码验证
	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
	}
	//限购
	order := model.OrderProduct{}
	count := order.Count("uid = ?", []interface{}{member.Id})
	if int(count)+this.Quantity > product.LimitBuy {
		return errors.New(fmt.Sprintf("限购%v份！", product.LimitBuy))
	}
	//优惠券分析
	var memberCoupon model.MemberCoupon
	var couponAmount decimal.Decimal
	if this.UseId != 0 {
		memberCoupon = model.MemberCoupon{
			Uid:   int64(member.Id),
			Id:    this.UseId,
			IsUse: 1,
		}
		if !memberCoupon.Get() {
			return errors.New("没有找到可用的优惠券信息！")
		}
		couponAmount = memberCoupon.Coupon.Price
	}
	//24小时内购买赠送可提现余额
	var registerGift bool
	lastOrder := model.OrderProduct{UId: member.Id}
	if !lastOrder.Get() {
		registerGift = true
	}
	//基础配置表
	config := model.SetBase{}
	config.Get()
	//购买不同分类的产品的订单处理
	switch this.Cate {
	case 1:
		//购买
		inc := &model.OrderProduct{
			UId:          member.Id,
			Pid:          product.Id,
			PayMoney:     amount.Sub(couponAmount),
			IsReturnTop:  1,
			AfterBalance: member.Balance.Sub(amount.Sub(couponAmount)),
			CreateTime:   time.Now().Unix(),
			UpdateTime:   time.Now().Unix(),
			IncomeRate:   product.IncomeRate,
			EndTime:      time.Now().Unix() + int64(product.Interval*86400),
		}
		err := inc.Insert()
		if err != nil {
			return err
		}
		//减去可投余额
		product.Current = product.Current.Add(amount)
		if product.Total.LessThanOrEqual(product.Current) {
			product.IsFinished = model.StatusOk
		}
		err = product.Update("current", "is_finished")
		if err != nil {
			logrus.Errorf("购买产品减去可投余额失败%v", err)
		}
		//加入账变记录
		trade := model.Trade{
			UId:        member.Id,
			TradeType:  1,
			ItemId:     inc.Id,
			Amount:     amount,
			Before:     member.Balance,
			After:      member.Balance.Sub(amount),
			Desc:       "购买产品",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err = trade.Insert()
		if err != nil {
			logrus.Errorf("购买产品加入账变记录失败%v", err)
		}
		//扣减可用余额
		member.Balance = member.Balance.Sub(amount.Sub(couponAmount))
		member.IsBuy = 1

		//优惠券使用记录
		if decimal.Zero.LessThan(couponAmount) {
			trade3 := model.Trade{
				UId:        member.Id,
				TradeType:  10,
				ItemId:     int(this.UseId),
				Amount:     couponAmount,
				Before:     member.Balance,
				After:      member.Balance,
				Desc:       "使用优惠券",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err = trade3.Insert()
			if err != nil {
				logrus.Errorf("使用优惠券记录 加入账变记录失败%v", err)
			}
			//更改优惠券状态
			memberCoupon.IsUse = 2
			err = memberCoupon.Update("is_use")
			if err != nil {
				logrus.Errorf("修改用户优惠券失败%v", err)
			}
		}
		//赠送
		if registerGift && decimal.Zero.LessThan(config.RegisterSend) {
			//赠送礼金 加入账变记录
			trade2 := model.Trade{
				UId:        member.Id,
				TradeType:  7,
				ItemId:     inc.Id,
				Amount:     config.RegisterSend,
				Before:     member.WithdrawBalance,
				After:      member.WithdrawBalance.Add(config.RegisterSend),
				Desc:       "第一次购买赠送礼金",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err = trade2.Insert()
			if err != nil {
				logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
			}
			member.WithdrawBalance = member.WithdrawBalance.Add(config.RegisterSend)
			member.PreIncome = member.PreIncome.Add(product.IncomeRate.Mul(decimal.NewFromInt(int64(product.Interval))))
			member.PreCapital = member.PreCapital.Add(product.Price)
			err = member.Update("balance", "withdraw_balance", "total_income")
			if err != nil {
				logrus.Errorf("更改会员余额信息失败%v", err)
			}
		}

		//检查是否有满送活动
		if product.IsCouponGift == model.StatusOk {
			full := model.CouponActivity{}
			if full.Find(amount) {
				//满送活动加入账变记录
				trade3 := model.Trade{
					UId:        member.Id,
					TradeType:  9,
					ItemId:     int(full.Coupon.Id),
					Amount:     full.Coupon.Price,
					Before:     member.Balance,
					After:      member.Balance,
					Desc:       fmt.Sprintf("赠送%v优惠券", full.Coupon.Price),
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade3.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}
				couponIns := model.MemberCoupon{
					Uid:      int64(member.Id),
					CouponId: full.Coupon.Id,
					IsUse:    1,
				}
				err = couponIns.Insert()
				if err != nil {
					logrus.Errorf("赠赠送优惠券记录失败%v %v", err, member.Id)
				}
				//更改会员当前余额
				//memberModel.Balance += full.Coupon.Price
				//err = memberModel.Update("balance")
				//if err != nil {
				//	logrus.Errorf("更改会员余额信息失败%v", err)
				//}
			}
		}
		//上级返佣
		if inc.IsReturnTop == 1 {
			//1级代理佣金计算
			this.ProxyRebate(&config, 1, inc)
			//2级代理佣金计算
			this.ProxyRebate(&config, 2, inc)
			//3级代理佣金计算
			this.ProxyRebate(&config, 3, inc)
		}

		//赠品分析
		if product.GiftId > 0 {
			giftModel := model.Product{
				Id:     product.GiftId,
				Status: 1,
				Type:   5,
			}

			//当赠品产品信息存在时
			if giftModel.Get() {
				//赠品金额
				//giftAmount := this.Amount.Mul(config.GiftRate)
				//赠品订单
				orderModel := &model.OrderProduct{
					UId:          member.Id,
					Pid:          giftModel.Id,
					PayMoney:     giftModel.Price,
					IsReturnTop:  1,
					AfterBalance: member.Balance,
					CreateTime:   time.Now().Unix(),
					UpdateTime:   time.Now().Unix(),
				}
				err := orderModel.Insert()
				if err != nil {
					//解锁
					redisLock.Unlock(lockKey)
					return err
				}

				//加入账变记录
				logModel := model.Trade{
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
				err = logModel.Insert()
				if err != nil {
					logrus.Errorf("赠送赠品加入账变记录失败%v", err)
				}
				//更改用户收益
				member.PreIncome = member.PreIncome.Add(giftModel.IncomeRate.Mul(decimal.NewFromInt(int64(giftModel.Interval))))
				member.PreCapital = member.PreCapital.Add(giftModel.Price)
			}
		}
		err = member.Update("balance", "is_buy", "pre_income", "pre_capital")
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
		}
		return nil
	case 2:
		//股权
		//p := model.Guquan{Id: int64(this.Id)}
		//if !p.Get(true) {
		//	//解锁
		//	redisLock.Unlock(lockKey)
		//	return errors.New("股权不存在！")
		//}
		//if p.PreStartTime > time.Now().Unix() {
		//	//解锁
		//	redisLock.Unlock(lockKey)
		//	return errors.New("股权预售时间未开始")
		//}
		//if p.PreEndTime < time.Now().Unix() {
		//	//解锁
		//	redisLock.Unlock(lockKey)
		//	return errors.New("股权预售时间已结束")
		//}
		//if this.Amount.LessThan(decimal.NewFromInt(p.LimitBuy)) {
		//	//解锁
		//	redisLock.Unlock(lockKey)
		//	return errors.New(fmt.Sprintf("购买金额必须大于%v！", p.LimitBuy))
		//}
		//
		////购买
		//inc := &model.OrderGuquan{
		//	UId:          member.Id,
		//	Pid:          int(p.Id),
		//	PayMoney:     this.Amount,
		//	Rate:         decimal.NewFromInt(1),
		//	AfterBalance: memberModel.Balance.Sub(this.Amount),
		//	CreateTime:   time.Now().Unix(),
		//	UpdateTime:   time.Now().Unix(),
		//}
		//err := inc.Insert()
		//if err != nil {
		//	//解锁
		//	redisLock.Unlock(lockKey)
		//	return err
		//}
		//
		////减去可投余额
		//p.OtherGuquan = p.OtherGuquan.Sub(this.Amount)
		//err = p.Update("other_guquan")
		//if err != nil {
		//	logrus.Errorf("购买产品减去可投余额失败%v", err)
		//}
		//
		////加入账变记录
		//trade := model.Trade{
		//	UId:        member.Id,
		//	TradeType:  2,
		//	ItemId:     inc.Id,
		//	Amount:     this.Amount,
		//	Before:     memberModel.Balance,
		//	After:      memberModel.Balance.Sub(this.Amount),
		//	Desc:       "购买股权",
		//	CreateTime: time.Now().Unix(),
		//	UpdateTime: time.Now().Unix(),
		//	IsFrontend: 1,
		//}
		//err = trade.Insert()
		//if err != nil {
		//	logrus.Errorf("购买股权加入账变记录失败%v", err)
		//}
		//
		////扣减余额
		//memberModel.Balance = memberModel.Balance.Mul(this.Amount)
		//memberModel.IsBuy = 1
		//memberModel.Guquan = memberModel.Guquan.Add(this.Amount)
		//err = memberModel.Update("balance", "is_buy", "guquan")
		//if err != nil {
		//	logrus.Errorf("更改会员余额信息失败%v", err)
		//}
		//
		//if isSendRigster {
		//	//获取会员当前最新余额信息
		//	memberModel.Get()
		//	//收盘状态分析
		//	isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
		//	if isRetreatStatus == true {
		//		//可用余额转换比例分析, 默认为90%
		//		if config.IncomeBalanceRate.LessThanOrEqual(decimal.Zero) {
		//			config.IncomeBalanceRate = decimal.NewFromFloat(0.9)
		//		}
		//		//可用余额,可提现余额分析
		//		balanceAmount := config.IncomeBalanceRate.Mul(config.RegisterSend)
		//		useBalanceAmount := config.RegisterSend.Sub(balanceAmount)
		//
		//		//赠送礼金 加入账变记录
		//		trade2 := model.Trade{
		//			UId:        member.Id,
		//			TradeType:  7,
		//			ItemId:     inc.Id,
		//			Amount:     useBalanceAmount,
		//			Before:     memberModel.WithdrawBalance,
		//			After:      memberModel.WithdrawBalance.Add(useBalanceAmount),
		//			Desc:       "第一次购买赠送礼金",
		//			CreateTime: time.Now().Unix(),
		//			UpdateTime: time.Now().Unix(),
		//			IsFrontend: 1,
		//		}
		//		err = trade2.Insert()
		//		if err != nil {
		//			logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		//		}
		//
		//		trade2 = model.Trade{
		//			UId:        member.Id,
		//			TradeType:  7,
		//			ItemId:     inc.Id,
		//			Amount:     balanceAmount,
		//			Before:     memberModel.Balance,
		//			After:      memberModel.Balance.Add(balanceAmount),
		//			Desc:       "第一次购买赠送礼金",
		//			CreateTime: time.Now().Unix(),
		//			UpdateTime: time.Now().Unix(),
		//			IsFrontend: 1,
		//		}
		//		err = trade2.Insert()
		//		if err != nil {
		//			logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		//		}
		//
		//		//更改会员当前余额信息
		//		memberModel.Balance = memberModel.Balance.Add(balanceAmount)
		//		memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(useBalanceAmount)
		//	} else {
		//		//赠送礼金 加入账变记录
		//		trade2 := model.Trade{
		//			UId:        member.Id,
		//			TradeType:  7,
		//			ItemId:     inc.Id,
		//			Amount:     config.RegisterSend,
		//			Before:     memberModel.Balance,
		//			After:      memberModel.Balance.Add(config.RegisterSend),
		//			Desc:       "第一次购买赠送礼金",
		//			CreateTime: time.Now().Unix(),
		//			UpdateTime: time.Now().Unix(),
		//			IsFrontend: 1,
		//		}
		//		err = trade2.Insert()
		//		if err != nil {
		//			logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
		//		}
		//		//更改会员余额
		//		memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(config.RegisterSend)
		//	}
		//
		//	//memberModel.TotalBalance = memberModel.TotalBalance.Add(config.RegisterSend)
		//	memberModel.TotalIncome = memberModel.TotalIncome.Add(config.RegisterSend)
		//	err = memberModel.Update("balance", "withdraw_balance", "total_income")
		//	if err != nil {
		//		logrus.Errorf("更改会员余额信息失败%v", err)
		//	}
		//}

	default:
		return errors.New("购买类型不存在")
	}
	return nil
}

// 代理返佣, 购买产品后,立即返佣
func (this *ProductBuy) ProxyRebate(c *model.SetBase, level int, productOrder *model.OrderProduct) {
	//1级代理佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	agent := model.MemberParents{
		Uid:   productOrder.UId,
		Level: level,
	}
	//当代理不存在时
	if !agent.Get() {
		return
	}

	var income decimal.Decimal
	var t int
	if level == 1 {
		income = c.OneSend.Mul(productOrder.PayMoney)
		//检测会员的订单
		ordersModel := model.OrderProduct{
			UId:         productOrder.UId,
			IsReturnTop: 2,
		}
		//会员第一次下单, 直接上级发放红包
		if !ordersModel.Get() {
			income = income.Add(c.OneSendMoney)
		}
		t = 18
	} else if level == 2 {
		income = c.TwoSend.Mul(productOrder.PayMoney)
		t = 19
	} else if level == 3 {
		income = c.ThreeSend.Mul(productOrder.PayMoney)
		t = 20
	}

	memberModel := model.Member{Id: agent.ParentId}
	//获取代理当前余额
	memberModel.Get()

	//收盘状态分析
	isRetreatStatus := common.ParseRetreatStatus(c.RetreatStartDate)
	if isRetreatStatus == true && level != 1 {
		trade := model.Trade{
			UId:        agent.ParentId,
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

		//memberModel.TotalBalance = memberModel.TotalBalance.Add(income)
		memberModel.Balance = memberModel.Balance.Add(income)
		memberModel.TotalIncome = memberModel.TotalIncome.Add(income)
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
			UId:        agent.ParentId,
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
		memberModel.TotalIncome = memberModel.TotalIncome.Add(income)
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
			if c.OneReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.OneReleaseRate = 1000
			}
			t2 = 25
			freeAmount = c.OneReleaseRate.Mul(productOrder.PayMoney)
		case 2:
			if c.TwoReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.TwoReleaseRate = 500
			}
			t2 = 26
			freeAmount = c.TwoReleaseRate.Mul(productOrder.PayMoney)
		case 3:
			if c.ThreeReleaseRate.LessThanOrEqual(decimal.Zero) {
				return
				//c.ThreeReleaseRate = 200
			}
			t2 = 27
			freeAmount = c.ThreeReleaseRate.Mul(productOrder.PayMoney)
		}
		//可用余额分析
		if memberModel.Balance.LessThan(freeAmount) {
			freeAmount = memberModel.Balance
		}

		trade2 := model.Trade{
			UId:        agent.ParentId,
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
			logrus.Errorf("%v级释放可用余额失败  代理UId %v err= &v", level, agent.ParentId, err)
		}

		memberModel.Balance = memberModel.Balance.Sub(freeAmount)
		memberModel.WithdrawBalance = memberModel.WithdrawBalance.Add(freeAmount)
		err = memberModel.Update("balance", "withdraw_balance")
		if err != nil {
			logrus.Errorf("%v级释放可用余额失败 代理UId %v 收益 %v  err= &v", level, agent.ParentId, freeAmount, err)
		}
	}
}

type BuyProductList struct {
	request.ProductBuyList
}

func (this BuyProductList) List(member *model.Member) *response.BuyListResp {
	if this.Page == 0 {
		this.Page = 1
	}
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	m := model.OrderProduct{}
	res := response.BuyListResp{}
	list, page := m.PageList("uid = ?", []interface{}{member.Id}, this.Page, this.PageSize)
	if len(list) == 0 {
		return &res
	}
	items := make([]response.BuyList, 0)
	for i := range list {
		//订单状态
		orderStatus := 1
		if list[i].IsReturnCapital == 1 {
			orderStatus = 2
		}
		product := model.Product{Id: m.Pid}
		if !product.Get() {
			continue
		}
		//每日收益
		income := list[i].PayMoney.Mul(list[i].IncomeRate)
		items = append(items, response.BuyList{
			Name:    list[i].Product.Name,
			BuyTime: int(list[i].CreateTime),
			Amount:  list[i].PayMoney,
			Status:  orderStatus,
			Income:  income,
			EndTime: list[i].EndTime,
		})
	}
	res.List = items
	res.Page = FormatPage(page)
	return &res
}

type BuyGuquanList struct {
	request.Request
}

func (this *BuyGuquanList) List(member *model.Member) *response.BuyGuquanResp {
	var res response.BuyGuquanResp
	m := model.OrderEquity{UId: member.Id}
	money := m.Sum()
	guquan := model.Equity{}
	guquan.Get(true)

	now := time.Now().Unix()
	if guquan.RecoverTime >= now {
		res.Status = "完成"
	}
	if guquan.OpenTime >= now {
		res.Status = "待回收"
	}
	if guquan.PreEndTime >= now {
		res.Status = "待发行"
	}

	res.Num = money
	res.Price = guquan.Price
	res.CreateTime = m.CreateTime
	weiMoney := money.Mul(decimal.NewFromInt(1).Sub(m.Rate)).Add(guquan.MissRate)
	huiMoney := money.Mul(m.Rate).Mul(guquan.SellRate)
	res.TotalPrice = weiMoney.Add(huiMoney)
	return &res

}

type BuyGuquanPageList struct {
	request.Pagination
}

// 获取用户购买股权列表
func (this *BuyGuquanPageList) PageList(member *model.Member) *response.BuyGuquanPageListResp {
	//参数分析
	if this.Page == 0 {
		this.Page = 1
	}
	if this.PageSize == 0 {
		this.PageSize = response.DefaultPageSize
	}

	//获取列表
	orderModel := model.OrderEquity{UId: member.Id}
	list, page := orderModel.PageList("uid=?", []interface{}{member.Id}, this.Page, this.PageSize)

	//获取股权信息
	guquan := model.Equity{}
	guquan.Get(true)

	//now := time.Now().Unix()
	//Status := ""
	//if now >= guquan.ReturnTime {
	//	Status = "完成"
	//} else if now >= guquan.OpenTime {
	//	Status = "待回收"
	//} else {
	//	Status = "待发行"
	//}

	res := make([]response.BuyGuquanList, 0)
	for _, v := range list {
		fmt.Println(v.PayMoney)
		//未中签回购金额
		//weiMoney := (v.PayMoney * int64(int(model.UNITY)-v.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(guquan.ReturnRate)) / int64(model.UNITY)
		////中签回购金额
		//huiMoney := (v.PayMoney * int64(v.Rate) / int64(model.UNITY)) * int64(guquan.ReturnLuckyRate) / int64(model.UNITY)
		//i := response.BuyGuquanList{
		//	Id:         v.Id,
		//	Num:        v.PayMoney / int64(model.UNITY),
		//	Price:      float64(guquan.Price),
		//	CreateTime: v.CreateTime,
		//	TotalPrice: float64(weiMoney + huiMoney),
		//	Status:     Status,
		//}
		//res = append(res, i)
	}

	return &response.BuyGuquanPageListResp{List: res, Page: FormatPage(page)}
}

type StockCertificate struct {
	request.StockCertificate
}

func (this *StockCertificate) GetInfo(member *model.Member) *response.StockCertificateResp {
	//参数分析
	if this.Id == 0 {
		return nil
	}
	//获取股权信息
	guquan := model.Equity{}
	guquan.Get(true)

	//now := time.Now().Unix()
	//if now >= guquan.ReturnTime {
	//	return nil
	//}

	//获取订单信息
	orderModel := model.OrderEquity{Id: this.Id, UId: member.Id}
	if !orderModel.Get() {
		return nil
	}

	//获取用户信息
	memberVerfiy := model.MemberVerified{UId: orderModel.UId}
	memberVerfiy.Get()

	//合同起始时
	startDate := time.Unix(int64(guquan.OpenTime), 0).Format("2006年01月02日")
	endDate := time.Unix(int64(guquan.RecoverTime), 0).Format("2006年01月02日")
	days := int(guquan.RecoverTime-guquan.OpenTime) / 86400
	createDate := time.Unix(int64(orderModel.CreateTime), 0).Format("2006年01月02日")
	//
	////中签回购金额
	//huiMoney := (orderModel.PayMoney * int64(orderModel.Rate) / int64(model.UNITY)) * int64(guquan.ReturnLuckyRate) / int64(model.UNITY)
	////未中签回购金额
	//weiMoney := (orderModel.PayMoney * int64(int(model.UNITY)-orderModel.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(guquan.ReturnRate)) / int64(model.UNITY)
	//
	////原始股权总金额
	//sourceAmount := float64(orderModel.PayMoney) / float64(model.UNITY) * float64(guquan.Price) / float64(model.UNITY)
	////中签回购利润
	//winProfit := float64(guquan.ReturnLuckyRate)*100/model.UNITY - 100
	////未中签加购利润
	//notWinProfit := float64(guquan.ReturnRate) * 100
	//
	////总股权数量
	//totalQuantity := orderModel.PayMoney / int64(model.UNITY)
	////中签股权数量
	//winQuantity := orderModel.PayMoney * int64(orderModel.Rate) / (int64(model.UNITY) * int64(model.UNITY))
	////未中签股权数量
	//notWinQuantity := totalQuantity - winQuantity

	return &response.StockCertificateResp{
		Id:         orderModel.Id,
		RealName:   memberVerfiy.RealName,
		IdCardNo:   memberVerfiy.IdNumber,
		StartDate:  startDate,
		EndDate:    endDate,
		CreateDate: createDate,
		Days:       days,

		//股权总数
		//Quantity: orderModel.PayMoney / int64(model.UNITY),
		//原订单价格
		//Price: float64(guquan.Price),
		//原始股权总金额
		//TotalAmount: sourceAmount,

		//中签股权数
		//WinQuantity: winQuantity,
		//中签回购利润
		//WinProfit: winProfit,
		//中签股权回购总金额
		//WinRepurchaseAmount: float64(huiMoney),

		//未中签股权数
		//NotWinQuantity: notWinQuantity,
		//未中签回购利润
		//NotWinProfit: notWinProfit,
		//未中签回购金额
		//NotWinRepurchaseAmount: float64(weiMoney),
	}
}
