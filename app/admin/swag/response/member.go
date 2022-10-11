package response

type MemberListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data MemberListData `json:"data"`
}
type MemberListData struct {
	List []MemberInfo `json:"list"`
	Page Page         `json:"page"`
}
type MemberInfo struct {
	ID               int     `json:"id"`                 //
	Username         string  `json:"username"`           //手机号
	TotalBalance     float64 `json:"total_balance"`      //余额
	Balance          float64 `json:"balance"`            //可用余额
	UseBalance       float64 `json:"use_balance"`        //可提余额
	Freeze           float64 `json:"freeze"`             //冻结金额
	ParentID         int     `json:"parent_id"`          //推荐人id
	AgentID          int     `json:"agent_id"`           //代理id
	IsReal           int     `json:"is_real"`            //是否实名
	RealName         string  `json:"real_name"`          //真实姓名
	InvestFreeze     float64 `json:"invest_freeze"`      //余额宝冻结金额
	InvestAmount     float64 `json:"invest_amount"`      //余额宝有效金额
	InvestIncome     float64 `json:"invest_income"`      //余额宝总收益
	Avatar           string  `json:"avatar"`             //头像
	Status           int     `json:"status"`             //帐号启用状态，1启用2禁用
	FundsStatus      int     `json:"funds_status"`       //帐1启用2禁用
	Level            int     `json:"level"`              //等级
	Score            int     `json:"score"`              //信誉分
	LastLoginTime    int64   `json:"last_login_time"`    //最后登录时间
	LastLoginIP      string  `json:"last_login_ip"`      //最后登录ip
	RegTime          int64   `json:"reg_time"`           //注册时间
	RegisterIP       string  `json:"register_ip"`        //注册ip
	Nickname         string  `json:"nickname"`           //昵称
	Mobile           string  `json:"mobile"`             //手机号
	Email            string  `json:"email"`              //邮箱
	Qq               string  `json:"qq"`                 //qq
	Wechat           string  `json:"wechat"`             //微信
	DisableLoginTime int64   `json:"disable_login_time"` //禁止登录时间
	DisableBetTime   int64   `json:"disable_bet_time"`   //禁止投注时间
	Code             string  `json:"code"`               //邀请码
	IsBuy            int     `json:"is_buy"`             //1=有效 2=无效
}

type Bank struct {
	//银行卡信息
	CardID     int    `json:"card_id"`
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
	ID         int    `json:"id"` //
	UID        int    `json:"uid"`
	Username   string `json:"username"`    //
	RealName   string `json:"real_name"`   //
	IDNumber   string `json:"id_number"`   //
	Mobile     string `json:"mobile"`      //
	Frontend   string `json:"frontend"`    //
	Backend    string `json:"backend"`     //
	Status     int    `json:"status"`      //1审核中 2通过 3驳回
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
