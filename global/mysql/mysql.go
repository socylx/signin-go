package mysql

import (
	"fmt"
	"log"
	"signin-go/global/config"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	log.Println("global.mysql.Init Start...")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
	)

	var err error
	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: newLogger(
				loggerConfig{
					SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
					LogLevel:                  logger.Info,            // 日志级别
					IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,                  // 禁用彩色打印
				},
			),
			QueryFields: true,
		})
	if err != nil {
		log.Fatalf("global.mysql.Init.gorm.Open() Error: %v", errors.Wrap(err, fmt.Sprintf("[DB connection failed] Database name: %s", config.Mysql.Database)))
	}
	DB.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("global.mysql.Init.DB.DB() Error: %v", err)
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(0)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(80)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * 60)

	// 使用插件
	DB.Use(&TracePlugin{})
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("global.mysql.Close.DB.DB() Error: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("global.mysql.Close.sqlDB.Close() Error: %v", err)
	}
}
