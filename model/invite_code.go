package model

import (
	"encoding/json"
	"finance/common"
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type InviteCode struct {
	ID         int    `gorm:"column:id;primary_key"`             //
	UID        int    `gorm:"column:uid"`                        //用户id
	Username   string `gorm:"column:username"`                   //
	AgentID    int    `gorm:"column:agent_id"`                   //代理id
	AgentName  string `gorm:"column:agent_name"`                 //
	Code       string `gorm:"column:code"`                       //邀请码
	RegCount   int    `gorm:"column:reg_count"`                  //注册人数
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (i *InviteCode) TableName() string {
	return "c_invite_code"
}

func (m *InviteCode) ExpireTime() time.Duration {
	return time.Hour * 24 * 7
}

func (i *InviteCode) Insert() error {
	res := global.DB.Create(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (i *InviteCode) Get() bool {
	if i.Code != "" {
		key := fmt.Sprintf(StringKeyInviteCode, i.Code)
		//取redis
		s := global.REDIS.Get(key).Val()
		if s != "" {
			err := json.Unmarshal([]byte(s), i)
			if err == nil {
				return true
			}
		}
	}
	//取数据库
	res := global.DB.Where(i).First(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	key := fmt.Sprintf(StringKeyInviteCode, i.Code)
	//同步redis
	bytes, err := json.Marshal(i)
	if err != nil {
		logrus.Error(res.Error)
	}
	global.REDIS.Set(key, string(bytes), i.ExpireTime())
	return true
}

func (i *InviteCode) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyInviteCode, i.ID)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	bytes, _ := json.Marshal(i)
	global.REDIS.Set(fmt.Sprintf(StringKeyInviteCode, i.Code), string(bytes), i.ExpireTime())
	return nil
}

func (i *InviteCode) Remove() error {
	res := global.DB.Delete(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	global.REDIS.Del(fmt.Sprintf(StringKeyInviteCode, i.Code))
	return nil
}

func (i InviteCode) Info() *InviteCode {
	return &InviteCode{
		ID:         i.ID,
		UID:        i.UID,
		AgentID:    i.AgentID,
		Code:       i.Code,
		RegCount:   i.RegCount,
		CreateTime: i.CreateTime,
	}
}

func (i *InviteCode) PageList(where string, args []interface{}, page, pageSize int) ([]InviteCode, common.Page) {
	res := make([]InviteCode, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64

	count := global.DB.Model(i).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Where(where, args...).Limit(pageUtil.PageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
