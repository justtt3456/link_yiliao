package response

type Banner struct {
	Image string `json:"image"` //图片
	Link  string `json:"link"`  //链接
	Type  int    `json:"type"`  //1=图片  2=视频
}
