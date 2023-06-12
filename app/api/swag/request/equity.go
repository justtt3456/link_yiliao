package request

type EquityBuyRequest struct {
	Id          int    `json:"id"`           //产品ID
	Quantity    int    `json:"quantity"`     //数量
	TransferPwd string `json:"transfer_pwd"` //交易密码
}
