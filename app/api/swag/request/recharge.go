package request

import "github.com/shopspring/decimal"

type RechargeCreate struct {
	Method  int             `json:"method"`  //充值方式id
	Amount  decimal.Decimal `json:"amount"`  //充值金额
	From    string          `json:"from"`    //付款账号
	To      int             `json:"to"`      //收款账号Id
	Voucher string          `json:"voucher"` //凭证图
	//ChannelId int             `json:"channel_id"` //三方通道id
}
type RechargeList struct {
	Status   int `form:"status"` //状态 1审核中 2通过 3驳回
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type RechargeMethodInfo struct {
	Code string `json:"code" form:"code"`
}
