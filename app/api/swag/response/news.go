package response

type NewsResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data NewsData `json:"data"`
}
type NewsData struct {
	List []News `json:"list"`
	Page Page   `json:"page"`
}
type News struct {
	Id         int    `json:"id"`          //
	Title      string `json:"title"`       // 标题
	Content    string `json:"content"`     // 内容
	Sort       int    `json:"sort"`        //
	Intro      string `json:"intro"`       //简介
	Cover      string `json:"cover"`       //封面图
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
