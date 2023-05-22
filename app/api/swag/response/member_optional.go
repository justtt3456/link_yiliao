package response

import "github.com/shopspring/decimal"

type MemberOptionalResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data MemberOptionalData `json:"data"`
}
type MemberOptionalData struct {
	List []MemberOptional `json:"list"`
}
type MemberOptional struct {
	Id             int             `json:"id"`
	Code           string          `json:"code"`
	Name           string          `json:"name"`
	Wave           decimal.Decimal `json:"wave"`
	OpenTime       string          `json:"open_time"`
	WeekendTrading int             `json:"weekend_trading"`
	Open           decimal.Decimal `json:"open"`
	Price          decimal.Decimal `json:"price"`
	Low            decimal.Decimal `json:"low"`
	High           decimal.Decimal `json:"high"`
	Change         decimal.Decimal `json:"change"`
	Vol            decimal.Decimal `json:"vol"`
}
