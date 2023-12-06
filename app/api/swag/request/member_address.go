package request

type MemberAddressCreate struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`   //预留手机号码
	Address string `json:"address"` //
	Other   string `json:"other"`   //
}
type MemberAddressUpdate struct {
	Id      int    `json:"id"` //
	Name    string `json:"name"`
	Phone   string `json:"phone"`   //预留手机号码
	Address string `json:"address"` //
	Other   string `json:"other"`   //
}
type MemberAddressRemove struct {
	Id int `json:"id"`
}
