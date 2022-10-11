package model

import (
	"finance/global"
	"github.com/sirupsen/logrus"
)

type Coupon struct {
	ID    int64 `gorm:"column:id;primary_key"` //
	Price int64 `gorm:"column:price"`          //优惠券面额
}

// TableName sets the insert table name for this struct type
func (c *Coupon) TableName() string {
	return "c_coupon"
}

func (this *Coupon) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Coupon) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Coupon) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Coupon) List() []Coupon {
	res := make([]Coupon, 0)
	tx := global.DB.Model(this).Where(this).Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
