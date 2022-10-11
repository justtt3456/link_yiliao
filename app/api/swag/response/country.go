package response

type CountryListResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data CountryData `json:"data"`
}
type CountryData struct {
	List []Country `json:"list"`
}
type Country struct {
	Name string `json:"zh_name"` //
	Code string `json:"code"`    //
}
