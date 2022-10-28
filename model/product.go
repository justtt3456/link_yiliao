package model

import (
	"encoding/json"
	"finance/common"
	"finance/global"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
)

type Product struct {
	ID              int             `gorm:"column:id;primary_key"` //
	Name            string          `gorm:"column:name"`           //产品名称
	Category        int             `gorm:"column:category"`       //分类id
	CreateTime      int64           `gorm:"column:create_time"`    //创建时间
	Status          int             `gorm:"column:status"`         //是否开启，1为开启，0为关闭
	Tag             int             `gorm:"column:tag"`            //1=热
	TimeLimit       int             `gorm:"column:time_limit"`     //投资期限 （天）
	IsRecommend     int             `gorm:"column:is_recommend"`   //是否推荐到首页 1是 2否
	Dayincome       int             `gorm:"column:day_income"`     //每日收益  千分比
	Price           int64           `gorm:"column:price"`          //价格  (最低买多少)
	TotalPrice      int64           `gorm:"column:total_price"`    //项目规模
	OtherPrice      int64           `gorm:"column:other_price"`    //可投余额
	MoreBuy         int             `gorm:"column:more_buy"`       //最多可以买多少份
	Desc            string          `gorm:"column:desc"`           //描述
	IsFinish        int             `gorm:"column:is_finish"`      //1=进行中  2=已投满
	IsManjian       int             `gorm:"column:is_manjian"`     //1=有满减  2=无满减
	BuyTimeLimit    int             `gorm:"column:buy_time_limit"` //产品限时多少天
	Progress        int             `gorm:"column:progress"`       //项目进度
	Type            int             `gorm:"column:type"`           //1=到期返本金 2=延迟反本金
	DelayTime       int             `gorm:"column:delay_time"`     //延迟多少天
	ProductCategory ProductCategory `gorm:"foreignKey:Category"`
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
	global.REDIS.HDel(HashKeyProduct, strconv.Itoa(this.ID))
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
	global.REDIS.HSet(HashKeyProduct, strconv.Itoa(this.ID), string(bytes))
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
