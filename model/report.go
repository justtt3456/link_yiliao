package model

import (
	"finance/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Report struct {
	ID                  int   `gorm:"column:id;primary_key"`             //
	RechargeCount       int   `gorm:"column:recharge_count"`             //充值次数
	RechargeCountTotal  int   `gorm:"column:recharge_count_total"`       //
	RechargeAmount      int64 `gorm:"column:recharge_amount"`            //充值金额
	RechargeAmountTotal int64 `gorm:"column:recharge_amount_total"`      //
	WithdrawCount       int   `gorm:"column:withdraw_count"`             //提现次数
	WithdrawCountTotal  int   `gorm:"column:withdraw_count_total"`       //
	WithdrawAmount      int64 `gorm:"column:withdraw_amount"`            //提现金额
	WithdrawAmountTotal int64 `gorm:"column:withdraw_amount_total"`      //
	BetCount            int   `gorm:"column:bet_count"`                  //投注次数
	BetCountTotal       int   `gorm:"column:bet_count_total"`            //
	BetAmount           int64 `gorm:"column:bet_amount"`                 //投注金额
	BetAmountTotal      int64 `gorm:"column:bet_amount_total"`           //
	BetResult           int64 `gorm:"column:bet_result"`                 //输赢
	BetResultTotal      int64 `gorm:"column:bet_result_total"`           //
	SysUp               int64 `gorm:"column:sys_up"`                     //系统上分
	SysUpTotal          int64 `gorm:"column:sys_up_total"`               //
	SysDown             int64 `gorm:"column:sys_down"`                   //系统下分
	SysDownTotal        int64 `gorm:"column:sys_down_total"`             //
	Freeze              int64 `gorm:"column:freeze"`                     //系统冻结
	FreezeTotal         int64 `gorm:"column:freeze_total"`               //
	Unfreeze            int64 `gorm:"column:unfreeze"`                   //系统解冻
	UnfreezeTotal       int64 `gorm:"column:unfreeze_total"`             //
	RegisterCount       int   `gorm:"column:register_count"`             //新增用户
	RegisterCountTotal  int   `gorm:"column:register_count_total"`       //
	Balance             int64 `gorm:"column:balance"`                    //总会员余额
	InvestAmount        int64 `gorm:"column:invest_amount"`              //余额宝可用余额
	InvestFreeze        int64 `gorm:"column:invest_freeze"`              //余额宝冻结余额
	InvestIncome        int64 `gorm:"column:invest_income"`              //余额宝收益
	CreateTime          int64 `gorm:"column:create_time"`                //
	UpdateTime          int64 `gorm:"column:update_time;autoUpdateTime"` //
}

// TableName sets the insert table name for this struct type
func (r *Report) TableName() string {
	return "c_report"
}
func (this *Report) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *Report) Get() bool {
	//取数据库
	res := global.DB.Where(this).Order("id desc").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *Report) Update() error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyReport, this.ID)
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
