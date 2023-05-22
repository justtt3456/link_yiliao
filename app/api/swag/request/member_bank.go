package request

type MemberBankCreate struct {
	BankName   string `json:"bank_name"`   //银行
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
}
type MemberBankUpdate struct {
	Id         int    `json:"id"`
	BankName   string `json:"bank_name"`   //银行
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
}
type MemberBankRemove struct {
	Id int `json:"id"`
}

type MemberUsdtCreate struct {
	Protocol string `json:"protocol"` //协议号
	Address  string `json:"address"`  //地址
}
