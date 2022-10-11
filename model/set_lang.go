package model

import (
	"encoding/json"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

type SetLang struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	Name       string `gorm:"column:name"`                       //语言名称
	Code       string `gorm:"column:code"`                       //英文简称
	Icon       string `gorm:"column:icon"`                       //语言图标
	IsDefault  int    `gorm:"column:is_default"`                 //是否默认语言
	Status     int    `gorm:"column:status"`                     //状态
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (s *SetLang) TableName() string {
	return "c_set_lang"
}

func (this *SetLang) List(isFront bool) []SetLang {
	list := make([]SetLang, 0)

	//取数据库
	res := global.DB.Model(this).Where(this).Find(&list)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil
	}
	return list
}
func (this *SetLang) Get() bool {
	//取redis
	s := global.REDIS.HGet(HashKeyLangConfig, strconv.Itoa(this.ID)).Val()
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
	global.REDIS.HSet(HashKeyLangConfig, strconv.Itoa(this.ID), string(bytes))
	return true
}
func (this *SetLang) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyLangConfig, this.ID)
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
	global.REDIS.HSet(HashKeyLangConfig, strconv.Itoa(this.ID), string(bytes))
	return nil
}
