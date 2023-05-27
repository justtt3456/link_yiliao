package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/model"
	"errors"
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
			Id:                    v.Id,
			Name:                  v.Name,
			Category:              v.Category,
			CategoryName:          v.ProductCategory.Name,
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
			GiftId:                v.GiftId,
			WithdrawThresholdRate: v.WithdrawThresholdRate,
			IsHot:                 v.IsHot,
			IsFinished:            v.IsFinished,
			IsCouponGift:          v.IsCouponGift,
			Sort:                  v.Sort,
			Status:                v.Status,
			CreateTime:            v.CreateTime,
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
	if this.Type == 0 {
		return errors.New("产品类型不能为空")
	}
	if this.Interval == 0 {
		return errors.New("投资期限不能为空")
	}
	if this.Img == "" {
		return errors.New("图片不能为空")
	}
	if this.IncomeRate.LessThanOrEqual(decimal.Zero) {
		return errors.New("收益率不能为空")
	}
	if this.Price.LessThanOrEqual(this.Price) {
		return errors.New("价格不能为空")
	}
	if this.Total.LessThanOrEqual(decimal.Zero) {
		return errors.New("项目规模不能为空")
	}
	if this.LimitBuy == 0 {
		return errors.New("限购数量不能为空")
	}
	if this.WithdrawThresholdRate.LessThanOrEqual(decimal.Zero) {
		return errors.New("提现额度比例不能为空")
	}
	if this.IsHot == 0 {
		return errors.New("是否热门不能为空")
	}
	if this.IsFinished == 0 {
		return errors.New("是否投满不能为空")
	}
	if this.IsCouponGift == 0 {
		return errors.New("是否有优惠券活动不能为空")
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.Type == 2 {
		if this.DelayTime <= 0 {
			return errors.New("延期时间必须大于0")
		}
	}
	//赠送产品Id分析
	if this.Type == 5 {
		this.GiftId = 0
	}
	if this.GiftId > 0 {
		giftModel := model.Product{
			Id:     this.GiftId,
			Type:   5,
			Status: 1,
		}
		if !giftModel.Get() {
			return errors.New("赠品不存在")
		}
	}
	m := model.Product{
		Name:                  this.Name,
		Category:              this.Category,
		Type:                  this.Type,
		Price:                 this.Price,
		Interval:              this.Interval,
		Img:                   this.Img,
		IncomeRate:            this.IncomeRate,
		LimitBuy:              this.LimitBuy,
		Total:                 this.Total,
		Current:               this.Current,
		Desc:                  this.Desc,
		DelayTime:             this.DelayTime,
		GiftId:                this.GiftId,
		WithdrawThresholdRate: this.WithdrawThresholdRate,
		IsHot:                 this.IsHot,
		IsFinished:            this.IsFinished,
		IsCouponGift:          this.IsCouponGift,
		Sort:                  this.Sort,
		Status:                this.Status,
	}
	return m.Insert()
}

type ProductUpdate struct {
	request.ProductUpdate
}

func (this ProductUpdate) Update() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	if this.Name == "" {
		return errors.New("产品名称不能为空")
	}
	if this.Category == 0 {
		return errors.New("产品类别不能为空")
	}
	if this.Type == 0 {
		return errors.New("产品类型不能为空")
	}
	if this.Interval == 0 {
		return errors.New("投资期限不能为空")
	}
	if this.Img == "" {
		return errors.New("图片不能为空")
	}
	if this.IncomeRate.LessThanOrEqual(decimal.Zero) {
		return errors.New("收益率不能为空")
	}
	if this.Price.LessThanOrEqual(this.Price) {
		return errors.New("价格不能为空")
	}
	if this.Total.LessThanOrEqual(decimal.Zero) {
		return errors.New("项目规模不能为空")
	}
	if this.LimitBuy == 0 {
		return errors.New("限购数量不能为空")
	}
	if this.WithdrawThresholdRate.LessThanOrEqual(decimal.Zero) {
		return errors.New("提现额度比例不能为空")
	}
	if this.IsHot == 0 {
		return errors.New("是否热门不能为空")
	}
	if this.IsFinished == 0 {
		return errors.New("是否投满不能为空")
	}
	if this.IsCouponGift == 0 {
		return errors.New("是否有优惠券活动不能为空")
	}
	if this.Desc == "" {
		return errors.New("描述不能为空")
	}
	if this.Status == 0 {
		return errors.New("状态不能为空")
	}
	if this.Type == 2 {
		if this.DelayTime <= 0 {
			return errors.New("延期时间必须大于0")
		}
	}
	//赠送产品Id分析
	if this.Type == 5 {
		this.GiftId = 0
	}
	if this.GiftId > 0 {
		giftModel := model.Product{
			Id:     this.GiftId,
			Type:   5,
			Status: 1,
		}
		if !giftModel.Get() {
			return errors.New("赠品不存在")
		}
	}
	m := model.Product{
		Id: this.Id,
	}
	if !m.Get() {
		return errors.New("产品不存在")
	}
	m.Name = this.Name
	m.Category = this.Category
	m.Type = this.Type
	m.Price = this.Price
	m.Interval = this.Interval
	m.Img = this.Img
	m.IncomeRate = this.IncomeRate
	m.LimitBuy = this.LimitBuy
	m.Total = this.Total
	m.Current = this.Current
	m.Desc = this.Desc
	m.DelayTime = this.DelayTime
	m.GiftId = this.GiftId
	m.WithdrawThresholdRate = this.WithdrawThresholdRate
	m.IsHot = this.IsHot
	m.IsFinished = this.IsFinished
	m.IsCouponGift = this.IsCouponGift
	m.Sort = this.Sort
	m.Status = this.Status
	return m.Update("name", "category", "type", "price", "interval", "img", "income_rate", "limit_buy", "total", "current", "desc", "delay_time", "gift_id", "withdraw_threshold_rate", "is_hot", "is_finished", "is_coupon_gift", "sort", "status")
}

type ProductUpdateStatus struct {
	request.ProductUpdateStatus
}

func (this ProductUpdateStatus) UpdateStatus() error {
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		Id: this.Id,
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
	if this.Id == 0 {
		return errors.New("参数错误")
	}
	m := model.Product{
		Id: this.Id,
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
			Id:   value.Id,
			Name: value.Name,
		}
		res = append(res, info)
	}
	//返回内容
	return response.ProductGiftOptions{List: res}
}
