package global

import (
	"china-russia/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	REDIS    *redis.Client
	CONFIG   config.Server
	Language string
)
