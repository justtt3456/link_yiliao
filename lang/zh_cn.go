package lang

var zh_cn map[string]string = map[string]string{
	//账单类型
	"Deposit":                         "入金",
	"Withdraw":                        "出金",
	"Buy product":                     "购买产品",
	"Product settlement":              "产品结算",
	"Deposit gift":                    "入金赠送",
	"System deposit":                  "系统存款",
	"Huijin gift":                     "汇金赠送",
	"Member benefits":                 "会员福利",
	"Protocol number cannot be empty": "协议号不能为空",
	"Address cannot be empty":         "地址不能为空",
	"System deduction":                "系统扣款",
	"System freezes":                  "系统冻结",
	"System unfreeze":                 "系统解冻",
	"Invest transfer in":              "投资转入",
	"Invest transfer out":             "投资转出",
	"Investment income":               "投资收益",

	//提示信息
	"Parameter error":                                              "参数错误",
	"Username cannot be empty":                                     "用户名不能为空",
	"Password cannot be empty":                                     "密码不能为空",
	"The two passwords are inconsistent":                           "两次密码不一致",
	"Please enter the correct phone number":                        "请输入正确的手机号码",
	"Invitation code cannot be empty":                              "邀请码不能为空",
	"Username already exists":                                      "用户名已存在",
	"Username does not exist":                                      "用户名不存在",
	"Registration function is closed":                              "注册功能关闭中",
	"Wrong invitation code":                                        "邀请码错误",
	"Incorrect username and password":                              "用户名密码错误",
	"User is forbidden to log in":                                  "用户禁止登录",
	"User is not logged in":                                        "用户未登录",
	"Mobile phone number cannot be empty":                          "手机号码不能为空",
	"The delivery address cannot be empty":                         "收货地址不能为空",
	"The detailed address cannot be empty":                         "详细地址不能为空",
	"Cardholder cannot be empty":                                   "持卡人不能为空",
	"Bank card number cannot be empty":                             "银行卡号不能为空",
	"Account bank cannot be empty":                                 "开户行不能为空",
	"Withdraw password cannot be empty":                            "出金密码不能为空",
	"Incorrect withdraw password":                                  "出金密码错误",
	"The original password cannot be empty":                        "原密码不能为空",
	"The original password is wrong":                               "原密码错误",
	"The original withdraw password cannot be empty":               "原出金密码不能为空",
	"The original withdraw password is wrong":                      "原出金密码错误",
	"Wrong deposit amount":                                         "入金金额错误",
	"Wrong amount":                                                 "金额错误",
	"The deposit voucher cannot be empty":                          "入金凭证不能为空",
	"Record does not exist":                                        "记录不存在",
	"Please deposit after %s":                                      "请在%s之后入金",
	"Please deposit before %s":                                     "请在%s之前入金",
	"Minimum deposit %.2f":                                         "最低入金%.2f",
	"Maximum deposit %.2f":                                         "最高入金%.2f",
	"The start time cannot be greater than the current time":       "开始时间不能大于当前时间",
	"The start time cannot be greater than the end time":           "开始时间不能大于结束时间",
	"Platform cannot be empty":                                     "平台不能为空",
	"No app version":                                               "无app版本",
	"Receiving address cannot be empty":                            "收款地址不能为空",
	"Insufficient account balance":                                 "账户余额不足",
	"Bank card error":                                              "银行卡错误",
	"Please withdraw after %s":                                     "请在%s之后出金",
	"Please withdraw before %s":                                    "请在%s之前出金",
	"Minimum withdraw %.2f":                                        "最低出金%.2f",
	"Maximum withdraw %.2f":                                        "最高出金%.2f",
	"System error, please contact the administrator":               "系统错误,请联系管理员",
	"The current user is forbidden to withdraw":                    "当前用户禁止出金",
	"Only one shipping address can be added":                       "只能添加一个收货地址",
	"Wrong deposit type":                                           "入金类型错误",
	"Wrong purchase type":                                          "购买类型错误",
	"Order number cannot be empty":                                 "订单号不能为空",
	"Order does not exist":                                         "订单不存在",
	"Platform name cannot be empty":                                "平台名称不能为空",
	"Platform name error":                                          "平台名称错误",
	"System configuration error, please contact the administrator": "系统配置错误,请联系管理员",
	"Payment account cannot be empty":                              "付款账号不能为空",
	"Receiving account cannot be empty":                            "收款账号不能为空",
	"Receiving account does not exist":                             "收款账号不存在",
	"Payment channel cannot be empty":                              "支付通道不能为空",
	"The payment channel does not exist":                           "支付通道不存在",
	"The payment does not exist":                                   "支付不存在",
	"Bank cannot be empty":                                         "银行不能为空",
	"Unable to bind bank card, please contact customer service":    "无法绑定银行卡,请联系客服",
	"Only %d bank cards can be bound":                              "只能绑定%d张银行卡",
	"Prohibit to modify bank card":                                 "禁止修改银行卡",
	"Bank card does not exist":                                     "银行卡不存在",
	"Please upload a picture within %dM":                           "请上传%dM内的图片",
	"Picture format error":                                         "图片格式错误",
	"Wrong withdrawal method":                                      "出金方式错误",
	"Configuration error":                                          "配置错误",
	"Prohibited to buy":                                            "禁止购买",
	"Insufficient transferable balance":                            "可转出余额不足",
	"Country code cannot be empty":                                 "国家编码不能为空",
	"Product does not exist":                                       "产品不存在",
	"Optional added":                                               "已添加自选",
	"You can only withdraw %d times per day":                       "每日只能出金%d次",
	"The current account has been frozen!":                         "当前账户已冻结!",
	"Real name cannot be empty":                                    "姓名不能为空",
	"ID number cannot be empty":                                    "身份证号码不能为空",
	"Phone number can not be blank":                                "手机号码不能为空",
	"The front of the ID card cannot be blank":                     "身份证正面不能为空",
	"The back of the ID card cannot be blank":                      "身份证背面不能为空",
	"Real name authentication already exists":                      "实名认证已存在",
	"Minimum purchase %d":                                          "最低购买%d",
	"Closed on weekends":                                           "周末休市中",
}
