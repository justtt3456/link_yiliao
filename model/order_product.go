package model

import (
	"finance/common"
	"finance/global"
	"time"

	"github.com/sirupsen/logrus"
)

type OrderProduct struct {
	ID              int     `gorm:"column:id;primary_key"`             //
	UID             int     `gorm:"column:uid"`                        //关联用户id
	Pid             int     `gorm:"column:pid"`                        //关联商品种类id
	PayMoney        int64   `gorm:"column:pay_money"`                  //购买付款金额(不含手续费)
	AfterBalance    int64   `gorm:"column:after_balance"`              //购买后余额
	CreateTime      int64   `gorm:"column:create_time;autoCreateTime"` //创建时间
	UpdateTime      int64   `gorm:"column:update_time;autoUpdateTime"` //系统开奖时间
	IsReturnTop     int64   `gorm:"column:is_return_top"`              //1=未返还上级 2=已反还上级
	IsReturnCapital int     `gorm:"column:is_return_capital"`          //是否返还本身 0:否 1:是
	IsReturnTeam    int     `gorm:"column:is_return_team"`             //是否已结算团队收益 0:否 1:是
	Member          Member  `gorm:"foreignKey:UID;"`                   //BeLongsTo 关联用户 自身外键UID
	Product         Product `gorm:"foreignKey:Pid;"`                   //BeLongsTo 关联商品 自身外键Pid
}

// TableName sets the insert table name for this struct type
func (o OrderProduct) TableName() string {
	return "c_product_order"
}

func (o *OrderProduct) ExpireTime() time.Duration {
	return time.Hour * 24 * 7
}

func (o *OrderProduct) Insert() error {
	res := global.DB.Create(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (this *OrderProduct) Get() bool {
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	return true
}

// 所有结算订单
func (this *OrderProduct) List(where string, args []interface{}) []OrderProduct {
	res := make([]OrderProduct, 0)
	tx := global.DB.Model(this).Where(where, args...).Order("id desc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 订单列表 使用joins联合查询 或使用Preload 根据需求决定
func (this *OrderProduct) PageList(where string, args []interface{}, page, pageSize int) ([]OrderProduct, common.Page) {
	res := make([]OrderProduct, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Joins("Product").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Joins("Product").Where(where, args...).Limit(pageSize).Offset(offset).Order(this.TableName() + ".id desc").Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	pageUtil.SetPage(pageSize, total)
	return res, pageUtil
}

func (o *OrderProduct) Update(col string, cols ...interface{}) error {
	res := global.DB.Select(col, cols...).Updates(o)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (this *OrderProduct) Count(where string, args []interface{}) int64 {
	var total int64
	tx := global.DB.Model(this).Where(where, args...).Count(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return int64(total)
}
func (this *OrderProduct) Sum(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}

func (this *OrderProduct) GetAll(today int64) []*OrderProduct {
	res := make([]*OrderProduct, 0)
	tx := global.DB.Model(this).Joins("Member").Joins("Product").Where("c_product_order.create_time < ?", today).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 获取有效的投资订单
func (this *OrderProduct) GetValidOrderList(today int64) []*OrderProduct {
	res := make([]*OrderProduct, 0)
	tx := global.DB.Model(this).Joins("Member").Joins("Product").Where("c_product_order.create_time < ? and c_product_order.is_return_capital = 0", today).Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return res
	}
	return res
}

// 获取团队订单用户ID列表(唯一)
func (this *OrderProduct) GetOrderUserIds(startTime, endTime int64) []int {
	res := make([]*OrderProduct, 0)
	var uids []int

	tx := global.DB.Model(this).Select("DISTINCT uid").Where("create_time >= ? and create_time <= ? and is_return_team = 0", startTime, endTime).Order("id asc").Find(&res)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return uids
	}

	//遍例订单列表
	if len(res) > 0 {
		for _, line := range res {
			uids = append(uids, line.UID)
		}
	}

	return uids
}

// 更改订单团队结算状态
func (this *OrderProduct) UpdateTeamSettleStatus(where string, args []interface{}) error {
	res := global.DB.Select("is_return_team").Where(where, args...).Updates(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
