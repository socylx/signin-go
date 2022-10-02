package logger

import (
	"log"
	"signin-go/global/config"
	"signin-go/global/time"

	"go.uber.org/zap"
)

var Logger *zap.Logger

var GORMLogger *zap.Logger

func Init() {
	log.Println("global.logger.Init Start...")

	var err error
	Logger, err = NewJSONLogger(
		// withDisableConsole(),
		withField("domain", "signin-go"),
		withTimeLayout(time.CSTLayout),
		withFileP(config.Server.ServerLogFile),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.Logger Error: %v", err)
	}

	GORMLogger, err = NewJSONLogger(
		// withDisableConsole(),
		withField("domain", "mysql"),
		withTimeLayout(time.CSTLayout),
		withFileP(config.Server.SQLLogFile),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.GORMLogger Error: %v", err)
	}
}

func Close() {
	_ = Logger.Sync()
	_ = GORMLogger.Sync()
}
