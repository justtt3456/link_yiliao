package request

type NewsPageListRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
	Category int `form:"category"`
}
