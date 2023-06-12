package request

type InviteCodeList struct {
	Code     string `json:"code"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
}
type InviteCodeCreate struct {
	AgentId int    `json:"agent_id"` //代理id
	Code    string `json:"code"`     //邀请码
}
type InviteCodeUpdate struct {
	Id   int    `json:"id"`   //
	Code string `json:"code"` //邀请码
}
type InviteCodeRemove struct {
	Id int `json:"id"`
}
