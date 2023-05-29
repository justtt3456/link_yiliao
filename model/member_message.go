package model

import (
	"china-russia/global"
	"github.com/sirupsen/logrus"
)

// 站内信
type MemberMessage struct {
	Id         int   `gorm:"column:id"`
	Uid        int   `gorm:"column:uid"`
	MessageId  int   `gorm:"column:message_id"`  //站内信id
	IsRead     int   `gorm:"column:is_read"`     //1=未读 2=已读
	CreateTime int64 `gorm:"column:create_time"` //创建日期
	UpdateTime int64 `gorm:"column:update_time"` //修改时间
}

func (MemberMessage) TableName() string {
	return "c_member_message"
}
func (this *MemberMessage) Count(where string, args []interface{}) int64 {
	var total int64
	count := global.DB.Model(this).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return 0
	}
	return total
}
