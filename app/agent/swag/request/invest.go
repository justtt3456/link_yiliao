package request

import "github.com/shopspring/decimal"

type InvestUpdate struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`            //余额宝名称
	Ratio          decimal.Decimal `json:"ratio"`           //利率 0.01=1%!
	FreezeDay      int             `json:"freeze_day"`      //冻结天数
	IncomeInterval int             `json:"income_interval"` //收益发放间隔天数
	Description    string          `json:"description"`     //余额宝说明
	Status         int             `json:"status"`          //余额宝开关，1开启，2关闭
	MinAmount      decimal.Decimal `json:"min_amount"`      //
}
type InvestOrder struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Username string `form:"username"`
}
type InvestIncome struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Username string `form:"username"`
}
