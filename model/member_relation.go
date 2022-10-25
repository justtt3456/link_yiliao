package model

import (
	"finance/common"
	"finance/global"
	"github.com/sirupsen/logrus"
)

type MemberRelation struct {
	UID     int    `gorm:"column:uid"`   //查询祖先
	Puid    int    `gorm:"column:puid"`  //查询后代
	Level   int64  `gorm:"column:level"` //代理层级
	Member  Member `gorm:"foreignKey:uid"`
	Member2 Member `gorm:"foreignKey:puid"`
}

func (m *MemberRelation) TableName() string {
	return "c_member_relation"
}

func (this *MemberRelation) Get() bool {
	//取数据库
	res := global.DB.Model(this).Joins("Member").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberRelation) Get2() bool {
	//取数据库
	res := global.DB.Model(this).Joins("Member2").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberRelation) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberRelation) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberRelation) InsertAll(result []MemberRelation) error {
	res := global.DB.Create(result)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

//查询祖先
func (this *MemberRelation) GetByUid() ([]MemberRelation, error) {
	var res []MemberRelation
	err := global.DB.Model(this).Where("uid = ?", this.UID).Find(&res)
	if err.Error != nil {
		logrus.Error(err.Error)
		return nil, err.Error
	}
	return res, nil
}

//查询后代
func (this *MemberRelation) GetByPuid(where string, args []interface{}, page, pageSize int) ([]MemberRelation, common.Page) {
	res := make([]MemberRelation, 0)
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
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Order("level").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

//查询后代
func (this *MemberRelation) GetByPuidAll(where string, args []interface{}) ([]*MemberRelation, int64) {

	res := make([]*MemberRelation, 0)

	var total int64
	count := global.DB.Model(this).Joins("Member").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, total
	}
	if total > 0 {
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, total
		}
	}
	return res, total
}
