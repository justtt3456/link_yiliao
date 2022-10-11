package model

import (
	"finance/common"
	"finance/global"
	"github.com/sirupsen/logrus"
)

type Manual struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	UID        int    `gorm:"column:uid"`                        //
	Username   string `gorm:"column:username"`                   //
	Type       int    `gorm:"column:type"`                       //1上分 2下分 3冻结 4解冻
	Amount     int64  `gorm:"column:amount"`                     //金额
	AdminID    int    `gorm:"column:admin_id"`                   //操作人
	AgentID    int    `gorm:"column:agent_id"`                   //操作人
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //创建时间
	Admin      Admin  `gorm:"foreignKey:AdminID"`
	Agent      Agent  `gorm:"foreignKey:AgentID"`
}

// TableName sets the insert table name for this struct type
func (g Manual) TableName() string {
	return "c_manual"
}
func (this *Manual) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Manual) Get() bool {
	res := global.DB.First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Manual) PageList(where string, args []interface{}, page, pageSize int) ([]Manual, common.Page) {
	res := make([]Manual, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Admin").Joins("Agent").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		find := global.DB.Model(this).Joins("Admin").Joins("Agent").Where(where, args...).Order("id desc").Offset(offset).Limit(pageSize).Find(&res)
		if find.Error != nil {
			logrus.Error(find.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
