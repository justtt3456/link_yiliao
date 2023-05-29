package model

import (
	"china-russia/global"
	"github.com/sirupsen/logrus"
)

type MemberCoupon struct {
	Id       int    `gorm:"column:id;primary_key"` //
	Uid      int    `gorm:"column:uid"`            //用户id
	CouponId int    `gorm:"column:coupon_id"`      //优惠券id
	IsUse    int    `gorm:"column:is_use"`         //1=未使用 2=已使用
	Coupon   Coupon `gorm:"foreignKey:CouponId"`   //
}

// TableName sets the insert table name for this struct type
func (c *MemberCoupon) TableName() string {
	return "c_member_coupon"
}

func (this *MemberCoupon) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberCoupon) Get() bool {
	//取数据库
	res := global.DB.Model(this).Joins("Coupon").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberCoupon) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberCoupon) List() []MemberCoupon {
	res := make([]MemberCoupon, 0)
	tx := global.DB.Model(this).Joins("Coupon").Where(this).Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
