package response

type MemberOptionalResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data MemberOptionalData `json:"data"`
}
type MemberOptionalData struct {
	List []MemberOptional `json:"list"`
}
type MemberOptional struct {
	ID             int     `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Wave           float64 `json:"wave"`
	OpenTime       string  `json:"open_time"`
	WeekendTrading int     `json:"weekend_trading"`
	Open           float64 `json:"open"`
	Price          float64 `json:"price"`
	Low            float64 `json:"low"`
	High           float64 `json:"high"`
	Change         float64 `json:"change"`
	Vol            float64 `json:"vol"`
}
