package request

type Request struct {
}
type Pagination struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
