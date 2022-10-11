package response

type MemberBankListResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data MemberBankList `json:"data"`
}
type MemberBankList struct {
	List []MemberBank `json:"list"`
}

type MemberBank struct {
	ID         int    `json:"id"` //
	BankName   string `json:"bank_name"`
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
}
