package request

import "github.com/shopspring/decimal"

type MedicineList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Name     string `form:"name"`
	Category int    `form:"category"`
}

type MedicineCreate struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`               //产品名称
	Price             decimal.Decimal `json:"price"`              //价格
	Img               string          `json:"img"`                //图片
	Desc              string          `json:"desc"`               //描述
	WithdrawThreshold decimal.Decimal `json:"withdraw_threshold"` //提现额度比例
	Interval          int             `json:"interval"`
	Sort              int             `json:"sort"`   //排序值
	Status            int             `json:"status"` //是否开启，1为开启，2为关闭
}
type MedicineUpdate struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`               //产品名称
	Price             decimal.Decimal `json:"price"`              //价格
	Img               string          `json:"img"`                //图片
	Desc              string          `json:"desc"`               //描述
	WithdrawThreshold decimal.Decimal `json:"withdraw_threshold"` //提现额度比例
	Interval          int             `json:"interval"`
	Sort              int             `json:"sort"`   //排序值
	Status            int             `json:"status"` //是否开启，1为开启，2为关闭
}
type MedicineUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态
}
type MedicineRemove struct {
	Id int `json:"id"`
}
