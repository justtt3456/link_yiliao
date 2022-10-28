package response

type ProductListResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ProductData `json:"data"`
}
type ProductData struct {
	List []Product `json:"list"`
	Page Page      `json:"page"`
}
type Product struct {
	ID           int     `json:"id"`             //
	Name         string  `json:"name"`           //产品名称
	Category     int     `json:"category"`       //类别
	CategoryName string  `json:"category_name"`  //类别名称
	CreateTime   int64   `json:"create_time"`    //创建时间
	Status       int     `json:"status"`         //是否开启，1为开启，0为关闭
	Tag          int     `json:"tag"`            //1=热 0=无
	TimeLimit    int     `json:"time_limit"`     //投资期限 （天）
	IsRecommend  int     `json:"is_recommend"`   //是否推荐到首页 1是 2否
	Dayincome    float64 `json:"dayincome"`      //每日收益  千分比
	Price        float64 `json:"price"`          //价格  (最低买多少)
	TotalPrice   float64 `json:"total_price"`    //项目规模
	OtherPrice   float64 `json:"other_price"`    //可投余额
	MoreBuy      int     `json:"more_buy"`       //最多可以买多少份
	Desc         string  `json:"desc"`           //描述
	IsFinish     int     `json:"is_finish"`      //1=进行中  2=已投满
	IsManjian    int     `json:"is_manjian"`     //1=有满减  2=无满减
	BuyTimeLimit int     `json:"buy_time_limit"` //产品限时多少天
	Progress     float64 `json:"progress"`       //进度百分比
	Type         int     `json:"type"`           //1=到期返本金 2=延迟反本金
	DelayTime    int     `json:"delay_time"`     //延迟多少天
}

type ProductRemoteListResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data ProductRemoteData `json:"data"`
}
type ProductRemoteData struct {
	List []ProductRemote `json:"list"`
}
type ProductRemote struct {
	Name string `json:"name"` //产品名称
	Code string `json:"code"` //产品代码
}
