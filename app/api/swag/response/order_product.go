package response

import "github.com/shopspring/decimal"

type OrderResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data OrderList `json:"data"`
}
type OrderList struct {
	List []Order `json:"list"`
	Page Page    `json:"page"`
}
type Order struct {
	Id          int             `json:"oid"`          //订单id
	PId         int             `json:"pid"`          //产品id
	ProductName string          `json:"product_name"` //产品名称
	ChooseType  int             `json:"choose_type"`  //购买类型
	PayMoney    decimal.Decimal `json:"pay_money"`    //购买金额
	DrawMoney   decimal.Decimal `json:"draw_money"`   //结算金额
	PayPrice    decimal.Decimal `json:"pay_price"`    //购买价格
	DrawPrice   decimal.Decimal `json:"draw_price"`   //结算价格
	Wave        decimal.Decimal `json:"wave"`         //浮动
	CreateTime  int64           `json:"create_time"`
	DrawTime    int64           `json:"draw_time"`
	Seconds     int             `json:"seconds"` //剩余秒数
}
