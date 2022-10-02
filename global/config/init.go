package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func PreInit() {
	log.Println("global.config.PreInit Start...")
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("global.config.PreInit Load config.ini Error: %v", err)
	}
	err = file.Section("Server").MapTo(Server)
	if err != nil {
		log.Fatalf("global.config.PreInit file.Section(%s).MapTo(object).Error: %v", "Server", err)
	}
}

func Init() {
	log.Println("global.config.Init Start...")

	objectMap := map[string]interface{}{
		"Server": Server,
		"Mysql":  Mysql,
		"Notify": Notify,
	}

	file, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("global.config.Init Load config.ini Error: %v", err)
	}

	for section, object := range objectMap {
		err = file.Section(section).MapTo(object)
		if err != nil {
			log.Fatalf("global.config.Init file.Section(%s).MapTo(object).Error: %v", "Server", err)
		}
	}
	Server.IsRelease = Server.Mode == gin.ReleaseMode
}
