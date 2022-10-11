package request

type BankList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
}
type BankCreate struct {
	BankName string `json:"bank_name"` //银行名称
	Sort     int    `json:"sort"`
	Status   int    `json:"status"` //状态
	Code     string `json:"code"`
	Lang     string `json:"lang"`
}
type BankUpdate struct {
	ID       int    `json:"id"`
	BankName string `json:"bank_name"` //银行名称
	Sort     int    `json:"sort"`
	Status   int    `json:"status"` //状态
	Code     string `json:"code"`
	Lang     string `json:"lang"`
}
type BankUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"` //状态
}
type BankRemove struct {
	ID int `json:"id"`
}
