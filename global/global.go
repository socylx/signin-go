package global

import (
	"log"
	"signin-go/global/config"
	"signin-go/global/logger"
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
}
