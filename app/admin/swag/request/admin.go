package request

//查询管理员参数
type AdminListRequest struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Name     string `json:"name" form:"name"` //账号
}

//添加管理员参数
type AdminInsertRequest struct {
	Username   string `json:"username"`    //账号
	Password   string `json:"password"`    //密码
	RePassword string `json:"re_password"` //重复密码
	Role       int    `json:"role"`        //角色
}

//修改管理员参数
type AdminUpdateRequest struct {
	AdminID  int    `json:"admin_id"`
	Password string `json:"password"` //密码
	Role     int    `json:"role"`     //角色
}
type AdminRemoveRequest struct {
	AdminID int `json:"admin_id"`
}
type AdminPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	RePassword  string `json:"re_password"`
}
type AdminGoogleRequest struct {
	AdminID    int    `json:"admin_id"`
	GoogleCode string `json:"google_code"`
}
