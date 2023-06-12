package request

type PermissionRequest struct {
	Role int `json:"role" form:"role"`
}
type PermissionSetRequest struct {
	Ids    []int `json:"ids"`     //权限id数组
	RoleId int   `json:"role_id"` //角色id
}

type PermissionCreateRequest struct {
	Backend  string `json:"backend"`  //后台路由
	Frontend string `json:"frontend"` //后台路由
	Label    string `json:"label"`    //名称
	PId      int    `json:"pid"`      //上级id
	IsBtn    int    `json:"is_btn"`   //是否按钮
	Sort     int    `json:"sort"`     //排序

}
type PermissionUpdateRequest struct {
	Id       int    `json:"id"`
	Backend  string `json:"backend"`  //后台路由
	Frontend string `json:"frontend"` //后台路由
	Label    string `json:"label"`    //名称
	PId      int    `json:"pid"`      //上级id
	IsBtn    int    `json:"is_btn"`   //是否按钮
	Sort     int    `json:"sort"`     //排序
}
type PermissionRemoveRequest struct {
	Id int `json:"id"`
}
