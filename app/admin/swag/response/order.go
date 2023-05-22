package response

import "github.com/shopspring/decimal"

type OrderListResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data OrderListData `json:"data"`
}
type OrderListData struct {
	List []OrderInfo `json:"list"`
	Page Page        `json:"page"`
}

type OrderInfo struct {
	Id            int             `json:"id"`              //
	UId           int             `json:"uid"`             //关联用户id
	Pid           int             `json:"pid"`             //关联商品种类id
	PayMoney      decimal.Decimal `json:"pay_money"`       //购买付款金额(不含手续费)
	Ratio         int             `json:"ratio"`           //对应赔率15=15%!
	PayPrice      decimal.Decimal `json:"pay_price"`       //购买时的产品价格
	Fee           decimal.Decimal `json:"fee"`             //手续费
	PayMoneyTotal decimal.Decimal `json:"pay_money_total"` //购买付款金额(含手续费)
	DrawPrice     decimal.Decimal `json:"draw_price"`      //开奖时的产品价格
	ChooseType    int             `json:"choose_type"`     //买的 1涨，2跌
	DrawResult    int             `json:"draw_result"`     //开奖结果，1中奖2未中
	DrawMoney     decimal.Decimal `json:"draw_money"`      //盈亏结果金额
	CtlType       int             `json:"ctl_type"`        //单控类型 0不控 1赢 2输
	Status        int             `json:"status"`          //状态0为未处理，1为已处理
	AfterBalance  decimal.Decimal `json:"after_balance"`   //购买后余额
	DrawTime      int64           `json:"draw_time"`       //结算开奖时间
	CreateTime    int64           `json:"create_time"`     //创建时间
	UpdateTime    int64           `json:"update_time"`     //系统开奖时间
	ProductName   string          `json:"product_name"`
	Username      string          `json:"username"`
}

type BuyList struct {
	Username string          `json:"username"` //用户名
	Uid      int             `json:"uid"`      //用户Id
	Name     string          `json:"name"`     //产品名字
	Status   int             `json:"status"`   //状态 1=进行中  2=结束
	BuyTime  int             `json:"buy_time"` //投资时间
	Amount   decimal.Decimal `json:"amount"`   //金额
}

type BuyListResp struct {
	List []BuyList `json:"list"`
	Page Page      `json:"page"`
}

type BuyGuquan struct {
	Id         int             `json:"id"`          //id
	Username   string          `json:"username"`    //用户名
	Uid        int             `json:"uid"`         //用户Id
	Num        int64           `json:"num"`         //股权数据量
	Price      decimal.Decimal `json:"price"`       //股权单价
	CreateTime int64           `json:"create_time"` //获得时间
	TotalPrice decimal.Decimal `json:"total_price"` //股权总价值
	Rate       decimal.Decimal `json:"rate"`        //中签率
}
type BuyGuquanResp struct {
	List []BuyGuquan `json:"list"`
	Page Page        `json:"page"`
}
