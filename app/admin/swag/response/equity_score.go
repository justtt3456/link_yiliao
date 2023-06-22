package response

import "github.com/shopspring/decimal"

type EquityScorePageListResponse struct {
	List []EquityScoreOrder `json:"list"`
	Page Page               `json:"page"`
}
type EquityScoreOrder struct {
	Id         int             `json:"id;"` //
	UId        int             `json:"uid"` //关联用户id
	Username   string          `json:"username"`
	PayMoney   decimal.Decimal `json:"pay_money"` //购买付款金额 =手数
	Rate       decimal.Decimal `json:"rate"`      //
	Interval   int             `json:"interval"`
	Status     int             `json:"status"`      //状态
	CreateTime int64           `json:"create_time"` //创建时间
	EndTime    int64           `json:"end_time"`
}
