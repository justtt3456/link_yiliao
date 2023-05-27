package route

import (
	_ "china-russia/app/api/docs"
	"china-russia/app/api/middleware"
	"china-russia/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Run() {
	r := gin.Default()
	r.Use(middleware.Filter())
	r.Use(middleware.CorsDomain(), middleware.Lang())
	//模板渲染
	r.LoadHTMLGlob("template/*")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/upload", "./upload")

	g := r.Group("api/v1")
	//公共路由
	PublicRouterApp.InitRouter(g)
	//鉴权路由
	ga := g.Group("").Use(middleware.Auth())
	AuthRouterApp.InitRouter(ga)

	address := fmt.Sprintf(":%d", global.CONFIG.System.ApiAddr)

	r.Run(address)
}
