package model

import (
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type WithdrawMethod struct {
	Id     int    `gorm:"column:id;primary_key"` //
	Name   string `gorm:"column:name"`           //
	Code   string `gorm:"column:code"`           //
	Icon   string `gorm:"column:icon"`           //
	Lang   string `gorm:"column:lang"`           //
	Status int    `gorm:"column:status"`         //
	Fee    int    `gorm:"column:fee"`            //
}

// TableName sets the insert table name for this struct type
func (w *WithdrawMethod) TableName() string {
	return "c_withdraw_method"
}
func (this *WithdrawMethod) Get() bool {
	//所属语言
	tx := global.DB.Where(this).First(this)
	if tx.Error != nil {
		return false
	}
	return true
}
func (this *WithdrawMethod) List() ([]WithdrawMethod, error) {
	//所属语言
	res := make([]WithdrawMethod, 0)
	tx := global.DB.Model(this).Where(this).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}
func (this *WithdrawMethod) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyWithdrawMethod, this.Id)
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
