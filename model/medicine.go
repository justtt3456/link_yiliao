package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type Medicine struct {
	Id                int             `gorm:"column:id"`
	Name              string          `gorm:"column:name"`                       //产品名称
	Price             decimal.Decimal `gorm:"column:price"`                      //价格
	Img               string          `gorm:"column:img"`                        //图片
	Desc              string          `gorm:"column:desc"`                       //描述
	WithdrawThreshold decimal.Decimal `gorm:"column:withdraw_threshold"`         //
	Interval          int             `gorm:"column:interval"`                   //投资期限 （天）
	Sort              int             `gorm:"column:sort"`                       //排序值
	Status            int             `gorm:"column:status"`                     //是否开启，1为开启，2为关闭
	CreateTime        int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
}

// TableName sets the insert table name for this struct type
func (p Medicine) TableName() string {
	return "c_medicine"
}

func (m *Medicine) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}

func (this *Medicine) List() []Medicine {
	res := make([]Medicine, 0)
	tx := global.DB.Where(this).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *Medicine) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Medicine) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Medicine) Remove() error {
	res := global.DB.Where(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this *Medicine) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Medicine) PageList(where string, args []interface{}, page, pageSize int) ([]Medicine, common.Page) {
	res := make([]Medicine, 0)
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
		tx := global.DB.Model(this).Where(where, args...).Limit(pageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
