package response

type BankResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []Bank `json:"list"`
	}
}
type Bank struct {
	Id       int    `json:"id"`        //
	BankName string `json:"bank_name"` //银行名称
}
