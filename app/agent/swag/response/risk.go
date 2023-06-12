package response

type RiskResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data RiskInfo `json:"data"`
}
type RiskInfo struct {
	Id        int    `json:"id"`         //
	WinList   string `json:"win_list"`   //包赢名单，以/分割
	LoseList  string `json:"lose_list"`  //包输名单，以/分割
	WcLine    int    `json:"wc_line"`    //风控启动最小值。下单金额达到次标准则执行风控规则
	WcRatio   string `json:"wc_ratio"`   //风控金额和概率，以/分割，如如100-500:45/501-1000:35
	LoseModel int    `json:"lose_model"` //亏损模式 1百分比亏损 2本金亏损
	LoseTime  string `json:"lose_time"`  //指定时间亏损 以/分割  23:00-08:00/18:00-19:00
	WinTime   string `json:"win_time"`   //指定时间盈利 以/分割  13:00-14:00/15:00-16:00
}
