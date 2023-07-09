package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Trade struct {
	Id         int             `gorm:"column:id;primary_key"`             //
	UId        int             `gorm:"column:uid"`                        //
	TradeType  int             `gorm:"column:trade_type"`                 //账单类型 1=购买产品  2=购买股权 3=充值 4=提现 5=可用转可提 6=可提转可用 7=注册买产品礼金 8=注册实名认证礼金 9=送优惠券 10=使用优惠券 11=余额宝转入 12=余额宝转出  13=余额宝收益 14=后台上分 15=后台下分 16=每日收益 17=股权收益 18=一级返佣 19=二级返佣 20=三级返佣
	ItemId     int             `gorm:"column:item_id"`                    //关联id
	Amount     decimal.Decimal `gorm:"column:amount"`                     //金额
	Before     decimal.Decimal `gorm:"column:before"`                     //
	After      decimal.Decimal `gorm:"column:after"`                      //
	Desc       string          `gorm:"column:desc"`                       //
	CreateTime int64           `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime int64           `gorm:"column:update_time;autoUpdateTime"` //
	IsFrontend int             `gorm:"column:is_frontend"`                //是否前端展示
	Member     Member          `gorm:"foreignKey:UId"`
}

// TableName sets the insert table name for this struct type
func (t Trade) TableName() string {
	return "c_trade"
}
func (this *Trade) Insert() error {
	res := global.DB.Create(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *Trade) CountByToday() (int64, error) {
	var total int64
	err := global.DB.Model(this).Where(this).Where("create_time >= ?", common.GetTodayZero()).Where("create_time <?", common.GetTodayZero()+3600*24).Count(&total).Error
	return total, err
}

func (this *Trade) PageList(where string, args []interface{}, page, pageSize int) ([]Trade, common.Page) {
	res := make([]Trade, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
func (this *Trade) Sum(where string, args []interface{}, field string) float64 {
	var total float64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
