package model

import (
	"finance/global"
	"github.com/sirupsen/logrus"
)

type Guquan struct {
	ID              int64 `gorm:"column:id;primary_key"`    //
	TotalGuquan     int64 `gorm:"column:total_guquan"`      //总股权数
	OtherGuquan     int64 `gorm:"column:other_guquan"`      //剩余
	ReleaseRate     int   `gorm:"column:release_rate"`      //释放百分比
	Price           int64 `gorm:"column:price"`             //价格
	LimitBuy        int64 `gorm:"column:limit_buy"`         //最低买多少股
	LuckyRate       int   `gorm:"column:lucky_rate"`        //中签率
	ReturnRate      int   `gorm:"column:return_rate"`       //未中签送的 百分比
	ReturnLuckyRate int   `gorm:"column:return_lucky_rate"` //中签回购  百分比
	PreStartTime    int64 `gorm:"column:pre_start_time"`    //预售开始时间
	PreEndTime      int64 `gorm:"column:pre_end_time"`      //预售结束时间
	OpenTime        int64 `gorm:"column:open_time"`         //发行时间
	ReturnTime      int64 `gorm:"column:return_time"`       //回收时间
	Status          int64 `gorm:"column:status"`            //1 = 开启 2 =关闭
}

func (g *Guquan) TableName() string {
	return "c_guquan"
}

func (this *Guquan) Get(isOpen bool) bool {
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

func (this *Guquan) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Guquan) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
