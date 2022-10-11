package model

import (
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Banner struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	Image      string `gorm:"column:image"`                      //图片
	Sort       int    `gorm:"column:sort"`                       //排序
	Link       string `gorm:"column:link"`                       //链接
	Status     int    `gorm:"column:status"`                     //
	Lang       string `gorm:"column:lang"`                       //
	Type       int    `gorm:"column:type"`                       //1=图片 2=视频
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (b *Banner) TableName() string {
	return "c_banner"
}
func (m *Banner) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *Banner) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Banner) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Banner) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyBanner, this.ID)
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
func (this *Banner) PageList(where string, args []interface{}, page, pageSize int) ([]Banner, common.Page) {
	res := make([]Banner, 0)
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
func (this *Banner) List() []Banner {
	res := make([]Banner, 0)
	////取redis
	//s := global.REDIS.HGet(HashKeyBanner, this.Lang).Val()
	//if s != "" {
	//	err := json.Unmarshal([]byte(s), &res)
	//	if err == nil {
	//		return res
	//	}
	//}
	//取数据库
	tx := global.DB.Where(this).Order("sort desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	//同步redis
	//banner, err := json.Marshal(res)
	//if err != nil {
	//	logrus.Error(err)
	//	return nil
	//}
	//global.REDIS.HSet(HashKeyBanner, this.Lang, string(banner))
	return res
}
func (this *Banner) Remove() error {
	res := global.DB.Where(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
