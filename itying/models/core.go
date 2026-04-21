package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func init() {
	// 配置日志记录器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢查询阈值，超过这个时间会标记为 SLOW SQL
			LogLevel:                  logger.Info,            // 设置为 Info 级别，打印所有 SQL
			IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,                   // 彩色输出
		},
	)

	dsn := "root:040412@tcp(47.109.80.234:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 应用日志配置
	})
	if err != nil {
		fmt.Println(err)
	}
}
