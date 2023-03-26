package response

type ConfigBaseResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ConfigBase `json:"data"`
}
type ConfigBase struct {
	ID           int     `json:"id"`             //
	AppName      string  `json:"app_name"`       //网站应用名称
	AppLogo      string  `json:"app_logo"`       //网站应用logo
	VerifiedSend float64 `json:"verified_send"`  //实名送金币
	RegisterSend float64 `json:"register_send"`  //注册24小时后  第一次充值送金币
	OneSend      float64 `json:"one_send"`       //一级奖励
	TwoSend      float64 `json:"two_send"`       //二级奖励
	ThreeSend    float64 `json:"three_send"`     //三级奖励
	OneSendMoeny float64 `json:"one_send_moeny"` //代理返佣基础值  （10）元
	SendDesc     string  `json:"send_desc"`      //奖励描述
	RegisterDesc string  `json:"register_desc"`  //注册好礼描述
	TeamDesc     string  `json:"team_desc"`      //团队奖励描述
	GiftRate     int     `json:"gift_rate"`      //赠品赠送比例
}
type ConfigFundsResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ConfigFunds `json:"data"`
}
type ConfigFunds struct {
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
	WithdrawFee         float64 `json:"withdraw_fee"`          //提现手续费
	WithdrawCount       int     `json:"withdraw_count"`        //每日提现次数
	ProductFee          int     `json:"product_fee"`           //购买产品手续费
	ProductQuickAmount  string  `json:"product_quick_amount"`  //购买产品快捷金额
	DayTurnMoneyNum     int64   `json:"day_turn_money_num"`    //每日 可用和可提互转次数
}
type ConfigBankResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data ConfigBankData `json:"data"`
}
type ConfigBankData struct {
	List []ConfigBank `json:"list"`
}
type ConfigBank struct {
	ID         int    `json:"id"`          //
	BankName   string `json:"bank_name"`   //银行名称
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
	Status     int    `json:"status"`      //
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
type ConfigAlipayResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data ConfigAlipayData `json:"data"`
}
type ConfigAlipayData struct {
	List []ConfigAlipay `json:"list"`
}
type ConfigAlipay struct {
	ID         int    `json:"id"`          //
	Account    string `json:"account"`     //支付宝账号
	RealName   string `json:"real_name"`   //真实姓名
	Status     int    `json:"status"`      //
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
type ConfigUsdtResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data ConfigUsdtData `json:"data"`
}
type ConfigUsdtData struct {
	List []ConfigUsdt `json:"list"`
}
type ConfigUsdt struct {
	ID         int    `json:"id"`          //
	Address    string `json:"address"`     //
	Status     int    `json:"status"`      //
	Proto      int    `json:"proto"`       //协议 1 ERC20 2 TRC20
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
type ConfigKfResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ConfigKfData `json:"data"`
}
type ConfigKfData struct {
	List []ConfigKf `json:"list"`
}
type ConfigKf struct {
	ID         int    `json:"id"`          //
	Name       string `json:"name"`        //
	StartTime  string `json:"start_time"`  //
	EndTime    string `json:"end_time"`    //
	Link       string `json:"link"`        //
	Key        string `json:"key"`         //
	Icon       string `json:"icon"`        //
	Status     int    `json:"status"`      //
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
type ConfigLangResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data ConfigLangData `json:"data"`
}
type ConfigLangData struct {
	List []ConfigLang `json:"list"`
}
type ConfigLang struct {
	ID         int    `json:"id"`          //
	Name       string `json:"name"`        //语言名称
	Code       string `json:"code"`        //英文简称
	Icon       string `json:"icon"`        //语言图标
	IsDefault  int    `json:"is_default"`  //是否默认语言
	Status     int    `json:"status"`      //状态
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
type ConfigRechargeMethodResponse struct {
	Code int                      `json:"code"`
	Msg  string                   `json:"msg"`
	Data ConfigRechargeMethodData `json:"data"`
}
type ConfigRechargeMethodData struct {
	List []ConfigRechargeMethod `json:"list"`
}
type ConfigRechargeMethod struct {
	ID     int    `json:"id"`     //ID
	Name   string `json:"name"`   //名称
	Code   string `json:"code"`   //code码
	Icon   string `json:"icon"`   //图片
	Lang   string `json:"lang"`   //语言  默认zh_cn
	Status int    `json:"status"` //状态 1=开启  2=关闭
}
type ConfigWithdrawMethodResponse struct {
	Code int                      `json:"code"`
	Msg  string                   `json:"msg"`
	Data ConfigWithdrawMethodData `json:"data"`
}
type ConfigWithdrawMethodData struct {
	List []ConfigWithdrawMethod `json:"list"`
}
type ConfigWithdrawMethod struct {
	ID     int    `json:"id"`     //
	Name   string `json:"name"`   //名称
	Code   string `json:"code"`   //code码
	Icon   string `json:"icon"`   //图片
	Status int    `json:"status"` //1=开启 2=关闭
	Fee    int    `json:"fee"`    //提现手续费百分比
}
