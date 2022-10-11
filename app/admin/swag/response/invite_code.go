package response

type InviteCodeListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data InviteCodeData `json:"data"`
}
type InviteCodeData struct {
	List []InviteCodeInfo `json:"list"`
	Page Page             `json:"page"`
}
type InviteCodeInfo struct {
	ID         int    `json:"id"`  //
	UID        int    `json:"uid"` //用户id
	Username   string `json:"username"`
	AgentID    int    `json:"agent_id"` //代理id
	AgentName  string `json:"agent_name"`
	Code       string `json:"code"`        //邀请码
	RegCount   int    `json:"reg_count"`   //注册人数
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
