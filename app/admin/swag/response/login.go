package response

type LoginResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data AdminInfo `json:"data"`
}
