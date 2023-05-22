package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"time"

	"github.com/sirupsen/logrus"
)

type Recharge struct {
	Id             int             `gorm:"column:id;primary_key"`             //
	OrderSn        string          `gorm:"column:order_sn"`                   //
	UId            int             `gorm:"column:uid"`                        //关联用户id
	Type           int             `gorm:"column:type"`                       //充值类别 1银行卡 2在线充值 3后台充值
	Amount         decimal.Decimal `gorm:"column:amount"`                     //充值金额
	RealAmount     decimal.Decimal `gorm:"column:real_amount"`                //实际到账金额
	From           string          `gorm:"column:from"`                       //付款账号
	To             string          `gorm:"column:to"`                         //收款账号
	Voucher        string          `gorm:"column:voucher"`                    //凭证图
	PaymentId      int             `gorm:"column:payment_id"`                 //三方支付id
	Status         int             `gorm:"column:status"`                     //状态，0待审核，1已审核
	UsdtAmount     decimal.Decimal `gorm:"column:usdt_amount"`                //usdt充值数量
	Operator       int             `gorm:"column:operator"`                   //操作的管理员
	Description    string          `gorm:"column:description"`                //订单备注
	SuccessTime    int64           `gorm:"column:success_time"`               //成功时间
	TradeSn        string          `gorm:"column:trade_sn"`                   //三方订单号
	UpdateTime     int64           `gorm:"column:update_time;autoCreateTime"` //审核时间
	CreateTime     int64           `gorm:"column:create_time;autoUpdateTime"` //创建时间
	Member         Member          `gorm:"foreignKey:UId"`
	RechargeMethod RechargeMethod  `gorm:"foreignKey:Type"`
	Payment        Payment         `gorm:"foreignKey:PaymentId"`
	Admin          Admin           `gorm:"foreignKey:Id"`
	MemberVerified MemberVerified  `gorm:"foreignKey:UId;references:UId"`
	ImageUrl       string          `gorm:"column:img_url"` //凭证图片网址
}

// TableName sets the insert table name for this struct type
func (r Recharge) TableName() string {
	return "c_recharge"
}

func (r *Recharge) Insert() error {
	res := global.DB.Create(r)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}

func (r *Recharge) ExpireTime() time.Duration {
	return time.Minute * 30
}

func (this *Recharge) GetPageList(where string, args []interface{}, page, pageSize int) ([]Recharge, common.Page) {

	res := make([]Recharge, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(this).Joins("Member").Joins("RechargeMethod").Joins("Admin").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	pageUtil.SetPage(pageSize, total)
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(this).Joins("Member").Joins("RechargeMethod").Joins("Admin").Joins("MemberVerified").Where(where, args...).
			Order(this.TableName() + ".id desc").Limit(pageUtil.PageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	return res, pageUtil
}

func (r *Recharge) Update(col string, cols ...interface{}) error {
	rds := Redis{}
	key := fmt.Sprintf(LockKeyRecharge, r.Id)
	if err := rds.Lock(key); err != nil {
		return err
	}
	defer rds.Unlock(key)
	res := global.DB.Select(col, cols...).Updates(r)

	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	bytes, _ := json.Marshal(r)
	global.REDIS.Set(fmt.Sprintf(StringKeyRecharge, r.Id), string(bytes), r.ExpireTime())
	return nil
}

func (r *Recharge) Count(where string, args []interface{}) int64 {
	var total int64
	tx := global.DB.Model(r).Joins("Member").Where(where, args...).Count(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return int64(total)
}
func (r *Recharge) Sum(where string, args []interface{}, field string) int64 {
	var total int64
	tx := global.DB.Model(r).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
func (this *Recharge) Get() bool {

	//取数据库
	res := global.DB.Where(this).Joins("Member").First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}

	return true
}

// 获取会员充值总人数
func (r *Recharge) GetMemberCount(where string, args []interface{}) int {
	var total int
	tx := global.DB.Model(r).Select("COUNT(DISTINCT uid)").Where(where, args...).First(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
