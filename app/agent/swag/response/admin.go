package response

// 管理员列表返回值
type AdminListResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data AdminListData `json:"data"`
}
type AdminListData struct {
	List []AdminInfo `json:"list"`
	Page Page        `json:"page"`
}

// 管理员信息返回值
type AdminItemResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data AdminInfo `json:"data"`
}

// 管理员信息
type AdminInfo struct {
	AdminId    int              `json:"admin_id"` //
	Username   string           `json:"username"` //
	Token      string           `json:"token"`    //
	Role       int              `json:"role"`     //
	RoleName   string           `json:"role_name"`
	CreateTime int64            `json:"create_time"` //
	UpdateTime int64            `json:"update_time"` //
	LoginIp    string           `json:"login_ip"`    //
	RegisterIp string           `json:"register_ip"` //
	Operator   int              `json:"operator"`    //
	Permission []PermissionTree `json:"permission"`
}

// 管理员谷歌验证码
type AdminGoogleResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data AdminGoogle `json:"data"`
}

// 管理员信息
type AdminGoogle struct {
	Username string `json:"username"` //
	Qrcode   string `json:"qrcode"`   //
}
