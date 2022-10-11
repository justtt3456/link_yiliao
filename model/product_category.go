package model

import (
	"finance/global"
	"github.com/sirupsen/logrus"
)

type ProductCategory struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	Name       string `gorm:"column:name"`                       //分类名称
	Status     int    `gorm:"column:status"`                     //状态
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (p ProductCategory) TableName() string {
	return "c_product_category"
}
func (this *ProductCategory) List(where string, args []interface{}) []ProductCategory {
	res := make([]ProductCategory, 0)
	//取数据库
	tx := global.DB.Model(this).Where(where, args...).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *ProductCategory) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *ProductCategory) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *ProductCategory) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this *ProductCategory) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
