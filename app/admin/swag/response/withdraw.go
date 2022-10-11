package response

type WithdrawResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data WithdrawData `json:"data"`
}
type WithdrawData struct {
	List []WithdrawInfo `json:"list"`
	Page Page           `json:"page"`
}
type WithdrawInfo struct {
	ID           int     `json:"id"`            //
	UID          int     `json:"uid"`           //关联用户id
	WithdrawType int     `json:"withdraw_type"` //提现类型1=银行卡
	MethodName   string  `json:"method_name"`
	BankName     string  `json:"bank_name"`    //关联银行名称
	BranchBank   string  `json:"branch_bank"`  //开户行
	RealName     string  `json:"real_name"`    //开户人
	CardNumber   string  `json:"card_number"`  //卡号
	BankPhone    string  `json:"bank_phone"`   //预留手机号码
	Amount       float64 `json:"amount"`       //实际到账金额
	Fee          float64 `json:"fee"`          //手续费
	TotalAmount  float64 `json:"total_amount"` //提现总额
	UsdtAmount   float64 `json:"usdt_amount"`  //提现总额
	Description  string  `json:"description"`  //审核备注
	Operator     int     `json:"operator"`     //操作管理员
	ViewStatus   int     `json:"view_status"`  //已读状态，0=未读，1=已读
	Status       int     `json:"status"`       //提现状态，0为未审核，1为已审核，2为已拒绝
	SuccessTime  int64   `json:"success_time"` //成功时间
	OrderSn      string  `json:"order_sn"`     //订单号
	PaymentID    int     `json:"payment_id"`   //三方支付id
	PaymentName  string  `json:"payment_name"` //三方支付名称
	TradeSn      string  `json:"trade_sn"`     //三方订单号
	CreateTime   int64   `json:"create_time"`  //
	UpdateTime   int64   `json:"update_time"`  //
	Username     string  `json:"username"`
}
