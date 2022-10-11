package response

type BankListResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data BannerData `json:"data"`
}
type BankData struct {
	List []BankInfo `json:"list"`
	Page Page       `json:"page"`
}
type BankInfo struct {
	ID         int    `json:"id"`        //
	BankName   string `json:"bank_name"` //银行名称
	Sort       int    `json:"sort"`
	Status     int    `json:"status"` //状态
	Lang       string `json:"lang"`
	Code       string `json:"code"`
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
