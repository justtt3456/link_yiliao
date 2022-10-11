package service

import "finance/model"

type FreezeHandle struct {
}

func (this FreezeHandle) Freeze(member model.Member, item int, amount float64, isfront int) error {
	//账单
	trade := model.Trade{
		UID:        member.ID,
		TradeType:  model.TradeTypeFreeze,
		ItemID:     item,
		Amount:     int64(amount * 100),
		Before:     member.Balance,
		After:      member.Balance + int64(amount*100),
		Desc:       "系统冻结",
		IsFrontend: isfront,
	}
	//更新余额
	member.Balance -= int64(amount * 100)
	member.Update("balance", "freeze")
	return trade.Insert()
}

func (this FreezeHandle) Unfreeze(member model.Member, item int, amount float64, isfront int) error {
	//账单
	trade := model.Trade{
		UID:        member.ID,
		TradeType:  model.TradeTypeUnfreeze,
		ItemID:     item,
		Amount:     int64(amount * 100),
		Before:     member.Balance,
		After:      member.Balance + int64(amount*100),
		Desc:       "系统解除冻结",
		IsFrontend: isfront,
	}
	//更新余额
	member.Balance += int64(amount * 100)
	member.Update("balance", "freeze")
	return trade.Insert()
}
