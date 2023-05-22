package response

type NoticeResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data NoticeData `json:"data"`
}
type NoticeData struct {
	List []Notice `json:"list"`
	Page Page     `json:"page"`
}
type Notice struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`       //标题
	Intro      string `json:"intro"`       //简介
	Content    string `json:"content"`     //内容
	Lang       string `json:"lang"`        //语言
	CreateTime int64  `json:"create_time"` //
}
type NoticeListResponse struct {
	List []NoticeInfo `json:"list"`
	Page Page         `json:"page"`
}
type NoticeInfo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`       //标题
	Content    string `json:"content"`     //内容
	CreateTime int64  `json:"create_time"` //
}
