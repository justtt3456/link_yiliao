package request

type WithdrawList struct {
	Status   int `form:"status"` //状态 1审核中 2通过 3驳回
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type WithdrawCreate struct {
	Method           int     `json:"method"`            //提现类型
	WithdrawPassword string  `json:"withdraw_password"` //交易密码
	TotalAmount      float64 `json:"total_amount"`      //提现总额
	ID               int     `json:"id"`                //银行卡ID
	//BankName   string `json:"bank_name"`   //关联银行名称
	//BranchBank string `json:"branch_bank"` //开户行
	//RealName   string `json:"real_name"`   //开户人
	//CardNumber string `json:"card_number"` //卡号
	//BankPhone  string `json:"bank_phone"`  //预留手机号码
}
