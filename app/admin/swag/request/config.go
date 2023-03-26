package request

type ConfigBaseUpdate struct {
	ID           int     `json:"id"`             //
	AppName      string  `json:"app_name"`       //网站应用名称
	AppLogo      string  `json:"app_logo"`       //网站应用logo
	VerifiedSend float64 `json:"verified_send"`  //实名送金币
	RegisterSend float64 `json:"register_send"`  //注册24小时后  第一次充值送金币
	OneSend      float64 `json:"one_send"`       //一级奖励
	TwoSend      float64 `json:"two_send"`       //二级奖励
	ThreeSend    float64 `json:"three_send"`     //三级奖励
	SendDesc     string  `json:"send_desc"`      //奖励描述
	RegisterDesc string  `json:"register_desc"`  //注册好礼描述
	TeamDesc     string  `json:"team_desc"`      //团队奖励描述
	OneSendMoeny float64 `json:"one_send_moeny"` //代理返佣基础值  （10）元
	GiftRate     float64 `json:"gift_rate"`      //赠品赠送比例
}

type ConfigFundsUpdate struct {
	ID                  int     `json:"id"`                    //
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
	WithdrawCount       int     `json:"withdraw_count"`        //每日提现次数
	ProductFee          int     `json:"product_fee"`           //购买产品手续费
	ProductQuickAmount  string  `json:"product_quick_amount"`  //购买产品快捷金额
	DayTurnMoneyNum     int64   `json:"day_turn_money_num"`    //每日 可用和可提互转次数
}
type ConfigBankCreate struct {
	BankName   string `json:"bank_name"`   //银行名称
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
	Status     int    `json:"status"`      //
}
type ConfigBankUpdate struct {
	ID         int    `json:"id"`          //
	BankName   string `json:"bank_name"`   //银行名称
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
	Status     int    `json:"status"`      //
}
type ConfigBankUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` // 1=启用 2=关闭
}
type ConfigBankRemove struct {
	ID int `json:"id"` //
}
type ConfigAlipayCreate struct {
	Account  string `json:"account"`   //支付宝账号
	RealName string `json:"real_name"` //真实姓名
	Lang     string `json:"lang"`
	Status   int    `json:"status"` //
}
type ConfigAlipayUpdate struct {
	ID       int    `json:"id"`        //
	Account  string `json:"account"`   //支付宝账号
	RealName string `json:"real_name"` //真实姓名
	Lang     string `json:"lang"`
	Status   int    `json:"status"` //
}
type ConfigAlipayUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` //
}
type ConfigAlipayRemove struct {
	ID int `json:"id"` //
}
type ConfigUsdtCreate struct {
	Address string `json:"address"` //
	Status  int    `json:"status"`  //
	Proto   int    `json:"proto"`   //协议 1 ERC20 2 TRC20
}
type ConfigUsdtUpdate struct {
	ID      int    `json:"id"`      //
	Address string `json:"address"` //
	Status  int    `json:"status"`  //
	Proto   int    `json:"proto"`   //协议 1 ERC20 2 TRC20
}
type ConfigUsdtUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` //
}
type ConfigUsdtRemove struct {
	ID int `json:"id"` //
}
type ConfigKfCreate struct {
	Name      string `json:"name"`       //
	StartTime string `json:"start_time"` //
	EndTime   string `json:"end_time"`   //
	Link      string `json:"link"`       //
	Key       string `json:"key"`        //
	Icon      string `json:"icon"`       //
	Status    int    `json:"status"`     //
}
type ConfigKfUpdate struct {
	ID        int    `json:"id"`         //
	Name      string `json:"name"`       //
	StartTime string `json:"start_time"` //
	EndTime   string `json:"end_time"`   //
	Link      string `json:"link"`       //
	Key       string `json:"key"`        //
	Icon      string `json:"icon"`       //
	Status    int    `json:"status"`     //
}
type ConfigKfUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` //
}
type ConfigKfRemove struct {
	ID int `json:"id"` //
}

type ConfigLangUpdate struct {
	ID        int    `json:"id"`         //
	Name      string `json:"name"`       //
	StartTime string `json:"start_time"` //
	EndTime   string `json:"end_time"`   //
	Link      string `json:"link"`       //
	Key       string `json:"key"`        //
	Icon      string `json:"icon"`       //
	Status    int    `json:"status"`     //
}
type ConfigLangUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` //
}
type ConfigRechargeMethodUpdate struct {
	ID   int    `json:"id"`   //
	Name string `json:"name"` //名称
	Icon string `json:"icon"` //图片
}
type ConfigWithdrawMethodUpdate struct {
	ID   int    `json:"id"`   //
	Name string `json:"name"` //名称
	Icon string `json:"icon"` //图片
}
type ConfigRechargeMethodUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` //
}
type ConfigWithdrawMethodUpdateStatus struct {
	ID     int `json:"id"`     //
	Status int `json:"status"` // 1=开启 2=关闭
}
