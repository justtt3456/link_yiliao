package request

type NoticeList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
}

type NoticeCreate struct {
	Title   string `json:"title"`   //标题
	Intro   string `json:"intro"`   //简介
	Content string `json:"content"` //内容
	Type    int    `json:"type"`    //类型 1滚动 2弹窗
	Lang    string `json:"lang"`    //语言
	Status  int    `json:"status"`  //
}
type NoticeUpdate struct {
	ID      int    `json:"id"`      //
	Title   string `json:"title"`   //标题
	Intro   string `json:"intro"`   //简介
	Content string `json:"content"` //内容
	Type    int    `json:"type"`    //类型 1滚动 2弹窗
	Lang    string `json:"lang"`    //语言
	Status  int    `json:"status"`  //
}
type NoticeUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"` //状态
}
type NoticeRemove struct {
	ID int `json:"id"`
}
