package response

import "github.com/shopspring/decimal"

type MedicineOrder struct {
	Username          string          `json:"username"`    //用户名
	Uid               int             `json:"uid"`         //用户Id
	Name              string          `json:"name"`        //产品名字
	CreateTime        int             `json:"create_time"` //投资时间
	Amount            decimal.Decimal `json:"amount"`      //金额
	RealName          string          `json:"real_name"`
	AgentName         string          `json:"agent_name"`
	Address           string          `json:"address"`
	WithdrawThreshold decimal.Decimal `json:"withdraw_threshold"` //
	Interval          int             `json:"interval"`
	Status            int             `json:"status"` //状态
	Current           int             `json:"current"`
}

type MedicineOrderPageListResponse struct {
	List        []MedicineOrder `json:"list"`
	Page        Page            `json:"page"`
	TotalAmount decimal.Decimal `json:"total_amount"`
}
