package request

type RechargeRequest struct {
	Id          int    `json:"id"`
	UId         int    `json:"uid" form:"uid"`               //用户id
	OrderSn     string `json:"order_sn" form:"order_sn"`     //订单号
	Username    string `json:"username" form:"username"`     //用户名
	StartTime   string `json:"start_time" form:"start_time"` //开始时间
	EndTime     string `json:"end_time" form:"end_time"`     //结束时间
	Page        int    `json:"page" form:"page"`             //页码
	PageSize    int    `json:"page_size" form:"page_size"`   //条数
	Status      int    `json:"status"  form:"status"`        //1待审核  2已审核 3驳回
	Description string `json:"description"`                  //备注
}
type RechargeListRequest struct {
	OrderSn   string `json:"order_sn"`
	Username  string `json:"username"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	Status    int    `json:"status"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}
type RechargeUpdateRequest struct {
	Ids         string `json:"ids"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
