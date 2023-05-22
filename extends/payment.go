package extends

import (
	"china-russia/common"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"sort"
)

type BaseParam struct {
	Url       string `json:"url"`       //网关地址
	Sign      string `json:"sign"`      //MD5签名
	Key       string `json:"key"`       //私钥
	AgentNo   string `json:"agentNo"`   //代理商商户号
	Timestamp int64  `json:"timestamp"` //时间戳，根据当前时间计算出的带毫秒的时间戳
}

type OrderParam struct {
	BaseParam   BaseParam
	OrderNo     string          `json:"orderNo"`     // 订单编号（唯一）
	Amount      decimal.Decimal `json:"amount"`      //订单金额（单位为分）
	Title       string          `json:"title"`       //测试	订单标题
	PaymentType string          `json:"paymentType"` //支付类型。BANK：网银支付；WECHAT：微信支付；ALIPAY：支付宝
	NotifyUrl   string          `json:"notifyUrl"`   //http://www.baidu.com	订单异步通知地址
}

type OrderReturn struct {
	Success bool        `json:"success"` // 接口请求是否成功，成功返回true，不成功返回false
	Code    interface{} `json:"code"`    //接口请求返回码，正常返回200，其他详见code返回值说明
	Msg     string      `json:"msg"`     //接口请求如果出错则为出错的说明，如果返回正确则为空
	Data    string      `json:"data"`    //接口请求返回的内容
}

func OrderXinMeng(order OrderParam) (*OrderReturn, error) {
	param := map[string]string{
		"orderNo":     order.OrderNo,
		"title":       order.Title,
		"paymentType": order.PaymentType,
		"notifyUrl":   order.NotifyUrl,
		"agentNo":     order.BaseParam.AgentNo,
		"timestamp":   fmt.Sprint(order.BaseParam.Timestamp),
		"amount":      order.Amount.String(),
	}
	//签名
	param["sign"] = Sign(param, order.BaseParam.Key)
	jsonParam, _ := json.Marshal(param)
	body, err := common.PostJson(order.BaseParam.Url, jsonParam, nil, nil, 3)
	if err != nil {
		return nil, err
	}
	var res OrderReturn
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Sign(s map[string]string, key string) string {
	var str string
	var keys []string
	for k := range s {

		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		str += v + "=" + s[v] + "&"
	}
	str += "key=" + key
	return common.Md5String(str)
}
