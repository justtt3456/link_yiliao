package controller

import (
	"bytes"
	"china-russia/model"
	"china-russia/pay"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
)

type NotifyController struct{}

func (this NotifyController) Notify(c *gin.Context) {
	payName := c.Param("payment")
	//回调参数
	notify := map[string]interface{}{}
	switch c.Request.Method {
	case "GET":
		for k, v := range c.Request.URL.Query() {
			notify[k] = v[0]
		}
	case "POST":
		c.Request.ParseMultipartForm(2 << 10)
		if len(c.Request.Form) > 0 {
			for k, v := range c.Request.Form {
				notify[k] = v[0]
			}
		} else {
			var bodyBytes []byte
			if c.Request.Body != nil {
				//读取参数
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			//重新赋值用于参数绑定
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			err := json.Unmarshal(bodyBytes, &notify)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
	log.Println("回调参数 : ", notify)
	payment := model.Payment{ClassName: payName}
	if !payment.Get() {
		log.Println(payName + "三方支付不存在")
		return
	}
	p := pay.NewPay(payment)
	if !p.VerifySign(notify) {
		c.String(http.StatusOK, p.Error())
		return
	}
	//获取参数
	data := p.ResponseData(notify)
	//验证状态
	if !p.OrderStatus(data) {
		c.String(http.StatusOK, p.Error())
		return
	}
	//验证类型
	if p.OrderType(data) != 1 {
		c.String(http.StatusOK, p.Error())
		return
	}
	item := model.Recharge{OrderSn: p.OrderSn(notify)}
	if !item.Get() {
		c.String(http.StatusOK, "订单不存在")
		log.Println("订单不存在", item.OrderSn)
		return
	}
	if item.Status != model.StatusReview {
		log.Println("当前状态无法修改")
		c.String(http.StatusOK, "订单状态错误")
		return
	}
	amount := decimal.NewFromFloat(p.RealMoney(data))
	if !item.Amount.Equal(amount) {
		log.Println("金额错误")
		c.String(http.StatusOK, "金额错误")
		return
	}
	member := model.Member{Id: item.UId}
	if !member.Get() {
		log.Println("用户不存在")
		return
	}
	err := this.recharge(member, item)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusOK, "业务错误")
		return
	}
	c.String(http.StatusOK, p.Success())
	return
}

func (this NotifyController) recharge(member model.Member, order model.Recharge) error {
	order.Status = model.StatusAccept
	order.Description = "支付回调"
	//更新状态 说明 操作者
	if err := order.Update("status", "description"); err != nil {
		return err
	}
	//账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: model.TradeTypeRecharge,
		ItemId:    order.Id,
		Amount:    order.Amount,
		Before:    member.Balance,
		After:     member.Balance.Add(order.Amount),
		Desc:      "支付回调",
	}
	err := trade.Insert()
	if err != nil {
		return err
	}
	//上分
	member.Balance = member.Balance.Add(order.Amount)
	member.TotalRecharge = member.TotalRecharge.Add(order.Amount)
	return member.Update("balance")
	//config := model.SetBase{}
	//if !config.Get() {
	//	return member.Update("balance")
	//}
	//if config.AutoUpgrade != 1 {
	//	return member.Update("balance")
	//}
	//等级检查
	//ul := model.UserLevel{}
	//if ul.GetLevelByAmount(user.RechargeAmount) {
	//	user.LevelId = ul.LevelId
	//	user.Update("balance", "level_id", "recharge_amount")
	//} else {
	//	user.Update("balance", "recharge_amount")
	//}
	//return nil
}
