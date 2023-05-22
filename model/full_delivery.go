package model

import (
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type FullDelivery struct {
	Id       int64           `gorm:"column:id;primary_key"` //
	Amount   decimal.Decimal `gorm:"column:amount"`         //满多少
	CouponId int64           `gorm:"column:coupon_id"`      //送什么优惠券
	Coupon   Coupon          `gorm:"foreignKey:CouponId"`   //
}

// TableName sets the insert table name for this struct type
func (c *FullDelivery) TableName() string {
	return "c_full_delivery"
}

func (this *FullDelivery) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *FullDelivery) Get() bool {
	//取数据库
	res := global.DB.Joins("Coupon").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *FullDelivery) Find(amount decimal.Decimal) bool {
	//取数据库
	res := global.DB.Model(this).Joins("Coupon").Where("amout <= ?", amount).Order("amout desc").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *FullDelivery) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *FullDelivery) List() []FullDelivery {
	res := make([]FullDelivery, 0)
	tx := global.DB.Model(this).Joins("Coupon").Where(this).Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}

func (this *FullDelivery) Del() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
