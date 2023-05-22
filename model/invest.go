package model

import (
	"china-russia/global"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Invest struct {
	Id             int             `gorm:"column:id;primary_key"`             //
	Name           string          `gorm:"column:name"`                       //余额宝名称
	Ratio          int             `gorm:"column:ratio"`                      //利率 0.01=1%!
	FreezeDay      int             `gorm:"column:freeze_day"`                 //冻结天数
	IncomeInterval int             `gorm:"column:income_interval"`            //收益发放间隔天数
	Status         int             `gorm:"column:status"`                     //余额宝开关，1开启，2关闭
	Description    string          `gorm:"column:description"`                //余额宝说明
	MinAmount      decimal.Decimal `gorm:"column:min_amount"`                 //
	CreateTime     int64           `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime     int64           `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (i *Invest) TableName() string {
	return "c_invest"
}
func (i *Invest) Insert() error {
	res := global.DB.Create(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Invest) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Invest) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMessage, this.Id)
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
