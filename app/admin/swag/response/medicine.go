package response

import "github.com/shopspring/decimal"

type MedicineListResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data MedicineData `json:"data"`
}
type MedicineData struct {
	List []Medicine `json:"list"`
	Page Page       `json:"page"`
}
type Medicine struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`               //产品名称
	Price             decimal.Decimal `json:"price"`              //价格
	Img               string          `json:"img"`                //图片
	Desc              string          `json:"desc"`               //描述
	WithdrawThreshold decimal.Decimal `json:"withdraw_threshold"` //提现额度比例
	Interval          int             `json:"interval"`
	Sort              int             `json:"sort"`        //排序值
	Status            int             `json:"status"`      //是否开启，1为开启，2为关闭
	CreateTime        int64           `json:"create_time"` //创建时间
}

type MedicineRemoteListResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data MedicineRemoteData `json:"data"`
}
type MedicineRemoteData struct {
	List []MedicineRemote `json:"list"`
}
type MedicineRemote struct {
	Name string `json:"name"` //产品名称
	Code string `json:"code"` //产品代码
}

type MedicineGiftOptions struct {
	List []MedicineGiftInfo `json:"list"`
}
type MedicineGiftInfo struct {
	Id   int    `json:"id"`   //产品Id
	Name string `json:"name"` //产品名称
}
