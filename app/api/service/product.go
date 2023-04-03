package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/common"
	"finance/global"
	"finance/lang"
	"finance/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
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
	acts := model.FullDelivery{}
	FullDelivery := acts.List()
	for i := range FullDelivery {
		act = append(act, response.ManSongActive{
			Amount: float64(FullDelivery[i].Amout) / model.UNITY,
			Price:  float64(FullDelivery[i].Coupon.Price) / model.UNITY,
			Id:     FullDelivery[i].Coupon.ID,
		})
	}
	for _, v := range list {
		//获取赠品产品名称
		giftName := ""
		if v.GiftId > 0 {
			giftModel := model.Product{
				ID: v.GiftId,
			}
			if giftModel.Get() {
				giftName = giftModel.Name
			}
		}

		i := response.Product{
			ID:           v.ID,
			Name:         v.Name,
			Category:     v.Category,
			CategoryName: v.ProductCategory.Name,
			Status:       v.Status,
			Tag:          v.Tag,
			TimeLimit:    v.TimeLimit,
			IsRecommend:  v.IsRecommend,
			Dayincome:    float64(v.Dayincome) / model.UNITY,
			Price:        float64(v.Price) / model.UNITY,
			TotalPrice:   float64(v.TotalPrice) / model.UNITY,
			OtherPrice:   float64(v.OtherPrice) / model.UNITY,
			MoreBuy:      v.MoreBuy,
			Desc:         v.Desc,
			CreateTime:   v.CreateTime,
			IsFinish:     v.IsFinish,
			IsManjian:    v.IsManjian,
			BuyTimeLimit: v.BuyTimeLimit,
			Type:         v.Type,
			DelayTime:    v.DelayTime,
			Progress:     float64(v.Progress) / model.UNITY,
			GiftName:     giftName,
		}
		if v.IsManjian == 1 {
			i.ManSongActive = act
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
	acts := model.FullDelivery{}
	FullDelivery := acts.List()
	for i := range FullDelivery {
		act = append(act, response.ManSongActive{
			Amount: float64(FullDelivery[i].Amout) / model.UNITY,
			Price:  float64(FullDelivery[i].Coupon.Price) / model.UNITY,
			Id:     FullDelivery[i].Coupon.ID,
		})
	}
	for _, v := range list {

		i := response.Product{
			ID:           v.ID,
			Name:         v.Name,
			Category:     v.Category,
			CategoryName: v.ProductCategory.Name,
			Status:       v.Status,
			Tag:          v.Tag,
			TimeLimit:    v.TimeLimit,
			IsRecommend:  v.IsRecommend,
			Dayincome:    float64(v.Dayincome) / model.UNITY,
			Price:        float64(v.Price) / model.UNITY,
			TotalPrice:   float64(v.TotalPrice) / model.UNITY,
			OtherPrice:   float64(v.OtherPrice) / model.UNITY,
			MoreBuy:      v.MoreBuy,
			Desc:         v.Desc,
			CreateTime:   v.CreateTime,
			IsFinish:     v.IsFinish,
			IsManjian:    v.IsManjian,
			BuyTimeLimit: v.BuyTimeLimit,
			Type:         v.Type,
			DelayTime:    v.DelayTime,
			Progress:     float64(v.Progress) / model.UNITY,
		}
		if v.IsManjian == 1 {
			i.ManSongActive = act
		}
		res = append(res, i)
	}
	return response.ProductListData{List: res, Page: FormatPage(page)}
}

type GetProduct struct {
	request.GetProduct
}

func (this GetProduct) GetOne() response.Product {

	m := model.Product{
		ID:     this.Id,
		Status: 1,
	}
	m.Get()

	act := make([]response.ManSongActive, 0)
	acts := model.FullDelivery{}
	FullDelivery := acts.List()
	for i := range FullDelivery {
		act = append(act, response.ManSongActive{
			Amount: float64(FullDelivery[i].Amout) / model.UNITY,
			Price:  float64(FullDelivery[i].Coupon.Price) / model.UNITY,
			Id:     FullDelivery[i].Coupon.ID,
		})
	}

	//获取当前项目进度
	startProgress := float64(m.Progress) / model.UNITY
	progress := 0.00
	if m.TotalPrice >= m.OtherPrice {
		usedAmount := m.TotalPrice - m.OtherPrice
		trueProgress, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", float64(usedAmount)/float64(m.TotalPrice)), 64)
		if startProgress > trueProgress {
			progress = startProgress
		} else {
			progress = trueProgress
		}
	} else {
		progress = 1.00
	}

	//获取赠品产品名称
	giftName := ""
	if m.GiftId > 0 {
		giftModel := model.Product{
			ID:     m.GiftId,
			Status: 1,
		}
		if giftModel.Get() {
			giftName = giftModel.Name
		}
	}

	res := response.Product{
		ID:           m.ID,
		Name:         m.Name,
		Category:     m.Category,
		CategoryName: m.ProductCategory.Name,
		Status:       m.Status,
		Tag:          m.Tag,
		TimeLimit:    m.TimeLimit,
		IsRecommend:  m.IsRecommend,
		Dayincome:    float64(m.Dayincome) / model.UNITY,
		Price:        float64(m.Price) / model.UNITY,
		TotalPrice:   float64(m.TotalPrice) / model.UNITY,
		OtherPrice:   float64(m.OtherPrice) / model.UNITY,
		MoreBuy:      m.MoreBuy,
		Desc:         m.Desc,
		CreateTime:   m.CreateTime,
		IsFinish:     m.IsFinish,
		IsManjian:    m.IsManjian,
		BuyTimeLimit: m.BuyTimeLimit,
		Progress:     progress,
		Type:         m.Type,
		GiftName:     giftName,
	}
	if res.IsManjian == 1 {
		res.ManSongActive = act
	}
	return res
}

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
	if this.ID == 0 {
		return errors.New(lang.Lang("Parameter error")), false
	}
	mo := model.MemberOptional{
		PID: this.ID,
		UID: member.ID,
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
			ID:   v.ID,
			Name: v.Name,
		}
		res = append(res, i)
	}
	return response.ProductCategoryData{List: res}
}

