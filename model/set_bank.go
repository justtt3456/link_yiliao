package model

import (
	"encoding/json"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

type SetBank struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	BankName   string `gorm:"column:bank_name"`                  //银行名称
	CardNumber string `gorm:"column:card_number"`                //卡号
	BranchBank string `gorm:"column:branch_bank"`                //开户行（开户所在地）
	RealName   string `gorm:"column:real_name"`                  //开户人
	Lang       string `gorm:"column:lang"`                       //
	Status     int    `gorm:"column:status"`                     //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (m *SetBank) TableName() string {
	return "c_set_bank"
}
func (this *SetBank) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.HSet(HashKeyBankConfig, strconv.Itoa(this.ID), string(bytes))
	return nil
}
func (this *SetBank) Get() bool {
	//取redis
	s := global.REDIS.HGet(HashKeyBankConfig, strconv.Itoa(this.ID)).Val()
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
	//同步redis
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.HSet(HashKeyBankConfig, strconv.Itoa(this.ID), string(bytes))
	return true
}

func (this *SetBank) List(isFront bool) []SetBank {
	list := make([]SetBank, 0)
	//取redis
	s := global.REDIS.HGetAll(HashKeyBankConfig).Val()
	if len(s) != 0 {
		for _, v := range s {
			item := SetBank{}
			err := json.Unmarshal([]byte(v), &item)
			if err != nil {
				logrus.Error(err)
				continue
			}
			if isFront {
				if item.Status == StatusOk {
					list = append(list, item)
				}
			} else {
				list = append(list, item)
			}
		}
		if list != nil {
			return list
		}
	}
	//取数据库
	res := global.DB.Where(this).Find(&list)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}
	//同步redis
	redisMap := map[string]interface{}{}
	for _, v := range list {
		marshal, _ := json.Marshal(v)
		redisMap[strconv.Itoa(v.ID)] = string(marshal)
	}
	global.REDIS.HMSet(HashKeyBankConfig, redisMap)
	return list
}
func (this *SetBank) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyBankConfig, this.ID)
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
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.HSet(HashKeyBankConfig, strconv.Itoa(this.ID), string(bytes))
	return nil
}
func (this *SetBank) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	//同步redis
	global.REDIS.HDel(HashKeyBankConfig, strconv.Itoa(this.ID))
	return nil
}
