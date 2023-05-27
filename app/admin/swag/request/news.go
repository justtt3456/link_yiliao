package request

type NewsList struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   int    `form:"status"`
	Lang     string `form:"lang"`
}

type NewsCreate struct {
	Title   string `json:"title"`   //
	Content string `json:"content"` //
	Status  int    `json:"status"`  //
	Sort    int    `json:"sort"`    //
	Intro   string `json:"intro"`   //
	Cover   string `json:"cover"`   //封面图
}
type NewsUpdate struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`   //
	Content string `json:"content"` //
	Status  int    `json:"status"`  //
	Sort    int    `json:"sort"`    //
	Intro   string `json:"intro"`   //
	Cover   string `json:"cover"`   //封面图
}
type NewsUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态
}
type NewsRemove struct {
	Id int `json:"id"`
}
