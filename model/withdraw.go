package model

import (
	"china-russia/common"
	"china-russia/global"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

type Withdraw struct {
	Id             int             `gorm:"column:id;primary_key"`             //
	UId            int             `gorm:"column:uid"`                        //关联用户id
	WithdrawType   int             `gorm:"column:withdraw_type"`              //提现类型1=银行卡
	BankName       string          `gorm:"column:bank_name"`                  //关联银行名称
	BankCode       string          `gorm:"column:bank_code"`                  //关联银行名称
	BranchBank     string          `gorm:"column:branch_bank"`                //开户行
	RealName       string          `gorm:"column:real_name"`                  //开户人
	CardNumber     string          `gorm:"column:card_number"`                //卡号
	BankPhone      string          `gorm:"column:bank_phone"`                 //预留手机号码
	Amount         decimal.Decimal `gorm:"column:amount"`                     //实际到账金额
	Fee            decimal.Decimal `gorm:"column:fee"`                        //手续费
	TotalAmount    decimal.Decimal `gorm:"column:total_amount"`               //提现总额
	UsdtAmount     decimal.Decimal `gorm:"column:usdt_amount"`                //提现总额
	Description    string          `gorm:"column:description"`                //审核备注
	Operator1      int             `gorm:"column:operator"`                   //操作管理员
	ViewStatus     int             `gorm:"column:view_status"`                //已读状态，0=未读，1=已读
	Status         int             `gorm:"column:status"`                     //提现状态，0为未审核，1为已审核，2为已拒绝
	SuccessTime    int64           `gorm:"column:success_time"`               //成功时间
	OrderSn        string          `gorm:"column:order_sn"`                   //订单号
	PaymentId      int             `gorm:"column:payment_id"`                 //三方支付id
	TradeSn        string          `gorm:"column:trade_sn"`                   //三方订单号
	CreateTime     int64           `gorm:"column:create_time;autoCreateTime"` //
	UpdateTime     int64           `gorm:"column:update_time;autoUpdateTime"` //
	Member         Member          `gorm:"foreignKey:UId"`
	Admin          Admin           `gorm:"foreignKey:Operator1"`
	Payment        Payment         `gorm:"foreignKey:PaymentId"`
	WithdrawMethod WithdrawMethod  `gorm:"foreignKey:WithdrawType"`
}

// TableName sets the insert table name for this struct type
func (w Withdraw) TableName() string {
	return "c_withdraw"
}

func (w *Withdraw) ExpireTime() time.Duration {
	return time.Minute * 30
}

func (w *Withdraw) Insert() error {
	res := global.DB.Create(w)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
func (w *Withdraw) GetPageList(where string, args []interface{}, page, pageSize int) ([]Withdraw, common.Page) {
	res := make([]Withdraw, 0)
	pageUtil := common.Page{
		Page: page,
	}
	var total int64
	count := global.DB.Model(w).Joins("Member").Joins("Admin").Joins("Payment").Joins("WithdrawMethod").Where(where, args...).Count(&total)
	if count.Error != nil {
		logrus.Error(count.Error)
		return res, pageUtil
	}
	pageUtil.SetPage(pageSize, total)
	if total > 0 {
		offset := (page - 1) * pageSize
		tx := global.DB.Model(w).Joins("Member").Joins("Admin").Joins("Payment").Joins("WithdrawMethod").Where(where, args...).Order(w.TableName() + ".id desc").Limit(pageUtil.PageSize).Offset(offset).Find(&res)
		if tx.Error != nil {
			logrus.Error(tx.Error)
			return res, pageUtil
		}
	}
	return res, pageUtil
}

func (w *Withdraw) Update(col string, cols ...interface{}) error {
	r := Redis{}
	key := fmt.Sprintf(LockKeyWithdraw, w.Id)

	if err := r.Lock(key); err != nil {
		return err
	}
	defer r.Unlock(key)

	res := global.DB.Select(col, cols...).Updates(w)

	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	//同步redis
	// if w.Id == nil {
	// 	global.REDIS.Del(fmt.Sprintf(StringKeyWithdraw, w.Id))
	// } else {
	bytes, _ := json.Marshal(w)

	global.REDIS.Set(fmt.Sprintf(StringKeyWithdraw, w.Id), string(bytes), w.ExpireTime())
	// }
	return nil
}

func (this *Withdraw) Count(where string, args []interface{}) int64 {
	var total int64
	tx := global.DB.Model(this).Joins("Member").Where(where, args...).Count(&total)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return int64(total)
}
func (this *Withdraw) Sum(where string, args []interface{}, field string) float64 {
	var total float64
	tx := global.DB.Model(this).Select("COALESCE(sum("+field+"),0)").Where(where, args...).Scan(&total)
	//COALESCE(SUM(column),0)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return 0
	}
	return total
}
func (this *Withdraw) Get() bool {
	if this.Id != 0 {
		key := fmt.Sprintf(StringKeyWithdraw, this.Id)
		//取redis
		s := global.REDIS.Get(key).Val()
		if s != "" {
			err := json.Unmarshal([]byte(s), this)
			if err == nil {
				return true
			}
		}
	}
	//取数据库
	res := global.DB.Where(this).First(this)
	if res.Error != nil {
		logrus.Error(res.Error)
		return false
	}
	key := fmt.Sprintf(StringKeyWithdraw, this.Id)
	//同步redis
	bytes, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
	}
	global.REDIS.Set(key, string(bytes), this.ExpireTime())
	return true
}
