package model

import (
	"china-russia/common"
	"china-russia/global"
	"time"

	"github.com/sirupsen/logrus"
)

type MemberAddress struct {
	Id      int    `gorm:"column:id;primary_key"` //
	UId     int    `gorm:"column:uid"`            //关联用户id
	Name    string `gorm:"column:name"`
	Phone   string `gorm:"column:phone"`
	Address string `gorm:"column:address"`
	Other   string `gorm:"column:other"`
}

// TableName sets the insert table name for this struct type
func (o MemberAddress) TableName() string {
	return "c_member_address"
}

func (o *MemberAddress) ExpireTime() time.Duration {
	return time.Hour * 24 * 7
}

func (o *MemberAddress) Insert() error {
	res := global.DB.Create(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberAddress) Get() bool {
	//取数据库
	res := global.DB.Where(this).Order("id desc").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberAddress) List() []MemberAddress {
	res := make([]MemberAddress, 0)
	tx := global.DB.Model(this).Where(this).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 订单列表 使用joins联合查询 或使用Preload 根据需求决定
func (this *MemberAddress) PageList(where string, args []interface{}, page, pageSize int) ([]MemberAddress, common.Page) {
	res := make([]MemberAddress, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

func (o *MemberAddress) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberAddress) Count(where string, args []interface{}) int64 {
	var total int64
	tx := global.DB.Model(this).Where(where, args...).Count(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return int64(total)
}
func (this *MemberAddress) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
