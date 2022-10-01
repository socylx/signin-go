package global

import (
	"log"
	"signin-go/global/config"

	"gopkg.in/ini.v1"
)

func PreInit() {
	log.Println("global.PreInit Start...")
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("global.PreInit Load config.ini Error: %v", err)
	}
	err = file.Section("Server").MapTo(config.Server)
	if err != nil {
		log.Fatalf("global.PreInit file.Section(%s).MapTo(object).Error: %v", "Server", err)
	}
}

func initConfig() {
	log.Println("global.initConfig Start...")

	objectMap := map[string]interface{}{
		"Server": config.Server,
	}

	file, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("global.initConfig Load config.ini Error: %v", err)
	}

	for section, object := range objectMap {
		err = file.Section(section).MapTo(object)
		if err != nil {
			log.Fatalf("global.initConfig Section(%s).MapTo(object).Error: %v", "Server", err)
		}
	}
}
func Init() {
	initConfig()
}
