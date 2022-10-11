package pay

import (
	"bytes"
	"encoding/json"
	"finance/common"
	"finance/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WePay struct {
	payment model.Payment
	error   string
}

func newWePay(payment model.Payment) PayInterface {
	return &WePay{
		payment: payment,
	}
}

func (this WePay) Recharge(param RechargeParam) PaymentResponse {
	//参数组装
	s := map[string]interface{}{
		"client_id":    this.payment.MerchantNo,
		"total_fee":    param.Amount,
		"out_trade_no": param.OrderNo,
		"notify_url":   this.payment.NotifyURL,
		"callback_url": "https://ooshanghai.top/",
	}
	//生成签名
	s["sign"] = this.createSign(s)
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
	if res["code"] != float64(0) {
		return PaymentResponse{10001, res["msg"].(string), nil}
	}
	data := res["data"].(map[string]interface{})
	return PaymentResponse{
		Code: 0,
		Msg:  "ok",
		Data: &PaymentData{
			Url: data["pay_url"].(string),
		},
	}
}
func (this WePay) createSign(m map[string]interface{}) string {
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
		paramStr += s + "&"
	}
	paramStr += "key=" + this.payment.Secret
	log.Println("加密前: ", paramStr)
	return strings.ToUpper(common.Md5String(paramStr))
}
func (this WePay) VerifySign(m map[string]interface{}) bool {
	sign := m["sign"].(string)
	delete(m, "sign")
	if sign != this.createSign(m) {
		this.error = "签名错误"
		return false
	}
	return true
}
func (this WePay) sendRequest(url string, data []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
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

func (this WePay) ResponseData(m map[string]interface{}) map[string]interface{} {
	return m
}
func (this WePay) OrderType(m map[string]interface{}) int {
	return 1
}
func (this WePay) OrderStatus(m map[string]interface{}) bool {
	status := m["pay_status"].(float64)
	if status != 1 {
		this.error = "状态错误"
		return false
	}
	return true
}
func (this WePay) OrderSn(m map[string]interface{}) string {
	return m["order_sn"].(string)
}
func (this WePay) TradeSn(m map[string]interface{}) string {
	return ""
}
func (this WePay) RealMoney(m map[string]interface{}) int64 {
	return int64(m["total_fee"].(float64))
}
func (this WePay) PayTime(m map[string]interface{}) int64 {
	return time.Now().Unix()
}
func (this WePay) Success() string {
	return "SUCCESS"
}
func (this WePay) Error() string {
	if this.error == "" {
		this.error = "未知错误"
	}
	return this.error
}
