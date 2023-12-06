package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
)

type Product struct {
	Id                    int             `gorm:"column:id"`
	Name                  string          `gorm:"column:name"`                       //产品名称
	Category              int             `gorm:"column:category"`                   //分类id
	Type                  int             `gorm:"column:type"`                       //1=到期返本金 2=延迟反本金
	Price                 decimal.Decimal `gorm:"column:price"`                      //价格
	Img                   string          `gorm:"column:img"`                        //图片
	Interval              int             `gorm:"column:interval"`                   //投资期限 （天）
	IncomeRate            decimal.Decimal `gorm:"column:income_rate"`                //每日收益率
	LimitBuy              int             `gorm:"column:limit_buy"`                  //限购数量
	Total                 decimal.Decimal `gorm:"column:total"`                      //项目规模
	Current               decimal.Decimal `gorm:"column:current"`                    //当前规模
	Desc                  string          `gorm:"column:desc"`                       //描述
	DelayTime             int             `gorm:"column:delay_time"`                 //延迟多少天
	GiftId                int             `gorm:"column:gift_id"`                    //赠送产品ID
	WithdrawThresholdRate decimal.Decimal `gorm:"column:withdraw_threshold_rate"`    //提现额度比例
	IsHot                 int             `gorm:"column:is_hot"`                     //是否热门
	IsFinished            int             `gorm:"column:is_finished"`                //是否已满
	IsCouponGift          int             `gorm:"column:is_coupon_gift"`             //是否赠送优惠券
	Sort                  int             `gorm:"column:sort"`                       //排序值
	Status                int             `gorm:"column:status"`                     //是否开启，1为开启，2为关闭
	YbAmount              decimal.Decimal `gorm:"column:yb_amount"`                  //医保卡抵扣金额
	YbGive                decimal.Decimal `gorm:"column:yb_give"`                    //医保卡赠送金额
	CreateTime            int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
	ProductCategory       ProductCategory `gorm:"foreignKey:Category"`
}

// TableName sets the insert table name for this struct type
func (p Product) TableName() string {
	return "c_product"
}

func (m *Product) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}

func (this *Product) IndexList() []Product {
	res := make([]Product, 0)

	tx := global.DB.Where(this).Joins("ProductCategory").
		Where(this.TableName()+".status = ? and "+this.TableName()+".lang = ?", StatusOk, global.Language).
		Where("is_recommend = ?", StatusOk).
		Where("ProductCategory.status = ? ", StatusOk).
		Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *Product) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Product) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Product) Remove() error {
	res := global.DB.Where(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	//同步redis
	global.REDIS.HDel(HashKeyProduct, strconv.Itoa(this.Id))
	return nil
}
func (this *Product) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.HSet(HashKeyProduct, strconv.Itoa(this.Id), string(bytes))
	return nil
}
func (this *Product) PageList(where string, args []interface{}, page, pageSize int) ([]Product, common.Page) {
	res := make([]Product, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("ProductCategory").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("ProductCategory").Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

func (this *Product) GiftList() []Product {
	list := make([]Product, 0)
	result := global.DB.Model(this).Select("id", "name").Where("type=?", 5).Find(&list)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}

	return list
}
