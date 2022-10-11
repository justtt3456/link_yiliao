package response

type GuquanResp struct {
	ID              int64   `json:"id"`                //
	TotalGuquan     int64   `json:"total_guquan"`      //总股权数
	OtherGuquan     int64   `json:"other_guquan"`      //剩余
	ReleaseRate     float64 `json:"release_rate"`      //释放百分比
	Price           float64 `json:"price"`             //价格
	LimitBuy        int64   `json:"limit_buy"`         //最低买多少股
	LuckyRate       float64 `json:"lucky_rate"`        //中签率
	ReturnRate      float64 `json:"return_rate"`       //未中签送的 百分比
	ReturnLuckyRate float64 `json:"return_lucky_rate"` //中签回购  百分比
	PreStartTime    int64   `json:"pre_start_time"`    //预售开始时间
	PreEndTime      int64   `json:"pre_end_time"`      //预售结束时间
	OpenTime        int64   `json:"open_time"`         //发行时间
	ReturnTime      int64   `json:"return_time"`       //回收时间
	Status          int64   `json:"status"`            //1 = 开启 2 =关闭
}
