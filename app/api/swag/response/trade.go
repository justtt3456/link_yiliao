package response

type TradeResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data TradeList `json:"data"`
}
type TradeList struct {
	List []Trade `json:"list"`
	Page Page    `json:"page"`
}
type Trade struct {
	ID         int     `json:"id"`         //id
	TradeType  int     `json:"type_trade"` //账单类型 '账单类型 1=购买餐品  2=购买股权 3=充值 4=提现 5=可用转可提 6=可提转可用 7=注册买产品礼金 8=注册实名认证礼金 9=送优惠券 10=使用优惠券 11=余额宝转入 12=余额宝转出  13=余额宝收益 14=后台上分 15=后台下分 16=每日收益 17=股权收益 18=一级返佣 19=二级返佣 20=三级返佣 21=团队收益'
	Amount     float64 `json:"amount"`     //金额
	Before     float64 `json:"before"`     //账变前余额
	After      float64 `json:"after"`      //账变后余额
	CreateTime int64   `json:"create_time"`
	Desc       string  `json:"desc"` //描述
}
