package ticker

import (
	"china-russia/app/task/repository"
	"time"
)

func Run() {
	//余额宝
	go invest()
	//余额宝接触冻结
	go unfreeze()
	//股权
	go equity()
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
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		i.Do()
	}
}
