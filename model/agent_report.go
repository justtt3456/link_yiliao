package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

type AgentReport struct {
	Id             int    `gorm:"column:id;primary_key"`             //
	Aid            int    `gorm:"column:aid"`                        //代理id
	Username       string `json:"username"`                          //代理名称
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
	Freeze         int64  `gorm:"column:freeze"`                     //系统冻结
	Unfreeze       int64  `gorm:"column:unfreeze"`                   //系统解冻
	RegisterCount  int    `json:"register_count"`
}

// TableName sets the insert table name for this struct type
func (a *AgentReport) TableName() string {
	return "c_agent_report"
}
func (m *AgentReport) GetPageList(where string, args []interface{}, page, pageSize int) ([]AgentReport, common.Page) {
	res := make([]AgentReport, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(m).Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	pageUtil.SetPage(pageSize, total)
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Where(where, args...).Limit(pageUtil.PageSize).Offset(offset).Order("id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}

	return res, pageUtil
}
func (this *AgentReport) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *AgentReport) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (this *AgentReport) Update() error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyAgentReport, this.Id)
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
