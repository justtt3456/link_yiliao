package ticker

import (
	"china-russia/app/task/repository"
	"github.com/robfig/cron/v3"
	"time"
)

func Run() {
	//余额宝
	go invest()
	//余额宝接触冻结
	go unfreeze()
	//股权
	go equity()
	//股权分
	go equityScore()
	//定时任务:收益结算
	go repository.InitCrontab()
	select {}
}

func invest() {
	i := repository.Invest{}
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		i.Do()
	}
}

func unfreeze() {
	i := repository.Unfreeze{}
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		i.Do()
	}
}

func equity() {
	i := repository.Equity{}
	ticker := time.NewTicker(time.Second * 64) //随机时间
	for {
		<-ticker.C
		i.Do()
	}
}
func equityScore() {
	i := repository.EquityScore{}
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 0 * * ?", func() {
		//id, err := c.AddFunc("0 * * * * *", func() {
		i.Do()
	})
	c.Start()
	select {}
}
