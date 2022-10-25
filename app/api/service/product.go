package service

import (
	"errors"
	"finance/app/api/swag/request"
	"finance/app/api/swag/response"
	"finance/common"
	"finance/lang"
	"finance/model"
	"fmt"
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
			Progress:     float64(v.Progress) / model.UNITY,
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
		Progress:     float64(m.Progress) / model.UNITY,
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
	if member.IsReal != 2 {
		return errors.New("请实名认证！")
	}

	if this.Amount <= 0 {
		return errors.New("购买金额必须大于0！")
	}

	if common.Md5String(this.TransferPwd+member.WithdrawSalt) != member.WithdrawPassword {
		return errors.New("交易密码错误")
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

	amount := int64(this.Amount * model.UNITY)
	if amount > member.Balance {
		return errors.New("余额不足,请先充值！")
	}
	switch this.Cate {
	case 1:
		//产品
		p := model.Product{ID: this.Id}
		if !p.Get() {
			return errors.New("产品不存在！")
		}
		if int64(this.Amount*model.UNITY) < p.Price {
			return errors.New(fmt.Sprintf("购买金额必须大于%v！", float64(p.Price)/model.UNITY))
		}

		//购买
		inc := &model.OrderProduct{
			UID:          member.ID,
			Pid:          p.ID,
			PayMoney:     amount,
			IsReturnTop:  1,
			AfterBalance: member.Balance - amount,
			CreateTime:   time.Now().Unix(),
			UpdateTime:   time.Now().Unix(),
		}
		err := inc.Insert()
		if err != nil {
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
			Amount:     amount,
			Before:     member.Balance,
			After:      member.Balance - amount,
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
		member.Balance -= amount
		member.IsBuy = 1

		if isSendRigster {
			//赠送礼金 加入账变记录
			trade2 := model.Trade{
				UID:        member.ID,
				TradeType:  7,
				ItemID:     inc.ID,
				Amount:     int64(config.RegisterSend),
				Before:     member.Balance,
				After:      member.Balance + int64(config.RegisterSend),
				Desc:       "第一次购买赠送礼金",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err = trade2.Insert()
			if err != nil {
				logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
			}
			member.UseBalance += int64(config.RegisterSend)
			member.TotalBalance += int64(config.RegisterSend)
			member.Income += int64(config.RegisterSend)
		}

		//检查是否有满送活动
		if p.IsManjian == 1 {
			full := model.FullDelivery{}
			if full.Find(amount) {
				//满送活动加入账变记录
				trade3 := model.Trade{
					UID:        member.ID,
					TradeType:  9,
					ItemID:     int(full.Coupon.ID),
					Amount:     full.Coupon.Price,
					Before:     0,
					After:      0,
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
				err := MemberCoupon.Insert()
				if err != nil {
					logrus.Errorf("赠赠送优惠券记录失败%v %v", err, member.ID)
				}
			}
		}
		//检查优惠券
		if this.UseId != 0 {
			MemberCoupon := model.MemberCoupon{
				Uid:   int64(member.ID),
				ID:    this.UseId,
				IsUse: 1,
			}
			if MemberCoupon.Get() {
				//使用优惠券记录
				trade3 := model.Trade{
					UID:        member.ID,
					TradeType:  10,
					ItemID:     int(this.UseId),
					Amount:     MemberCoupon.Coupon.Price,
					Before:     member.Balance,
					After:      member.Balance + MemberCoupon.Coupon.Price,
					Desc:       "使用优惠券",
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					IsFrontend: 1,
				}
				err = trade3.Insert()
				if err != nil {
					logrus.Errorf("使用优惠券记录 加入账变记录失败%v", err)
				}
				MemberCoupon.IsUse = 2
				err = MemberCoupon.Update("is_use")
				if err != nil {
					logrus.Errorf("修改用户优惠券失败%v", err)
				}
				member.Balance += MemberCoupon.Coupon.Price
				member.TotalBalance += MemberCoupon.Coupon.Price
			}
		}
		member.WillIncome += amount * int64(p.Dayincome*p.TimeLimit) / int64(model.UNITY)

	case 2:
		//股权
		p := model.Guquan{ID: int64(this.Id)}
		if !p.Get(true) {
			return errors.New("股权不存在！")
		}
		if p.PreStartTime > time.Now().Unix() {
			return errors.New("股权预售时间未开始")
		}
		if p.PreEndTime < time.Now().Unix() {
			return errors.New("股权预售时间已结束")
		}
		if int64(this.Amount) < p.LimitBuy {
			return errors.New(fmt.Sprintf("购买金额必须大于%v！", p.LimitBuy))
		}

		//购买
		inc := &model.OrderGuquan{
			UID:          member.ID,
			Pid:          int(p.ID),
			PayMoney:     amount,
			Rate:         int(model.UNITY),
			AfterBalance: member.Balance - amount,
			CreateTime:   time.Now().Unix(),
			UpdateTime:   time.Now().Unix(),
		}
		err := inc.Insert()
		if err != nil {
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
			Before:     member.Balance,
			After:      member.Balance - amount,
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
		member.Balance -= amount
		member.IsBuy = 1
		member.Guquan += amount / int64(model.UNITY)

		if isSendRigster {
			//赠送礼金 加入账变记录
			trade2 := model.Trade{
				UID:        member.ID,
				TradeType:  7,
				ItemID:     inc.ID,
				Amount:     int64(config.RegisterSend),
				Before:     member.Balance,
				After:      member.Balance + int64(config.RegisterSend),
				Desc:       "第一次购买赠送礼金",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				IsFrontend: 1,
			}
			err = trade2.Insert()
			if err != nil {
				logrus.Errorf("赠送礼金 加入账变记录失败%v", err)
			}
			member.UseBalance += int64(config.RegisterSend)
			member.TotalBalance += int64(config.RegisterSend)
			member.Income += int64(config.RegisterSend)
		}

	default:
		return errors.New("购买类型不存在")
	}

	return member.Update("balance", "total_balance", "use_balance", "is_buy", "income", "wll_income", "guquan")
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
		items = append(items, response.BuyList{
			Name:    list[i].Product.Name,
			BuyTime: int(list[i].CreateTime),
			Amount:  float64(list[i].PayMoney) / model.UNITY,
			Status:  1,
		})
	}
	res.List = items
	res.Page = FormatPage(page)
	return &res
}

type BuyGuquanList struct {
	request.Request
}

func (this BuyGuquanList) List(member *model.Member) *response.BuyGuquanResp {
	var res response.BuyGuquanResp
	m := model.OrderGuquan{UID: member.ID}
	m.Get()
	guquan := model.Guquan{}
	guquan.Get(true)
	money, err := m.Sum()
	if err != nil || money == 0 {
		return nil
	}
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
