package request

import "github.com/shopspring/decimal"

type ProductKline struct {
	Code     string `form:"code"`
	Interval int    `form:"interval"` //1 1分钟 2 5分钟 3 15分钟 4 30分钟 5 1小时 6 1天
}
type ProductOption struct {
	Id int `json:"id" form:"id"`
}
type ProductBuy struct {
	Id       int             `json:"id" form:"id"` //产品id
	Interval int             `json:"interval"`     //购买时长 秒
	Ratio    int             `json:"ratio"`        //购买收益率
	Type     int             `json:"type"`         //购买类型 1涨 2跌
	Amount   decimal.Decimal `json:"amount"`       //购买金额
}

type ProductList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Category int    `form:"category"`
	Name     string `form:"name"`
}
type ProductOptional struct {
	Id int `json:"id" form:"id"`
}

type ProductBuyList struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
