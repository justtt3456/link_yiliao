package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
)

type PayChannel struct {
	Id         int     `gorm:"column:id;primary_key"`             //
	Name       string  `gorm:"column:name"`                       //支付方式名称
	PaymentId  int     `gorm:"column:payment_id"`                 //第三方名称
	Code       string  `gorm:"column:code"`                       //支付编码
	Min        int64   `gorm:"column:min"`                        //最小值
	Max        int64   `gorm:"column:max"`                        //最大值
	Status     int     `gorm:"column:status"`                     //状态
	Category   int     `gorm:"column:category"`                   //分类(所属支付方式,预留待使用)
	Sort       int     `gorm:"column:sort"`                       //排序值
	Icon       string  `gorm:"column:icon"`                       //图标
	Fee        int     `gorm:"column:fee"`                        //手续费
	Lang       string  `gorm:"column:lang"`                       //语言
	CreateTime int64   `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64   `gorm:"column:update_time;autoUpdateTime"` //
	Payment    Payment `gorm:"foreignKey:PaymentId"`              //外键
}

// TableName sets the insert table name for this struct type
func (p *PayChannel) TableName() string {
	return "c_pay_channel"
}
func (m *PayChannel) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *PayChannel) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	if this.Status == StatusOk {
		//同步redis
		bytes, err := json.Marshal(this)
		if err != nil {
			log.Println(err)
		}
		global.REDIS.HSet(HashKeyPayChannel, strconv.Itoa(this.Id), string(bytes))
	}
	return nil
}
func (this *PayChannel) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyPayChannel, this.Id)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	if this.Status != StatusOk {
		global.REDIS.HDel(HashKeyPayment, strconv.Itoa(this.Id))
	} else {
		//同步redis
		bytes, err := json.Marshal(this)
		if err != nil {
			log.Println(err)
		}
		global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	}
	return nil
}
func (this *PayChannel) Get() bool {
	//if this.Id != 0 {
	//	//取redis
	//	s := global.REDIS.HGet(HashKeyPayment, strconv.Itoa(this.Id)).Val()
	//	if s != "" {
	//		err := json.Unmarshal([]byte(s), this)
	//		if err == nil {
	//			return true
	//		}
	//	}
	//}
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	////同步redis
	//bytes, err := json.Marshal(this)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return true
}
func (this PayChannel) List() []PayChannel {
	res := make([]PayChannel, 0)
	tx := global.DB.Model(this).Joins("Payment").Where(this).Where("`Payment`.`type` = ?", this.Payment.Type).Order("sort desc").Scan(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *PayChannel) PageList(where string, args []interface{}, page, pageSize int) ([]PayChannel, common.Page) {
	res := make([]PayChannel, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Payment").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Payment").Where(where, args...).Order("sort desc").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *PayChannel) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		return res.Error
	}
	global.REDIS.HDel(HashKeyPayment, strconv.Itoa(this.Id))
	return nil
}
