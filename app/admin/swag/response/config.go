package response

import "github.com/shopspring/decimal"

type ConfigBaseResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ConfigBase `json:"data"`
}
type ConfigBase struct {
	Id                int             `json:"id"`                  //
	AppName           string          `json:"app_name"`            //网站应用名称
	AppLogo           string          `json:"app_logo"`            //网站应用logo
	VerifiedSend      decimal.Decimal `json:"verified_send"`       //实名送金币
	RegisterSend      decimal.Decimal `json:"register_send"`       //注册24小时后  第一次充值送金币
	OneSend           decimal.Decimal `json:"one_send"`            //一级奖励
	TwoSend           decimal.Decimal `json:"two_send"`            //二级奖励
	ThreeSend         decimal.Decimal `json:"three_send"`          //三级奖励
	OneSendMoney      decimal.Decimal `json:"one_send_money"`      //代理返佣基础值  （10）元
	SendDesc          string          `json:"send_desc"`           //奖励描述
	RegisterDesc      string          `json:"register_desc"`       //注册好礼描述
	TeamDesc          string          `json:"team_desc"`           //团队奖励描述
	GiftRate          decimal.Decimal `json:"gift_rate"`           //赠品赠送比例
	SignRewards       decimal.Decimal `json:"sign_rewards"`        //签到奖励
	OneReleaseRate    decimal.Decimal `json:"one_release_rate"`    //一级代理释放比
	TwoReleaseRate    decimal.Decimal `json:"two_release_rate"`    //二级代理释放比
	ThreeReleaseRate  decimal.Decimal `json:"three_release_rate"`  //三级代理释放比
	IncomeBalanceRate decimal.Decimal `json:"income_balance_rate"` //收益转可用余额比例
	EquityStartDate   int64           `json:"equity_start_date"`   //股权分开启时间
	EquityRate        decimal.Decimal `json:"equity_rate"`         //股权分额度比例
	EquityInterval    int             `json:"equity_interval"`     //股权分收益天数
	EquityIncomeRate  decimal.Decimal `json:"equity_income_rate"`  //股权分收益比例
	DownloadUrl       string          `json:"download_url"`        //下载链接
	//UsdtBuyRate       decimal.Decimal `json:"usdt_buy_rate"`       //usdt买汇率
	//UsdtSellRate      decimal.Decimal `json:"usdt_sell_rate"`      //usdt卖汇率
}
type ConfigFundsResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ConfigFunds `json:"data"`
}
type ConfigFunds struct {
	Id                  int             `json:"id"`                    //
	RechargeStartTime   string          `json:"recharge_start_time"`   //充值开始时间
	RechargeEndTime     string          `json:"recharge_end_time"`     //充值结束时间
	RechargeMinAmount   decimal.Decimal `json:"recharge_min_amount"`   //充值最小金额
	RechargeMaxAmount   decimal.Decimal `json:"recharge_max_amount"`   //充值最大金额
	RechargeFee         decimal.Decimal `json:"recharge_fee"`          //充值手续费(百分比)
	RechargeQuickAmount string          `json:"recharge_quick_amount"` //快捷充值金额
	WithdrawStartTime   string          `json:"withdraw_start_time"`   //提现开始时间
	WithdrawEndTime     string          `json:"withdraw_end_time"`     //提现结束时间
	MustPassword        int             `json:"must_password"`         //是否必须体现密码
	PasswordFreeze      int             `json:"password_freeze"`       //提现密码错误冻结次数
	WithdrawMinAmount   decimal.Decimal `json:"withdraw_min_amount"`   //提现最小金额
	WithdrawMaxAmount   decimal.Decimal `json:"withdraw_max_amount"`   //提现最大金额
	WithdrawFee         decimal.Decimal `json:"withdraw_fee"`          //提现手续费
	WithdrawCount       int             `json:"withdraw_count"`        //每日提现次数
	ProductFee          decimal.Decimal `json:"product_fee"`           //购买产品手续费
	ProductQuickAmount  string          `json:"product_quick_amount"`  //购买产品快捷金额
	DayTurnMoneyNum     int64           `json:"day_turn_money_num"`    //每日 可用和可提互转次数
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
	Id         int    `json:"id"`          //
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
	Id         int    `json:"id"`          //
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
	Id         int    `json:"id"`          //
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
	Id         int    `json:"id"`          //
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
	Id         int    `json:"id"`          //
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
	Id     int    `json:"id"`     //Id
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
	Id     int    `json:"id"`     //
	Name   string `json:"name"`   //名称
	Code   string `json:"code"`   //code码
	Icon   string `json:"icon"`   //图片
	Status int    `json:"status"` //1=开启 2=关闭
	Fee    int    `json:"fee"`    //提现手续费百分比
}
