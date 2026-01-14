package mysqlConfig

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, // 打印所有 SQL
			Colorful:      true,
		},
	)

	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 打印 SQL
	})

	if err != nil {
		panic("❌ 数据库连接失败: " + err.Error()) // 直接 crash，避免后续 nil 使用
	}
	fmt.Println("✅ 数据库连接成功!")
}
