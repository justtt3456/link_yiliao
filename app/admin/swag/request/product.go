package request

import "github.com/shopspring/decimal"

type ProductList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Name     string `form:"name"`
	Category int    `form:"category"`
}

type ProductCreate struct {
	Id                    int             `json:"id"`
	Name                  string          `json:"name"`                    //产品名称
	Category              int             `json:"category"`                //分类id
	Type                  int             `json:"type"`                    //1=到期返本金 2=延迟反本金
	Price                 decimal.Decimal `json:"price"`                   //价格
	Img                   string          `json:"img"`                     //图片
	Interval              int             `json:"interval"`                //投资期限 （天）
	IncomeRate            decimal.Decimal `json:"income_rate"`             //每日收益率
	LimitBuy              int             `json:"limit_buy"`               //限购数量
	Total                 decimal.Decimal `json:"total"`                   //项目规模
	Current               decimal.Decimal `json:"current"`                 //当前规模
	Desc                  string          `json:"desc"`                    //描述
	Progress              int             `json:"progress"`                //项目进度
	DelayTime             int             `json:"delay_time"`              //延迟多少天
	GiftId                int             `json:"gift_id"`                 //赠送产品ID
	WithdrawThresholdRate decimal.Decimal `json:"withdraw_threshold_rate"` //提现额度比例
	IsHot                 int             `json:"is_hot"`                  //是否热门
	IsFinished            int             `gorm:"column:is_finished"`      //是否已满
	IsCouponGift          int             `gorm:"column:is_coupon_gift"`   //是否赠送优惠券
	Sort                  int             `json:"sort"`                    //排序值
	Status                int             `json:"status"`                  //是否开启，1为开启，2为关闭
}
type ProductUpdate struct {
	Id                    int             `json:"id"`
	Name                  string          `json:"name"`                    //产品名称
	Category              int             `json:"category"`                //分类id
	Type                  int             `json:"type"`                    //1=到期返本金 2=延迟反本金
	Price                 decimal.Decimal `json:"price"`                   //价格
	Img                   string          `json:"img"`                     //图片
	Interval              int             `json:"interval"`                //投资期限 （天）
	IncomeRate            decimal.Decimal `json:"income_rate"`             //每日收益率
	LimitBuy              int             `json:"limit_buy"`               //限购数量
	Total                 decimal.Decimal `json:"total"`                   //项目规模
	Current               decimal.Decimal `json:"current"`                 //当前规模
	Desc                  string          `json:"desc"`                    //描述
	Progress              int             `json:"progress"`                //项目进度
	DelayTime             int             `json:"delay_time"`              //延迟多少天
	GiftId                int             `json:"gift_id"`                 //赠送产品ID
	WithdrawThresholdRate decimal.Decimal `json:"withdraw_threshold_rate"` //提现额度比例
	IsHot                 int             `json:"is_hot"`                  //是否热门
	IsFinished            int             `gorm:"column:is_finished"`      //是否已满
	IsCouponGift          int             `gorm:"column:is_coupon_gift"`   //是否赠送优惠券
	Sort                  int             `json:"sort"`                    //排序值
	Status                int             `json:"status"`                  //是否开启，1为开启，2为关闭
}
type ProductUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态
}
type ProductRemove struct {
	Id int `json:"id"`
}

type GiftProductOptions struct {
}
