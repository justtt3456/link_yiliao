package request

type EquityScorePageList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Username string `form:"username"`
	Uid      int    `form:"uid"`
}
