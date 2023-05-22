package model

import (
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type SetFunds struct {
	Id                  int             `gorm:"column:id;primary_key"`        //
	RechargeStartTime   string          `gorm:"column:recharge_start_time"`   //充值开始时间
	RechargeEndTime     string          `gorm:"column:recharge_end_time"`     //充值结束时间
	RechargeMinAmount   decimal.Decimal `gorm:"column:recharge_min_amount"`   //充值最小金额
	RechargeMaxAmount   decimal.Decimal `gorm:"column:recharge_max_amount"`   //充值最大金额
	RechargeFee         int             `gorm:"column:recharge_fee"`          //充值手续费(百分比)
	RechargeQuickAmount string          `gorm:"column:recharge_quick_amount"` //快捷充值金额
	WithdrawStartTime   string          `gorm:"column:withdraw_start_time"`   //提现开始时间
	WithdrawEndTime     string          `gorm:"column:withdraw_end_time"`     //提现结束时间
	MustPassword        int             `gorm:"column:must_password"`         //是否必须体现密码
	PasswordFreeze      int             `gorm:"column:password_freeze"`       //提现密码错误冻结次数
	WithdrawMinAmount   decimal.Decimal `gorm:"column:withdraw_min_amount"`   //提现最小金额
	WithdrawMaxAmount   decimal.Decimal `gorm:"column:withdraw_max_amount"`   //提现最大金额
	WithdrawFee         decimal.Decimal `gorm:"column:withdraw_fee"`          //提现手续费
	WithdrawCount       int             `gorm:"column:withdraw_count"`        //每日提现次数
	ProductFee          int             `gorm:"column:product_fee"`           //购买产品手续费
	ProductQuickAmount  string          `gorm:"column:product_quick_amount"`  //购买产品快捷金额
	DayTurnMoneyNum     int64           `gorm:"column:day_turn_money_num"`    //每日 可用和可提互转次数
}

// TableName sets the insert table name for this struct type
func (s *SetFunds) TableName() string {
	return "c_set_funds"
}
func (this SetFunds) Insert() error {
	err := global.DB.Create(this)
	if err != nil {
		logrus.Error(err.Error)
		return err.Error
	}
	return nil
}
func (this *SetFunds) Get() bool {
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *SetFunds) Update() error {
	//全部更新
	res := global.DB.Save(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
