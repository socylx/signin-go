package global

import (
	"log"
	"signin-go/global/config"

	"gopkg.in/ini.v1"
)

func initConfig() {
	log.Println("initConfig start...")

	c, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("initConfig Error: %v", err)
	}

	err = c.Section("Server").MapTo(config.Server)
	if err != nil {
		log.Fatalf("MapTo(config.MapTo): [%s]: %v", "Server", err)
	}
}

func Init() {
	initConfig()
}
