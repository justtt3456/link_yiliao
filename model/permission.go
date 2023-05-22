package model

import (
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Permission struct {
	Id       int    `gorm:"column:id;primary_key"` //
	Backend  string `gorm:"column:backend"`        //
	Frontend string `gorm:"column:frontend"`       //
	Label    string `gorm:"column:label"`          //
	Pid      int    `gorm:"column:pid"`            //
	IsBtn    int    `gorm:"column:is_btn"`         //是否按钮 1是
	Sort     int    `gorm:"column:sort"`           //排序
}

// TableName sets the insert table name for this struct type
func (p *Permission) TableName() string {
	return "c_permission"
}
func (this *Permission) List() []Permission {
	res := make([]Permission, 0)
	tx := global.DB.Model(this).Order("sort desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (this *Permission) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	//bytes, err := json.Marshal(this)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return nil
}
func (this *Permission) Get() bool {
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
	//同步redis
	//bytes, err := json.Marshal(this)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return true
}
func (this *Permission) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyPermission, this.Id)
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
	//bytes, err := json.Marshal(this)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return nil
}
func (this *Permission) Remove() error {
	if this.Id > 0 || this.Pid > 0 {
		res := global.DB.Where(this).Delete(this)
		if res.Error != nil {
			return res.Error
		}
		m := Permission{Pid: this.Id}
		if len(m.List()) > 0 {
			m.Remove()
		}
	}
	//同步redis
	//global.REDIS.HDel(HashKeyPayment, strconv.Itoa(this.Id))
	return nil
}
