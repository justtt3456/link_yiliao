package response

type RechargeResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data RechargeData `json:"data"`
}
type RechargeData struct {
	List []RechargeInfo `json:"list"`
	Page Page           `json:"page"`
}
type RechargeInfo struct {
	ID          int     `json:"id"`          //
	OrderSn     string  `json:"order_sn"`    //
	UID         int     `json:"uid"`         //关联用户id
	Type        int     `json:"type"`        //充值类别 1银行卡 2在线充值 3后台充值
	Amount      float64 `json:"amount"`      //充值金额
	RealAmount  float64 `json:"real_amount"` //实际到账金额
	From        string  `json:"from"`        //付款账号
	To          string  `json:"to"`          //收款账号
	Voucher     string  `json:"voucher"`     //凭证图
	PaymentID   int     `json:"payment_id"`  //三方支付id
	Status      int     `json:"status"`      //状态，0待审核，1已审核
	UsdtAmount  float64 `json:"usdt_amount"` //usdt充值数量
	Operator    int     `json:"operator"`    //操作的管理员
	Description string  `json:"description"` //订单备注
	UpdateTime  int64   `json:"update_time"` //审核时间
	CreateTime  int64   `json:"create_time"` //创建时间
	Username    string  `json:"username"`
	MethodName  string  `json:"method_name"`
	PaymentName string  `json:"payment_name"`
	SuccessTime int64   `json:"success_time"` //成功时间
	TradeSn     string  `json:"trade_sn"`     //三方订单号
}
