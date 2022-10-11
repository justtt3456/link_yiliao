package response

type Coupon struct {
	Id    int64   `json:"id"`    //id
	Price float64 `json:"price"` //面额
}

type CouponResp struct {
	List []Coupon `json:"list"`
}

type Active struct {
	Id       int64   `json:"id"`        //活动ID
	Price    float64 `json:"price"`     //送多少
	Amout    float64 `json:"amout"`     //满多少
	CouponId int64   `json:"coupon_id"` //优惠券ID
}

type ActiveResp struct {
	List []Active `json:"list"`
}
