package response

import "github.com/shopspring/decimal"

type ProductResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data ProductData `json:"data"`
}
type ProductData struct {
	Category []ProductCategory `json:"category"`
}
type ProductCategory struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Product []Product `json:"product"`
}
type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`     //产品名称
	Category int    `json:"category"` //分类id
	//CategoryName string          `json:"category_name"`
	Type                  int             `json:"type"`                    //1=到期返本金 2=延迟反本金
	Price                 decimal.Decimal `json:"price"`                   //价格
	Img                   string          `json:"img"`                     //图片
	Interval              int             `json:"interval"`                //投资期限 （天）
	IncomeRate            decimal.Decimal `json:"income_rate"`             //每日收益率
	LimitBuy              int             `json:"limit_buy"`               //限购数量
	Total                 decimal.Decimal `json:"total"`                   //项目规模
	Current               decimal.Decimal `json:"current"`                 //当前规模
	Desc                  string          `json:"desc"`                    //描述
	DelayTime             int             `json:"delay_time"`              //延迟多少天
	GiftName              string          `json:"gift_name"`               //赠送产品
	WithdrawThresholdRate decimal.Decimal `json:"withdraw_threshold_rate"` //提现额度比例
	IsHot                 int             `json:"is_hot"`                  //是否热门
	IsFinished            int             `json:"is_finished"`             //是否投满
	IsCouponGift          int             `json:"is_coupon_gift"`          //是否赠送优惠券
	//Sort                  int             `json:"sort"`                    //排序值
	Status int `json:"status"` //是否开启，1为开启，2为关闭
	//CreateTime            int64           `json:"create_time"`             //创建时间
}
type ManSongActive struct {
	Amount decimal.Decimal `json:"amount"` //满多少钱
	Price  decimal.Decimal `json:"price"`  //送多少钱的优惠券
	Id     int             `json:"id"`     //优惠券Id
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
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type EquityListResp struct {
	Id           int             `json:"id"`
	Total        int64           `json:"total"`          //总股权数
	Current      int64           `json:"current"`        //当前权数
	Price        decimal.Decimal `json:"price"`          //价格
	MinBuy       int             `json:"min_buy"`        //最低买多少股
	HitRate      decimal.Decimal `json:"hit_rate"`       //中签率
	MissRate     decimal.Decimal `json:"miss_rate"`      //未中签送的 百分比
	SellRate     decimal.Decimal `json:"sell_rate"`      //中签回购  百分比
	PreStartTime int64           `json:"pre_start_time"` //预售开始时间
	PreEndTime   int64           `json:"pre_end_time"`   //预售结束时间
	OpenTime     int64           `json:"open_time"`      //发行时间
	RecoverTime  int64           `json:"recover_time"`   //回收时间
	Status       int64           `json:"status"`         //1 = 开启 2 =关闭
}

type BuyList struct {
	Name     string          `json:"name"`     //产品名字
	Status   int             `json:"status"`   //状态 1=进行中  2=结束
	BuyTime  int             `json:"buy_time"` //投资时间
	Amount   decimal.Decimal `json:"amount"`   //金额
	Income   decimal.Decimal `json:"income"`   //每日收益
	EndTime  int64           `json:"end_time"` //到期时间
	Interval int             `json:"interval"` //投资时间
	IsGift   int             `json:"is_gift"`  //是否赠品
}

type BuyListResp struct {
	List []BuyList `json:"list"`
	Page Page      `json:"page"`
}

type BuyGuquanResp struct {
	Num        decimal.Decimal `json:"num"`         //股权数据量
	Price      decimal.Decimal `json:"price"`       //股权单价
	CreateTime int64           `json:"create_time"` //获得时间
	TotalPrice decimal.Decimal `json:"total_price"` //股权总价值
	Status     string          `json:"status"`      //发行中  回购中  完成
}

type BuyGuquanPageListResp struct {
	List []BuyGuquanList `json:"list"`
	Page Page            `json:"page"`
}

type BuyGuquanList struct {
	Id         int             `json:"id"`          //订单Id
	Num        int             `json:"num"`         //股权数据量
	Price      decimal.Decimal `json:"price"`       //股权单价
	CreateTime int64           `json:"create_time"` //获得时间
	TotalPrice decimal.Decimal `json:"total_price"` //股权总价值
	Status     int             `json:"status"`      //发行中  回购中  完成
}

type StockCertificateResp struct {
	Id                     int             `json:"id"`                        //订单Id
	RealName               string          `json:"real_name"`                 //会员真实姓名
	IdCardNo               string          `json:"id_card_no"`                //会员身份证号
	StartDate              string          `json:"start_date"`                //合同开始时间
	EndDate                string          `json:"end_date"`                  //合同结束时间
	CreateDate             string          `json:"signing_date"`              //签约时间
	Days                   int             `json:"days"`                      //合同天数
	Price                  decimal.Decimal `json:"price"`                     //单价
	Quantity               int64           `json:"quantity"`                  //股权总数量
	TotalAmount            decimal.Decimal `json:"total_amount"`              //购买总金额
	WinQuantity            int64           `json:"win_quantity"`              //中签数量
	WinProfit              decimal.Decimal `json:"win_profit"`                //中签回购利润
	WinRepurchaseAmount    decimal.Decimal `json:"win_repurchase_amount"`     //中签回购金额
	NotWinQuantity         int64           `json:"not_win_quantity"`          //未中签数量
	NotWinProfit           decimal.Decimal `json:"not_win_profit"`            //未中签回购利润
	NotWinRepurchaseAmount decimal.Decimal `json:"not_win_repurchase_amount"` //未中签回购金额
}
