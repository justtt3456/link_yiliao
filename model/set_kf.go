package model

import (
	"encoding/json"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

type SetKf struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	Name       string `gorm:"column:name"`                       //
	StartTime  string `gorm:"column:start_time"`                 //
	EndTime    string `gorm:"column:end_time"`                   //
	Link       string `gorm:"column:link"`                       //
	Key        string `gorm:"column:key"`                        //
	Icon       string `gorm:"column:icon"`                       //
	Status     int    `gorm:"column:status"`                     //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (s *SetKf) TableName() string {
	return "c_set_kf"
}
func (this *SetKf) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyKfConfig, this.ID)
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
	global.REDIS.HSet(HashKeyKfConfig, strconv.Itoa(this.ID), string(bytes))
	return nil
}
func (this *SetKf) Get() bool {
	//取redis
	s := global.REDIS.HGet(HashKeyKfConfig, strconv.Itoa(this.ID)).Val()
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
	global.REDIS.HSet(HashKeyKfConfig, strconv.Itoa(this.ID), string(bytes))
	return true
}

//前台显示可用
func (this *SetKf) List(isFront bool) []SetKf {
	list := make([]SetKf, 0)
	////取redis
	//s := global.REDIS.HGetAll(HashKeyKfConfig).Val()
	//if len(s) != 0 {
	//	for _, v := range s {
	//		item := SetKf{}
	//		err := json.Unmarshal([]byte(v), &item)
	//		if err != nil {
	//			logrus.Error(err)
	//			continue
	//		}
	//		if isFront {
	//			if item.Status == StatusOk {
	//				list = append(list, item)
	//			}
	//		} else {
	//			list = append(list, item)
	//		}
	//	}
	//	if list != nil {
	//		return list
	//	}
	//}
	//取数据库
	res := global.DB.Model(this).Where(this).Find(&list)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}
	////同步redis
	//redisMap := map[string]interface{}{}
	//for _, v := range list {
	//	marshal, _ := json.Marshal(v)
	//	redisMap[strconv.Itoa(v.ID)] = string(marshal)
	//}
	//global.REDIS.HMSet(HashKeyKfConfig, redisMap)
	return list
}
