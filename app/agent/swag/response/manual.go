package response

import "github.com/shopspring/decimal"

type ManualListResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ManualData `json:"data"`
}
type ManualData struct {
	List []ManualInfo `json:"list"`
	Page Page         `json:"page"`
}
type ManualInfo struct {
	Id         int             `json:"id"`       //
	UserId     int             `json:"user_id"`  //
	Username   string          `json:"username"` //
	Type       int             `json:"type"`
	Amount     decimal.Decimal `json:"amount"`      //金额
	AdminName  string          `json:"admin_name"`  //操作人
	AgentName  string          `json:"agent_name"`  //操作人
	CreateTime int64           `json:"create_time"` //创建时间
}
