package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/app/api/swag/response"
	"china-russia/common"
	"china-russia/extends"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"china-russia/pay"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"

	//"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type RechargeCreate struct {
	request.RechargeCreate
}

func (this RechargeCreate) checkError() error {
	//if this.Amount == 0 {
	//	return errors.New(lang.Lang("Wrong deposit amount"))
	//}
	if this.Method == 0 {
		return errors.New(lang.Lang("Wrong deposit type"))
	}
	funds := model.SetFunds{}
	if !funds.Get() {
		return errors.New(lang.Lang("System configuration error, please contact the administrator"))
	}
	//充值金额和时间
	if this.Amount.LessThan(funds.RechargeMinAmount) {
		return errors.New(fmt.Sprintf(lang.Lang("Minimum deposit %.2f"), funds.RechargeMinAmount.InexactFloat64()))
	}
	if funds.RechargeMaxAmount.LessThan(this.Amount) {
		return errors.New(fmt.Sprintf(lang.Lang("Maximum deposit %.2f"), funds.RechargeMaxAmount.InexactFloat64()))
	}
	now := time.Now().Unix()
	startTime := common.TimeToUnix(funds.RechargeStartTime)
	endTime := common.TimeToUnix(funds.RechargeEndTime)
	if now < startTime {
		return errors.New(fmt.Sprintf(lang.Lang("Please deposit after %s"), funds.RechargeStartTime))
	}
	if now > endTime {
		return errors.New(fmt.Sprintf(lang.Lang("Please deposit before %s"), funds.RechargeEndTime))
	}
	//获取支付方式
	method := model.RechargeMethod{Id: this.Method}
	if !method.Get() {
		return errors.New(lang.Lang("Wrong deposit type"))
	}
	switch method.Code {
	case "bank":
		//if this.From == "" {
		//	return errors.New(lang.Lang("Payment account cannot be empty"))
		//}
		//if this.To == 0 {
		//	return errors.New(lang.Lang("Receiving account cannot be empty"))
		//}
		//充值凭证
		//if this.Voucher == "" {
		//	return errors.New(lang.Lang("The deposit voucher cannot be empty"))
		//}
		//获取收款银行卡
		//bank := model.SetBank{
		//	Id: this.To,
		//}
		//if !bank.Get() {
		//	return errors.New(lang.Lang("Receiving account does not exist"))
		//}
		if this.Voucher == "" {
			return errors.New(lang.Lang("credential image must be required!"))
		}
	case "paymentAlipay":
		//if this.ChannelId == 0 {
		//	return errors.New(lang.Lang("Payment channel cannot be empty"))
		//}
	case "paymentWx":
		//if this.ChannelId == 0 {
		//	return errors.New(lang.Lang("Payment channel cannot be empty"))
		//}
	}

	return nil
}
func (this RechargeCreate) Create(member model.Member, ip string) (*response.RechargeCreate, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}
	base := model.SetBase{}
	base.Get()
	//获取支付方式
	method := model.RechargeMethod{Id: this.Method}
	method.Get()
	switch method.Code {
	case "bank": //银行卡
		bank := model.SetBank{
			Id: this.To,
		}
		bank.Get()
		_, err := this.create(member, bank.CardNumber, decimal.Zero, 0)
		if err != nil {
			return nil, err
		} else {
			return &response.RechargeCreate{}, nil
		}
	case "usdt":
		usdt := model.SetUsdt{
			Id: this.To,
		}
		usdt.Get()
		_, err := this.create(member, usdt.Address, this.Amount, 0)
		if err != nil {
			return nil, err
		} else {
			return &response.RechargeCreate{}, nil
		}
	case "paymentAlipay", "paymentWx":
		//三方支付
		channel := model.PayChannel{MethodId: this.Method, Status: model.StatusOk}
		if !channel.Get() {
			return nil, errors.New(lang.Lang("The payment channel does not exist"))
		}
		//充值金额
		if this.Amount.LessThan(channel.Min) {
			return nil, errors.New(fmt.Sprintf(lang.Lang("Minimum deposit %.2f"), channel.Min.InexactFloat64()))
		}
		if channel.Max.LessThan(this.Amount) {
			return nil, errors.New(fmt.Sprintf(lang.Lang("Maximum recharge %.2f"), channel.Max.InexactFloat64()))
		}
		p := model.Payment{
			Id: channel.PaymentId,
		}
		if !p.Get() {
			return nil, errors.New(lang.Lang("The payment does not exist"))
		}
		order, err := this.create(member, "", decimal.Zero, p.Id)
		if err != nil {
			return nil, err
		}
		//三方支付
		payment := pay.NewPay(p)
		param := pay.RechargeParam{
			OrderNo: order.OrderSn,
			Amount:  this.Amount,
			Channel: channel.Code,
			Other:   ip,
		}
		recharge := payment.Recharge(param)
		if recharge.Code != 0 {
			return nil, errors.New(recharge.Msg)
		}
		return &response.RechargeCreate{
			Code: 0,
			Msg:  "",
			Data: response.RechargeUrl{Url: recharge.Data.Url},
		}, nil
		payOrder := extends.OrderParam{
			BaseParam: extends.BaseParam{
				Url:       strings.TrimSpace(p.RechargeURL) + "/payment",
				Key:       p.Secret,
				AgentNo:   p.MerchantNo,
				Timestamp: time.Now().Unix() * 1000,
			},
			OrderNo: order.OrderSn,
			//Amount:      float64(order.Amount / 100), //分
			Title:       "充值",
			PaymentType: channel.Code,
			NotifyUrl:   p.NotifyURL,
		}
		res, err := extends.OrderXinMeng(payOrder)
		if err != nil {
			return nil, err
		}
		if !res.Success {
			return nil, errors.New("第三方下单失败")
		}
		return &response.RechargeCreate{
			Code: 200,
			Msg:  "成功",
			Data: response.RechargeUrl{
				res.Data,
			},
		}, nil
	default:
		return &response.RechargeCreate{}, errors.New(lang.Lang("Wrong recharge type"))
	}

}
func (this RechargeCreate) create(member model.Member, to string, usdtAmount decimal.Decimal, payment int) (model.Recharge, error) {
	recharge := model.Recharge{
		OrderSn:    common.OrderSn(),
		UId:        member.Id,
		Type:       this.Method,
		Amount:     this.Amount,
		From:       this.From,
		To:         to,
		Voucher:    this.Voucher,
		UsdtAmount: usdtAmount,
		PaymentId:  payment,
		Status:     model.StatusReview,
	}
	err := recharge.Insert()
	return recharge, err
}

