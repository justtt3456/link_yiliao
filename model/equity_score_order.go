package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type EquityScoreOrder struct {
	Id         int             `gorm:"column:id;primary_key"` //
	UId        int             `gorm:"column:uid"`            //关联用户id
	PayMoney   decimal.Decimal `gorm:"column:pay_money"`      //购买付款金额 =手数
	Rate       decimal.Decimal `gorm:"column:rate"`           //
	Interval   int             `gorm:"column:interval"`
	Status     int             `gorm:"column:status"`                     //状态
	CreateTime int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
	UpdateTime int64           `gorm:"column:update_time;autoUpdateTime"` //系统开奖时间
	EndTime    int64           `gorm:"column:end_time"`
	Member     Member          `gorm:"foreignKey:UId;"` //BeLongsTo 关联用户 自身外键UId
}

// TableName sets the insert table name for this struct type
func (o EquityScoreOrder) TableName() string {
	return "c_equity_score_order"
}

func (o *EquityScoreOrder) Insert() error {
	res := global.DB.Create(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *EquityScoreOrder) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (o *EquityScoreOrder) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (o *EquityScoreOrder) Count() (int64, error) {
	var count int64
	err := global.DB.Model(o).Where(o).Count(&count)
	if err.Error != nil {
		logrus.Error(err.Error)
		return 0, err.Error
	}
	return count, nil
}

func (this *EquityScoreOrder) SumScore(where string, args []interface{}, field string) decimal.Decimal {
	var total float64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return decimal.Zero
	}
	return decimal.NewFromFloat(total)
}

func (this *EquityScoreOrder) List(where string, args []interface{}) []EquityScoreOrder {
	res := make([]EquityScoreOrder, 0)
	tx := global.DB.Model(this).Joins("Member").Where(where, args...).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}

// 订单列表 使用joins联合查询 或使用Preload 根据需求决定
func (this *EquityScoreOrder) PageList(where string, args []interface{}, page, pageSize int) ([]EquityScoreOrder, common.Page) {
	res := make([]EquityScoreOrder, 0)
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
		tx := global.DB.Model(this).Joins("Member").Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
