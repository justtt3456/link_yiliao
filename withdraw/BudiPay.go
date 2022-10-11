package withdraw

import (
	"bytes"
	"finance/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type BudiPay struct {
	payment model.Payment
	error   string
}

func newBudiPay(payment model.Payment) PayInterface {
	return &BudiPay{
		payment: payment,
	}
}

func (this BudiPay) Withdraw(param WithdrawParam) PaymentResponse {
	return PaymentResponse{
		Code: 0,
		Msg:  "",
		Data: PaymentData{},
	}
}
func (this BudiPay) createSign(m map[string]interface{}) string {
	return ""
}
func (this BudiPay) VerifySign(m map[string]interface{}) bool {
	this.error = "签名错误"
	return true
}
func (this BudiPay) sendRequest(url string, data []byte) []byte {
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

func (this BudiPay) ResponseData(m map[string]interface{}) map[string]interface{} {
	return m
}
func (this BudiPay) OrderType(m map[string]interface{}) int {
	this.error = "类型错误"
	return 1
}
func (this BudiPay) OrderStatus(m map[string]interface{}) bool {
	this.error = "状态错误"
	return true
}
func (this BudiPay) OrderSn(m map[string]interface{}) string {
	return ""
}
func (this BudiPay) TradeSn(m map[string]interface{}) string {
	return ""
}
func (this BudiPay) RealMoney(m map[string]interface{}) int64 {
	return 0
}
func (this BudiPay) PayTime(m map[string]interface{}) int64 {
	return time.Now().Unix()
}
func (this BudiPay) Success() string {
	return "SUCCESS"
}
func (this BudiPay) Error() string {
	if this.error == "" {
		this.error = "未知错误"
	}
	return this.error
}

//
//// 加密
//func (this BudiPay) rsaEncrypt(origData string, privateKey string) string {
//	priKey := `-----BEGIN PRIVATE KEY-----
//` +
//		privateKey +
//		`
//-----END PRIVATE KEY-----`
//	encryptString, err := gorsa.PriKeyEncrypt(origData, priKey)
//	if err != nil {
//		return ""
//	}
//
//	return encryptString
//}
//
//// 解密
//func (this BudiPay) rsaDecrypt(ciphertext string, publicKey string) (string, error) {
//	pubKey := `-----BEGIN PUBLIC KEY-----
//` +
//		publicKey +
//		`
//-----END PUBLIC KEY-----`
//	decryptString, err := gorsa.PublicDecrypt(ciphertext, pubKey)
//	if err != nil {
//		return "", err
//	}
//	fmt.Println("解密", decryptString)
//	return decryptString, nil
//}
