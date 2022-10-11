package model

import (
	"encoding/json"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type ProductPrice struct {
	ID         int     `gorm:"column:id"`                         //
	Code       string  `gorm:"column:code"`                       //code
	Open       float64 `gorm:"column:open"`                       //开盘价
	Close      float64 `gorm:"column:close"`                      //收盘价
	Price      float64 `gorm:"column:price"`                      //当前价
	High       float64 `gorm:"column:high"`                       //最高价
	Low        float64 `gorm:"column:low"`                        //最低价
	Vol        float64 `gorm:"column:vol"`                        //交易量
	Change     float64 `gorm:"column:change"`                     //涨幅百分比
	CreateTime int64   `gorm:"column:create_time;autoCreateTime"` //创建时间
	UpdateTime int64   `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (p *ProductPrice) TableName() string {
	return "c_product_price"
}
func (m *ProductPrice) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}

func (this *ProductPrice) Get() bool {
	if this.Code == "" {
		return false
	}
	//取redis
	key := fmt.Sprintf(StringKeyProductPrice, this.Code)
	s := global.REDIS.Get(key).Val()
	if s != "" {
		err := json.Unmarshal([]byte(s), this)
		if err == nil {
			return true
		}
	}
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.Set(key, string(bytes), this.ExpireTime())
	return true
}
