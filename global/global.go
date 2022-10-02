package global

import (
	"log"
	"signin-go/global/config"
	"signin-go/global/logger"
	"signin-go/global/mysql"
	"signin-go/global/time"
	"sync"
)

var Once sync.Once

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
}

func Close() {
	mysql.Close()
	logger.Close()
}
