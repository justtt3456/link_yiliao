package response

import "github.com/shopspring/decimal"

type InvestIndexResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data InvestIndexData `json:"data"`
}
type InvestIndexData struct {
	Info   InvestInfo       `json:"info"`   //余额宝信息
	Income InvestIncomeData `json:"income"` //收益记录
	Member InvestMember     `json:"member"` //用户余额宝信息
}

type InvestMember struct {
	Balance         decimal.Decimal `json:"balance"`          //余额
	TotalIncome     decimal.Decimal `json:"total_income"`     //总收益
	YesterdayIncome decimal.Decimal `json:"yesterday_income"` //昨日收益
}
type InvestInfo struct {
	Name           string          `json:"name"`            //余额宝名称
	Ratio          decimal.Decimal `json:"ratio"`           //利率 年化收益百分比
	FreezeDay      int             `json:"freeze_day"`      //冻结天数
	IncomeInterval int             `json:"income_interval"` //收益发放间隔天数
	Status         int             `json:"status"`          //余额宝开关，1开启，2关闭
	Description    string          `json:"description"`     //余额宝说明
}
type InvestIncomeResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data InvestIncomeData `json:"data"`
}
type InvestIncomeData struct {
	List []InvestIncome `json:"list"`
	Page Page           `json:"page"`
}
type InvestIncome struct {
	Income     decimal.Decimal `json:"income"`      //余额宝奖励金额
	Balance    decimal.Decimal `json:"balance"`     //余额宝当前余额
	CreateTime int64           `json:"create_time"` //生成时间
}

type InvestOrderResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data InvestOrderData `json:"data"`
}
type InvestOrderData struct {
	List []InvestOrder `json:"list"`
	Page Page          `json:"page"`
}
type InvestOrder struct {
	Type         int             `json:"type"`          //转入转出类型 1转入 2转出
	Amount       decimal.Decimal `json:"amount"`        //转入转出金额
	CreateTime   int64           `json:"create_time"`   //投入时间
	UnfreezeTime int64           `json:"unfreeze_time"` //冻结结束时间
	IncomeTime   int64           `json:"income_time"`   //可以发放奖励的首次时间
	Balance      decimal.Decimal `json:"balance"`       //当前余额宝余额
}
