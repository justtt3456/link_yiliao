package model

const (
	//由于gorm或probuf等原因,过滤非零值,所以状态不使用0
	//状态 1开启true 2关闭false
	StatusOk    = 1
	StatusClose = 2
	//订单状态 1审核中null 2通过true 3驳回false
	StatusReview   = 1
	StatusAccept   = 2
	StatusRollback = 3
	//账变类型
	TradeTypeRecharge     = 1  //充值
	TradeTypeWithdraw     = 2  //提现
	TradeTypeBuyProduct   = 3  //购买
	TradeTypeDrawProduct  = 4  //结算
	TradeTypeRechargeGift = 5  //充值赠送
	TradeTypeSystemInc    = 6  //入金
	TradeTypeSystemInc1   = 13 //汇金赠送
	TradeTypeSystemInc2   = 14 //会员福利
	TradeTypeSystemDec    = 7  //系统扣款
	TradeTypeFreeze       = 8  //冻结
	TradeTypeUnfreeze     = 9  //解冻
	TradeTypeFundsIn      = 10 //余额宝转入
	TradeTypeFundsOut     = 11 //余额宝转出
	TradeTypeFundsIncome  = 12 //余额宝收益
)
