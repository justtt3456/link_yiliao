package model

import (
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MemberBank struct {
	Id         int    `gorm:"column:id;primary_key"`             //
	UId        int    `gorm:"column:uid"`                        //关联用户id
	BankName   string `gorm:"column:bank_name"`                  //银行卡
	CardNumber string `gorm:"column:card_number"`                //卡号
	Province   string `gorm:"column:province"`                   //省份
	City       string `gorm:"column:city"`                       //市
	BranchBank string `gorm:"column:branch_bank"`                //开户行（开户所在地）
	RealName   string `gorm:"column:real_name"`                  //开户人
	IdNumber   string `gorm:"column:id_number"`                  //身份证号码
	BankPhone  string `gorm:"column:bank_phone"`                 //预留手机号码
	IsDefault  int    `gorm:"column:is_default"`                 //默认银行卡
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"` //
	Member     Member `gorm:"foreignKey:UId"`
}

// TableName sets the insert table name for this struct type
func (m *MemberBank) TableName() string {
	return "c_member_bank"
}
func (this *MemberBank) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberBank) Get() bool {
	res := global.DB.Model(this).Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *MemberBank) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMemberBank, this.Id)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberBank) ListByAgent(where string, args []interface{}) []MemberBank {
	res := make([]MemberBank, 0)
	tx := global.DB.Model(this).Joins("Member").Where(this).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}
func (this *MemberBank) List() []MemberBank {
	res := make([]MemberBank, 0)
	tx := global.DB.Model(this).Where(this).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}
func (this *MemberBank) Remove() error {
	res := global.DB.Delete(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
