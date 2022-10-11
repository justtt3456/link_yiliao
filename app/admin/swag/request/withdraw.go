package request

type WithdrawListRequest struct {
	UID       int    `json:"uid" form:"uid"`               //用户ID
	Username  string `json:"username" form:"username"`     //用户
	StartTime string `json:"start_time" form:"start_time"` //开始时间
	EndTime   string `json:"end_time" form:"end_time"`     //结束时间
	Status    int    `json:"status" form:"status"`         //1为未审核，2为已审核，3为已拒绝
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
}
type WithdrawUpdateRequest struct {
	ID          int    `json:"id"`
	Status      int    `json:"status"`      //2为已审核，3为已拒绝
	Description string `json:"description"` //备注
	Operator    int    `json:"operator"`    //操作人ID
}
type WithdrawUpdateInfoRequest struct {
	ID         int    `json:"id"`
	BankName   string `json:"bank_name"`
	BranchBank string `json:"branch_bank"`
	RealName   string `json:"real_name"`
	CardNumber string `json:"card_number"`
}
