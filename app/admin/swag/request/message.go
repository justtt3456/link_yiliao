package request

type MessageList struct {
	Status   int `json:"status"`
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
type MessageCreate struct {
	UID     int    `json:"uid"`     //
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
	Status  int    `json:"status"`
}
type MessageUpdate struct {
	ID      int    `json:"id"`
	UID     int    `json:"uid"`     //
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
	Status  int    `json:"status"`
}

type MessageUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
type MessageRemove struct {
	ID int `json:"id"`
}
