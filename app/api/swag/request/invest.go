package request

type InvestOrder struct {
	Amount float64 `json:"amount"` //转入转出金额
	Type   int     `json:"type"`   //类型 1转入 2转出

}
