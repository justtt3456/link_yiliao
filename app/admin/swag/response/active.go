package response

import "github.com/shopspring/decimal"

type Coupon struct {
	Id    int64           `json:"id"`    //id
	Price decimal.Decimal `json:"price"` //面额
}

type CouponResp struct {
	List []Coupon `json:"list"`
}

type Active struct {
	Id       int64           `json:"id"`        //活动Id
	Price    decimal.Decimal `json:"price"`     //送多少
	Amout    decimal.Decimal `json:"amout"`     //满多少
	CouponId int64           `json:"coupon_id"` //优惠券Id
}

type ActiveResp struct {
	List []Active `json:"list"`
}
