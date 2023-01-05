package main

import (
	"finance/app/api/route"
	"finance/dao"
	"finance/global"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//	@title		api
//	@version	2.0
//	@host		52.76.81.121:8000
//	@BasePath	/api/v1
func main() {
	//初始化viper
	global.VP = global.Viper()
	//初始化log
	global.Log()
	//dao连接
	global.DB = dao.Gorm()
	global.REDIS = dao.Redis()
	//设置默认中文
	global.Language = "zh_cn"

	//路由初始化
	route.Run()
}
