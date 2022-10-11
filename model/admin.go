package model

import (
	"encoding/json"
	"finance/common"
	"finance/global"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

type Admin struct {
	ID         int    `gorm:"column:id;primary_key"` //
	Username   string `gorm:"column:username"`       //
	Password   string `gorm:"column:password"`       //
	Salt       string `gorm:"column:salt"`           //
	Role       int    `gorm:"column:role"`           //
	Token      string `gorm:"column:token"`          //
	LoginIp    string `gorm:"column:login_ip"`
	RegisterIp string `gorm:"column:register_ip"`
	Operator   int    `gorm:"column:operator"`
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
	GoogleAuth string `gorm:"column:google_auth"`                //
	RoleTab    Role   `gorm:"foreignKey:Role"`                   //
}

// TableName sets the insert table name for this struct type
func (a *Admin) TableName() string {
	return "c_admin"
}

func (m *Admin) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *Admin) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Admin) Get() bool {
	if this.ID != 0 {
		key := fmt.Sprintf(StringKeyAdmin, this.ID)
		//取redis
		s := global.REDIS.Get(key).Val()
		if s != "" {
			err := json.Unmarshal([]byte(s), this)
			if err == nil {
				return true
			}
		}
	}
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	key := fmt.Sprintf(StringKeyAdmin, this.ID)
	//同步redis
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.Set(key, string(bytes), this.ExpireTime())
	return true
}
func (this *Admin) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyAdmin, this.ID)
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
	if this.Token == "" {
		global.REDIS.Del(fmt.Sprintf(StringKeyAdmin, this.ID))
	} else {
		bytes, _ := json.Marshal(this)
		global.REDIS.Set(fmt.Sprintf(StringKeyAdmin, this.ID), string(bytes), this.ExpireTime())
	}
	return nil
}

func (this *Admin) Info() {
	//加密token
	//jwtService := extends.JwtUtils{}
	//token := jwtService.NewToken(this.ID, this.Token)

}

func (this *Admin) PageList(where string, args []interface{}, page, pageSize int) ([]Admin, common.Page) {
	res := make([]Admin, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64

	count := global.DB.Model(this).Joins("RoleTab").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("RoleTab").Where(where, args...).Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

func (this *Admin) Remove() error {
	res := global.DB.Model(this).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
