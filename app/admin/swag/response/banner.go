package response

type BannerListResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data BannerData `json:"data"`
}
type BannerData struct {
	List []BannerInfo `json:"list"`
	Page Page         `json:"page"`
}
type BannerInfo struct {
	Id         int    `json:"id"`
	Lang       string `json:"lang"`
	Image      string `json:"image"`
	Link       string `json:"link"`
	Sort       int    `json:"sort"` //排序
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Type       int    `json:"type"` //1=图片 2=视频
}
