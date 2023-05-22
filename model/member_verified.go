package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MemberVerified struct {
	Id         int    `gorm:"column:id;primary_key"`             //
	UId        int    `gorm:"column:uid"`                        //
	RealName   string `gorm:"column:real_name"`                  //
	IdNumber   string `gorm:"column:id_number"`                  //
	Mobile     string `gorm:"column:mobile"`                     //
	Frontend   string `gorm:"column:frontend"`                   //
	Backend    string `gorm:"column:backend"`                    //
	Status     int    `gorm:"column:status"`                     //1审核中 2通过 3驳回
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
	Member     Member `gorm:"foreignKey:UId"`
}

// TableName sets the insert table name for this struct type
func (m MemberVerified) TableName() string {
	return "c_member_verified"
}
func (this *MemberVerified) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberVerified) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberVerified) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMemberVerified, this.Id)
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
func (this *MemberVerified) PageList(where string, args []interface{}, page, pageSize int) ([]MemberVerified, common.Page) {
	res := make([]MemberVerified, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Limit(pageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *MemberVerified) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
