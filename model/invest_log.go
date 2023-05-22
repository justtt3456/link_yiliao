package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type InvestLog struct {
	Id         int             `gorm:"column:id;primary_key"`             //
	UId        int             `gorm:"column:uid"`                        //关联用户id
	Income     decimal.Decimal `gorm:"column:income"`                     //余额宝奖励金额
	Balance    decimal.Decimal `gorm:"column:balance"`                    //余额宝余额
	CreateTime int64           `gorm:"column:create_time;autoCreateTime"` //生成时间
	Member     Member          `gorm:"foreignKey:UId"`
}

// TableName sets the insert table name for this struct type
func (i *InvestLog) TableName() string {
	return "c_invest_log"
}
func (i *InvestLog) Insert() error {
	res := global.DB.Create(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *InvestLog) PageList(where string, args []interface{}, page, pageSize int) ([]InvestLog, common.Page) {
	res := make([]InvestLog, 0)
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
func (r *InvestLog) YesterdayIncome(uid int) decimal.Decimal {
	var total decimal.Decimal
	zero := common.GetTodayZero()
	where := "uid = ? and create_time >= ?"
	args := []interface{}{uid, zero}
	tx := global.DB.Model(r).Select("income").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return decimal.Zero
	}
	return total
}
func (r *InvestLog) Sum(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(r).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
