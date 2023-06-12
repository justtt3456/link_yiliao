package request

type LoginRequest struct {
	Username   string `json:"username"`    //用户名
	Password   string `json:"password"`    //密码
	GoogleCode string `json:"google_code"` //谷歌验证码
}
