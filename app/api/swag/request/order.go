package request

type Order struct {
	Status   int `form:"status"` //状态 1未结算 2已结算
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type BuyReq struct {
	Cate        int     `json:"cate"`         //1=买产品  2=买股权
	Id          int     `json:"id"`           //产品ID
	Amount      float64 `json:"amount"`       //金额
	UseId       int64   `json:"use_id"`       //使用优惠券传的id
	TransferPwd string  `json:"transfer_pwd"` //交易密码
}