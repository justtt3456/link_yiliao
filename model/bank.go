package model

import (
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Bank struct {
	ID         int    `gorm:"column:id;primary_key" json:"id"`   //
	BankName   string `gorm:"column:bank_name" json:"bank_name"` //银行名称
	Sort       int    `gorm:"column:sort" json:"sort"`
	Status     int    `gorm:"column:status" json:"status"` //状态
	Lang       string `gorm:"column:lang" json:"lang"`
	Code       string `gorm:"column:code" json:"code"`
	CreateTime int64  `gorm:"column:create_time;autoCreateTime" json:"create_time"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime" json:"update_time"` //
}

// TableName sets the insert table name for this struct type
func (this *Bank) TableName() string {
	return "c_bank"
}
func (this *Bank) PageList(where string, args []interface{}, page, pageSize int) ([]Bank, common.Page) {
	res := make([]Bank, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Order("sort desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Bank) List() []Bank {
	res := make([]Bank, 0)
	tx := global.DB.Model(this).Where(this).Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
func (this *Bank) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Bank) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Bank) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyBank, this.ID)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Bank) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
