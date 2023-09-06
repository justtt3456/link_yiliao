package route

import (
	"china-russia/app/api/controller"
	"china-russia/app/api/controller/v1"
	"china-russia/common"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

var PublicRouterApp = new(PublicRouter)

func (s PublicRouter) InitRouter(r *gin.RouterGroup) {
	//注册登录
	c := v1.LoginController{}
	r.Use(common.Session("topgoer"))
	r.POST("login", c.Login)
	r.POST("register", c.Register)
	r.POST("sendCode", c.SendSms)
	r.GET("captcha", func(c *gin.Context) {
		common.Captcha(c, 4)
	})
	//支付回调地址
	n := v1.NotifyController{}
	r.POST("notify_xinmeng", n.NotifyXinMeng)
	//配置
	config := v1.ConfigController{}
	r.GET("config", config.List)
	//获取版本升级
	upgrade := v1.UpgradeController{}
	r.GET("upgrade", upgrade.Version)
	//客服
	kf := v1.KfController{}
	r.GET("kf", kf.Redirect)
	//回调
	notify := controller.NotifyController{}
	r.GET("notify/:payment", notify.Notify)
	r.POST("notify/:payment", notify.Notify)
	news := v1.NewsController{}
	r.GET("news/page_list", news.PageList)
}
