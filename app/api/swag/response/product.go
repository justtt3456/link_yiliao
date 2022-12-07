package response

type ProductResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ProductData `json:"data"`
}
type ProductData struct {
	Category []ProductCategory `json:"category"`
}
type ProductCategory struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Product []Product `json:"product"`
}
type Product struct {
	ID            int             `json:"id"`              //产品ID
	Name          string          `json:"name"`            //产品名字
	Category      int             `json:"category"`        //分类ID
	CategoryName  string          `json:"category_name"`   //分类名
	Status        int             `json:"status"`          //是否开启，1为开启，0为关闭
	Tag           int             `json:"tag"`             //1=热
	TimeLimit     int             `json:"time_limit"`      //投资期限 （天）
	IsRecommend   int             `json:"is_recommend"`    //是否推荐到首页 1是 2否
	Dayincome     float64         `json:"day_income"`      //每日收益
	Price         float64         `json:"price"`           //价格  (最低买多少)
	TotalPrice    float64         `json:"total_price"`     //项目规模
	OtherPrice    float64         `json:"OtherPrice"`      //可投余额
	MoreBuy       int             `json:"more_buy"`        //最多可以买多少份
	Desc          string          `json:"desc"`            //描述
	IsFinish      int             `json:"is_finish"`       //1=进行中  2=已投满
	IsManjian     int             `json:"is_manjian"`      //1=有满送活动  2=无满送活动
	BuyTimeLimit  int             `json:"buy_time_limit"`  //产品限时多少天
	Progress      float64         `json:"progress"`        //进度百分比
	Type          int             `json:"type"`            //1=到期返本金 2=延迟反本金
	DelayTime     int             `json:"delay_time"`      //延迟多少天
	CreateTime    int64           `json:"create_time"`     //创建时间
	ManSongActive []ManSongActive `json:"man_song_active"` //满送活动
}
type ManSongActive struct {
	Amount float64 `json:"amount"` //满多少钱
	Price  float64 `json:"price"`  //送多少钱的优惠券
	Id     int64   `json:"id"`     //优惠券ID
}

type ProductOptionResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data ProductOption `json:"data"`
}
type ProductOption struct {
	Interval []int `json:"interval"` //秒
	Ratio    []int `json:"ratio"`    //收益百分比
	Quick    []int `json:"quick"`    //快捷金额
	Fee      int   `json:"fee"`      //购买手续费 百分比
}

type ProductListResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data ProductListData `json:"data"`
}
type ProductListData struct {
	List []Product `json:"list"`
	Page Page      `json:"page"`
}

type ProductCategoryResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data ProductCategoryData `json:"data"`
}
type ProductCategoryData struct {
	List []ProductCategoryItem `json:"list"`
}
type ProductCategoryItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GuquanListResp struct {
	Id              int64   `json:"id"`
	TotalGuquan     int64   `json:"total_guquan"`      //总股权数
	OtherGuquan     int64   `json:"other_guquan"`      //剩余权数
	ReleaseRate     float64 `json:"release_rate"`      //释放百分比
	Price           float64 `json:"price"`             //价格
	LimitBuy        int64   `json:"limit_buy"`         //最低买多少股
	LuckyRate       float64 `json:"lucky_rate"`        //中签率
	ReturnRate      float64 `json:"return_rate"`       //未中签送的 百分比
	ReturnLuckyRate float64 `json:"return_lucky_rate"` //中签回购  百分比
	PreStartTime    int64   `json:"pre_start_time"`    //预售开始时间
	PreEndTime      int64   `json:"pre_end_time"`      //预售结束时间
	OpenTime        int64   `json:"open_time"`         //发行时间
	ReturnTime      int64   `json:"return_time"`       //回收时间
	Status          int64   `json:"status"`            //1 = 开启 2 =关闭
}

type BuyList struct {
	Name    string  `json:"name"`     //产品名字
	Status  int     `json:"status"`   //状态 1=进行中  2=结束
	BuyTime int     `json:"buy_time"` //投资时间
	Amount  float64 `json:"amount"`   //金额
}

type BuyListResp struct {
	List []BuyList `json:"list"`
	Page Page      `json:"page"`
}

type BuyGuquanResp struct {
	Num        int64   `json:"num"`         //股权数据量
	Price      float64 `json:"price"`       //股权单价
	CreateTime int64   `json:"create_time"` //获得时间
	TotalPrice float64 `json:"total_price"` //股权总价值
	Status     string  `json:"status"`      //发行中  回购中  完成
}

type BuyGuquanPageListResp struct {
	List []BuyGuquanList `json:"list"`
	Page Page            `json:"page"`
}

type BuyGuquanList struct {
	ID         int     `json:"id"`          //订单ID
	Num        int64   `json:"num"`         //股权数据量
	Price      float64 `json:"price"`       //股权单价
	CreateTime int64   `json:"create_time"` //获得时间
	TotalPrice float64 `json:"total_price"` //股权总价值
	Status     string  `json:"status"`      //发行中  回购中  完成
}

type StockCertificateResp struct {
	ID                     int     `json:"id"`                        //订单ID
	RealName               string  `json:"real_name"`                 //会员真实姓名
	IdCardNo               string  `json:"id_card_no"`                //会员身份证号
	StartDate              string  `json:"start_date"`                //合同开始时间
	EndDate                string  `json:"end_date"`                  //合同结束时间
	CreateDate             string  `json:"signing_date"`              //签约时间
	Days                   int     `json:"days"`                      //合同天数
	Price                  float64 `json:"price"`                     //单价
	Quantity               int64   `json:"quantity"`                  //股权总数量
	TotalAmount            float64 `json:"total_amount"`              //购买总金额
	WinQuantity            int64   `json:"win_quantity"`              //中签数量
	WinProfit              float64 `json:"win_profit"`                //中签回购利润
	WinRepurchaseAmount    float64 `json:"win_repurchase_amount"`     //中签回购金额
	NotWinQuantity         int64   `json:"not_win_quantity"`          //未中签数量
	NotWinProfit           float64 `json:"not_win_profit"`            //未中签回购利润
	NotWinRepurchaseAmount float64 `json:"not_win_repurchase_amount"` //未中签回购金额
}
