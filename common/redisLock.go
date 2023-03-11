package common

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

type RedisLock struct {
	RedisClient *redis.Client
}

//加锁
func (this *RedisLock) Lock(lockKey string) bool {
	//Redis乐观锁缓存时间默认为3秒
	success, err := this.RedisClient.SetNX(lockKey, "locking", 3*time.Second).Result()
	if err != nil {
		logrus.Error(err)
		return false
	}
	if !success {
		return false
	}
	return true
}

//解锁
func (this *RedisLock) Unlock(lockKey string) bool {
	ok, err := this.RedisClient.Del(lockKey).Result()
	if err != nil {
		logrus.Error(err)
		return false
	}
	if ok != 1 {
		return false
	}
	return true
}
