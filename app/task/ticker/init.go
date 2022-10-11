package ticker

import (
	"finance/app/task/repository"
	"time"
)

func Run() {
	//余额宝
	go invest()

	//余额宝接触冻结
	go unfreeze()
	//股权
	go guquan()

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

func guquan() {
	i := repository.Guquan{}
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		i.Do()
	}
}
