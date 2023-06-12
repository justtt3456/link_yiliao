package response

import "github.com/shopspring/decimal"

type Coupon struct {
	Id    int             `json:"id"`    //id
	Price decimal.Decimal `json:"price"` //面额
}

type CouponResp struct {
	List []Coupon `json:"list"`
}

type Active struct {
	Id       int             `json:"id"`        //活动Id
	Price    decimal.Decimal `json:"price"`     //送多少
	Amount   decimal.Decimal `json:"amount"`    //满多少
	CouponId int             `json:"coupon_id"` //优惠券Id
}

type ActiveResp struct {
	List []Active `json:"list"`
}
