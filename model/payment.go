package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
)

type Payment struct {
	Id             int    `gorm:"column:id;primary_key"`             //
	PayName        string `gorm:"column:pay_name"`                   //支付方式名称
	RechargeURL    string `gorm:"column:recharge_url"`               //充值提交地址
	WithdrawURL    string `gorm:"column:withdraw_url"`               //提现提交地址
	NotifyURL      string `gorm:"column:notify_url"`                 //回调地址
	MerchantNo     string `gorm:"column:merchant_no"`                //商户号
	Secret         string `gorm:"column:secret"`                     //密钥
	PriKey         string `gorm:"column:pri_key"`                    //私钥
	PubKey         string `gorm:"column:pub_key"`                    //公钥
	ClassName      string `gorm:"column:class_name"`                 //类名
	Type           int    `gorm:"column:type"`                       //1=微信  2=支付宝
	WithdrawStatus int    `gorm:"column:withdraw_status"`            //是否启用代付 1是2否
	CreateTime     int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime     int64  `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (p *Payment) TableName() string {
	return "c_payment"
}
func (m *Payment) ExpireTime() time.Duration {
	return time.Hour * 24 * 30
}
func (this *Payment) Get() bool {
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
	////同步redis
	//bytes, err := json.Marshal(this)
	//if err != nil {
	//	log.Println(err)
	//}
	//global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return true
}

func (this *Payment) Insert() error {
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
	global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return nil
}

func (this *Payment) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyPayment, this.Id)
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
	global.REDIS.HSet(HashKeyPayment, strconv.Itoa(this.Id), string(bytes))
	return nil
}
func (this Payment) List() []Payment {
	res := make([]Payment, 0)
	tx := global.DB.Model(this).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return res
}
func (this *Payment) PageList(where string, args []interface{}, page, pageSize int) ([]Payment, common.Page) {
	res := make([]Payment, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		find := global.DB.Model(this).Where(where, args...).Order("id desc").Offset(offset).Limit(pageSize).Find(&res)
		if find.Error != nil {
			logrus.Error(find.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Payment) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		return res.Error
	}
	//同步redis
	global.REDIS.HDel(HashKeyPayment, strconv.Itoa(this.Id))
	return nil
}
