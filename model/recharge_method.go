package model

import (
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type RechargeMethod struct {
	Id     int    `gorm:"column:id;primary_key"` //
	Name   string `gorm:"column:name"`           //
	Code   string `gorm:"column:code"`           //
	Icon   string `gorm:"column:icon"`           //
	Lang   string `gorm:"column:lang"`           //
	Status int    `gorm:"column:status"`         //
}

// TableName sets the insert table name for this struct type
func (r *RechargeMethod) TableName() string {
	return "c_recharge_method"
}
func (this *RechargeMethod) Get() bool {
	//所属语言
	tx := global.DB.Where(this).First(this)
	if tx.Error != nil {
		return false
	}
	return true
}
func (this *RechargeMethod) List() ([]RechargeMethod, error) {
	//所属语言
	res := make([]RechargeMethod, 0)
	tx := global.DB.Model(this).Where(this).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}
func (this *RechargeMethod) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyRechargeMethod, this.Id)
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
