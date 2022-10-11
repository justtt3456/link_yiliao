package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST = "imap.163.com"
	// 端口
	SMTP_MAIL_PORT = "465"
	// 发送邮件用户账号
	SMTP_MAIL_USER = "zhongmeishengtai@163.com"
	// 授权密码
	SMTP_MAIL_PWD = "UJYLHEQLBOARODJS"
	// 发送邮件昵称
	SMTP_MAIL_NICKNAME = "test"

	//短信地址
	ADDRESS = "http://api.smsbao.com/wsms"
	//账号
	ACCOUNT = "hsj16899"
	//密码
	PASSWORD = "9e0421a5ad2d1c74554beabab773cdc5"
)

func File_get_contents(url string, t int) (s []byte, err error) {
	//读取文件
	if t == 1 {
		s, err = ioutil.ReadFile(url)
		return
	}
	//读取网络文件
	if t == 2 {
		res, err := http.Get(url)
		if err != nil {
			return []byte(""), err
		}
		s, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		return s, err
	}
	err = errors.New("参数错误")
	return

}

type BaiduIP struct {
	Srcid      string
	ResultCode string
	status     string
	QueryID    string
	Data       []Data `json:"data"`
}
type Data struct {
	Location string `json:"location"`
}

func main() {
	ip := "193.37.32.5"
	url := "https://sp1.baidu.com/8aQDcjqpAAV3otqbppnN2DJv/api.php?query=" + ip + "&co=&resource_id=5809&t=1655103416828&ie=utf8&oe=gbk&cb=op_aladdin_callback&format=json&tn=baidu&cb=jQuery110207088687929741351_1655099545142&_=1655099545306"
	s, _ := File_get_contents(url, 2)
	a := strings.Split(string(s), "(")
	if len(a) == 2 {
		res := a[1][:len(a[1])-1]
		r := BaiduIP{}
		err := json.Unmarshal([]byte(res), &r)
		if err != nil {

		} else {
			if len(r.Data) > 0 {
				fmt.Println(r.Data[0].Location)
			} else {
			}
		}
	}
	//a := 1083.5
	//amount :=decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(100.5))
	//c,_ := amount.Div(decimal.NewFromFloat(100)).Float64()

	//var statustr = map[string]string{
	//	"0":  "SMS sent successfully",
	//	"-1": "Incomplete parameters",
	//	"-2": "The server space is not supported, please confirm that curl or fsocket is supported, and contact your space provider to solve or replace the space!",
	//	"30": "wrong password",
	//	"40": "Account does not exist",
	//	"41": "Insufficient balance",
	//	"42": "Account expired",
	//	"43": "IP address restrictions",
	//	"50": "Content contains sensitive words",
	//	"51": "Wrong format of phone number",
	//}
	//code := "123456"
	//sign := "【ZM】 Your verification code is {%v}"
	//c := fmt.Sprintf(sign, code)
	//u := url.Values{}
	//u.Set("u", ACCOUNT)
	//u.Set("p", PASSWORD)
	//u.Set("m", "+6113666663456")
	//u.Set("c", c)
	//urls := ADDRESS + "?" + u.Encode()
	//fmt.Println(urls)
	//res := SendRequest(urls)
	//fmt.Println(statustr[string(res)])
	//声明err, subject,body类型，并为address，auth以及contentType赋值，
	//subeject是主题，body是邮件内容, address是收件人

	//var err error
	//var subject, body string
	//subject = "test1"
	//body = "test222"
	//address := "258963123@protonmail.com"
	//auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	//contentType := "Content-Type: text/html; charset=UTF-8"
	//
	////要发送的消息，可以直接写在[]bytes里，但是看着太乱，因此使用格式化
	//s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
	//	address, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
	//msg := []byte(s)
	//
	////邮件服务地址格式是"host:port"，因此把addr格式化为这个格式，直接赋值也行。
	//addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
	//
	////发送邮件
	//err = smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{address}, msg)
	//
	//if err != nil {
	//	fmt.Println(13213,err)
	//} else {
	//	fmt.Println("send email succeed")
	//}
}

func SendRequest(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("token", this.apiKey)
	response, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return responseData
}
