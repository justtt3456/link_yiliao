package dao

import (
	"china-russia/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
