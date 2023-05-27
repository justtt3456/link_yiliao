package request

import "github.com/shopspring/decimal"

type AddCoupon struct {
	Price decimal.Decimal `json:"price"` //面额
}

type AddActive struct {
	Amount   decimal.Decimal `json:"amount"`    //满多少
	CouponId int             `json:"coupon_id"` //送的优惠券Id
}

type DelActive struct {
	Id int `json:"id"` //活动Id
}
