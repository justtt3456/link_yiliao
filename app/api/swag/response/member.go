package response

type MemberResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Member `json:"data"`
}
type Member struct {
	ID                  int      `json:"id"`                    //
	Username            string   `json:"username"`              //手机号
	TotalBalance        float64  `json:"total_balance"`         //可用余额
	Balance             float64  `json:"balance"`               //可用余额
	UseBalance          float64  `json:"use_balance"`           //可提现余额
	ParentID            int      `json:"parent_id"`             //推荐人id
	IsReal              int      `json:"is_real"`               //是否实名 0未实名 1审核中 2通过 3驳回
	RealName            string   `json:"real_name"`             //真实姓名
	InvestFreeze        float64  `json:"invest_freeze"`         //余额宝冻结金额
	InvestAmount        float64  `json:"invest_amount"`         //余额宝有效金额
	InvestIncome        float64  `json:"invest_income"`         //余额宝总收益
	Avatar              string   `json:"avatar"`                //头像
	Status              int      `json:"status"`                //帐号状态 1启用2禁用
	FundsStatus         int      `json:"funds_status"`          //资金状态 1启用2禁用
	Level               int      `json:"level"`                 //等级
	Score               int      `json:"score"`                 //信誉分
	LastLoginTime       int64    `json:"last_login_time"`       //最后登录时间
	LastLoginIP         string   `json:"last_login_ip"`         //最后登录ip
	RegTime             int64    `json:"reg_time"`              //注册时间
	HasWithdrawPassword int      `json:"has_withdraw_password"` //是否设置提现密码
	RegisterIP          string   `json:"register_ip"`           //注册ip
	Token               string   `json:"token"`                 //token盐
	Nickname            string   `json:"nickname"`              //昵称
	Mobile              string   `json:"mobile"`                //手机号
	Email               string   `json:"email"`                 //邮箱
	Qq                  string   `json:"qq"`                    //qq
	Wechat              string   `json:"wechat"`                //微信
	InviteCode          string   `json:"invite_code"`           //邀请码
	Coupon              []Coupon `json:"coupon"`                //用户有的优惠券
	Income              float64  `json:"income"`                //总收益
	Guquan              int64    `json:"guquan"`                //股权
}

type Coupon struct {
	UseId int64   `json:"use_id"` //使用优惠券传的id
	Id    int64   `json:"id"`     //优惠券Id
	Price float64 `json:"price"`  //优惠券面额
}

type MyTeam struct {
	ID       int    `json:"id"`       //
	Username string `json:"username"` //用户名
	Level    int    `json:"level"`    //层级
	RegTime  int64  `json:"reg_time"` //注册时间
}

type MyTeamList struct {
	List []MyTeam `json:"list"`
	Page Page     `json:"page"`
}
