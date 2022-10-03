package global

import (
	"log"
	"signin-go/global/config"
	"signin-go/global/logger"
	"signin-go/global/mysql"
	"signin-go/global/redis"
	"signin-go/global/time"
)

func Init() {
	// todo
	// log.SetOutput()
	log.Println("-----------------------------")

	// 严格执行顺序
	config.Init()
	log.Println("-----------------------------")

	time.Init()
	log.Println("-----------------------------")

	logger.Init()
	log.Println("-----------------------------")

	mysql.Init()
	log.Println("-----------------------------")

	redis.Init()
	log.Println("-----------------------------")
}

func Close() {
	redis.Close()
	mysql.Close()
	logger.Close()
}
