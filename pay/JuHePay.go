package pay

import (
	"china-russia/common"
	"china-russia/model"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type JuHePay struct {
	payment model.Payment
	error   string
}

func newJuHePay(payment model.Payment) PayInterface {
	return &JuHePay{
		payment: payment,
	}
}

func (this JuHePay) Recharge(param RechargeParam) PaymentResponse {
	//参数组装
	amount := param.Amount.InexactFloat64()
	s := map[string]interface{}{
		"pay_memberid":    this.payment.MerchantNo,
		"pay_amount":      amount,
		"pay_bankcode":    param.Channel,
		"pay_orderid":     param.OrderNo,
		"pay_notifyurl":   this.payment.NotifyURL,
		"pay_callbackurl": this.payment.NotifyURL,
	}
	//生成签名
	s["pay_md5sign"] = this.createSign(s)
	log.Println("用户下单ip: ", s["pay_ip"])
	log.Println("加密后: ", s["pay_md5sign"])
	j, _ := json.Marshal(s)
	//发送请求
	resp := this.sendRequest(this.payment.RechargeURL, j)
	log.Println("三方返回：", string(resp))
	//接收请求
	res := map[string]interface{}{}
	err := json.Unmarshal(resp, &res)
	if err != nil {
		fmt.Println(err)
		return PaymentResponse{10001, err.Error(), nil}
	}
	if res["code"] != float64(200) {
		return PaymentResponse{10001, res["msg"].(string), nil}
	}
	data := res["data"].(string)
	return PaymentResponse{
		Code: 0,
		Msg:  "ok",
		Data: &PaymentData{
			Url: data,
		},
	}
}
func (this JuHePay) createSign(m map[string]interface{}) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramStr string
	for _, k := range keys {
		var s string
		switch m[k].(type) {
		case string:
			s = m[k].(string)
			break
		case int:
			s = strconv.Itoa(m[k].(int))
			break
		case float64:
			s = strconv.Itoa(int(m[k].(float64)))
		default:
			s = ""
		}
		paramStr += k + "=" + s + "&"
	}
	paramStr += "key=" + this.payment.Secret
	log.Println("加密前: ", paramStr)
	return strings.ToUpper(common.Md5String(paramStr))
}
func (this JuHePay) VerifySign(m map[string]interface{}) bool {
	sign := m["sign"].(string)
	delete(m, "sign")
	if sign != this.createSign(m) {
		this.error = "签名错误"
		return false
	}
	return true
}
func (this JuHePay) sendRequest(uri string, data []byte) []byte {
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	paramStr := url.Values{}
	for k, v := range m {
		var s string
		switch v.(type) {
		case string:
			s = v.(string)
			break
		case int:
			s = strconv.Itoa(v.(int))
			break
		case float64:
			s = strconv.Itoa(int(v.(float64)))
		default:
			s = ""
		}
		paramStr.Add(k, s)
	}

	fmt.Println(paramStr)
	resp, err := http.PostForm(uri, paramStr)

	if err != nil {
		// handle error
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return responseData
}

func (this JuHePay) ResponseData(m map[string]interface{}) map[string]interface{} {
	return m
}
func (this JuHePay) OrderType(m map[string]interface{}) int {
	return 1
}
func (this JuHePay) OrderStatus(m map[string]interface{}) bool {
	status := m["returncode"].(string)
	if status != "00" {
		this.error = "状态错误"
		return false
	}
	return true
}
func (this JuHePay) OrderSn(m map[string]interface{}) string {
	return m["orderid"].(string)
}
func (this JuHePay) TradeSn(m map[string]interface{}) string {
	return m["transaction_id"].(string)
}
func (this JuHePay) RealMoney(m map[string]interface{}) float64 {
	float, _ := strconv.ParseFloat(m["amount"].(string), 10)
	return float
}
func (this JuHePay) PayTime(m map[string]interface{}) int64 {
	return time.Now().Unix()
}
func (this JuHePay) Success() string {
	return "OK"
}
func (this JuHePay) Error() string {
	if this.error == "" {
		this.error = "未知错误"
	}
	return this.error
}
