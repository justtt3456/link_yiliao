package request

type ProductCategoryList struct {
	Status int `form:"status"`
}

type ProductCategoryCreate struct {
	Name   string `json:"name"`   //产品名称
	Status int    `json:"status"` //是否开启，1为开启，0为关闭
	Lang   string `json:"lang"`   //固定传 zh_cn
}
type ProductCategoryUpdate struct {
	Id     int    `json:"id"`     //
	Name   string `json:"name"`   //产品名称
	Status int    `json:"status"` //是否开启，1为开启，0为关闭
	Lang   string `json:"lang"`
}
type ProductCategoryUpdateStatus struct {
	Id     int `json:"id"`
	Status int `json:"status"` //状态
}
type ProductCategoryRemove struct {
	Id int `json:"id"`
}
