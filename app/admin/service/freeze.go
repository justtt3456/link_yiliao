package service

import "china-russia/model"

type FreezeHandle struct {
}

func (this FreezeHandle) Freeze(member model.Member, item int, amount float64, isfront int) error {
	//账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: model.TradeTypeFreeze,
		ItemId:    item,
		//Amount:     int64(amount * 100),
		Before: member.Balance,
		//After:      member.Balance + int64(amount*100),
		Desc:       "系统冻结",
		IsFrontend: isfront,
	}
	//更新余额
	//member.Balance -= int64(amount * 100)
	member.Update("balance", "freeze")
	return trade.Insert()
}

func (this FreezeHandle) Unfreeze(member model.Member, item int, amount float64, isfront int) error {
	//账单
	trade := model.Trade{
		UId:       member.Id,
		TradeType: model.TradeTypeUnfreeze,
		ItemId:    item,
		//Amount:     int64(amount * 100),
		Before: member.Balance,
		//After:      member.Balance + int64(amount*100),
		Desc:       "系统解除冻结",
		IsFrontend: isfront,
	}
	//更新余额
	//member.Balance += int64(amount * 100)
	member.Update("balance", "freeze")
	return trade.Insert()
}
