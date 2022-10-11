package request

type TradeRequest struct {
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	Username  string `json:"username" form:"username"`
	TradeType int    `json:"trade_type" form:"trade_type"` //账单类型 1=购买餐品  2=购买股权 3=充值 4=提现 5=可用转可提 6=可提转可用 7=注册买产品礼金 8=注册实名认证礼金 9=送优惠券 10=使用优惠券 11=余额宝转入 12=余额宝转出  13=余额宝收益 14=后台上分 15=后台下分
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}
