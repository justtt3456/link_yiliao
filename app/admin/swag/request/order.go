package request

import "github.com/shopspring/decimal"

type OrderListRequest struct {
	ProductName string `form:"product_name" json:"product_name"` //产品名字
	Username    string `form:"username" json:"username"`         //用户名字
	Uid         int    `form:"uid" json:"uid"`                   //用户Id
	StartTime   string `form:"start_time" json:"start_time"`
	EndTime     string `form:"end_time" json:"end_time"`
	Page        int    `form:"page" json:"page"`
	PageSize    int    `form:"page_size" json:"page_size"`
}

type OrderUpdate struct {
	Id   int             `json:"id"`
	Rate decimal.Decimal `json:"rate"` //中签率
}

type OrderCommission struct {
	Hash      string `form:"hash"`       //哈唏
	StartTime int64  `form:"start_time"` //开始时间
	EndTime   int64  `form:"end_time"`   //结束时间
}
