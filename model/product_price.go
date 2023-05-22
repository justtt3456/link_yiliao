package model

import (
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type ProductPrice struct {
	Id         int             `gorm:"column:id"`                         //
	Code       string          `gorm:"column:code"`                       //code
	Open       decimal.Decimal `gorm:"column:open"`                       //开盘价
	Close      decimal.Decimal `gorm:"column:close"`                      //收盘价
	Price      decimal.Decimal `gorm:"column:price"`                      //当前价
	High       decimal.Decimal `gorm:"column:high"`                       //最高价
	Low        decimal.Decimal `gorm:"column:low"`                        //最低价
	Vol        decimal.Decimal `gorm:"column:vol"`                        //交易量
	Change     decimal.Decimal `gorm:"column:change"`                     //涨幅百分比
	CreateTime int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
	UpdateTime int64           `gorm:"column:update_time;autoUpdateTime"` //
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
