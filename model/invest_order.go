package model

import (
	"china-russia/common"
	"china-russia/global"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type InvestOrder struct {
	Id             int             `gorm:"column:id;primary_key"`             //
	UId            int             `gorm:"column:uid"`                        //关联用户id
	Type           int             `gorm:"column:type"`                       //转入转出类型 1转入 2转出
	Amount         decimal.Decimal `gorm:"column:amount"`                     //转入转出金额
	Rate           int             `gorm:"column:rate"`                       //收益比例
	CreateTime     int64           `gorm:"column:create_time;autoCreateTime"` //投入时间
	UnfreezeTime   int64           `gorm:"column:unfreeze_time"`              //冻结结束时间
	IncomeTime     int64           `gorm:"column:income_time"`                //可以发放奖励的首次时间
	Balance        decimal.Decimal `gorm:"column:balance"`                    //余额宝余额
	UnfreezeStatus int             `gorm:"column:unfreeze_status"`            //解冻状态 1已解冻 2冻结中
	Member         Member          `gorm:"foreignKey:UId"`
}

// TableName sets the insert table name for this struct type
func (i *InvestOrder) TableName() string {
	return "c_invest_order"
}
func (i *InvestOrder) Insert() error {
	res := global.DB.Create(i)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *InvestOrder) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}
func (this *InvestOrder) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyMessage, this.Id)
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
func (this *InvestOrder) PageList(where string, args []interface{}, page, pageSize int) ([]InvestOrder, common.Page) {
	res := make([]InvestOrder, 0)
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
func (this *InvestOrder) List(where string, args []interface{}) []InvestOrder {
	res := make([]InvestOrder, 0)
	tx := global.DB.Model(this).Where(where, args...).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return nil
	}
	return res
}
func (r *InvestOrder) Sum(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(r).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
