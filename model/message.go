package model

import (
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Message struct {
	ID         int    `gorm:"column:id;primary_key"` //
	UID        int    `gorm:"column:uid"`            //
	Title      string `gorm:"column:title"`          //标题
	Content    string `gorm:"column:content"`        //内容
	Status     int    `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //创建日期
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //修改时间
}

// TableName sets the insert table name for this struct type
func (m *Message) TableName() string {
	return "c_message"
}

func (this *Message) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Message) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Message) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMessage, this.ID)
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
func (this *Message) PageList(where string, args []interface{}, page, pageSize int) ([]Message, common.Page) {
	res := make([]Message, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Message) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
