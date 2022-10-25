package request

type OrderListRequest struct {
	ProductName string `form:"product_name" json:"product_name"` //产品名字
	Username    string `form:"username" json:"username"`         //用户名字
	Uid         int    `form:"uid" json:"uid"`                   //用户ID
	StartTime   string `form:"start_time" json:"start_time"`
	EndTime     string `form:"end_time" json:"end_time"`
	Page        int    `form:"page" json:"page"`
	PageSize    int    `form:"page_size" json:"page_size"`
}

type OrderUpdate struct {
	ID      int `json:"id"`
	Rate float64 `json:"rate"` //中签率
}
