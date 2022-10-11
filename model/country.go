package model

import (
	"finance/global"
	"github.com/sirupsen/logrus"
)

type Country struct {
	ID        int    `gorm:"column:id;primary_key"` //
	ZhName    string `gorm:"column:zh_name"`        //
	EnName    string `gorm:"column:en_name"`        //
	Code      string `gorm:"column:code"`           //
	Lang      string `gorm:"column:lang"`           //
	Currency  string `gorm:"column:currency"`       //
	Sort      int    `gorm:"column:sort"`           //
	IsReg     int    `gorm:"column:is_reg"`         //
	IsDefault int    `gorm:"column:is_default"`     //
}

// TableName sets the insert table name for this struct type
func (c *Country) TableName() string {
	return "c_country"
}

func (this *Country) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Country) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Country) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Country) List() []Country {
	res := make([]Country, 0)
	tx := global.DB.Model(this).Where(this).Order("is_default asc,sort desc").Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
