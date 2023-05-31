package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type OrderEquity struct {
	Id           int             `gorm:"column:id;primary_key"`             //
	UId          int             `gorm:"column:uid"`                        //关联用户id
	Pid          int             `gorm:"column:pid"`                        //关联商品种类id
	PayMoney     decimal.Decimal `gorm:"column:pay_money"`                  //购买付款金额 =手数
	AfterBalance decimal.Decimal `gorm:"column:after_balance"`              //购买后余额
	Rate         decimal.Decimal `gorm:"column:rate"`                       //中签率
	CreateTime   int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
	UpdateTime   int64           `gorm:"column:update_time;autoUpdateTime"` //系统开奖时间
	Member       Member          `gorm:"foreignKey:UId;"`                   //BeLongsTo 关联用户 自身外键UId
	Equity       Equity          `gorm:"foreignKey:Pid;"`                   //BeLongsTo 关联商品 自身外键Pid
}

// TableName sets the insert table name for this struct type
func (o OrderEquity) TableName() string {
	return "c_equity_order"
}

func (o *OrderEquity) Insert() error {
	res := global.DB.Create(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *OrderEquity) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

func (o *OrderEquity) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (o *OrderEquity) Count() (int64, error) {
	var count int64
	err := global.DB.Model(o).Where(o).Count(&count)
	if err.Error != nil {
		logrus.Error(err.Error)
		return 0, err.Error
	}
	return count, nil
}

func (o *OrderEquity) Sum() decimal.Decimal {
	var count decimal.Decimal
	err := global.DB.Model(o).Where(o).Pluck("COALESCE(SUM(pay_money),0) as count", &count)
	if err.Error != nil {
		logrus.Error(err.Error)
		return decimal.Zero
	}
	return count
}

func (this *OrderEquity) Sum2(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}

func (this *OrderEquity) List(where string, args []interface{}) []*OrderEquity {
	res := make([]*OrderEquity, 0)
	tx := global.DB.Model(this).Joins("Guquan").Joins("Member").Where(where, args...).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 订单列表 使用joins联合查询 或使用Preload 根据需求决定
func (this *OrderEquity) PageList(where string, args []interface{}, page, pageSize int) ([]OrderEquity, common.Page) {
	res := make([]OrderEquity, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Joins("Equity").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Joins("Equity").Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}
