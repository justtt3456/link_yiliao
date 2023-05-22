package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type News struct {
	Id         int    `gorm:"column:id;primary_key"`             //
	Title      string `gorm:"column:title"`                      //
	Content    string `gorm:"column:content"`                    //
	Status     int    `gorm:"column:status"`                     //
	Sort       int    `gorm:"column:sort"`                       //
	Intro      string `gorm:"column:intro"`                      //
	Cover      string `gorm:"column:cover"`                      //封面图
	Lang       string `gorm:"column:lang"`                       //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (n *News) TableName() string {
	return "c_news"
}
func (this *News) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *News) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *News) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyNews, this.Id)
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
func (this *News) PageList(where string, args []interface{}, page, pageSize int) ([]News, common.Page) {
	res := make([]News, 0)
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
func (this *News) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
