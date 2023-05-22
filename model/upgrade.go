package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Upgrade struct {
	Id          int    `gorm:"column:id;primary_key"`             //
	Platform    string `gorm:"column:platform"`                   //平台
	Version     string `gorm:"column:version"`                    //版本
	DownloadURL string `gorm:"column:download_url"`               //下载地址
	MustUpgrade int    `gorm:"column:must_upgrade"`               //强制更新 1是
	UpgradeDesc string `gorm:"column:upgrade_desc"`               //更新说明
	Status      int    `gorm:"column:status"`                     //状态 1启用
	CreateTime  int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime  int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (u *Upgrade) TableName() string {
	return "c_upgrade"
}
func (this *Upgrade) GetLastVersion() bool {
	tx := global.DB.Where(this).Order("version desc").First(this)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return false
	}
	return true
}
func (this *Upgrade) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Upgrade) Get() bool {
	res := global.DB.Where(this).Order("version desc").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Upgrade) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyUpgrade, this.Id)
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

func (this *Upgrade) List() []Upgrade {
	list := make([]Upgrade, 0)
	err := global.DB.Model(this).Where(this).Find(&list)
	if err != nil {
		logrus.Error(err.Error)
		return list
	}
	return nil
}
func (this *Upgrade) PageList(where string, args []interface{}, page, pageSize int) ([]Upgrade, common.Page) {
	res := make([]Upgrade, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Upgrade) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
