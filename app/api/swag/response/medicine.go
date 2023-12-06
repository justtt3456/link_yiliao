package response

import "github.com/shopspring/decimal"

type MedicineResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data MedicineData `json:"data"`
}
type MedicineData struct {
	Category []MedicineCategory `json:"category"`
}
type MedicineCategory struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Medicine []Medicine `json:"Medicine"`
}
type Medicine struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`               //产品名称
	Price             decimal.Decimal `json:"price"`              //价格
	Img               string          `json:"img"`                //图片
	Desc              string          `json:"desc"`               //描述
	WithdrawThreshold decimal.Decimal `json:"withdraw_threshold"` //提现额度比例
	Interval          int             `json:"interval"`           //投资期限 （天）

}

type MedicineListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data MedicineListData `json:"data"`
}
type MedicineListData struct {
	List []Medicine `json:"list"`
	Page Page       `json:"page"`
}

type MedicineCategoryResponse struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg"`
	Data MedicineCategoryData `json:"data"`
}
type MedicineCategoryData struct {
	List []MedicineCategoryItem `json:"list"`
}
type MedicineCategoryItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MedicineBuyList struct {
	Name       string          `json:"name"`        //产品名字
	CreateTime int64           `json:"create_time"` //投资时间
	Amount     decimal.Decimal `json:"amount"`      //金额
	Address    string          `json:"address"`
}

type MedicineBuyListResp struct {
	List []MedicineBuyList `json:"list"`
	Page Page              `json:"page"`
}
