package v1

import (
	"finance/app/api/controller"
	"finance/extends"
	"finance/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NotifyController struct {
	controller.Controller
}

type XinMenReq struct {
	OrderNo   string  `json:"orderNo" form:"orderNo"`     //订单编号
	TradeNo   string  `json:"tradeNo" form:"tradeNo"`     //平台流水号
	Amount    float64 `json:"amount" form:"amount"`       //金额（单位分）
	PayTime   string  `json:"payTime" form:"payTime"`     //支付时间 2018-01-01 10:10:10
	ActType   int     `json:"actType" form:"actType"`     //回调类型（1、支付；2、代付）
	Rate      int     `json:"rate" form:"rate"`           //手续费（单位分）
	Status    string  `json:"status" form:"status"`       //支付结果
	Timestamp int64   `json:"timestamp" form:"timestamp"` //时间戳
	Sign      string  `json:"sign" form:"sign"`           //签名
}

func (this NotifyController) NotifyXinMeng(c *gin.Context) {
	s := XinMenReq{}
	if err := c.ShouldBind(&s); err != nil {
		logrus.Errorf("接收参数失败%v", err)
		return
	}
	logrus.Infof("接收参数%v", s)
	if s.OrderNo == "" {
		logrus.Errorf("接收参数失败")
		return
	}
	//查询订单是否存在
	o := model.Recharge{OrderSn: s.OrderNo}
	if !o.Get() {
		logrus.Errorf("订单不存在%v", s.OrderNo)
		return
	}
	//查询用户是否存在
	m := model.Member{ID: o.UID}
	if !m.Get() {
		logrus.Errorf("用户不存在%v", o.UID)
		return
	}

	//获取私钥
	p := model.Payment{ID: o.PaymentID}
	if !p.Get() {
		logrus.Errorf("支付通道不存在%v", o.PaymentID)
		return
	}
	//验证签名
	param := map[string]string{
		"orderNo":   s.OrderNo,
		"tradeNo":   s.TradeNo,
		"amount":    fmt.Sprint(s.Amount),
		"payTime":   s.PayTime,
		"actType":   fmt.Sprint(s.ActType),
		"rate":      fmt.Sprint(s.Rate),
		"status":    s.Status,
		"timestamp": fmt.Sprint(s.Timestamp),
	}
	if s.Sign != extends.Sign(param, p.Secret) {
		logrus.Errorf("验证签名失败%v  sign=%v   mysign = %v ", param, s.Sign, extends.Sign(param, p.Secret))
		return
	}

	//处理充值
	if s.Status != "SUCCESS" {
		logrus.Errorf("充值失败状态%v", s.Status)
		return
	}
	o.Status = 2
	o.UpdateTime = s.Timestamp
	o.SuccessTime = s.Timestamp
	err := o.Update("status", "update_time", "success_time")
	if err != nil {
		logrus.Errorf("修改订单状态失败%v", err)
		return
	}

	m.Balance += int64(s.Amount * 100)
	m.TotalBalance += int64(s.Amount * 100)

	err = m.Update("balance", "total_balance")
	if err != nil {
		logrus.Errorf("用户价钱失败%v", err)
		return
	}
	c.String(http.StatusOK, "SUCCESS")
	return
}
