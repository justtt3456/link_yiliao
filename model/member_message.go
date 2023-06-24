package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

// 站内信
type MemberMessage struct {
	Id         int     `gorm:"column:id"`
	Uid        int     `gorm:"column:uid"`
	MessageId  int     `gorm:"column:message_id"`                 //站内信id
	IsRead     int     `gorm:"column:is_read"`                    //1=未读 2=已读
	CreateTime int64   `gorm:"column:create_time;autoCreateTime"` //创建日期
	UpdateTime int64   `gorm:"column:update_time;autoUpdateTime"` //修改时间
	Message    Message `gorm:"foreignKey:MessageId"`
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
func (this *MemberMessage) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberMessage) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberMessage) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMessage, this.Id)
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
func (this *MemberMessage) PageList(where string, args []interface{}, page, pageSize int) ([]MemberMessage, common.Page) {
	res := make([]MemberMessage, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Joins("Message").Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
