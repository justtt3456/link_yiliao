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

type YiPay struct {
	payment model.Payment
	error   string
}

func newYiPay(payment model.Payment) PayInterface {
	return &YiPay{
		payment: payment,
	}
}

func (this YiPay) Recharge(param RechargeParam) PaymentResponse {
	//参数组装
	amount := param.Amount.String()
	s := map[string]interface{}{
		"partnerid": this.payment.MerchantNo,
		"amount":    amount,
		//"pay_bankcode": param.Channel,
		"orderid":   param.OrderNo,
		"notifyurl": this.payment.NotifyURL,
	}
	//生成签名
	s["sign"] = this.createSign(s)
	//s["pay_ip"] = param.Other
	//log.Println("用户下单ip: ", s["pay_ip"])
	log.Println("加密后: ", s["sign"])
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
	if res["status"] != float64(200) {
		return PaymentResponse{10001, res["msg"].(string), nil}
	}
	data := res["url"].(string)
	return PaymentResponse{
		Code: 0,
		Msg:  "ok",
		Data: &PaymentData{
			Url: data,
		},
	}
}
func (this YiPay) createSign(m map[string]interface{}) string {
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
		if s == "" {
			continue
		}
		paramStr += k + "=" + s + "&"
	}
	paramStr += "key=" + this.payment.Secret
	log.Println("加密前: ", paramStr)
	return strings.ToLower(common.Md5String(paramStr))
}
func (this YiPay) VerifySign(m map[string]interface{}) bool {
	sign := m["sign"].(string)
	delete(m, "sign")
	if sign != this.createSign(m) {
		this.error = "签名错误"
		return false
	}
	return true
}
func (this YiPay) sendRequest(uri string, data []byte) []byte {
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

func (this YiPay) ResponseData(m map[string]interface{}) map[string]interface{} {
	return m
}
func (this YiPay) OrderType(m map[string]interface{}) int {
	return 1
}
func (this YiPay) OrderStatus(m map[string]interface{}) bool {
	status := m["status"].(string)
	if status != "1" {
		this.error = "状态错误"
		return false
	}
	return true
}
func (this YiPay) OrderSn(m map[string]interface{}) string {
	return m["orderid"].(string)
}
func (this YiPay) TradeSn(m map[string]interface{}) string {
	return m["transaction_id"].(string)
}
func (this YiPay) RealMoney(m map[string]interface{}) float64 {
	float, _ := strconv.ParseFloat(m["amount"].(string), 10)
	return float
}
func (this YiPay) PayTime(m map[string]interface{}) int64 {
	return time.Now().Unix()
}
func (this YiPay) Success() string {
	return `{"code":200,"msg":"ok"}`
}
func (this YiPay) Error() string {
	if this.error == "" {
		this.error = "未知错误"
	}
	return this.error
}
