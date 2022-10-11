package response

type MessageListResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data MessageData `json:"data"`
}
type MessageData struct {
	List []MessageInfo `json:"list"`
	Page Page          `json:"page"`
}
type MessageInfo struct {
	ID         int    `json:"id"`      //
	UID        int    `json:"uid"`     //
	Title      string `json:"title"`   //标题
	Content    string `json:"content"` //内容
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"` //创建日期
	UpdateTime int64  `json:"update_time"` //修改时间
}
