package response

import "github.com/shopspring/decimal"

type TradeResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data TradeData `json:"data"`
}
type TradeData struct {
	List []TradeInfo `json:"list"`
	Page Page        `json:"page"`
}
type TradeInfo struct {
	Tid          int             `json:"tid"`           //
	Username     string          `json:"username"`      //
	TradeType    int             `json:"trade_type"`    // 账单类型 1=购买产品  2=购买股权 3=充值 4=提现 5=可用转可提 6=可提转可用 7=注册买产品礼金 8=注册实名认证礼金 9=送优惠券 10=使用优惠券 11=余额宝转入 12=余额宝转出  13=余额宝收益 14=后台上分 15=后台下分
	Amount       decimal.Decimal `json:"amount"`        //
	BeforeAmount decimal.Decimal `json:"before_amount"` //
	AfterAmount  decimal.Decimal `json:"after_amount"`  //
	Desc         string          `json:"desc"`          //
	CreateTime   int64           `json:"create_time"`   //修改时间
}
