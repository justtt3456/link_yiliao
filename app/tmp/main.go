package main

import (
	"china-russia/app/task/repository"
	"china-russia/dao"
	"china-russia/global"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	//初始化viper
	global.Viper()
	//初始化log
	global.Log()
	//dao连接
	global.DB = dao.Gorm()
	global.REDIS = dao.Redis()
	service := repository.Award{}
	service.Run()
}
