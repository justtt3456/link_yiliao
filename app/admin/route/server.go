package route

import (
	"china-russia/app/admin/controller"
	"china-russia/app/admin/controller/v1"
	_ "china-russia/app/admin/docs"
	"china-russia/app/admin/middleware"
	"china-russia/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Run() {
	r := gin.Default()
	r.Use(middleware.CorsDomain())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/upload", "./upload")
	g := r.Group("admin/api")
	/******************************************************************************************************************/
	//公共路由
	login := v1.LoginController{}
	g.POST("login", login.Login)
	//提示音
	sound := v1.SoundController{}
	g.GET("getSound", sound.Data)
	notify := controller.NotifyController{}
	g.POST("notify/:payment", notify.Notify)
	g.GET("notify/:payment", notify.Notify)
	/******************************************************************************************************************/
	//鉴权路由
	ga := g.Group("").Use(middleware.Auth())
	//首页
	index := v1.IndexController{}
	ga.GET("index", index.Report)
	//图片上传
	upload := v1.UploadController{}
	ga.POST("upload/image", upload.UploadImage)
	//管理员
	admin := v1.AdminController{}
	ga.GET("admin/list", admin.List)
	ga.POST("admin/create", admin.Create)
	ga.POST("admin/update", admin.Update)
	ga.POST("admin/remove", admin.Remove)
	ga.POST("logout", admin.Logout)
	ga.POST("admin/password", admin.Password)
	ga.POST("admin/google", admin.Google)
	//角色
	role := v1.RoleController{}
	ga.GET("role/list", role.List)
	ga.POST("role/create", role.Create)
	ga.POST("role/update", role.Update)
	ga.POST("role/remove", role.Remove)
	//权限
	permission := v1.PermissionController{}
	ga.GET("/permission/list", permission.List)
	ga.POST("permission/create", permission.Create)
	ga.POST("permission/update", permission.Update)
	ga.POST("permission/remove", permission.Remove)

	//用户
	member := v1.MemberController{}
	ga.GET("member/page_list", member.PageList)
	//ga.POST("member/create", member.Create)
	ga.POST("member/update", member.Update)
	ga.POST("member/update_password", member.UpdatePassword)
	ga.POST("member/update_status", member.UpdateStatus)
	ga.POST("member/remove", member.Remove)
	ga.POST("member/team", member.Team)
	ga.POST("member/sendCoupon", member.SendCoupon)
	ga.POST("member/getCode", member.GetCode)

	//用户银行卡
	ga.GET("member/bankcard/list", member.BankCardList)
	ga.POST("member/bankcard/update", member.UpdateBankCard)
	ga.POST("member/bankcard/remove", member.RemoveBankCard)
	//用户实名认证
	ga.GET("member/verified/page_list", member.VerifiedPageList)
	ga.POST("member/verified/update", member.UpdateVerified)
	ga.POST("member/verified/remove", member.RemoveVerified)

	//充值
	recharge := v1.RechargeController{}
	ga.GET("recharge/page_list", recharge.PageList)
	ga.POST("recharge/update", recharge.Update)
	//提现
	withdraw := v1.WithdrawController{}
	ga.GET("withdraw/page_list", withdraw.PageList)
	ga.POST("withdraw/update", withdraw.Update)
	ga.POST("withdraw/update_info", withdraw.UpdateInfo)
	//账单
	trade := v1.TradeController{}
	ga.GET("trade/page_list", trade.PageList)
	//配置
	config := v1.ConfigController{}
	ga.GET("config/base", config.Base)
	ga.POST("config/base/update", config.BaseUpdate)
	ga.GET("config/funds", config.Funds)
	ga.POST("config/funds/update", config.FundsUpdate)
	//收款银行卡
	ga.GET("config/bank/list", config.BankList)
	ga.POST("config/bank/create", config.BankCreate)
	ga.POST("config/bank/update", config.BankUpdate)
	ga.POST("config/bank/update_status", config.BankUpdateStatus)
	ga.POST("config/bank/remove", config.BankRemove)
	//客服
	ga.GET("config/kf/list", config.KfList)
	ga.POST("config/kf/update", config.KfUpdate)
	ga.POST("config/kf/update_status", config.KfUpdateStatus)
	//充值 提现方式
	ga.GET("config/recharge_method/list", config.RechargeMethodList)
	ga.POST("config/recharge_method/update", config.RechargeMethodUpdate)
	ga.POST("config/recharge_method/update_status", config.RechargeMethodUpdateStatus)
	ga.GET("config/withdraw_method/list", config.WithdrawMethodList)
	ga.POST("config/withdraw_method/update", config.WithdrawMethodUpdate)
	ga.POST("config/withdraw_method/update_status", config.WithdrawMethodUpdateStatus)
	//banner
	banner := v1.BannerController{}
	ga.GET("banner/page_list", banner.PageList)
	ga.POST("banner/create", banner.Create)
	ga.POST("banner/update", banner.Update)
	ga.POST("banner/update_status", banner.UpdateStatus)
	ga.POST("banner/remove", banner.Remove)
	//人工充提
	manual := v1.ManualController{}
	//ga.GET("manual/page_list", manual.PageList)
	ga.POST("manual/handle", manual.Handle)

	//公告notice
	notice := v1.NoticeController{}
	ga.GET("notice/page_list", notice.PageList)
	ga.POST("notice/create", notice.Create)
	ga.POST("notice/update", notice.Update)
	ga.POST("notice/update_status", notice.UpdateStatus)
	ga.POST("notice/remove", notice.Remove)
	//站内信message
	msg := v1.MessageController{}
	ga.GET("message/page_list", msg.PageList)
	ga.POST("message/create", msg.Create)
	ga.POST("message/update", msg.Update)
	ga.POST("message/update_status", msg.UpdateStatus)
	ga.POST("message/remove", msg.Remove)

	//支付
	payment := v1.PaymentController{}
	ga.GET("payment/page_list", payment.PageList)
	ga.GET("payment/list", payment.List)
	ga.POST("payment/create", payment.Create)
	ga.POST("payment/update", payment.Update)
	ga.POST("payment/remove", payment.Remove)
	//支付通道
	channel := v1.PayChannelController{}
	ga.GET("pay_channel/page_list", channel.PageList)
	ga.POST("pay_channel/create", channel.Create)
	ga.POST("pay_channel/update", channel.Update)
	ga.POST("pay_channel/update_status", channel.UpdateStatus)
	ga.POST("pay_channel/remove", channel.Remove)
	//投资理财
	invest := v1.InvestController{}
	ga.GET("invest", invest.Index)
	ga.POST("invest/update", invest.Update)
	ga.GET("invest/order/page_list", invest.OrderPageList)
	ga.GET("invest/income/page_list", invest.IncomePageList)
	//产品分类
	pc := v1.ProductCategoryController{}
	ga.GET("product_category/list", pc.List)
	ga.POST("product_category/create", pc.Create)
	ga.POST("product_category/update", pc.Update)
	ga.POST("product_category/update_status", pc.UpdateStatus)
	ga.POST("product_category/remove", pc.Remove)
	//产品
	product := v1.ProductController{}
	ga.GET("product/page_list", product.PageList)
	ga.POST("product/create", product.Create)
	ga.POST("product/update", product.Update)
	ga.POST("product/update_status", product.UpdateStatus)
	ga.POST("product/remove", product.Remove)
	//获取赠品产品列表
	ga.GET("product/gift_options", product.GiftOptions)

	//股权
	guquan := v1.GuquanController{}
	ga.GET("guquan/list", guquan.List)
	ga.POST("guquan/update", guquan.Update)

	//订单管理
	order := v1.OrderController{}
	ga.GET("order/product_list", order.PageList)
	ga.GET("order/guquan_list", order.GuQuanPageList)
	ga.POST("order/guquan_update", order.Update)

	//公司简介和用户须知
	help := v1.HelpController{}
	ga.GET("help/page_list", help.PageList)
	ga.POST("help/create", help.Create)
	ga.POST("help/update", help.Update)
	ga.POST("help/update_status", help.UpdateStatus)
	ga.POST("help/remove", help.Remove)

	//活动
	active := v1.ActiveController{}
	ga.GET("active/couponList", active.CouponList)
	ga.POST("active/addCoupon", active.AddCoupon)
	ga.GET("active/page_list", active.PageList)
	ga.POST("active/addActive", active.AddActive)
	ga.POST("active/delActive", active.DelActive)
	address := fmt.Sprintf(":%d", global.CONFIG.System.AdminAddr)

	//新闻
	news := v1.NewsController{}
	ga.GET("news/page_list", news.PageList)
	ga.POST("news/create", news.Create)
	ga.POST("news/update", news.Update)
	ga.POST("news/update_status", news.UpdateStatus)
	ga.POST("news/remove", news.Remove)

	agent := v1.AgentController{}
	ga.GET("agent/list", agent.List)
	ga.GET("agent/page_list", agent.PageList)
	ga.POST("agent/create", agent.Create)
	ga.POST("agent/update", agent.Update)
	ga.POST("agent/update_status", agent.UpdateStatus)

	ga.GET("config/lang/list", config.LangList)
	r.Run(address)
}
