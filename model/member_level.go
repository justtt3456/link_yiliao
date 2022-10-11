package model

import (
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MemberLevel struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"` //
	Name string `gorm:"column:name" json:"name"`         //等级名称
	Img  string `gorm:"column:img" json:"img"`           //图标
}

// TableName sets the insert table name for this struct type
func (m *MemberLevel) TableName() string {
	return "c_member_level"
}
func (this *MemberLevel) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberLevel) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyLevel, this.ID)
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
func (this *MemberLevel) List() []MemberLevel {
	res := make([]MemberLevel, 0)
	tx := global.DB.Model(this).Where(this).Order("id asc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
