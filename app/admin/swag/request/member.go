package request

type MemberCreate struct {
	Username string  `json:"username"` //用户名
	Password string  `json:"password"` //密码
	Balance  float64 `json:"balance"`  //余额
}
type MemberList struct {
	ID        int    `json:"id" form:"id"`
	Mobile    string `json:"mobile" form:"mobile"`
	Username  string `json:"username" form:"username"` //用户名
	RealName  string `json:"real_name" form:"real_name"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
}
type MemberUpdate struct {
	ID          int    `json:"id"`
	Description string `json:"description"` //备注
}
type MemberUpdatePassword struct {
	ID          int    `json:"id"`
	Password    string `json:"password"`     //密码
	PayPassword string `json:"pay_password"` // 交易密码
}
type MemberUpdateBankCard struct {
	ID         int    `json:"id"`
	RealName   string `json:"real_name"`
	BankName   string `json:"bank_name"`
	CardNumber string `json:"card_number"`
}
type MemberUpdateStatus struct {
	ID     int    `json:"id" form:"id"`
	Type   string `json:"type"` //login =禁止登录  funds =冻结资金
	Status int    `json:"status" form:"status"`
}
type MemberVerifiedList struct {
	Username string `json:"username" form:"username"` //用户名
	Status   int    `json:"status" form:"status"`     //状态 1审核中 2通过 3驳回
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
type MemberVerifiedUpdate struct {
	ID     int `json:"id"`
	Status int `json:"status" form:"status"` // 2通过 3驳回
}
type MemberVerifiedRemove struct {
	ID int `json:"id"`
}

type MemberTeamReq struct {
	Page     int  `json:"page" form:"page"`
	PageSize int  `json:"page_size" form:"page_size"`
	UserId   int  `json:"user_id"` //用户id
	Level    *int `json:"level"`   //用户层级（手动输入即可  1代表1级）
}

type SendCouponReq struct {
	Ids      string `json:"ids"`       //用户ID  用,隔开
	CouponId int64  `json:"coupon_id"` //券的ID
}

type GetCodeReq struct {
	Mobile string `json:"mobile"` //手机号
}
