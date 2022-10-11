package response

type RoleListResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data RoleData `json:"data"`
}
type RoleData struct {
	List []RoleInfo `json:"list"`
	Page Page       `json:"page"`
}
type RoleInfo struct {
	RoleID      int              `json:"role_id"`     //
	RoleName    string           `json:"role_name"`   //
	Status      int              `json:"status"`      //
	CreateTime  int64            `json:"create_time"` //创建时间
	UpdateTime  int64            `json:"update_time"` //修改时间
	Permissions []PermissionTree `json:"permissions"`
}
