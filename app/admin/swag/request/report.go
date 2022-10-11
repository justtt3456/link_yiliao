package request

type ReportSum struct {
	Username  string `form:"username"`
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	StartTime string `form:"start_time"` //开始日期
	EndTime   string `form:"end_time"`   //结束日期
}
