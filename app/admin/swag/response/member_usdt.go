package response

type MemberUsdtListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data MemberUsdtList `json:"data"`
}
type MemberUsdtList struct {
	List []MemberUsdt `json:"list"`
}

type MemberUsdt struct {
	Id       int    `json:"id"` //
	Address  string `json:"address"`
	Protocol string `json:"protocol"`
}
