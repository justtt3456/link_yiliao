package response

import "github.com/shopspring/decimal"

type MemberListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data MemberListData `json:"data"`
}
type MemberListData struct {
	List                []MemberInfo    `json:"list"`
	TotalSumProduct     decimal.Decimal `json:"total_sum_product"`
	TotalSumBalance     decimal.Decimal `json:"total_sum_balance"`
	TotalSumUseBalance  decimal.Decimal `json:"total_sum_use_balance"`
	TotalSumIncome      decimal.Decimal `json:"total_sum_income"`
	Page                Page            `json:"page"`
	TotalMemberCount    int             `json:"count"`                 //团队总人数
	TotalRechargeAmount decimal.Decimal `json:"total_recharge_amount"` //团队总充值金额
	TodayRechargeAmount decimal.Decimal `json:"today_recharge_amount"` //今日团队总充值金额
	TotalWithdrawAmount decimal.Decimal `json:"total_withdraw_amount"` //团队总提现金额
	TodayWithdrawAmount decimal.Decimal `json:"today_withdraw_amount"` //今日团队总提现金额
	TotalRechargeCount  int             `json:"total_recharge_count"`  //团队充值总人数
	TodayRechargeCount  int             `json:"today_recharge_count"`  //今日团队充值总人数
}
type MemberInfo struct {
	Id                 int             `json:"id"`                   //
	Username           string          `json:"username"`             //手机号
	TotalBalance       decimal.Decimal `json:"total_balance"`        //余额
	Balance            decimal.Decimal `json:"balance"`              //可用余额
	UseBalance         decimal.Decimal `json:"withdraw_balance"`     //可提余额
	Freeze             decimal.Decimal `json:"freeze"`               //冻结金额
	ParentId           int             `json:"parent_id"`            //推荐人id
	AgentId            int             `json:"agent_id"`             //代理id
	IsReal             int             `json:"is_real"`              //是否实名
	RealName           string          `json:"real_name"`            //真实姓名
	InvestFreeze       decimal.Decimal `json:"invest_freeze"`        //余额宝冻结金额
	InvestAmount       decimal.Decimal `json:"invest_amount"`        //余额宝有效金额
	InvestIncome       decimal.Decimal `json:"invest_income"`        //余额宝总收益
	Avatar             string          `json:"avatar"`               //头像
	Status             int             `json:"status"`               //帐号启用状态，1启用2禁用
	FundsStatus        int             `json:"funds_status"`         //帐1启用2禁用
	Level              int             `json:"level"`                //等级
	Score              int             `json:"score"`                //信誉分
	LastLoginTime      int64           `json:"last_login_time"`      //最后登录时间
	LastLoginIP        string          `json:"last_login_ip"`        //最后登录ip
	RegTime            int64           `json:"reg_time"`             //注册时间
	RegisterIP         string          `json:"register_ip"`          //注册ip
	DisableLoginTime   int64           `json:"disable_login_time"`   //禁止登录时间
	DisableBetTime     int64           `json:"disable_bet_time"`     //禁止投注时间
	Code               string          `json:"code"`                 //邀请码
	IsBuy              int             `json:"is_buy"`               //1=有效 2=无效
	TopId              int             `json:"top_id"`               //上级Id
	TopName            string          `json:"top_name"`             //上级名字
	ProductOrderAmount decimal.Decimal `json:"product_order_amount"` //投注金额
}

type Bank struct {
	//银行卡信息
	CardId     int    `json:"card_id"`
	RealName   string `json:"real_name"`   //
	BankName   string `json:"bank_name"`   //
	CardNumber string `json:"card_number"` //
}

type MemberVerifiedListResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data MemberVerifiedData `json:"data"`
}
type MemberVerifiedData struct {
	List []MemberVerified `json:"list"`
	Page Page             `json:"page"`
}
type MemberVerified struct {
	Id         int    `json:"id"` //
	UId        int    `json:"uid"`
	Username   string `json:"username"`    //
	RealName   string `json:"real_name"`   //
	IdNumber   string `json:"id_number"`   //
	Mobile     string `json:"mobile"`      //
	Frontend   string `json:"frontend"`    //
	Backend    string `json:"backend"`     //
	Status     int    `json:"status"`      //1审核中 2通过 3驳回
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
