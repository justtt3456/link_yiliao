package response

import "github.com/shopspring/decimal"

type PayChannelListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data PayChannelData `json:"data"`
}
type PayChannelData struct {
	List []PayChannel `json:"list"`
	Page Page         `json:"page"`
}
type PayChannel struct {
	Id          int             `json:"id"`           //
	Name        string          `json:"name"`         //支付方式名称
	PaymentId   int             `json:"payment_id"`   //第三方id
	PaymentName string          `json:"payment_name"` //第三方名称
	Code        string          `json:"code"`         //支付编码
	Min         decimal.Decimal `json:"min"`          //最小值
	Max         decimal.Decimal `json:"max"`          //最大值
	Status      int             `json:"status"`       //状态
	Category    int             `json:"category"`     //分类(所属支付方式)
	Sort        int             `json:"sort"`         //排序值
	Icon        string          `json:"icon"`         //图标
	Fee         int             `json:"fee"`          //手续费
	Lang        string          `json:"lang"`
	CreateTime  int64           `json:"create_time"`
	UpdateTime  int64           `json:"update_time"`
}
