package response

import "github.com/shopspring/decimal"

type ProductListResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ProductData `json:"data"`
}
type ProductData struct {
	List []Product `json:"list"`
	Page Page      `json:"page"`
}
type Product struct {
	Id                    int             `json:"id"`
	Name                  string          `json:"name"`     //产品名称
	Category              int             `json:"category"` //分类id
	CategoryName          string          `json:"category_name"`
	Type                  int             `json:"type"`                    //1=到期返本金 2=延迟反本金
	Price                 decimal.Decimal `json:"price"`                   //价格
	Img                   string          `json:"img"`                     //图片
	Interval              int             `json:"interval"`                //投资期限 （天）
	IncomeRate            decimal.Decimal `json:"income_rate"`             //每日收益率
	LimitBuy              int             `json:"limit_buy"`               //限购数量
	Total                 decimal.Decimal `json:"total"`                   //项目规模
	Current               decimal.Decimal `json:"current"`                 //当前规模
	Desc                  string          `json:"desc"`                    //描述
	DelayTime             int             `json:"delay_time"`              //延迟多少天
	GiftId                int             `json:"gift_id"`                 //赠送产品ID
	WithdrawThresholdRate decimal.Decimal `json:"withdraw_threshold_rate"` //提现额度比例
	IsHot                 int             `json:"is_hot"`                  //是否热门
	IsFinished            int             `json:"is_finished"`             //是否投满
	IsCouponGift          int             `json:"is_coupon_gift"`          //是否赠送优惠券
	Sort                  int             `json:"sort"`                    //排序值
	Status                int             `json:"status"`                  //是否开启，1为开启，2为关闭
	CreateTime            int64           `json:"create_time"`             //创建时间
}

type ProductRemoteListResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data ProductRemoteData `json:"data"`
}
type ProductRemoteData struct {
	List []ProductRemote `json:"list"`
}
type ProductRemote struct {
	Name string `json:"name"` //产品名称
	Code string `json:"code"` //产品代码
}

type ProductGiftOptions struct {
	List []ProductGiftInfo `json:"list"`
}
type ProductGiftInfo struct {
	Id   int    `json:"id"`   //产品Id
	Name string `json:"name"` //产品名称
}
