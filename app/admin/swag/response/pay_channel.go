package response

type PayChannelListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data PayChannelData `json:"data"`
}
type PayChannelData struct {
	List []PayChannel `json:"list"`
	Page Page         `json:"page"`
}
type PayChannel struct {
	ID          int     `json:"id"`           //
	Name        string  `json:"name"`         //支付方式名称
	PaymentID   int     `json:"payment_id"`   //第三方id
	PaymentName string  `json:"payment_name"` //第三方名称
	Code        string  `json:"code"`         //支付编码
	Min         float64 `json:"min"`          //最小值
	Max         float64 `json:"max"`          //最大值
	Status      int     `json:"status"`       //状态
	Category    int     `json:"category"`     //分类(所属支付方式)
	Sort        int     `json:"sort"`         //排序值
	Icon        string  `json:"icon"`         //图标
	Fee         int     `json:"fee"`          //手续费
	Lang        string  `json:"lang"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}
