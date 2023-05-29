package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/sirupsen/logrus"
)

type MemberParents struct {
	Uid            int            `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                         //查询祖先
	ParentId       int            `gorm:"column:parent_id" db:"parent_id" json:"parent_id" form:"parent_id"` //查询后代
	Level          int            `gorm:"column:level" db:"level" json:"level" form:"level"`
	Member         Member         `gorm:"foreignKey:uid"`
	Parent         Member         `gorm:"foreignKey:parent_id"`
	MemberVerified MemberVerified `gorm:"foreignKey:Uid;references:UId"`
}

func (m *MemberParents) TableName() string {
	return "c_member_parents"
}

func (this *MemberParents) Get() bool {
	//取数据库
	res := global.DB.Model(this).Joins("Member").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberParents) Get2() bool {
	//取数据库
	res := global.DB.Model(this).Joins("Parent").Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberParents) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberParents) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MemberParents) InsertAll(result []MemberParents) error {
	res := global.DB.Create(result)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

// 查询祖先
func (this *MemberParents) GetByUid() ([]MemberParents, error) {
	var res []MemberParents
	err := global.DB.Model(this).Where("uid = ?", this.Uid).Order("level asc").Find(&res)
	if err.Error != nil {
		logrus.Error(err.Error)
		return nil, err.Error
	}
	return res, nil
}

// 查询后代
func (this *MemberParents) GetByPuid(where string, args []interface{}, page, pageSize int) ([]MemberParents, common.Page) {
	res := make([]MemberParents, 0)
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
		tx := global.DB.Model(this).Joins("Member").Joins("MemberVerified").Where(where, args...).Order(this.TableName() + ".uid ASC").Order("level").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

// 查询后代
func (this *MemberParents) GetByPuidAll(where string, args []interface{}) ([]*MemberParents, int64) {

	res := make([]*MemberParents, 0)

	var total int64
	count := global.DB.Model(this).Joins("Member").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, total
	}
	if total > 0 {
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Order(this.TableName() + ".uid ASC").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, total
		}
	}
	return res, total
}

// 根据下线会员Id获取团队代理Id列表
func (this *MemberParents) GetTeamLeaderIds(userIds []int) []int {
	res := make([]*MemberParents, 0)
	var proxyIds []int

	tx := global.DB.Model(this).Select("DISTINCT puid").Where("uid in ?", userIds).Where("level > 0").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return proxyIds
	}
	//当查询记录为空虹
	if len(res) == 0 {
		return proxyIds
	}

	for _, lines := range res {
		proxyIds = append(proxyIds, lines.ParentId)
	}
	return proxyIds
}

// 查询后代, 注:前台使用, 前台/后台的排序不一样,所以前台显示的数据使用独立的函数
func (this *MemberParents) GetChildListByParentId(where string, args []interface{}, page, pageSize int) ([]MemberParents, common.Page) {
	res := make([]MemberParents, 0)
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
		tx := global.DB.Model(this).Joins("Member").Joins("MemberVerified").Where(where, args...).Order("level").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
