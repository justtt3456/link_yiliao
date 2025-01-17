package model

import (
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type SetBase struct {
	Id                int             `gorm:"column:id;primary_key"`      //
	AppName           string          `gorm:"column:app_name"`            //网站应用名称
	AppLogo           string          `gorm:"column:app_logo"`            //网站应用logo
	VerifiedSend      decimal.Decimal `gorm:"column:verified_send"`       //实名送金币
	RegisterSend      decimal.Decimal `gorm:"column:register_send"`       //注册24小时后  第一次充值送金币
	OneSend           decimal.Decimal `gorm:"column:one_send"`            //一级奖励
	OneSendMoney      decimal.Decimal `gorm:"column:one_send_money"`      //三级代理享受现金奖励
	TwoSend           decimal.Decimal `gorm:"column:two_send"`            //二级奖励
	ThreeSend         decimal.Decimal `gorm:"column:three_send"`          //三级奖励
	SendDesc          string          `gorm:"column:send_desc"`           //奖励描述
	RegisterDesc      string          `gorm:"column:register_desc"`       //注册好礼描述
	TeamDesc          string          `gorm:"column:team_desc"`           //团队奖励描述
	GiftRate          decimal.Decimal `gorm:"column:gift_rate"`           //赠品赠送比例
	EquityStartDate   int64           `gorm:"column:equity_start_date"`   //股权分开始日期
	EquityRate        decimal.Decimal `gorm:"column:equity_rate"`         //股权分对应提现额度比例
	EquityInterval    int             `gorm:"column:equity_interval"`     //股权分收益时长
	EquityIncomeRate  decimal.Decimal `gorm:"column:equity_income_rate"`  //股权分收益比例
	OneReleaseRate    decimal.Decimal `gorm:"column:one_release_rate"`    //一级代理释放比例
	TwoReleaseRate    decimal.Decimal `gorm:"column:two_release_rate"`    //二级代理释放比例
	ThreeReleaseRate  decimal.Decimal `gorm:"column:three_release_rate"`  //三级代理释放比例
	IncomeBalanceRate decimal.Decimal `gorm:"column:income_balance_rate"` //收益转可用余额比例
	SignRewards       decimal.Decimal `gorm:"column:sign_rewards"`        //签到奖励
	DownloadUrl       string          `gorm:"column:download_url"`        //下载链接
	UsdtBuyRate       decimal.Decimal `gorm:"column:usdt_buy_rate"`       //usdt买汇率
	UsdtSellRate      decimal.Decimal `gorm:"column:usdt_sell_rate"`      //usdt卖汇率
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
