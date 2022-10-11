package model

import (
	"finance/app/api/swag/response"
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Notice struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	Title      string `gorm:"column:title"`                      //标题
	Intro      string `gorm:"column:intro"`                      //简介
	Content    string `gorm:"column:content"`                    //内容
	Type       int    `gorm:"column:type"`                       //类型 1滚动 2弹窗
	Lang       string `gorm:"column:lang"`                       //语言
	Status     int    `gorm:"column:status"`                     //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (n *Notice) TableName() string {
	return "c_notice"
}
func (m *Notice) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *Notice) Roll() *response.Notice {
	notice := new(response.Notice)
	//取数据库
	res := global.DB.Model(this).Where("type = ? and status = ? and lang = ? ", 1, StatusOk, global.Language).Order("id desc").First(notice)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}
	return notice
}
func (this *Notice) Pop() *response.Notice {
	notice := new(response.Notice)
	//取数据库
	res := global.DB.Model(this).Where("type = ? and status = ? and lang = ? ", 2, StatusOk, global.Language).Order("id desc").First(notice)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}

	return notice
}
func (this *Notice) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Notice) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Notice) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyNotice, this.ID)
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
func (this *Notice) PageList(where string, args []interface{}, page, pageSize int) ([]Notice, common.Page) {
	res := make([]Notice, 0)
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
func (this *Notice) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
