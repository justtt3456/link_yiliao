package model

import (
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Equity struct {
	Id           int64           `gorm:"column:id;primary_key"`    //
	Total        int64           `gorm:"column:total"`             //总股权数
	Current      int64           `gorm:"column:current"`           //当前数量
	ReleaseRate  decimal.Decimal `gorm:"column:release_rate"`      //释放百分比
	Price        decimal.Decimal `gorm:"column:price"`             //价格
	MinBuy       int64           `gorm:"column:min_buy"`           //最低买多少股
	HitRate      decimal.Decimal `gorm:"column:hit_rate"`          //中签率
	MissRate     decimal.Decimal `gorm:"column:miss_rate"`         //未中签送的 百分比
	SellRate     decimal.Decimal `gorm:"column:return_lucky_rate"` //中签回购  百分比
	PreStartTime int64           `gorm:"column:pre_start_time"`    //预售开始时间
	PreEndTime   int64           `gorm:"column:pre_end_time"`      //预售结束时间
	OpenTime     int64           `gorm:"column:open_time"`         //发行时间
	RecoverTime  int64           `gorm:"column:recover_time"`      //回收时间
	Status       int64           `gorm:"column:status"`            //1 = 开启 2 =关闭
}

func (g *Equity) TableName() string {
	return "c_equity"
}

func (this *Equity) Get(isOpen bool) bool {
	//取数据库
	if isOpen {
		res := global.DB.Where("status = 1").First(this)
		if res.Error != nil {
			logrus.Error(res.Error)
			return false
		}
	} else {
		res := global.DB.First(this)
		if res.Error != nil {
			logrus.Error(res.Error)
			return false
		}
	}

	return true
}

func (this *Equity) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Equity) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
