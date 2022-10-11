package model

import (
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Role struct {
	RoleID     int    `gorm:"column:role_id;primary_key"`        //
	RoleName   string `gorm:"column:role_name"`                  //
	Status     int    `gorm:"column:status"`                     //
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "c_role"
}
func (this *Role) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Role) Get() bool {
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *Role) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyRole, this.RoleID)
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
	//if this.AuthToken == "" {
	//	global.REDIS.Del(fmt.Sprintf(StringKeyMember, this.RoleID))
	//} else {
	//	bytes, _ := json.Marshal(this)
	//	global.REDIS.Set(fmt.Sprintf(StringKeyMember, this.RoleID), string(bytes), time.Hour*24*30)
	//}
	return nil
}
func (this *Role) PageList(page, pageSize int) ([]Role, common.Page) {
	res := make([]Role, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		find := global.DB.Where(this).Order("role_id desc").Offset(offset).Limit(pageSize).Find(&res)
		if find.Error != nil {
			logrus.Error(find.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Role) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
