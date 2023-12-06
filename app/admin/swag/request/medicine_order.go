package request

type MedicineOrderListRequest struct {
	MedicineName string `form:"medicine_name" json:"medicine_name"` //产品名字
	Username     string `form:"username" json:"username"`           //用户名字
	RealName     string `form:"real_name" json:"real_name"`         //用户名字
	AgentName    string `form:"agent_name" json:"agent_name"`       //用户名字
	Uid          int    `form:"uid" json:"uid"`                     //用户Id
	StartTime    string `form:"start_time" json:"start_time"`
	EndTime      string `form:"end_time" json:"end_time"`
	Page         int    `form:"page" json:"page"`
	PageSize     int    `form:"page_size" json:"page_size"`
}
