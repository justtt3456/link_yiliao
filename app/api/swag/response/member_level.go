package response

type MemberLevelResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []MemberLevel `json:"list"`
	}
}
type MemberLevel struct {
	ID   int    `json:"id"`   //
	Name string `json:"name"` //
	Img  string `json:"img"`
}
