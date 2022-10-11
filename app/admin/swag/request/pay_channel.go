package request

type PayChannelListRequest struct {
	Name      string `form:"name" json:"name"`
	PaymentID int    `form:"payment_id" json:"payment_id"`
	Page      int    `form:"page" json:"page"`
	PageSize  int    `form:"page_size" json:"page_size"`
}
type PayChannelCreateRequest struct {
	Name      string  `json:"name"`       //支付方式名称
	PaymentID int     `json:"payment_id"` //第三方名称
	Code      string  `json:"code"`       //支付编码
	Min       float64 `json:"min"`        //最小值
	Max       float64 `json:"max"`        //最大值
	Status    int     `json:"status"`     //状态
	Category  int     `json:"category"`   //分类(所属支付方式)
	Sort      int     `json:"sort"`       //排序值
	Icon      string  `json:"icon"`       //图标
	Fee       int     `json:"fee"`        //手续费
	Lang      string  `json:"lang"`
}
type PayChannelUpdateRequest struct {
	ID        int     `json:"id"`         //
	Name      string  `json:"name"`       //支付方式名称
	PaymentID int     `json:"payment_id"` //第三方名称
	Code      string  `json:"code"`       //支付编码
	Min       float64 `json:"min"`        //最小值
	Max       float64 `json:"max"`        //最大值
	Status    int     `json:"status"`     //状态
	Category  int     `json:"category"`   //分类(所属支付方式)
	Sort      int     `json:"sort"`       //排序值
	Icon      string  `json:"icon"`       //图标
	Fee       int     `json:"fee"`        //手续费
	Lang      string  `json:"lang"`
}
type PayChannelUpdateStatusRequest struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
type PayChannelRemoveRequest struct {
	ID int `json:"id"`
}
