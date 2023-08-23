package pay

import (
	"china-russia/model"
	"github.com/shopspring/decimal"
)

type PayInterface interface {
	//发起充值代收
	Recharge(RechargeParam) PaymentResponse
	//创建签名
	createSign(map[string]interface{}) string
	//验证签名
	VerifySign(map[string]interface{}) bool
	//发送请求
	sendRequest(string, []byte) []byte
	//成功响应
	Success() string
	//失败响应
	Error() string
	//获取响应数据
	ResponseData(map[string]interface{}) map[string]interface{}
	//获取订单类型 1充值/2提现
	OrderType(map[string]interface{}) int
	//获取订单状态
	OrderStatus(map[string]interface{}) bool
	//获取系统订单号
	OrderSn(map[string]interface{}) string
	//获取三方订单号
	TradeSn(map[string]interface{}) string
	//获取实际金额
	RealMoney(map[string]interface{}) float64
	//获取支付时间
	PayTime(map[string]interface{}) int64
}
type PaymentResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data *PaymentData `json:"data"`
}
type PaymentData struct {
	Url string `json:"url"`
}

// 代收充值参数
type RechargeParam struct {
	OrderNo string          //订单号
	Amount  decimal.Decimal //金额
	Symbol  string          //币种
	Channel string          //通道名称
	Other   string          //额外参数
}

// 代付提现参数
type WithdrawParam struct {
	OrderNo     string          //订单号
	Amount      decimal.Decimal //金额
	Symbol      string          //币种
	Channel     string          //通道名称
	Holder      string          //收款人
	Account     string          //收款账号
	AccountType string          //收款账号类型
	Bank        string          //银行名称
	Mobile      string          //手机号
	Other       string          //额外参数
}

func NewPay(payment model.Payment) PayInterface {
	switch payment.ClassName {
	case "WePay":
		return newWePay(payment)
	case "AzfPay":
		return newAzfPay(payment)
	case "JuHePay":
		return newJuHePay(payment)
	default:
		return newWePay(payment)
	}
}
