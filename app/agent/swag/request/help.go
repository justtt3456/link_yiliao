package request

type HelpList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
	Category int    `form:"category"`
}

type HelpCreate struct {
	Title    string `json:"title"`   //
	Content  string `json:"content"` //
	Lang     string `json:"lang"`    //
	Sort     int    `json:"sort"`
	Status   int    `json:"status"` //
	Category int    `json:"category"`
}
type HelpUpdate struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`   //
	Content  string `json:"content"` //
	Lang     string `json:"lang"`    //
	Sort     int    `json:"sort"`
	Status   int    `json:"status"` //
	Category int    `json:"category"`
}
type HelpUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态
}
type HelpRemove struct {
	Id int `json:"id"`
}
