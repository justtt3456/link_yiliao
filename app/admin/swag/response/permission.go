package response

type PermissionListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data []PermissionTree `json:"data"`
}
type PermissionTree struct {
	PermissionInfo
	Children []PermissionTree `json:"children"`
}
type PermissionInfo struct {
	ID       int    `json:"id"`
	PID      int    `json:"pid"`
	Label    string `json:"label"`
	Frontend string `json:"frontend"`
	Backend  string `json:"backend"`
	IsBtn    int    `json:"is_btn"`
	Checked  bool   `json:"checked"`
	Sort     int    `json:"sort"`
}
