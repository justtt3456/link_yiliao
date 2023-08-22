package request

import "github.com/shopspring/decimal"

type MemberInfo struct {
	Avatar   string `json:"avatar"`   //头像
	Nickname string `json:"nickname"` //昵称
	//Mobile   string `json:"mobile"`   //手机号
	Email  string `json:"email"`  //邮箱
	Qq     string `json:"qq"`     //qq
	Wechat string `json:"wechat"` //微信
}
type MemberPassword struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	PasswordConfirm string `json:"password_confirm"`
}
type MemberVerified struct {
	RealName string `json:"real_name"`
	IdNumber string `json:"id_number"`
	Mobile   string `json:"mobile"`
	Frontend string `json:"frontend"`
	Backend  string `json:"backend"`
}

type MemberTransfer struct {
	Type        int             `json:"type"`         // //1=U转R  2=R转U
	Amount      decimal.Decimal `json:"amount"`       // 金额
	TransferPwd string          `json:"transfer_pwd"` // 交易密码
}
