package request

import "github.com/shopspring/decimal"

type GuquanReq struct {
	Id              int             `json:"id"`                //
	TotalGuquan     int64           `json:"total_guquan"`      //总股权数
	OtherGuquan     int64           `json:"other_guquan"`      //剩余
	ReleaseRate     decimal.Decimal `json:"release_rate"`      //释放百分比
	Price           decimal.Decimal `json:"price"`             //价格
	LimitBuy        int             `json:"limit_buy"`         //最低买多少股
	LuckyRate       decimal.Decimal `json:"lucky_rate"`        //中签率
	ReturnRate      decimal.Decimal `json:"return_rate"`       //未中签送的 百分比
	ReturnLuckyRate decimal.Decimal `json:"return_lucky_rate"` //中签回购  百分比
	PreStartTime    string          `json:"pre_start_time"`    //预售开始时间
	PreEndTime      string          `json:"pre_end_time"`      //预售结束时间
	OpenTime        string          `json:"open_time"`         //发行时间
	ReturnTime      string          `json:"return_time"`       //回收时间
	Status          int64           `json:"status"`            //1 = 开启 2 =关闭
}
