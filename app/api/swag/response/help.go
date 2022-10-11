package response

type HelpListResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data HelpData `json:"data"`
}
type HelpData struct {
	List []Help `json:"list"`
}
type Help struct {
	ID         int    `json:"id"`          //
	Title      string `json:"title"`       //
	Content    string `json:"content"`     //
	CreateTime int64  `json:"create_time"` //
}
