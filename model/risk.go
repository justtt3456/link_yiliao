package model

import (
	"encoding/json"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

type Risk struct {
	ID        int    `gorm:"column:id;primary_key"` //
	WinList   string `gorm:"column:win_list"`       //包赢名单，以|分割
	LoseList  string `gorm:"column:lose_list"`      //包输名单，以|分割
	WcLine    int    `gorm:"column:wc_line"`        //风控启动最小值。下单金额达到次标准则执行风控规则
	WcRatio   string `gorm:"column:wc_ratio"`       //风控金额和概率，以|分割，如0-100:50|100-200:30
	LoseModel int    `gorm:"column:lose_model"`     //亏损模式 1百分比亏损 2本金亏损
	LoseTime  string `gorm:"column:lose_time"`      //指定时间亏损 以/分割  23:00-08:00/18:00-19:00
	WinTime   string `gorm:"column:win_time"`       //指定时间盈利 以/分割  13:00-14:00/15:00-16:00
}

// TableName sets the insert table name for this struct type
func (r *Risk) TableName() string {
	return "c_risk"
}
func (this *Risk) Get() bool {
	s := global.REDIS.Get(StringKeyRisk).Val()
	if s != "" {
		err := json.Unmarshal([]byte(s), this)
		if err == nil {
			return true
		}
	}
	tx := global.DB.Model(this).First(this)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return false
	}
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.Set(StringKeyRisk, string(bytes), -1)
	return true
}
func (this *Risk) Update() error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyRisk, this.ID)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	bytes, _ := json.Marshal(this)
	global.REDIS.Set(StringKeyRisk, string(bytes), -1)
	return nil
}
