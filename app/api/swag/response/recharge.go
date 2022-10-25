package response

type RechargeCreate struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data RechargeUrl `json:"data"`
}
type RechargeUrl struct {
	Url string `json:"url"`
}
type RechargeResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data RechargeList `json:"data"`
}
type RechargeList struct {
	List []Recharge `json:"list"`
	Page Page       `json:"page"`
}
type Recharge struct {
	ID         int     `json:"id"`          //
	OrderSn    string  `json:"order_sn"`    //订单号
	Type       int     `json:"type"`        //充值类别
	TypeName   string  `json:"type_name"`   //充值类别
	Amount     float64 `json:"amount"`      //充值金额
	RealAmount float64 `json:"real_amount"` //实际到账金额
	Remarks    string  `json:"remarks"`     //备注
	From       string  `json:"from"`        //付款账号
	To         string  `json:"to"`          //收款账号
	Voucher    string  `json:"voucher"`     //凭证图
	Status     int     `json:"status"`      //1待审核  2已审核 3驳回
	UpdateTime int64   `json:"update_time"` //审核时间
	CreateTime int64   `json:"create_time"` //创建时间
}
type RechargeMethodResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data RechargeMethodData `json:"data"`
}
type RechargeMethodData struct {
	List []RechargeMethod `json:"list"`
}
type RechargeMethod struct {
	ID   int               `json:"id"`   //
	Name string            `json:"name"` //
	Code string            `json:"code"` //
	Icon string            `json:"icon"` //
	Info []map[string]interface{} `json:"info"` //
}
