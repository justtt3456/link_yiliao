package route

import (
	"china-russia/app/api/controller/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

var AuthRouterApp = new(AuthRouter)

func (s AuthRouter) InitRouter(r gin.IRoutes) {
	//用户
	m := v1.MemberController{}
	r.GET("member/info", m.Info)
	r.POST("member/password", m.UpdatePassword)
	r.POST("member/pay_password", m.UpdatePayPassword)
	//r.POST("member/update", m.UpdateInfo)
	r.POST("member/verified", m.Verified)
	r.POST("logout", m.Logout)
	r.GET("member/team", m.Team)
	r.POST("member/transfer", m.Transfer)
	r.POST("member/member_coupon", m.MemberCoupon)

	//银行卡
	memberBank := v1.MemberBankController{}
	r.POST("member_bank/create", memberBank.Create)
	r.POST("member_bank/update", memberBank.Update)
	r.POST("member_bank/remove", memberBank.Remove)
	r.GET("member_bank/list", memberBank.List)

	//产品列表
	product := v1.ProductController{}
	r.GET("product/category", product.Category)
	r.GET("product/page_list", product.PageList)
	//r.GET("product/recommend", product.Recommend)
	//r.GET("product/getproduct", product.Getproduct)
	r.GET("product/guquan", product.Guquan)
	r.POST("product/buy", product.Buy)
	r.POST("product/buy_equity", product.BuyEquity)
	r.GET("product/buy_list", product.BuyList)
	r.GET("product/buy_guquan_list", product.BuyGuquanList)
	//获取股权证书内容
	r.GET("product/stock_certificate", product.StockCertificate)
	//订单上级返佣数据修复
	r.GET("product/commission_repair", product.CommissionRepair)

	//充值
	recharge := v1.RechargeController{}
	r.GET("recharge/method", recharge.Method)
	r.GET("recharge/method_info", recharge.MethodInfo)
	r.POST("recharge/create", recharge.Create)
	r.GET("recharge/page_list", recharge.PageList)
	//提现
	withdraw := v1.WithdrawController{}
	r.POST("withdraw/create", withdraw.Create)
	r.GET("withdraw/page_list", withdraw.PageList)
	r.GET("withdraw/method", withdraw.Method)
	//图片上传
	upload := v1.UploadController{}
	r.POST("upload/image", upload.UploadImage)
	//余额宝
	invest := v1.InvestController{}
	r.GET("invest", invest.Index)
	r.POST("invest/transfer", invest.Transfer)
	r.GET("invest/order", invest.Order)
	//站内信
	msg := v1.MessageController{}
	r.GET("message/page_list", msg.PageList)
	r.POST("message/read", msg.Read)
	//首页
	index := v1.IndexController{}
	r.GET("index", index.Index)
	//滚动公告列表
	notice := v1.NoticeController{}
	r.GET("notice/notice_list", notice.NoticeList)
	//公司简介和推荐奖励
	help := v1.HelpController{}
	r.GET("help/list", help.List)
	//账单
	Trade := v1.TradeController{}
	r.GET("trade/page_list", Trade.PageList)
	r.GET("trade/income_list", Trade.IncomeList)

	//签到
	sign := v1.SignController{}
	r.GET("sign/sign", sign.Sign)
}
