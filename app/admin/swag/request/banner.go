package request

type BannerList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
}

type BannerCreate struct {
	Lang  string `json:"lang"`
	Image string `json:"image"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
	Type  int    `json:"type"` //1=图片 2=视频
}
type BannerUpdate struct {
	Id    int    `json:"id"`
	Lang  string `json:"lang"`
	Image string `json:"image"`
	Link  string `json:"link"`
	Sort  int    `json:"sort"`
	Type  int    `json:"type"` //1=图片 2=视频
}
type BannerUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态 1=开启 2=关闭
}
type BannerRemove struct {
	Id int `json:"id"`
}
