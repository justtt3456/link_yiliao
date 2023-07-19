package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/logic"
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
		giftName := ""
		if v.GiftId > 0 {
			giftModel := model.Product{
				Id: v.GiftId,
			}
			if giftModel.Get() {
				giftName = giftModel.Name
			}
		}

		i := response.Product{
			Id:                    v.Id,
			Name:                  v.Name,
			Category:              v.Category,
			Type:                  v.Type,
			Price:                 v.Price,
			Img:                   v.Img,
			Interval:              v.Interval,
			IncomeRate:            v.IncomeRate,
			LimitBuy:              v.LimitBuy,
			Total:                 v.Total,
			Current:               v.Current,
			Desc:                  v.Desc,
			DelayTime:             v.DelayTime,
			GiftName:              giftName,
			WithdrawThresholdRate: v.WithdrawThresholdRate,
			IsHot:                 v.IsHot,
			IsFinished:            v.IsFinished,
			IsCouponGift:          v.IsCouponGift,
			Status:                v.Status,
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
	if member.Balance.LessThan(amount) {
		return errors.New("余额不足,请先充值！")
	}
	//交易密码验证
	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
	}
	//限购
	order := model.OrderProduct{}
	count := order.Sum("uid = ? and pid = ?", []interface{}{member.Id, this.Id}, "quantity")
	if int(count)+this.Quantity > product.LimitBuy {
		return errors.New(fmt.Sprintf("限购%v份！", product.LimitBuy))
	}
	//优惠券分析
	var memberCoupon model.MemberCoupon
	if this.UseId != 0 {
		memberCoupon = model.MemberCoupon{
			Uid:   member.Id,
			Id:    int(this.UseId),
			IsUse: 1,
		}
		if !memberCoupon.Get() {
			return errors.New("没有找到可用的优惠券信息！")
		}
		//amount = amount.Sub(memberCoupon.Coupon.Price)
	}
	//24小时内购买赠送可提现余额
	var isFirst bool
	lastOrder := model.OrderProduct{UId: member.Id}
	if !lastOrder.Get() && time.Now().Unix() < member.RegTime+86400 {
		isFirst = true
	}
	buyLogic := logic.NewProductBuyLogic()
	err := buyLogic.ProductBuy(this.Cate, member, product, amount, this.Quantity, memberCoupon, isFirst)
	if err != nil {
		return err
	}
	return nil
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
	for _, v := range list {
		//订单状态
		orderStatus := 1
		if v.IsReturnCapital == 1 {
			orderStatus = 2
		}
		product := model.Product{Id: m.Pid}
		if !product.Get() {
			continue
		}
		//每日收益
		income := v.PayMoney.Mul(v.IncomeRate).Div(decimal.NewFromInt(100)).Round(2)
		gift := 2
		if v.Product.Type == 4 { //4为赠品
			gift = 1
		}
		items = append(items, response.BuyList{
			Name:     v.Product.Name,
			BuyTime:  int(v.CreateTime),
			Amount:   v.PayMoney,
			Status:   orderStatus,
			Income:   income,
			EndTime:  v.EndTime,
			Interval: v.Product.Interval,
			IsGift:   gift,
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
	equity := model.Equity{}
	equity.Get(true)

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
		i := response.BuyGuquanList{
			Id:         v.Id,
			Num:        v.Quantity,
			Price:      v.Price,
			CreateTime: v.CreateTime,
			Status:     v.Status,
		}
		switch v.Status {
		case 3:
			//未中签回购金额
			missIncome := v.PayMoney.Mul(equity.MissRate).Div(decimal.NewFromInt(100))
			i.TotalPrice = v.PayMoney.Add(missIncome)
		case 2, 4:
			//中签回购金额
			hitIncome := v.PayMoney.Mul(equity.SellRate).Div(decimal.NewFromInt(100)).Round(2)
			i.TotalPrice = v.PayMoney.Add(hitIncome)
		default:
			i.TotalPrice = v.PayMoney
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
	guquan := model.Equity{}
	guquan.Get(true)

	now := time.Now().Unix()
	if now < guquan.OpenTime {
		return nil
	}

	//获取订单信息
	orderModel := model.OrderEquity{Id: this.Id, UId: member.Id}
	if !orderModel.Get() {
		return nil
	}
	if orderModel.Status == model.StatusReview || orderModel.Status == model.StatusRollback {
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
	//中签回购金额
	huiMoney := orderModel.PayMoney.Mul(guquan.SellRate).Div(decimal.NewFromInt(100)).Round(2).Add(orderModel.PayMoney)
	//未中签回购金额
	weiMoney := orderModel.PayMoney.Mul(guquan.MissRate).Div(decimal.NewFromInt(100)).Round(2).Add(orderModel.PayMoney)

	//原始股权总金额
	sourceAmount := orderModel.PayMoney.Mul(guquan.Price)
	//中签回购利润
	winProfit := guquan.SellRate
	//未中签加购利润
	notWinProfit := guquan.MissRate

	//总股权数量
	totalQuantity := orderModel.PayMoney.Div(guquan.Price).IntPart()
	//中签股权数量
	winQuantity := orderModel.PayMoney.Div(guquan.Price).IntPart()
	//未中签股权数量
	notWinQuantity := totalQuantity
	//
	return &response.StockCertificateResp{
		Id:         orderModel.Id,
		RealName:   memberVerfiy.RealName,
		IdCardNo:   memberVerfiy.IdNumber,
		StartDate:  startDate,
		EndDate:    endDate,
		CreateDate: createDate,
		Days:       days,

		//股权总数
		Quantity: orderModel.PayMoney.Div(guquan.Price).IntPart(),
		//原订单价格
		Price: guquan.Price,
		//原始股权总金额
		TotalAmount: sourceAmount,

		//中签股权数
		WinQuantity: winQuantity,
		//中签回购利润
		WinProfit: winProfit,
		//中签股权回购总金额
		WinRepurchaseAmount: huiMoney,

		//未中签股权数
		NotWinQuantity: notWinQuantity,
		//未中签回购利润
		NotWinProfit: notWinProfit,
		//未中签回购金额
		NotWinRepurchaseAmount: weiMoney,
	}
}
