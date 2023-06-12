package response

import "github.com/shopspring/decimal"

type InvestResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data InvestInfo `json:"data"`
}
type InvestInfo struct {
	Id             int             `json:"id"`              //
	Name           string          `json:"name"`            //余额宝名称
	Ratio          decimal.Decimal `json:"ratio"`           //利率 0.01=1%!
	FreezeDay      int             `json:"freeze_day"`      //冻结天数
	IncomeInterval int             `json:"income_interval"` //收益发放间隔天数
	Status         int             `json:"status"`          //余额宝开关，1开启，2关闭
	DealTime       int             `json:"deal_time"`       //最后一次执行时间，防止刷单漏洞
	Description    string          `json:"description"`     //余额宝说明
	MinAmount      decimal.Decimal `json:"min_amount"`      //
	CreateTime     int64           `json:"create_time"`     //
	UpdateTime     int64           `json:"update_time"`     //
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
	Id           int             `json:"id"`  //
	UId          int             `json:"uid"` //关联用户id
	Username     string          `json:"username"`
	Type         int             `json:"type"`        //转入转出类型 1转入 2转出
	Amount       decimal.Decimal `json:"amount"`      //转入转出金额
	CreateTime   int64           `json:"in_time"`     //投入时间
	UnfreezeTime int64           `json:"out_time"`    //冻结结束时间
	IncomeTime   int64           `json:"income_time"` //可以发放奖励的首次时间
	Balance      decimal.Decimal `json:"balance"`     //余额宝余额
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
	Id         int             `json:"id"`          //
	UId        int             `json:"uid"`         //关联用户id
	Username   string          `json:"username"`    //
	Income     decimal.Decimal `json:"income"`      //余额宝奖励金额
	Balance    decimal.Decimal `json:"balance"`     //余额宝余额
	CreateTime int64           `json:"create_time"` //生成时间
}
