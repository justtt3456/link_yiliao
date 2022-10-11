package common

type PageInterface interface {
	SetPage(int, int64)
}

func NewPage(page int) PageInterface {
	return &Page{
		Page: page,
	}
}

type Page struct {
	Page     int `json:"page"`      //当前页
	PageSize int `json:"page_size"` //每页数量
	Record   int `json:"record"`    //总记录数
	Total    int `json:"total"`     //总页数
}

const (
	MinPageSize     = 5
	MaxPageSize     = 100
	DefaultPageSize = 15
)

func (this *Page) SetPage(pageSize int, totalRecord int64) {
	if this.Page == 0 {
		return
	}
	this.PageSize = pageSize
	total := int(totalRecord)
	this.Record = total
	if total%pageSize == 0 {
		this.Total = total / pageSize
	} else {
		this.Total = total/pageSize + 1
	}
	return
}
