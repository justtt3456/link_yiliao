package model

import (
	"finance/common"
	"finance/global"
	"fmt"

	"github.com/sirupsen/logrus"
)

type MemberReport struct {
	ID             int    `gorm:"column:id;primary_key"`             //
	UID            int    `gorm:"column:uid"`                        //
	Username       string `gorm:"column:username"`                   //用户名
	RechargeCount  int    `gorm:"column:recharge_count"`             //充值次数
	RechargeAmount int64  `gorm:"column:recharge_amount"`            //充值金额
	WithdrawCount  int    `gorm:"column:withdraw_count"`             //提现次数
	WithdrawAmount int64  `gorm:"column:withdraw_amount"`            //提现金额
	BetCount       int    `gorm:"column:bet_count"`                  //投注次数
	BetAmount      int64  `gorm:"column:bet_amount"`                 //投注金额
	BetResult      int64  `gorm:"column:bet_result"`                 //输赢
	SysUp          int64  `gorm:"column:sys_up"`                     //系统上分
	SysDown        int64  `gorm:"column:sys_down"`                   //系统下分
	CreateTime     int64  `gorm:"column:create_time"`                //
	UpdateTime     int64  `gorm:"column:update_time;autoUpdateTime"` //
	Aid            int    `gorm:"column:aid"`                        //代理id
	Freeze         int64  `gorm:"column:freeze"`                     //系统冻结
	Unfreeze       int64  `gorm:"column:unfreeze"`                   //系统解冻
}

// TableName sets the insert table name for this struct type
func (m *MemberReport) TableName() string {
	return "c_member_report"
}

// GetPageList get member report list
func (m *MemberReport) PageList(where string, args []interface{}, page, pageSize int) ([]MemberReport, common.Page) {
	res := make([]MemberReport, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(&m).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	pageUtil.SetPage(pageSize, total)
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Where(where, args...).Order("id desc").Limit(pageUtil.PageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	return res, pageUtil
}
func (this *MemberReport) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MemberReport) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *MemberReport) Update() error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMemberReport, this.ID)
	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)
	res := global.DB.Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

//func (this *MemberReport) Sum(where string, args []interface{}) MemberReport {
//	var res MemberReport
//	tx := global.DB.Model(this).Select("sum(recharge_count) recharge_count,sum(recharge_amount) recharge_amount,sum(withdraw_count) withdraw_count,sum(withdraw_amount) withdraw_amount,sum(bet_count) bet_count,sum(bet_amount) bet_amount,sum(bet_result) bet_result,sum(sys_up) sys_up,sum(sys_down) sys_down,sum(freeze) freeze,sum(unfreeze) unfreeze,uid,username").Where(where, args...).First(&res)
//	if tx.Error != nil {
//		logrus.Error(tx.Error)
//		return res
//	}
//	return res
//}
