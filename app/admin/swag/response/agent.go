package response

type AgentListResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data AgentData `json:"data"`
}
type AgentData struct {
	List []AgentInfo `json:"list"`
	Page Page        `json:"page"`
}
type AgentInfo struct {
	ID         int    `json:"id"`          //
	Name       string `json:"name"`        //
	ParentID   int    `json:"parent_id"`   //父级id 为0时则为组长
	ParentName string `json:"parent_name"` //组长
	GroupName  string `json:"group_name"`  //小组名称
	Status     int    `json:"status"`      //帐号启用状态，1启用2禁用
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
