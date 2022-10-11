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
	Lang    string `json:"lang"`    //
}
type NewsUpdate struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`   //
	Content string `json:"content"` //
	Status  int    `json:"status"`  //
	Sort    int    `json:"sort"`    //
	Intro   string `json:"intro"`   //
	Cover   string `json:"cover"`   //封面图
	Lang    string `json:"lang"`    //
}
type NewsUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"` //状态
}
type NewsRemove struct {
	ID int `json:"id"`
}
