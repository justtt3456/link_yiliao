package request

type Register struct {
	Username         string `json:"username"`          //用户名
	Password         string `json:"password"`          //密码
	RePassword       string `json:"re_password"`       //重复密码
	InviteCode       string `json:"invite_code"`       //邀请码
	WithdrawPassword string `json:"withdraw_password"` //提现密码
	Code             string `json:"code"`              //验证码
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Code     string `json:"code"`
}

type Help struct {
	Category int `json:"category" form:"category"` //1公司简介2推荐奖励
}

type SendCode struct {
	Username string `json:"username"` //用户名（手机号）
	Type     int    `json:"type"`     //1=注册
}
