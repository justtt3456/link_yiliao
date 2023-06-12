package request

type RoleListRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
type RoleCreateRequest struct {
	RoleName string `json:"role_name"`
	Ids      []int  `json:"ids"`
}
type RoleUpdateRequest struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	Ids      []int  `json:"ids"`
}
type RoleRemoveRequest struct {
	RoleId int `json:"role_id" form:"role_id"`
}
