package model

import (
	"china-russia/global"
	"github.com/sirupsen/logrus"
)

type MemberLoginLog struct {
	Id        int    `gorm:"column:id;primary_key"` //
	UId       int    `gorm:"column:uid"`            //
	Username  string `gorm:"column:username"`       //
	LoginIP   string `gorm:"column:login_ip"`       //
	LoginTime int64  `gorm:"column:login_time"`     //
}

// TableName sets the insert table name for this struct type
func (m *MemberLoginLog) TableName() string {
	return "c_member_login_log"
}
func (this *MemberLoginLog) Insert() error {
	tx := global.DB.Create(this)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return tx.Error
	}
	return nil
}
