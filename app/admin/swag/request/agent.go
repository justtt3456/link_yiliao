package request

type AgentList struct {
	ParentID int    `form:"parent_id"`
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
}
type AgentCreate struct {
	Name      string `json:"name"` //
	Password  string `json:"password"`
	ParentID  int    `json:"parent_id"`  //父级id 为0时则为组长
	GroupName string `json:"group_name"` //小组名称
	Status    int    `json:"status"`     //帐号启用状态，1启用2禁用
}
type AgentUpdate struct {
	ID        int    `json:"id"`         //
	ParentID  int    `json:"parent_id"`  //父级id 为0时则为组长
	GroupName string `json:"group_name"` //小组名称
	Password  string `json:"password"`
	Status    int    `json:"status"` //帐号启用状态，1启用2禁用
}
type AgentUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"` //状态
}
type AgentRemove struct {
	ID int `json:"id"`
}
