package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Help struct {
	Id         int    `gorm:"column:id;primary_key"` //
	Title      string `gorm:"column:title"`          //
	Content    string `gorm:"column:content"`        //
	Category   int    `gorm:"column:category"`
	Lang       string `gorm:"column:lang"` //
	Sort       int    `gorm:"column:sort"`
	Status     int    `gorm:"column:status"`                     //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (h *Help) TableName() string {
	return "c_help"
}

func (this *Help) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Help) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Help) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyHelp, this.Id)
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
func (this *Help) PageList(where string, args []interface{}, page, pageSize int) ([]Help, common.Page) {
	res := make([]Help, 0)
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
func (this *Help) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this *Help) List() []Help {
	res := make([]Help, 0)
	tx := global.DB.Model(this).Where(this).Order("sort desc").Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
