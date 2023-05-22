package request

import "github.com/shopspring/decimal"

type AddCoupon struct {
	Price decimal.Decimal `json:"price"` //面额
}

type AddActive struct {
	Amout    decimal.Decimal `json:"amout"`     //满多少
	CouponId int64           `json:"coupon_id"` //送的优惠券Id
}

type DelActive struct {
	Id int64 `json:"id"` //活动Id
}
