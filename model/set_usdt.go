package model

import (
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

type SetUsdt struct {
	Id         int    `gorm:"column:id;primary_key"`             //
	Address    string `gorm:"column:address"`                    //
	Status     int    `gorm:"column:status"`                     //
	Proto      string `gorm:"column:proto"`                      //协议 ERC20 TRC20
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (d *SetUsdt) TableName() string {
	return "c_set_usdt"
}
func (this *SetUsdt) Insert() error {
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
	global.REDIS.HSet(HashKeyUsdtConfig, strconv.Itoa(this.Id), string(bytes))

	return nil
}
func (this *SetUsdt) Get() bool {
	//取redis
	s := global.REDIS.HGet(HashKeyUsdtConfig, strconv.Itoa(this.Id)).Val()
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
	global.REDIS.HSet(HashKeyUsdtConfig, strconv.Itoa(this.Id), string(bytes))
	return true
}

func (this *SetUsdt) List(isFront bool) []SetUsdt {
	list := make([]SetUsdt, 0)
	//取redis
	s := global.REDIS.HGetAll(HashKeyUsdtConfig).Val()
	if len(s) != 0 {
		for _, v := range s {
			item := SetUsdt{}
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
	res := global.DB.Model(this).Where(this).Find(&list)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}
	//同步redis
	redisMap := map[string]interface{}{}
	for _, v := range list {
		marshal, _ := json.Marshal(v)
		redisMap[strconv.Itoa(v.Id)] = string(marshal)
	}
	global.REDIS.HMSet(HashKeyUsdtConfig, redisMap)
	return list
}
func (this *SetUsdt) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyUsdtConfig, this.Id)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.HSet(HashKeyUsdtConfig, strconv.Itoa(this.Id), string(bytes))
	return nil
}
func (this *SetUsdt) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	//同步redis
	global.REDIS.HDel(HashKeyUsdtConfig, strconv.Itoa(this.Id))
	return nil
}
