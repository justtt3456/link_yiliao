package main

import (
	"china-russia/dao"
	"china-russia/global"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	global.Viper()
	//初始化log
	global.Log()
	//dao连接
	global.DB = dao.Gorm()
	global.REDIS = dao.Redis()

	recharge := make([]map[string]interface{}, 0)
	global.DB.Raw("SELECT m.username,sum(usdt_amount) usdt_amount,sum(amount) amount FROM `c_recharge` r left join c_member m on r.uid = m.id where  r.status = 2 group by uid  ORDER BY `amount` DESC").Scan(&recharge)
	export(recharge, "recharge", "会员充值")
	withdraw := make([]map[string]interface{}, 0)
	global.DB.Raw("SELECT m.username,sum(usdt_amount) usdt_amount,sum(total_amount) amount FROM `c_withdraw` w left join c_member m on w.uid = m.id where w.status = 2 group by uid  ORDER BY `amount` DESC").Scan(&withdraw)
	export(withdraw, "withdraw", "会员提现")
	//获取所有代理
	//agentModel := model.Agent{}
	//agents := agentModel.List("", nil)
	//for _, v := range agents {
	////代理充值
	//recharge := make([]map[string]interface{}, 0)
	//global.DB.Raw("SELECT m.username,r.amount,r.usdt_amount,r.create_time FROM `c_recharge` r left join c_member m on r.uid = m.id where r.status = 2 and uid in (select id from c_member where agent_id = ?) order by r.id desc", v.Id).Scan(&recharge)
	//export(recharge, "recharge", v.Account)
	////代理提现
	//withdraw := make([]map[string]interface{}, 0)
	//global.DB.Raw("SELECT m.username,w.total_amount amount,w.usdt_amount,w.create_time FROM `c_withdraw` w left join c_member m on w.uid = m.id where w.status = 2 and uid in (select id from c_member where agent_id = ?) order by w.id desc", v.Id).Scan(&withdraw)
	//export(withdraw, "withdraw", v.Account)
	//}

}
func export(data []map[string]interface{}, directory string, name string) {
	f := excelize.NewFile() // 设置单元格的值
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "用户名")
	f.SetCellValue("Sheet1", "B1", "金额")
	f.SetCellValue("Sheet1", "C1", "U金额")

	line := 1
	// 循环写入数据
	for _, v := range data {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v["username"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v["amount"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v["usdt_amount"])
		//var createTime string
		//switch v["create_time"].(type) {
		//case int:
		//	createTime = time.Unix(int64(v["create_time"].(int)), 0).Format("2006-01-02 15:04:05")
		//case int64:
		//	createTime = time.Unix(v["create_time"].(int64), 0).Format("2006-01-02 15:04:05")
		//case uint64:
		//	i := v["create_time"].(uint64)
		//	createTime = time.Unix(int64(i), 0).Format("2006-01-02 15:04:05")
		//}
		//f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), createTime)
	}

	// 保存文件
	if err := f.SaveAs("export/" + directory + "/" + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
