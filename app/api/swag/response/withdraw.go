package response

import "github.com/shopspring/decimal"

type WithdrawListResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data WithdrawData `json:"data"`
}
type WithdrawData struct {
	List []Withdraw `json:"list"`
	Page Page       `json:"page"`
}
type Withdraw struct {
	Id          int             `json:"id"` //
	OrderSn     string          `json:"order_sn"`
	Type        int             `json:"type"` //提现类型1=银行卡
	TypeName    string          `json:"type_name"`
	BankName    string          `json:"bank_name"`    //关联银行名称
	BranchBank  string          `json:"branch_bank"`  //开户行
	RealName    string          `json:"real_name"`    //开户人
	CardNumber  string          `json:"card_number"`  //卡号
	BankPhone   string          `json:"bank_phone"`   //预留手机号码
	Amount      decimal.Decimal `json:"amount"`       //实际到账金额
	Fee         decimal.Decimal `json:"fee"`          //手续费
	TotalAmount decimal.Decimal `json:"total_amount"` //提现总额
	Description string          `json:"description"`  //审核备注
	Status      int             `json:"status"`       //提现状态，1为未审核，2为已审核，3为已拒绝
	CreateTime  int64           `json:"create_time"`  //创建时间
	UpdateTime  int64           `json:"update_time"`  //审核时间
}
type WithdrawMethodResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data WithdrawMethodData `json:"data"`
}
type WithdrawMethodData struct {
	List []WithdrawMethod `json:"list"`
}
type WithdrawMethod struct {
	Id   int             `json:"id"`   //
	Name string          `json:"name"` //
	Code string          `json:"code"` //
	Icon string          `json:"icon"` //
	Fee  decimal.Decimal `json:"fee"`
}
