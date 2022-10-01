package logger

import (
	"log"
	"signin-go/global/config"
	"signin-go/global/time"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	log.Println("global.logger.Init Start...")

	var err error
	Logger, err = newJSONLogger(
		// withDisableConsole(),
		withField("domain", "signin-go"),
		withTimeLayout(time.CSTLayout),
		withFileP(config.Server.ServerLogFile),
	)
	if err != nil {
		log.Fatalf("global.logger.Init Error: %v", err)
	}
}
