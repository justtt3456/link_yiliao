package response

type MemberBankResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data MemberBankList `json:"data"`
}
type MemberBankList struct {
	List []MemberBank `json:"list"`
}

type MemberBank struct {
	Id         int    `json:"id"`      //
	BankId     int    `json:"bank_id"` //银行id
	BankName   string `json:"bank_name"`
	CardNumber string `json:"card_number"` //卡号
	Province   string `json:"province"`    //省份
	City       string `json:"city"`        //市
	BranchBank string `json:"branch_bank"` //开户行（开户所在地）
	RealName   string `json:"real_name"`   //开户人
	IdNumber   string `json:"id_number"`   //身份证号码
	BankPhone  string `json:"bank_phone"`  //预留手机号码
	IsDefault  int    `json:"is_default"`  //默认银行卡
}

type MemberUsdt struct {
	Id       int    `json:"id"`       //
	Address  string `json:"address"`  //地址
	Protocol string `json:"protocol"` //协议
}

type MemberUsdtList struct {
	List []MemberUsdt `json:"list"`
}
