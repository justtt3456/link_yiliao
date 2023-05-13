package main

import (
	"finance/dao"
	"finance/global"
	"finance/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	//初始化viper
	global.VP = global.Viper()
	//初始化log
	global.Log()
	//dao连接
	global.DB = dao.Gorm()
	global.REDIS = dao.Redis()
	updateUserBalance()
	updateUserOrder()
}

func updateUserBalance() {
	list := make([]model.Member, 0)
	tx := global.DB.Model(model.Member{}).Where("use_balance > 0").Find(&list)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return
	}
	for _, v := range list {
		global.DB.Model(v).Updates(map[string]interface{}{"balance": gorm.Expr("balance  + use_balance"), "use_balance": 0})
	}
}
func updateUserOrder() {
	list := make([]model.OrderProduct, 0)
	tx := global.DB.Model(model.OrderProduct{}).Where("is_return_capital = 0").Find(&list)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return
	}
	for _, v := range list {
		product := model.Product{ID: v.Pid}
		if !product.Get() {
			continue
		}
		//更新产品天数收益率
		v.IncomeRate = product.Dayincome / 5
		v.EndTime = v.CreateTime + int64(product.TimeLimit*5*86400)
		v.Update("income_rate", "end_time")
	}
}
