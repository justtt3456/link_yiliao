package global

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func Viper() {
	var config string

	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		config = "config.yaml"
		fmt.Printf("您正在使用默认值,config的路径为%v\n", config)
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}
}
