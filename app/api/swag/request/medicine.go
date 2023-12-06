package request

type MedicineList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Name     string `form:"name"`
}
type MedicineOrder struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type MedicineBuyReq struct {
	Id          int    `json:"id"` //产品Id
	Quantity    int    `json:"quantity"`
	AddressId   int    `json:"address_id"`
	TransferPwd string `json:"transfer_pwd"` //交易密码
}
