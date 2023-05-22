package response

type MessageResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data MessageData `json:"data"`
}
type MessageData struct {
	List []Message `json:"list"`
	Page Page      `json:"page"`
}
type Message struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`       //标题
	Content    string `json:"content"`     //内容
	CreateTime int64  `json:"create_time"` //
}
