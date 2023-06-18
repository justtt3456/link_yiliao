package service

import (
	"china-russia/app/api/swag/request"
	"china-russia/common"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

type EquityServiceBuy struct {
	request.EquityBuyRequest
}

func (this EquityServiceBuy) Buy(member *model.Member) error {
	//添加Redis乐观锁
	lockKey := fmt.Sprintf("equity_buy:%d", member.Id)
	redisLock := common.RedisLock{RedisClient: global.REDIS}
	if !redisLock.Lock(lockKey) {
		return errors.New(lang.Lang("During data processing, Please try again later"))
	}
	defer redisLock.Unlock(lockKey)
	//股权
	p := model.Equity{Id: this.Id}
	if !p.Get(true) {
		return errors.New("股权不存在！")
	}

	if p.PreStartTime > time.Now().Unix() {
		return errors.New("股权预售时间未开始")
	}
	if p.PreEndTime < time.Now().Unix() {
		return errors.New("股权预售时间已结束")
	}
	if p.MinBuy > this.Quantity {
		return errors.New(fmt.Sprintf("购买股权数量必须大于%v！", p.MinBuy))
	}
	amount := decimal.NewFromInt(int64(this.Quantity)).Mul(p.Price)
	if member.Balance.LessThan(amount) {
		return errors.New("余额不足,请先充值！")
	}
	//购买
	inc := &model.OrderEquity{
		UId:          member.Id,
		Pid:          p.Id,
		Price:        p.Price,
		Quantity:     this.Quantity,
		PayMoney:     amount,
		Rate:         decimal.NewFromInt(100),
		AfterBalance: member.Balance.Sub(amount),
		Status:       model.StatusReview,
	}
	err := inc.Insert()
	if err != nil {
		return err
	}
	//减去可投余额
	p.Current += int64(this.Quantity)
	err = p.Update("current")
	if err != nil {
		logrus.Errorf("购买产品减去可投余额失败%v", err)
		return err
	}
	//加入账变记录
	trade := model.Trade{
		UId:        member.Id,
		TradeType:  2,
		ItemId:     inc.Id,
		Amount:     amount,
		Before:     member.Balance,
		After:      member.Balance.Sub(amount),
		Desc:       "购买股权",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsFrontend: 1,
	}
	err = trade.Insert()
	if err != nil {
		logrus.Errorf("购买股权加入账变记录失败%v", err)
		return err
	}
	//扣减余额
	member.Balance = member.Balance.Sub(amount)
	member.Update("balance")
	if err != nil {
		logrus.Errorf("更改会员余额信息失败%v", err)
		return err
	}
	return nil
}
