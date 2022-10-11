package global

import (
	"finance/config"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	REDIS    *redis.Client
	CONFIG   config.Server
	VP       *viper.Viper
	Language string
	Ws       websocket.Upgrader
)
