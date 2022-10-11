package response

type ManualListResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ManualData `json:"data"`
}
type ManualData struct {
	List []ManualInfo `json:"list"`
	Page Page         `json:"page"`
}
type ManualInfo struct {
	ID         int     `json:"id"`       //
	UserID     int     `json:"user_id"`  //
	Username   string  `json:"username"` //
	Type       int     `json:"type"`
	Amount     float64 `json:"amount"`      //金额
	AdminName  string  `json:"admin_name"`  //操作人
	AgentName  string  `json:"agent_name"`  //操作人
	CreateTime int64   `json:"create_time"` //创建时间
}
