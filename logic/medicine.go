package logic

import (
	"china-russia/global"
	"china-russia/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type medicineBuyLogic struct {
	tx     *gorm.DB
	config *model.SetBase
}

func NewMedicineBuyLogic() *medicineBuyLogic {
	//基础配置表
	config := model.SetBase{}
	config.Get()
	return &medicineBuyLogic{
		tx:     global.DB.Begin(),
		config: &config,
	}
}
func (this medicineBuyLogic) MedicineBuy(member *model.Member, medicine model.Medicine, amount decimal.Decimal, addressId int, quantity int) error {
	if this.tx == nil {
		this.tx = global.DB
	}
	if this.config == nil {
		config := model.SetBase{}
		config.Get()
		this.config = &config
	}
	var err error
	defer func() {
		if err != nil {
			this.tx.Rollback()
		} else {
			this.tx.Commit()
		}
	}()
	//订单入库
	order, err := this.createOrder(member, medicine, amount, addressId, quantity)
	if err != nil {
		return err
	}
	//账变记录
	err = this.createTrade(member, order, amount)
	if err != nil {
		return err
	}
	//扣减可用余额 记录已购状态
	member.Balance = member.Balance.Sub(amount)
	//用户可提额度增加
	//member.WithdrawThreshold = member.WithdrawThreshold.Add(medicine.WithdrawThresholdRate.Mul(amount).Div(decimal.NewFromInt(100)).Round(2))
	member.EquityScore += int(amount.IntPart())
	err = this.tx.Select("balance", "equity_score").Updates(member).Error
	if err != nil {
		logrus.Errorf("更改会员余额信息失败%v", err)
		return err
	}
	return nil
}
func (this medicineBuyLogic) createOrder(member *model.Member, medicine model.Medicine, amount decimal.Decimal, addressId int, quantity int) (*model.MedicineOrder, error) {
	//购买
	inc := &model.MedicineOrder{
		UId:               member.Id,
		Pid:               medicine.Id,
		WithdrawThreshold: medicine.WithdrawThreshold.Mul(decimal.NewFromInt(int64(quantity))),
		Interval:          medicine.Interval,
		Status:            model.StatusOk,
		PayMoney:          amount,
		AfterBalance:      member.Balance.Sub(amount),
		AddressId:         addressId,
		CreateTime:        time.Now().Unix(),
	}
	err := this.tx.Create(&inc).Error
	if err != nil {
		return nil, err
	}
	return inc, nil
}
func (this medicineBuyLogic) createTrade(member *model.Member, order *model.MedicineOrder, amount decimal.Decimal) error {
	//加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  1,
		ItemId:     order.Id,
		Amount:     amount,
		Before:     member.Balance,
		After:      member.Balance.Sub(amount),
		Desc:       "购买药品",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err := this.tx.Create(&trade).Error
	if err != nil {
		logrus.Errorf("购买药品加入账变记录失败%v", err)
		return err
	}
	return nil
}
