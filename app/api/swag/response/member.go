package response

import "github.com/shopspring/decimal"

type MemberResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Member `json:"data"`
}
type Member struct {
	Id       int    `json:"id"`       //
	Username string `json:"username"` //手机号
	//TotalBalance        decimal.Decimal `json:"total_balance"`         //可用余额
	Balance             decimal.Decimal `json:"balance"`               //可用余额
	WithdrawBalance     decimal.Decimal `json:"withdraw_balance"`      //可提现余额
	WithdrawThreshold   decimal.Decimal `json:"withdraw_threshold"`    //可提现额度
	IsReal              int             `json:"is_real"`               //是否实名 0未实名 1审核中 2通过 3驳回
	RealName            string          `json:"real_name"`             //真实姓名
	InvestFreeze        decimal.Decimal `json:"invest_freeze"`         //余额宝冻结金额
	InvestAmount        decimal.Decimal `json:"invest_amount"`         //余额宝有效金额
	InvestIncome        decimal.Decimal `json:"invest_income"`         //余额宝总收益
	Avatar              string          `json:"avatar"`                //头像
	Status              int             `json:"status"`                //帐号状态 1启用2禁用
	FundsStatus         int             `json:"funds_status"`          //资金状态 1启用2禁用
	Level               int             `json:"level"`                 //等级
	Score               int             `json:"score"`                 //信誉分
	LastLoginTime       int64           `json:"last_login_time"`       //最后登录时间
	LastLoginIP         string          `json:"last_login_ip"`         //最后登录ip
	RegTime             int64           `json:"reg_time"`              //注册时间
	HasWithdrawPassword int             `json:"has_withdraw_password"` //是否设置提现密码
	RegisterIP          string          `json:"register_ip"`           //注册ip
	Token               string          `json:"token"`                 //token盐
	InviteCode          string          `json:"invite_code"`           //邀请码
	Coupon              []Coupon        `json:"coupon"`                //用户有的优惠券
	//Income decimal.Decimal `json:"income"` //总收益
	//Guquan              int64           `json:"guquan"`                //股权
	Message    int64           `json:"message"`     //站内信总条数
	PreIncome  decimal.Decimal `json:"pre_income"`  //待收益
	PreCapital decimal.Decimal `json:"pre_capital"` //待收本金
}

type Coupon struct {
	UseId int             `json:"use_id"` //使用优惠券传的id
	Id    int             `json:"id"`     //优惠券Id
	Price decimal.Decimal `json:"price"`  //优惠券面额
}

type MyTeam struct {
	Id       int             `json:"id"`        //
	Username string          `json:"username"`  //用户名
	Level    int             `json:"level"`     //层级
	RegTime  int64           `json:"reg_time"`  //注册时间
	Income   decimal.Decimal `json:"income"`    //收益
	RealName string          `json:"real_name"` //实名姓名
}

type MyTeamList struct {
	List           []MyTeam        `json:"list"`
	RegisterMember int             `json:"register_member"` //注册人数
	BuyMember      int64           `json:"buy_member"`      //激活人数
	TotalRecharge  decimal.Decimal `json:"total_recharge"`  //总计充值
	TotalRebate    decimal.Decimal `json:"total_rebate"`    //总返佣
	Page           Page            `json:"page"`
}