type RechargeList struct {
	request.RechargeList
}

func (this RechargeList) PageList(member model.Member) response.RechargeList {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > response.MaxPageSize || this.PageSize < response.MinPageSize {
		this.PageSize = response.DefaultPageSize
	}
	m := model.Recharge{}
	where, args, _ := this.getWhere(member.Id)
	list, page := m.GetPageList(where, args, this.Page, this.PageSize)
	return response.RechargeList{List: this.formatList(list), Page: FormatPage(page)}
}
func (this RechargeList) formatList(lists []model.Recharge) []response.Recharge {
	res := make([]response.Recharge, 0)
	for _, v := range lists {
		i := response.Recharge{
			Id:          v.Id,
			OrderSn:     v.OrderSn,
			Type:        v.Type,
			TypeName:    v.RechargeMethod.Name,
			Amount:      v.Amount,
			RealAmount:  v.RealAmount,
			From:        v.From,
			To:          v.To,
			Voucher:     v.Voucher,
			Status:      v.Status,
			UpdateTime:  v.UpdateTime,
			CreateTime:  v.CreateTime,
			Description: v.Description,
		}
		res = append(res, i)
	}
	return res
}
func (this RechargeList) getWhere(uid int) (string, []interface{}, error) {
	where := map[string]interface{}{
		model.Recharge{}.TableName() + ".uid": uid,
	}
	if this.Status > 0 {
		where[model.Recharge{}.TableName()+".status"] = this.Status
		//where["o.draw_time >"] = time.Now().Unix()
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals, nil
}

type RechargeMethod struct {
}

func (this RechargeMethod) List() *response.RechargeMethodData {

	m := model.RechargeMethod{
		Lang:   global.Language,
		Status: model.StatusOk,
	}
	methods, err := m.List()
	if err != nil {
		return nil
	}
	return &response.RechargeMethodData{
		List: this.formatList(methods),
	}
}
func (this RechargeMethod) formatList(lists []model.RechargeMethod) []response.RechargeMethod {
	res := make([]response.RechargeMethod, 0)
	for _, v := range lists {
		ress := make([]map[string]interface{}, 0)
		if v.Code == "bank" {
			m := model.SetBank{
				Status: model.StatusOk,
			}
			list := m.List(true)
			for _, v := range list {
				item := map[string]interface{}{
					"id":          v.Id,
					"bank_name":   v.BankName,
					"card_number": v.CardNumber,
					"real_name":   v.RealName,
					"branch_bank": v.BranchBank,
				}
				ress = append(ress, item)
			}
		} else if v.Code == "usdt" {
			m := model.SetUsdt{
				Status: model.StatusOk,
			}
			list := m.List(true)
			for _, u := range list {
				item := map[string]interface{}{
					"id":      u.Id,
					"address": u.Address,
					"proto":   u.Proto,
				}
				ress = append(ress, item)
			}
		}
		i := response.RechargeMethod{
			Id:   v.Id,
			Name: v.Name,
			Code: v.Code,
			Icon: v.Icon,
			Info: ress,
		}
		res = append(res, i)
	}
	return res
}

//type RechargeMethodInfo struct {
//	request.RechargeMethodInfo
//}
//
//func (this RechargeMethodInfo) Info() []map[string]interface{} {
//	res := make([]map[string]interface{}, 0)
//	switch this.Code {
//	case "kf": // 客服直充 返回客服充值链接
//		m := model.SetKf{
//			Status: model.StatusOk,
//		}
//		list := m.List(true)
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":         v.Id,
//				"name":       v.Name,
//				"start_time": v.StartTime,
//				"end_time":   v.EndTime,
//				"link":       v.Link,
//			}
//			res = append(res, item)
//		}
//		break
//	case "bank": //银行卡充值 返回收款银行卡信息
//		m := model.SetBank{
//			Status: model.StatusOk,
//		}
//		list := m.List(true)
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":          v.Id,
//				"bank_name":   v.BankName,
//				"card_number": v.CardNumber,
//				"real_name":   v.RealName,
//				"branch_bank": v.BranchBank,
//			}
//			res = append(res, item)
//		}
//		break
//	case "alipay": //支付宝充值 返回收款支付宝信息
//		m := model.SetAlipay{
//			Status: model.StatusOk,
//		}
//		list := m.List(true)
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":        v.Id,
//				"account":   v.Account,
//				"real_name": v.RealName,
//			}
//			res = append(res, item)
//		}
//		break
//	case "usdt": //usdt充值 返回usdt收款信息
//		m := model.SetUsdt{
//			Status: model.StatusOk,
//		}
//		list := m.List(true)
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":      v.Id,
//				"address": v.Address,
//				"proto":   v.Proto,
//			}
//			res = append(res, item)
//		}
//		break
//	case "paymentAlipay": // 三方支付 返回三方支付信息
//		m := model.PayChannel{
//			Status:  model.StatusOk,
//			Payment: model.Payment{Type: 2},
//		}
//		list := m.List()
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":   v.Id,
//				"name": v.Name,
//				"min":  v.Min,
//				"max":  v.Max,
//				"icon": v.Icon,
//			}
//			res = append(res, item)
//		}
//		break
//	case "paymentWx": // 三方支付 返回三方支付信息
//		m := model.PayChannel{
//			Status:  model.StatusOk,
//			Lang:    global.Language,
//			Payment: model.Payment{Type: 1},
//		}
//		list := m.List()
//		for _, v := range list {
//			item := map[string]interface{}{
//				"id":   v.Id,
//				"name": v.Name,
//				"min":  v.Min,
//				"max":  v.Max,
//				"icon": v.Icon,
//			}
//			res = append(res, item)
//		}
//		break
//	default:
//		return nil
//	}
//	return res
//}
