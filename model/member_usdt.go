package model

import (
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MemberUsdt struct {
	Id       int    `gorm:"column:id;primary_key"` //
	UId      int    `gorm:"column:uid"`            //关联用户id
	Protocol string `gorm:"column:protocol"`       //协议
	Address  string `gorm:"column:address"`        //地址
}

// TableName sets the insert table name for this struct type
func (m *MemberUsdt) TableName() string {
	return "c_member_usdt"
}
func (this *MemberUsdt) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberUsdt) Get() bool {
	res := global.DB.Model(this).Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *MemberUsdt) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMemberBank, this.Id)
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

func (this *MemberUsdt) List() []MemberUsdt {
	res := make([]MemberUsdt, 0)
	tx := global.DB.Model(this).Where(this).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}
func (this *MemberUsdt) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
