package model

import (
	"finance/global"
	"github.com/sirupsen/logrus"
)

type SetBase struct {
	ID           int    `gorm:"column:id;primary_key"` //
	AppName      string `gorm:"column:app_name"`       //网站应用名称
	AppLogo      string `gorm:"column:app_logo"`       //网站应用logo
	VerifiedSend int    `gorm:"column:verified_send"`  //实名送金币
	RegisterSend int    `gorm:"column:register_send"`  //注册24小时后  第一次充值送金币
	OneSend      int    `gorm:"column:one_send"`       //一级奖励
	OneSendMoeny int64  `gorm:"column:one_send_moeny"` //三级代理享受现金奖励
	TwoSend      int    `gorm:"column:two_send"`       //二级奖励
	ThreeSend    int    `gorm:"column:three_send"`     //三级奖励
	SendDesc     string `gorm:"column:send_desc"`      //奖励描述
	RegisterDesc string `gorm:"column:register_desc"`  //注册好礼描述
	TeamDesc     string `gorm:"column:team_desc"`      //团队奖励描述
	GiftRate     int    `gorm:"column:gift_rate"`      //赠品赠送比例
}

// TableName sets the insert table name for this struct type
func (s *SetBase) TableName() string {
	return "c_set_base"
}
func (this SetBase) Insert() error {
	err := global.DB.Create(this)
	if err != nil {
		logrus.Error(err.Error)
		return err.Error
	}
	return nil
}
func (this *SetBase) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *SetBase) Update() error {
	//全部更新
	res := global.DB.Save(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
