package request

import "github.com/shopspring/decimal"

type InvestOrder struct {
	Amount decimal.Decimal `json:"amount"` //转入转出金额
	Type   int             `json:"type"`   //类型 1转入 2转出

}
