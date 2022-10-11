package request

type ManualRequest struct {
	UID        int     `json:"uid"`
	Amount     float64 `json:"amount" `
	IsFrontend int     `json:"is_frontend" ` //1=展示  2=不展示
	Handle     int     `json:"handle"`       //操作 1上分可用余额  2下分可用余额  3下分可提现余额
}
type ManualListRequest struct {
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	Username  string `json:"username" form:"username"`
	Type      int    `json:"type" form:"type"` // 1上分 2下分 3冻结 4解冻
	StartTime int    `json:"start_time" form:"start_time"`
	EndTime   int    `json:"end_time" form:"end_time"`
}
