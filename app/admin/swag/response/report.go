package response

type ReportResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ReportData `json:"data"`
}
type ReportData struct {
	RechargeAmount      float64 `json:"recharge_amount"`       //今日充值
	RechargeAmountTotal float64 `json:"recharge_amount_total"` //总充值
	WithdrawAmount      float64 `json:"withdraw_amount"`       //今日提现
	WithdrawAmountTotal float64 `json:"withdraw_amount_total"` //总提现
	RegCount            int64   `json:"reg_count"`             //今日注册
	RegCountTotal       int64   `json:"reg_count_total"`       //总注册
	RegBuyCount         int64   `json:"reg_buy_count"`         //今日有效会员
	RegBuyCountTotal    int64   `json:"reg_buy_count_total"`   //总有效会员
	SendMoney           float64 `json:"send_money"`            //今日用户总收益
	SendMoneyTotal      float64 `json:"send_money_total"`      //用户总收益
}

type MemberReportResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data MemberReportData `json:"data"`
}
type MemberReportData struct {
	List []MemberReport `json:"list"`
	Page Page           `json:"page"`
}
type MemberReport struct {
	ID             int     `json:"id"`              //
	UID            int     `json:"uid"`             //
	Username       string  `json:"username"`        //用户名
	RechargeCount  int     `json:"recharge_count"`  //充值次数
	RechargeAmount float64 `json:"recharge_amount"` //充值金额
	WithdrawCount  int     `json:"withdraw_count"`  //提现次数
	WithdrawAmount float64 `json:"withdraw_amount"` //提现金额
	BetCount       int     `json:"bet_count"`       //投注次数
	BetAmount      float64 `json:"bet_amount"`      //投注金额
	BetResult      float64 `json:"bet_result"`      //输赢
	SysUp          float64 `json:"sys_up"`          //系统上分
	SysDown        float64 `json:"sys_down"`        //系统下分
	Freeze         float64 `json:"freeze"`          //系统冻结
	Unfreeze       float64 `json:"unfreeze"`        //系统解冻
	CreateTime     int64   `json:"create_time"`     //
	UpdateTime     int64   `json:"update_time"`     //
}
type AgentReportResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data AgentReportData `json:"data"`
}
type AgentReportData struct {
	List []AgentReport `json:"list"`
	Page Page          `json:"page"`
}
type AgentReport struct {
	ID             int     `json:"id"`              //
	Aid            int     `json:"aid"`             //代理id
	Username       string  `json:"username"`        //代理名称
	RechargeCount  int     `json:"recharge_count"`  //充值次数
	RechargeAmount float64 `json:"recharge_amount"` //充值金额
	WithdrawCount  int     `json:"withdraw_count"`  //提现次数
	WithdrawAmount float64 `json:"withdraw_amount"` //提现金额
	BetCount       int     `json:"bet_count"`       //投注次数
	BetAmount      float64 `json:"bet_amount"`      //投注金额
	BetResult      float64 `json:"bet_result"`      //输赢
	SysUp          float64 `json:"sys_up"`          //系统上分
	SysDown        float64 `json:"sys_down"`        //系统下分
	Freeze         float64 `json:"freeze"`          //系统冻结
	Unfreeze       float64 `json:"unfreeze"`        //系统解冻
	RegisterCount  int     `json:"register_count"`
	CreateTime     int64   `json:"create_time"` //
	UpdateTime     int64   `json:"update_time"` //
}