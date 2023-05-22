package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MemberOptional struct {
	Id         int     `gorm:"column:id;primary_key"` //
	UId        int     `gorm:"column:uid"`            //
	PId        int     `gorm:"column:pid"`
	CreateTime int64   `gorm:"column:create_time;autoCreateTime"` //
	Product    Product `gorm:"foreignKey:PId"`
}

func (h *MemberOptional) TableName() string {
	return "c_member_optional"
}

func (this *MemberOptional) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberOptional) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberOptional) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMemberOptional, this.Id)
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
func (this *MemberOptional) PageList(where string, args []interface{}, page, pageSize int) ([]MemberOptional, common.Page) {
	res := make([]MemberOptional, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Order("sort id").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *MemberOptional) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this *MemberOptional) List() []MemberOptional {
	res := make([]MemberOptional, 0)
	tx := global.DB.Model(this).Joins("Product").Where(this).Order("id desc").Find(&res)
	if tx.Error != nil {
		return nil
	}
	return res
}
