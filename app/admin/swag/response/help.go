package response

type HelpListResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data HelpData `json:"data"`
}
type HelpData struct {
	List []Help `json:"list"`
	Page Page   `json:"page"`
}
type Help struct {
	ID         int    `json:"id"`      //
	Title      string `json:"title"`   //
	Content    string `json:"content"` //
	Lang       string `json:"lang"`    //
	Sort       int    `json:"sort"`
	Status     int    `json:"status"` //
	Category   int    `json:"category"`
	CreateTime int64  `json:"create_time"` //
	UpdateTime int64  `json:"update_time"` //
}
