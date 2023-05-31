package request

type Trade struct {
	Type      int    `form:"type"` //账单类型 1=购买产品  2=购买股权 3=充值 4=提现 5=可用转可提 6=可提转可用 7=注册买产品礼金 8=注册实名认证礼金 9=送优惠券 10=使用优惠券 11=余额宝转入 12=余额宝转出  13=余额宝收益
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}
