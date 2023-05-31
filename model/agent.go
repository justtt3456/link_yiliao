package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Agent struct {
	Id         int    `gorm:"column:id;primary_key"`             //
	Account    string `gorm:"column:account"`                    //
	Password   string `gorm:"column:password"`                   //
	Salt       string `gorm:"column:salt"`                       //
	Token      string `gorm:"column:token"`                      //token盐
	Status     int    `gorm:"column:status"`                     //帐号启用状态，1启用2禁用
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (a *Agent) TableName() string {
	return "c_agent"
}

func (a *Agent) ExpireTime() time.Duration {
	return time.Hour * 24
}

func (a *Agent) Insert() error {
	res := global.DB.Create(a)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (a *Agent) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyAgent, a.Id)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(a)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	// if a.Id == nil {
	// 	global.REDIS.Del(fmt.Sprintf(StringKeyAgent, a.Name))
	// } else {
	bytes, _ := json.Marshal(a)
	global.REDIS.Set(fmt.Sprintf(StringKeyAgent, a.Id), string(bytes), a.ExpireTime())
	// }
	return nil
}

func (a *Agent) Remove() error {
	res := global.DB.Delete(a)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	global.REDIS.Del(fmt.Sprintf(StringKeyAgent, a.Id))
	return nil
}

func (a *Agent) Get() bool {
	//if a.Id != 0 {
	//	key := fmt.Sprintf(StringKeyAgent, a.Id)
	//	//取redis
	//	s := global.REDIS.Get(key).Val()
	//	if s != "" {
	//		err := json.Unmarshal([]byte(s), a)
	//		if err == nil {
	//			return true
	//		}
	//	}
	//}
	//取数据库
	res := global.DB.Where(a).First(a)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}

	//同步redis
	//bytes, err := json.Marshal(a)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.Set(fmt.Sprintf(StringKeyAgent, a.Id), string(bytes), a.ExpireTime())
	return true
}

func (a Agent) Info() *Agent {
	return &Agent{
		Id:         a.Id,
		Account:    a.Account,
		CreateTime: a.CreateTime,
		UpdateTime: a.UpdateTime,
	}
}

func (a *Agent) PageList(where string, args []interface{}, page, pageSize int) ([]Agent, common.Page) {
	res := make([]Agent, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64

	count := global.DB.Model(a).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(a).Where(where, args...).Limit(pageUtil.PageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (a *Agent) List(where string, args []interface{}) []Agent {
	res := make([]Agent, 0)
	tx := global.DB.Model(a).Where(where, args...).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}
