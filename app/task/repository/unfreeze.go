package repository

import (
	"china-russia/model"
	"time"
)

type Unfreeze struct {
}

func (this Unfreeze) Do() {
	now := int(time.Now().Unix())
	out := model.InvestOrder{}
	outWhere := "type = ? and unfreeze_status = ? and unfreeze_time <= ?"
	outArgs := []interface{}{1, 2, now}
	outList := out.List(outWhere, outArgs)
	for _, iv := range outList {
		this.unfreeze(iv)
	}
}

// 解除冻结余额宝余额
func (this Unfreeze) unfreeze(order model.InvestOrder) {
	member := model.Member{Id: order.UId}
	if !member.Get() {
		return
	}
	order.UnfreezeStatus = 1
	order.Update("unfreeze_status")
	member.InvestAmount += order.Amount
	member.InvestFreeze -= order.Amount
	member.Update("invest_amount", "invest_freeze")
}
