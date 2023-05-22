package request

type MessageList struct {
	Status   int `json:"status"`
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
type MessageCreate struct {
	UId     int    `json:"uid"`     //
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
	Status  int    `json:"status"`
}
type MessageUpdate struct {
	Id      int    `json:"id"`
	UId     int    `json:"uid"`     //
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
	Status  int    `json:"status"`
}

type MessageUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}
type MessageRemove struct {
	Id int `json:"id"`
}