type GuQuanList struct {
	request.Request
}

func (this GuQuanList) List() *response.GuquanListResp {

	m := model.Guquan{}
	if !m.Get(true) {
		return nil
	}
	return &response.GuquanListResp{
		Id:              m.ID,
		TotalGuquan:     m.TotalGuquan,
		OtherGuquan:     m.OtherGuquan,
		ReleaseRate:     float64(m.ReleaseRate) / model.UNITY,
		Price:           float64(m.Price) / model.UNITY,
		LimitBuy:        m.LimitBuy,
		LuckyRate:       float64(m.LuckyRate) / model.UNITY,
		ReturnRate:      float64(m.LuckyRate) / model.UNITY,
		ReturnLuckyRate: float64(m.ReturnLuckyRate) / model.UNITY,
		PreStartTime:    m.PreStartTime,
		PreEndTime:      m.PreEndTime,
		OpenTime:        m.OpenTime,
		ReturnTime:      m.ReturnTime,
		Status:          m.Status,
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
	//产品ID
	if this.Id <= 0 {
		return errors.New("产品ID格式不正确！")
	}
	//金额分析
	if this.Amount <= 0 {
		return errors.New("购买金额必须大于0！")
	}
	//交易密码验证
	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
	}

	//金额分析
	amount := int64(this.Amount * model.UNITY)

	//优惠券分析
	var couponAmount int64
	if this.UseId != 0 {
		MemberCoupon := model.MemberCoupon{
			Uid:   int64(member.ID),
			ID:    this.UseId,
			IsUse: 1,
		}
		if !MemberCoupon.Get() {
			return errors.New("没有找到可用的优惠券信息！")
		}
		couponAmount = MemberCoupon.Coupon.Price
	}

	//添加Redis乐观锁
	lockKey := fmt.Sprintf("redisLock:api:submitProductOrder:memberId_%s_productId_%d_amount_%d", member.ID, this.Id, amount)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	lockResult := redisLock.Lock(lockKey)
	if !lockResult {
		return errors.New(lang.Lang("During data processing, Please try again later"))
	}

	//获取会员当前最新余额
	memberModel := model.Member{
		ID: member.ID,
	}
	memberModel.Get()

	//所需用户金额
	needAmount := amount - couponAmount
	if needAmount > memberModel.Balance {
		//解锁
		redisLock.Unlock(lockKey)
		return errors.New("余额不足,请先充值！")
	}

	//检查用户是否在注册24小时内第一次购买产品或股权
	var isSendRigster bool
	if member.RegTime+24*3600 >= time.Now().Unix() {
		guquan := model.OrderGuquan{UID: member.ID}
		product := model.OrderProduct{UID: member.ID}
		guquanNum, _ := guquan.Count()
		productNum := product.Count("uid = ?", []interface{}{member.ID})
		if guquanNum == 0 && productNum == 0 {
			isSendRigster = true
		}
	}

	//基础配置表
	config := model.SetBase{}
	config.Get()

	//购买不同分类的产品的订单处理
	switch this.Cate {
	case 1:
		//产品
		p := model.Product{ID: this.Id, Status: 1}
		if !p.Get() {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("产品不存在！")
		}
		//赠品分析
		if p.Type == 5 {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("产品为赠品！不能购买")
		}
		if int64(this.Amount*model.UNITY) < p.Price {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New(fmt.Sprintf("购买金额必须大于%v！", float64(p.Price)/model.UNITY))
		}
		if int(this.Amount) > p.MoreBuy {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New(fmt.Sprintf("购买金额必须小于%v！", p.MoreBuy))
		}

		money := model.OrderProduct{}
		wheres := "uid = ? and pid=?"
		agrss := []interface{}{member.ID, this.Id}
		pmoney := money.Sum(wheres, agrss, "pay_money")
		if int(float64((pmoney+int64(this.Amount*model.UNITY)))/model.UNITY) > p.MoreBuy {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New(fmt.Sprintf("您购买的产品已达到上限%v", p.MoreBuy))
		}
		if p.OtherPrice < int64(this.Amount*model.UNITY) {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("项目可投余额不足")
		}

		//购买
		inc := &model.OrderProduct{
			UID:          member.ID,
			Pid:          p.ID,
			PayMoney:     amount,
			IsReturnTop:  1,
			AfterBalance: memberModel.Balance - needAmount,
			CreateTime:   time.Now().Unix(),
			UpdateTime:   time.Now().Unix(),
		}
		err := inc.Insert()
		if err != nil {
			//解锁
			redisLock.Unlock(lockKey)
			return err
		}

		//减去可投余额
		p.OtherPrice -= amount
		err = p.Update("other_price")
		if err != nil {
			logrus.Errorf("购买产品减去可投余额失败%v", err)
		}

		//加入账变记录
		trade := model.Trade{
			UID:        member.ID,
			TradeType:  1,
			ItemID:     inc.ID,
			Amount:     needAmount,
			Before:     memberModel.Balance,
			After:      memberModel.Balance - needAmount,
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
		memberModel.Balance -= needAmount
		memberModel.IsBuy = 1
		err = memberModel.Update("balance", "is_buy")
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
		}

		//优惠券使用记录
		if couponAmount > 0 {
			//使用优惠券记录
			trade3 := model.Trade{
				UID:        member.ID,
				TradeType:  10,
				ItemID:     int(this.UseId),
				Amount:     couponAmount,
				Before:     memberModel.Balance,
				After:      memberModel.Balance,
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
			MemberCoupon2 := model.MemberCoupon{
				ID: this.UseId,
			}
			MemberCoupon2.IsUse = 2
			err = MemberCoupon2.Update("is_use")
			if err != nil {
				logrus.Errorf("修改用户优惠券失败%v", err)
			}
		}

		if isSendRigster {
			//获取会员当前最新余额信息
			memberModel.Get()
			//收盘状态分析
			isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
			if isRetreatStatus == true {
				//可用余额转换比例分析, 默认为90%
				if config.IncomeBalanceRate == 0 {
					config.IncomeBalanceRate = 9000
				}
				//可用余额,可提现余额分析
				balanceAmount := int64(config.IncomeBalanceRate) / int64(model.UNITY) * int64(config.RegisterSend)
				useBalanceAmount := int64(config.RegisterSend) - balanceAmount

				//赠送礼金 加入账变记录
				trade2 := model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     useBalanceAmount,
					Before:     memberModel.UseBalance,
					After:      memberModel.UseBalance + useBalanceAmount,
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}

				trade2 = model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     balanceAmount,
					Before:     memberModel.Balance,
					After:      memberModel.Balance + balanceAmount,
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}

				//更改会员当前余额信息
				memberModel.Balance += balanceAmount
				memberModel.UseBalance += useBalanceAmount
			} else {
				//赠送礼金 加入账变记录
				trade2 := model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     int64(config.RegisterSend),
					Before:     memberModel.UseBalance,
					After:      memberModel.UseBalance + int64(config.RegisterSend),
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}

				//更改会员当前余额信息
				memberModel.Balance += 0
				memberModel.UseBalance += int64(config.RegisterSend)
			}

			memberModel.TotalBalance += int64(config.RegisterSend)
			memberModel.Income += int64(config.RegisterSend)
			err = memberModel.Update("total_balance", "balance", "use_balance", "income")
			if err != nil {
				logrus.Errorf("更改会员余额信息失败%v", err)
			}
		}

		//检查是否有满送活动
		if p.IsManjian == 1 {
			//获取会员当前最新余额
			memberModel.Get()
			full := model.FullDelivery{}
			if full.Find(amount) {
				//满送活动加入账变记录
				trade3 := model.Trade{
					UID:        member.ID,
					TradeType:  9,
					ItemID:     int(full.Coupon.ID),
					Amount:     full.Coupon.Price,
					Before:     memberModel.Balance,
					After:      memberModel.Balance + full.Coupon.Price,
					Desc:       "赠送优惠券",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade3.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}
				MemberCoupon := model.MemberCoupon{
					Uid:      int64(member.ID),
					CouponId: full.Coupon.ID,
					IsUse:    1,
				}
				err = MemberCoupon.Insert()
				if err != nil {
					logrus.Errorf("赠赠送优惠券记录失败%v %v", err, member.ID)
				}
				//更改会员当前余额
				memberModel.Balance += full.Coupon.Price
				err = memberModel.Update("balance")
				if err != nil {
					logrus.Errorf("更改会员余额信息失败%v", err)
				}
			}
		}

		//更改用户收益
		memberModel.WillIncome += amount * int64(p.Dayincome*p.TimeLimit) / int64(model.UNITY)
		err = memberModel.Update("wll_income")
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
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
		if p.GiftId > 0 {
			giftModel := model.Product{
				ID:     p.GiftId,
				Status: 1,
				Type:   5,
			}

			//当赠品产品信息存在时
			if giftModel.Get() {
				//获取会员当前最新余额信息
				memberModel.Get()
				//赠品金额 @todo
				giftAmount := amount * int64(config.GiftRate) / int64(model.UNITY)

				//赠品订单
				orderModel := &model.OrderProduct{
					UID:          member.ID,
					Pid:          giftModel.ID,
					PayMoney:     giftAmount,
					IsReturnTop:  1,
					AfterBalance: memberModel.Balance,
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
					UID:        member.ID,
					TradeType:  1,
					ItemID:     orderModel.ID,
					Amount:     giftAmount,
					Before:     memberModel.Balance,
					After:      memberModel.Balance,
					Desc:       "赠品:购买产品赠送",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = logModel.Insert()
				if err != nil {
					logrus.Errorf("赠送赠品加入账变记录失败%v", err)
				}

				//更改用户收益
				memberModel.WillIncome += giftAmount * int64(giftModel.Dayincome*p.TimeLimit) / int64(model.UNITY)
				err = memberModel.Update("wll_income")
				if err != nil {
					logrus.Errorf("赠送赠品更改会员收益信息失败%v", err)
				}
			}
		}

		//解锁
		redisLock.Unlock(lockKey)
		return nil

	case 2:
		//股权
		p := model.Guquan{ID: int64(this.Id)}
		if !p.Get(true) {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("股权不存在！")
		}
		if p.PreStartTime > time.Now().Unix() {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("股权预售时间未开始")
		}
		if p.PreEndTime < time.Now().Unix() {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New("股权预售时间已结束")
		}
		if int64(this.Amount) < p.LimitBuy {
			//解锁
			redisLock.Unlock(lockKey)
			return errors.New(fmt.Sprintf("购买金额必须大于%v！", p.LimitBuy))
		}

		//购买
		inc := &model.OrderGuquan{
			UID:          member.ID,
			Pid:          int(p.ID),
			PayMoney:     amount,
			Rate:         int(model.UNITY),
			AfterBalance: memberModel.Balance - amount,
			CreateTime:   time.Now().Unix(),
			UpdateTime:   time.Now().Unix(),
		}
		err := inc.Insert()
		if err != nil {
			//解锁
			redisLock.Unlock(lockKey)
			return err
		}

		//减去可投余额
		p.OtherGuquan -= int64(this.Amount)
		err = p.Update("other_guquan")
		if err != nil {
			logrus.Errorf("购买产品减去可投余额失败%v", err)
		}

		//加入账变记录
		trade := model.Trade{
			UID:        member.ID,
			TradeType:  2,
			ItemID:     inc.ID,
			Amount:     amount,
			Before:     memberModel.Balance,
			After:      memberModel.Balance - amount,
			Desc:       "购买股权",
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			IsFrontend: 1,
		}
		err = trade.Insert()
		if err != nil {
			logrus.Errorf("购买股权加入账变记录失败%v", err)
		}

		//扣减余额
		memberModel.Balance -= amount
		memberModel.IsBuy = 1
		memberModel.Guquan += amount / int64(model.UNITY)
		err = memberModel.Update("balance", "is_buy", "guquan")
		if err != nil {
			logrus.Errorf("更改会员余额信息失败%v", err)
		}

		if isSendRigster {
			//获取会员当前最新余额信息
			memberModel.Get()
			//收盘状态分析
			isRetreatStatus := common.ParseRetreatStatus(config.RetreatStartDate)
			if isRetreatStatus == true {
				//可用余额转换比例分析, 默认为90%
				if config.IncomeBalanceRate == 0 {
					config.IncomeBalanceRate = 9000
				}
				//可用余额,可提现余额分析
				balanceAmount := int64(config.IncomeBalanceRate) / int64(model.UNITY) * int64(config.RegisterSend)
				useBalanceAmount := int64(config.RegisterSend) - balanceAmount

				//赠送礼金 加入账变记录
				trade2 := model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     useBalanceAmount,
					Before:     memberModel.UseBalance,
					After:      memberModel.UseBalance + useBalanceAmount,
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}

				trade2 = model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     balanceAmount,
					Before:     memberModel.Balance,
					After:      memberModel.Balance + balanceAmount,
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}

				//更改会员当前余额信息
				memberModel.Balance += balanceAmount
				memberModel.UseBalance += useBalanceAmount
			} else {
				//赠送礼金 加入账变记录
				trade2 := model.Trade{
					UID:        member.ID,
					TradeType:  7,
					ItemID:     inc.ID,
					Amount:     int64(config.RegisterSend),
					Before:     memberModel.Balance,
					After:      memberModel.Balance + int64(config.RegisterSend),
					Desc:       "第一次购买赠送礼金",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade2.Insert()
				if err != nil {
					logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
				}
				//更改会员余额
				memberModel.Balance += 0
				memberModel.UseBalance += int64(config.RegisterSend)
			}

			memberModel.TotalBalance += int64(config.RegisterSend)
			memberModel.Income += int64(config.RegisterSend)
			err = memberModel.Update("total_balance", "balance", "use_balance", "income")
			if err != nil {
				logrus.Errorf("更改会员余额信息失败%v", err)
			}
		}

	default:
		//解锁
		redisLock.Unlock(lockKey)
		return errors.New("购买类型不存在")
	}
	//解锁
	redisLock.Unlock(lockKey)
	return nil
}

