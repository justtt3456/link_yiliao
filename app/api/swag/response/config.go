package response

type ConfigResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Config `json:"data"`
}
type Config struct {
	Base  Base   `json:"base"`
	Funds Funds  `json:"funds"`
	Kf    []Kf   `json:"kf"`
	Lang  []Lang `json:"lang"`
}
type Base struct {
	AppName      string  `json:"app_name"`      //网站应用名称
	AppLogo      string  `json:"app_logo"`      //网站应用logo
	VerifiedSend float64 `json:"verified_send"` //实名送金币
	RegisterSend float64 `json:"register_send"` //注册24小时后  第一次充值送金币
	OneSend      float64 `json:"one_send"`      //一级奖励
	TwoSend      float64 `json:"two_send"`      //二级奖励
	ThreeSend    float64 `json:"three_send"`    //三级奖励
	SendDesc     string  `json:"send_desc"`     //奖励描述
	RegisterDesc string  `json:"register_desc"` //注册好礼描述
	TeamDesc     string  `json:"team_desc"`     //团队奖励描述

}
type Funds struct {
	RechargeStartTime   string  `json:"recharge_start_time"`   //充值开始时间
	RechargeEndTime     string  `json:"recharge_end_time"`     //充值结束时间
	RechargeMinAmount   float64 `json:"recharge_min_amount"`   //充值最小金额
	RechargeMaxAmount   float64 `json:"recharge_max_amount"`   //充值最大金额
	RechargeFee         int     `json:"recharge_fee"`          //充值手续费(百分比)
	RechargeQuickAmount string  `json:"recharge_quick_amount"` //快捷充值金额
	WithdrawStartTime   string  `json:"withdraw_start_time"`   //提现开始时间
	WithdrawEndTime     string  `json:"withdraw_end_time"`     //提现结束时间
	MustPassword        int     `json:"must_password"`         //是否必须体现密码
	PasswordFreeze      int     `json:"password_freeze"`       //提现密码错误冻结次数
	WithdrawMinAmount   float64 `json:"withdraw_min_amount"`   //提现最小金额
	WithdrawMaxAmount   float64 `json:"withdraw_max_amount"`   //提现最大金额
	WithdrawFee         int     `json:"withdraw_fee"`          //提现手续费
	ProductFee          int     `json:"product_fee"`           //购买产品手续费
	ProductQuickAmount  string  `json:"product_quick_amount"`  //购买产品快捷金额
	KfRecharge          int     `json:"kf_recharge"`           //客服充值(联系客服)
}
type Kf struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`       //
	StartTime string `json:"start_time"` //
	EndTime   string `json:"end_time"`   //
	Link      string `json:"link"`       //
}
type Lang struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`       //语言名称
	Code      string `json:"code"`       //英文简称
	Icon      string `json:"icon"`       //语言图标
	IsDefault int    `json:"is_default"` //是否默认语言
}
