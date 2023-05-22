package model

import (
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type IP struct {
	Id int    `gorm:"column:id;primary_key"` //
	IP string `gorm:"column:ip"`             //

}

// TableName sets the insert table name for this struct type
func (i *IP) TableName() string {
	return "c_white_ip"
}
func (this *IP) Get() bool {
	//取redis
	s := global.REDIS.Get(StringKeyWhiteIP).Val()
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
	return true
}
func (this *IP) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyWhiteIP, this.Id)
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
	bytes, _ := json.Marshal(this)
	global.REDIS.Set(StringKeyWhiteIP, string(bytes), -1)
	return nil
}