// 代理返佣, 购买产品后,立即返佣
func (this *ProductBuy) ProxyRebate(c *model.SetBase, level int64, productOrder *model.OrderProduct) {
	//1级代理佣金计算  18=一级返佣 19=二级返佣 20=三级返佣
	agent := model.MemberRelation{
		UID:   productOrder.UID,
		Level: level,
	}
	//当代理不存在时
	if !agent.Get() {
		return
	}

	var income int64
	var t int
	if level == 1 {
		income = int64(c.OneSend) * productOrder.PayMoney / int64(model.UNITY)
		//检测会员的订单
		ordersModel := model.OrderProduct{
			UID:         productOrder.UID,
			IsReturnTop: 2,
		}
		//会员第一次下单, 直接上级发放红包
		if !ordersModel.Get() {
			income += c.OneSendMoeny
		}
		t = 18
	} else if level == 2 {
		income = int64(c.TwoSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 19
	} else if level == 3 {
		income = int64(c.ThreeSend) * productOrder.PayMoney / int64(model.UNITY)
		t = 20
	}

	memberModel := model.Member{ID: agent.Puid}
	//获取代理当前余额
	memberModel.Get()

	trade := model.Trade{
		UID:        agent.Puid,
		TradeType:  t,
		ItemID:     productOrder.UID,
		Amount:     income,
		Before:     memberModel.UseBalance,
		After:      memberModel.UseBalance + income,
		Desc:       fmt.Sprintf("%v级返佣", level),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := trade.Insert()
	if err != nil {
		logrus.Errorf("%v级返佣收益存入账单失败  用户ID %v err= &v", level, productOrder.UID, err)
	}

	memberModel.TotalBalance += income
	memberModel.UseBalance += income
	memberModel.Income += income
	err = memberModel.Update("total_balance", "use_balance", "income")
	if err != nil {
		logrus.Errorf("%v级返佣收益修改余额失败 用户ID %v 收益 %v  err= &v", level, productOrder.UID, income, err)
	}

	//修改产品状态
	productOrder.IsReturnTop = 2
	err = productOrder.Update("is_return_top")
	if err != nil {
		logrus.Errorf("修改产品状态失败   订单ID %v err= &v", productOrder.ID, err)
	}
}

type BuyProducList struct {
	request.ProductBuyList
}

func (this BuyProducList) List(member *model.Member) *response.BuyListResp {
	if this.Page == 0 {
		this.Page = 1
	}
	if this.PageSize == 0 {
		this.PageSize = 10
	}
	m := model.OrderProduct{}
	res := response.BuyListResp{}
	list, page := m.PageList("uid = ?", []interface{}{member.ID}, this.Page, this.PageSize)
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
		items = append(items, response.BuyList{
			Name:    list[i].Product.Name,
			BuyTime: int(list[i].CreateTime),
			Amount:  float64(list[i].PayMoney) / model.UNITY,
			Status:  orderStatus,
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
	m := model.OrderGuquan{UID: member.ID}
	money, err := m.Sum()
	if err != nil || money == 0 {
		return nil
	}

	guquan := model.Guquan{}
	guquan.Get(true)

	now := time.Now().Unix()
	if guquan.ReturnTime >= now {
		res.Status = "完成"
	}
	if guquan.OpenTime >= now {
		res.Status = "待回收"
	}
	if guquan.PreEndTime >= now {
		res.Status = "待发行"
	}

	res.Num = money / int64(model.UNITY)
	res.Price = float64(guquan.Price) / model.UNITY
	res.CreateTime = m.CreateTime
	weiMoney := (money * int64(int(model.UNITY)-m.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(guquan.ReturnRate)) / int64(model.UNITY)
	huiMoney := (money * int64(m.Rate) / int64(model.UNITY)) * int64(guquan.ReturnLuckyRate) / int64(model.UNITY)
	res.TotalPrice = float64(weiMoney+huiMoney) / model.UNITY
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
	orderModel := model.OrderGuquan{UID: member.ID}
	list, page := orderModel.PageList("uid=?", []interface{}{member.ID}, this.Page, this.PageSize)

	//获取股权信息
	guquan := model.Guquan{}
	guquan.Get(true)

	now := time.Now().Unix()
	Status := ""
	if now >= guquan.ReturnTime {
		Status = "完成"
	} else if now >= guquan.OpenTime {
		Status = "待回收"
	} else {
		Status = "待发行"
	}

	res := make([]response.BuyGuquanList, 0)
	for _, v := range list {
		//未中签回购金额
		weiMoney := (v.PayMoney * int64(int(model.UNITY)-v.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(guquan.ReturnRate)) / int64(model.UNITY)
		//中签回购金额
		huiMoney := (v.PayMoney * int64(v.Rate) / int64(model.UNITY)) * int64(guquan.ReturnLuckyRate) / int64(model.UNITY)
		i := response.BuyGuquanList{
			ID:         v.ID,
			Num:        v.PayMoney / int64(model.UNITY),
			Price:      float64(guquan.Price) / model.UNITY,
			CreateTime: v.CreateTime,
			TotalPrice: float64(weiMoney+huiMoney) / model.UNITY,
			Status:     Status,
		}
		res = append(res, i)
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
	guquan := model.Guquan{}
	guquan.Get(true)

	//now := time.Now().Unix()
	//if now >= guquan.ReturnTime {
	//	return nil
	//}

	//获取订单信息
	orderModel := model.OrderGuquan{ID: this.Id, UID: member.ID}
	if !orderModel.Get() {
		return nil
	}

	//获取用户信息
	memberVerfiy := model.MemberVerified{UID: orderModel.UID}
	memberVerfiy.Get()

	//合同起始时
	startDate := time.Unix(int64(guquan.OpenTime), 0).Format("2006年01月02日")
	endDate := time.Unix(int64(guquan.ReturnTime), 0).Format("2006年01月02日")
	days := int(guquan.ReturnTime-guquan.OpenTime) / 86400
	createDate := time.Unix(int64(orderModel.CreateTime), 0).Format("2006年01月02日")

	//中签回购金额
	huiMoney := (orderModel.PayMoney * int64(orderModel.Rate) / int64(model.UNITY)) * int64(guquan.ReturnLuckyRate) / int64(model.UNITY)
	//未中签回购金额
	weiMoney := (orderModel.PayMoney * int64(int(model.UNITY)-orderModel.Rate) / int64(model.UNITY)) * (int64(model.UNITY) + int64(guquan.ReturnRate)) / int64(model.UNITY)

	//原始股权总金额
	sourceAmount := float64(orderModel.PayMoney) / float64(model.UNITY) * float64(guquan.Price) / float64(model.UNITY)
	//中签回购利润
	winProfit := float64(guquan.ReturnLuckyRate)*100/model.UNITY - 100
	//未中签加购利润
	notWinProfit := float64(guquan.ReturnRate) * 100 / model.UNITY

	//总股权数量
	totalQuantity := orderModel.PayMoney / int64(model.UNITY)
	//中签股权数量
	winQuantity := orderModel.PayMoney * int64(orderModel.Rate) / (int64(model.UNITY) * int64(model.UNITY))
	//未中签股权数量
	notWinQuantity := totalQuantity - winQuantity

	return &response.StockCertificateResp{
		ID:         orderModel.ID,
		RealName:   memberVerfiy.RealName,
		IdCardNo:   memberVerfiy.IDNumber,
		StartDate:  startDate,
		EndDate:    endDate,
		CreateDate: createDate,
		Days:       days,

		//股权总数
		Quantity: orderModel.PayMoney / int64(model.UNITY),
		//原订单价格
		Price: float64(guquan.Price) / model.UNITY,
		//原始股权总金额
		TotalAmount: sourceAmount,

		//中签股权数
		WinQuantity: winQuantity,
		//中签回购利润
		WinProfit: winProfit,
		//中签股权回购总金额
		WinRepurchaseAmount: float64(huiMoney) / model.UNITY,

		//未中签股权数
		NotWinQuantity: notWinQuantity,
		//未中签回购利润
		NotWinProfit: notWinProfit,
		//未中签回购金额
		NotWinRepurchaseAmount: float64(weiMoney) / model.UNITY,
	}
}
