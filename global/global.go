package global

import (
	"gsteps-go/global/config"
	"gsteps-go/global/logger"
	"gsteps-go/global/mongo"
	"gsteps-go/global/mysql"
	"gsteps-go/global/redis"
	"gsteps-go/global/time"
	"log"
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

	mongo.Init()
	log.Println("-----------------------------")
}

func Close() {
	mongo.Close()
	redis.Close()
	mysql.Close()
	logger.Close()
}
