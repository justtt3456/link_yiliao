package response

type MemberLevelListResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data MemberLevelData `json:"data"`
}
type MemberLevelData struct {
	List []MemberLevelInfo `json:"list"`
}
type MemberLevelInfo struct {
	Id   int    `json:"id"`   //
	Name string `json:"name"` //等级名称
	Img  string `json:"img"`  //图标
}
