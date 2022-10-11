package request

type AddCoupon struct {
	Price float64 `json:"price"` //面额
}

type AddActive struct {
	Amout    float64 `json:"amout"`     //满多少
	CouponId int64   `json:"coupon_id"` //送的优惠券ID
}

type DelActive struct {
	Id int64 `json:"id"` //活动ID
}
