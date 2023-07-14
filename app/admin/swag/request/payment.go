package request

type PaymentListRequest struct {
	PayName  string `form:"pay_name" json:"pay_name"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
}
type PaymentAddRequest struct {
	PayName        string `json:"pay_name"`        //支付方式名称
	RechargeURL    string `json:"submit_url"`      //充值提交地址
	WithdrawURL    string `json:"query_url"`       //提现提交地址
	NotifyURL      string `json:"notify_url"`      //回调地址
	MerchantNo     string `json:"merchant_no"`     //商户号
	Secret         string `json:"secret"`          //密钥
	PriKey         string `json:"pri_key"`         //私钥
	PubKey         string `json:"pub_key"`         //公钥
	ClassName      string `json:"class_name"`      //类名
	WithdrawStatus int    `json:"withdraw_status"` //是否启用代付 1是2否
}
type PaymentUpdateRequest struct {
	Id             int    `json:"id"`
	PayName        string `json:"pay_name"`        //支付方式名称
	RechargeURL    string `json:"submit_url"`      //充值提交地址
	WithdrawURL    string `json:"query_url"`       //提现提交地址
	NotifyURL      string `json:"notify_url"`      //回调地址
	MerchantNo     string `json:"merchant_no"`     //商户号
	Secret         string `json:"secret"`          //密钥
	PriKey         string `json:"pri_key"`         //私钥
	PubKey         string `json:"pub_key"`         //公钥
	ClassName      string `json:"class_name"`      //类名
	WithdrawStatus int    `json:"withdraw_status"` //是否启用代付 1是2否
}
type PaymentRemoveRequest struct {
	Id int `json:"id"`
}
