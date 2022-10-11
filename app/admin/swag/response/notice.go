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
	ID         int    `json:"id"`          //
	Title      string `json:"title"`       //标题
	Intro      string `json:"intro"`       //简介
	Content    string `json:"content"`     //内容
	Type       int    `json:"type"`        //类型 1滚动 2弹窗
	Lang       string `json:"lang"`        //语言
	Status     int    `json:"status"`      //
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
