package route

import (
	"china-russia/app/agent/controller/v1"
	//_ "china-russia/app/agent/docs"
	"china-russia/app/agent/middleware"
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
	g := r.Group("agent/api")
	/******************************************************************************************************************/
	//公共路由
	login := v1.LoginController{}
	g.POST("login", login.Login)
	/******************************************************************************************************************/
	//鉴权路由
	ga := g.Group("").Use(middleware.Auth())
	//首页
	index := v1.IndexController{}
	ga.GET("index", index.Report)
	//图片上传
	upload := v1.UploadController{}
	ga.POST("upload/image", upload.UploadImage)

	//用户
	member := v1.MemberController{}
	ga.GET("member/page_list", member.PageList)
	ga.POST("member/team", member.Team)
	ga.POST("logout", member.Logout)
	//用户银行卡
	ga.GET("member/bankcard/list", member.BankCardList)

	//用户实名认证
	ga.GET("member/verified/page_list", member.VerifiedPageList)

	//充值
	recharge := v1.RechargeController{}
	ga.GET("recharge/page_list", recharge.PageList)

	//提现
	withdraw := v1.WithdrawController{}
	ga.GET("withdraw/page_list", withdraw.PageList)

	//账单
	trade := v1.TradeController{}
	ga.GET("trade/page_list", trade.PageList)

	//投资理财
	invest := v1.InvestController{}
	ga.GET("invest/order/page_list", invest.OrderPageList)
	ga.GET("invest/income/page_list", invest.IncomePageList)

	//订单管理
	order := v1.OrderController{}
	ga.GET("order/product_list", order.PageList)
	ga.GET("order/guquan_list", order.GuQuanPageList)

	//agent := v1.AgentController{}
	//ga.GET("agent/list", agent.List)
	//ga.GET("agent/page_list", agent.PageList)
	//ga.POST("agent/create", agent.Create)
	//ga.POST("agent/update", agent.Update)
	//ga.POST("agent/update_status", agent.UpdateStatus)

	address := fmt.Sprintf(":%d", global.CONFIG.System.AgentAddr)
	r.Run(address)
}
