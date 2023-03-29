package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type ProductList struct {
	request.ProductList
}

func (this ProductList) PageList() response.ProductData {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Product{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.Product, 0)
	for _, v := range list {
		i := response.Product{
			ID:           v.ID,
			Name:         v.Name,
			Category:     v.Category,
			CategoryName: v.ProductCategory.Name,
			CreateTime:   v.CreateTime,
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
			IsFinish:     v.IsFinish,
			IsManjian:    v.IsManjian,
			BuyTimeLimit: v.BuyTimeLimit,
			DelayTime:    v.DelayTime,
			Type:         v.Type,
			Progress:     float64(v.Progress) / model.UNITY,
			GiftId:       v.GiftId,
		}
		res = append(res, i)
	}
	return response.ProductData{List: res, Page: FormatPage(page)}
}
func (this ProductList) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where[model.Product{}.TableName()+".name"] = this.Name
	}
	if this.Category > 0 {
		where[model.Product{}.TableName()+".category"] = this.Category
	}
	if this.Status > 0 {
		where[model.Product{}.TableName()+".status"] = this.Status
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}

type ProductCreate struct {
	request.ProductCreate
}

func (this ProductCreate) Create() error {
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Category == 0 {
		return errors.New("产品类别不能为空")
	}

	if this.TimeLimit == 0 {
		return errors.New("产品投资期限不能为空")
	}
	if this.Dayincome == 0 {
		return errors.New("每日收益不能为空")
	}
	if this.Price == 0 {
		return errors.New("最低买入不能为空")
	}
	if this.TotalPrice == 0 {
		return errors.New("项目规模不能为空")
	}
	if this.OtherPrice == 0 {
		return errors.New("可投余额不能为空")
	}

	if this.MoreBuy == 0 {
		this.MoreBuy = 99999
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.IsRecommend == 0 {
		return errors.New("是否推荐到首页不能为空")
	}

	if this.IsFinish == 0 {
		return errors.New("是否投资完毕不能为空")
	}

	if this.IsManjian == 0 {
		return errors.New("是否有满减活动不能为空")
	}
	if this.BuyTimeLimit == 0 {
		return errors.New("产品限时多少天不能为空")
	}
	if this.Type == 0 {
		return errors.New("延期类型不能为空")
	}
	if this.Type == 2 {
		if this.DelayTime <= 0 {
			return errors.New("延期时间必须大于0")
		}
	}
	//赠送产品ID分析
	if this.Type == 5 {
		this.GiftId = 0
	}
	if this.GiftId > 0 {
		giftModel := model.Product{
			ID:     this.GiftId,
			Type:   5,
			Status: 1,
		}
		if !giftModel.Get() {
			return errors.New("赠品不存在")
		}
	}

	ex := model.Product{
		Name: this.Name,
	}
	if ex.Get() {
		return errors.New("产品已存在")
	}

	//高精度浮点计算
	unity := decimal.NewFromFloat(model.UNITY)
	//年利率计算
	dayIncome := decimal.NewFromFloat(this.Dayincome)
	//单价
	price := decimal.NewFromFloat(this.Price)
	//总金额
	totalPrice := decimal.NewFromFloat(this.TotalPrice)
	//可投余额
	otherPrice := decimal.NewFromFloat(this.OtherPrice)

	m := model.Product{
		Name:         this.Name,
		Category:     this.Category,
		CreateTime:   this.CreateTime,
		Status:       this.Status,
		Tag:          this.Tag,
		TimeLimit:    this.TimeLimit,
		IsRecommend:  this.IsRecommend,
		Dayincome:    int(dayIncome.Mul(unity).IntPart()),
		Price:        price.Mul(unity).IntPart(),
		TotalPrice:   totalPrice.Mul(unity).IntPart(),
		OtherPrice:   otherPrice.Mul(unity).IntPart(),
		MoreBuy:      this.MoreBuy,
		Desc:         this.Desc,
		IsFinish:     this.IsFinish,
		IsManjian:    this.IsManjian,
		BuyTimeLimit: this.BuyTimeLimit,
		Type:         this.Type,
		DelayTime:    this.DelayTime,
		Progress:     int(this.Progress * model.UNITY),
		GiftId:       this.GiftId,
	}
	return m.Insert()
}

type ProductUpdate struct {
	request.ProductUpdate
}

func (this ProductUpdate) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Category == 0 {
		return errors.New("产品类别不能为空")
	}

	if this.TimeLimit == 0 {
		return errors.New("产品投资期限不能为空")
	}
	if this.Dayincome == 0 {
		return errors.New("每日收益不能为空")
	}
	if this.Price == 0 {
		return errors.New("最低买入不能为空")
	}
	if this.TotalPrice == 0 {
		return errors.New("项目规模不能为空")
	}
	if this.OtherPrice < 0 {
		return errors.New("可投余额不能小于0")
	}
	if this.MoreBuy == 0 {
		this.MoreBuy = 99999
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.IsRecommend == 0 {
		return errors.New("是否推荐到首页不能为空")
	}

	if this.IsFinish == 0 {
		return errors.New("是否投资完毕不能为空")
	}

	if this.IsManjian == 0 {
		return errors.New("是否有满减活动不能为空")
	}
	if this.BuyTimeLimit == 0 {
		return errors.New("产品限时多少天不能为空")
	}
	if this.Type == 0 {
		return errors.New("延期类型不能为空")
	}
	if this.Type == 2 {
		if this.DelayTime <= 0 {
			return errors.New("延期时间必须大于0")
		}
	}
	//赠送产品ID分析
	if this.Type == 5 {
		this.GiftId = 0
	}
	if this.GiftId > 0 {
		giftModel := model.Product{
			ID:     this.GiftId,
			Type:   5,
			Status: 1,
		}
		if !giftModel.Get() {
			return errors.New("赠品不存在")
		}
	}

	m := model.Product{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}

	//高精度浮点计算
	unity := decimal.NewFromFloat(model.UNITY)
	//年利率计算
	dayIncome := decimal.NewFromFloat(this.Dayincome)
	//单价
	price := decimal.NewFromFloat(this.Price)
	//总金额
	totalPrice := decimal.NewFromFloat(this.TotalPrice)
	//可投余额
	otherPrice := decimal.NewFromFloat(this.OtherPrice)

	m.Name = this.Name
	m.Category = this.Category
	m.Status = this.Status
	m.TimeLimit = this.TimeLimit
	m.Dayincome = int(dayIncome.Mul(unity).IntPart())
	m.Price = price.Mul(unity).IntPart()
	m.TotalPrice = totalPrice.Mul(unity).IntPart()
	m.OtherPrice = otherPrice.Mul(unity).IntPart()
	m.MoreBuy = this.MoreBuy
	m.Desc = this.Desc
	m.Status = this.Status
	m.IsRecommend = this.IsRecommend
	m.IsFinish = this.IsFinish
	m.IsManjian = this.IsManjian
	m.BuyTimeLimit = this.BuyTimeLimit
	m.Tag = this.Tag
	m.Type = this.Type
	m.DelayTime = this.DelayTime
	m.Progress = int(this.Progress * model.UNITY)
	m.GiftId = this.GiftId

	return m.Update("name", "progress", "buy_time_limit", "category", "create_time", "status", "tag", "time_limit", "is_recommend", "day_income", "price", "total_price", "other_price", "more_buy", "desc", "is_finish", "is_manjian", "gift_id")
}

type ProductUpdateStatus struct {
	request.ProductUpdateStatus
}

func (this ProductUpdateStatus) UpdateStatus() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		ID: this.ID,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}
	m.Status = this.Status
	return m.Update("status")
}

type ProductRemove struct {
	request.ProductRemove
}

func (this ProductRemove) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		ID: this.ID,
	}
	return m.Remove()
}

type GiftProductOptions struct {
	request.GiftProductOptions
}

func (this GiftProductOptions) GiftList() response.ProductGiftOptions {
	productModel := model.Product{}
	//获取列表内容
	list := productModel.GiftList()
	res := make([]response.ProductGiftInfo, 0)
	for _, value := range list {
		info := response.ProductGiftInfo{
			ID:   value.ID,
			Name: value.Name,
		}
		res = append(res, info)
	}
	//返回内容
	return response.ProductGiftOptions{List: res}
}
