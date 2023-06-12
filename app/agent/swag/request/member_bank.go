package request

type MemberBankList struct {
	UId int `json:"uid" form:"uid"`
}
type MemberBankCreate struct {
	BankId     int    `json:"bank_id"`     //银行id
	CardNumber string `json:"card_number"` //卡号
	Province   string `json:"province"`    //省份
	City       string `json:"city"`        //市
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
	IdNumber   string `json:"id_number"`   //身份证号码
	BankPhone  string `json:"bank_phone"`  //预留手机号码
	IsDefault  int    `json:"is_default"`  //默认银行卡
}
type MemberBankUpdate struct {
	Id         int    `json:"id"`
	BankName   string `json:"bank_name"`   //银行名称
	CardNumber string `json:"card_number"` //卡号
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
}
type MemberBankRemove struct {
	Id int `json:"id"`
}
