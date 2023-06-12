package response

type ProductCategoryListResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data ProductCategoryData `json:"data"`
}
type ProductCategoryData struct {
	List []ProductCategory `json:"list"`
}
type ProductCategory struct {
	Id         int    `json:"id"` //
	RemoteId   int    `json:"remote_id"`
	Name       string `json:"name"` //产品名称
	Lang       string `json:"lang"`
	Status     int    `json:"status"`      //是否开启，1为开启，0为关闭
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"`
}

type ProductCategoryRemoteListResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data ProductCategoryData `json:"data"`
}
type ProductCategoryRemoteData struct {
	List []ProductCategoryRemote `json:"list"`
}
type ProductCategoryRemote struct {
	Id   int    `json:"id"`   //
	Name string `json:"name"` //产品名称
}
