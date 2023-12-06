package model

import (
	"china-russia/common"
	"china-russia/global"
	"github.com/shopspring/decimal"
	"time"

	"github.com/sirupsen/logrus"
)

type MedicineOrder struct {
	Id                int             `gorm:"column:id;primary_key"`     //
	UId               int             `gorm:"column:uid"`                //关联用户id
	Pid               int             `gorm:"column:pid"`                //关联商品种类id
	WithdrawThreshold decimal.Decimal `gorm:"column:withdraw_threshold"` //
	Interval          int             `gorm:"column:interval"`
	Status            int             `gorm:"column:status"` //状态
	Current           int             `gorm:"column:current"`
	PayMoney          decimal.Decimal `gorm:"column:pay_money"`     //购买付款金额(不含手续费)
	AfterBalance      decimal.Decimal `gorm:"column:after_balance"` //购买后余额
	AddressId         int             `gorm:"column:address_id"`
	CreateTime        int64           `gorm:"column:create_time;autoCreateTime"` //创建时间
	Member            Member          `gorm:"foreignKey:UId;"`                   //BeLongsTo 关联用户 自身外键UId
	Medicine          Medicine        `gorm:"foreignKey:Pid;"`                   //BeLongsTo 关联商品 自身外键Pid
	Address           MemberAddress   `gorm:"foreignKey:AddressId;"`
}

// TableName sets the insert table name for this struct type
func (o MedicineOrder) TableName() string {
	return "c_medicine_order"
}

func (o *MedicineOrder) ExpireTime() time.Duration {
	return time.Hour * 24 * 7
}

func (o *MedicineOrder) Insert() error {
	res := global.DB.Create(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *MedicineOrder) Get() bool {
	//取数据库
	res := global.DB.Where(this).Order("id desc").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

// 所有结算订单
func (this *MedicineOrder) List(where string, args []interface{}) []MedicineOrder {
	res := make([]MedicineOrder, 0)
	tx := global.DB.Model(this).Where(where, args...).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 订单列表 使用joins联合查询 或使用Preload 根据需求决定
func (this *MedicineOrder) PageList(where string, args []interface{}, page, pageSize int) ([]MedicineOrder, common.Page) {
	res := make([]MedicineOrder, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Joins("Medicine").Joins("Address").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Joins("Medicine").Joins("Address").Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

func (o *MedicineOrder) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *MedicineOrder) Count(where string, args []interface{}) int64 {
	var total int64
	tx := global.DB.Model(this).Where(where, args...).Count(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return int64(total)
}
func (this *MedicineOrder) Sum(where string, args []interface{}, field string) float64 {
	var total float64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}

func (this *MedicineOrder) GetAll(today int64) []*MedicineOrder {
	res := make([]*MedicineOrder, 0)
	tx := global.DB.Model(this).Joins("Member").Joins("Medicine").Where("c_medicine_order.create_time < ?", today).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}
