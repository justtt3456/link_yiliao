package dao

import (
	"finance/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func Redis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		logrus.Info("redis connect ping response:", pong)
		return client
	}
}
