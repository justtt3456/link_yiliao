package response

type MemberAddressResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data MemberAddressList `json:"data"`
}
type MemberAddressList struct {
	List []MemberAddress `json:"list"`
}

type MemberAddress struct {
	Id      int    `json:"id"` //
	Name    string `json:"name"`
	Phone   string `json:"phone"`   //预留手机号码
	Address string `json:"address"` //
	Other   string `json:"other"`   //
}
