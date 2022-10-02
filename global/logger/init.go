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
		WithField("domain", "signin-go"),
		WithTimeLayout(time.CSTLayout),
		WithFileP(config.Server.ServerLogFile),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.Logger Error: %v", err)
	}

	GORMLogger, err = NewJSONLogger(
		// withDisableConsole(),
		WithField("domain", "mysql"),
		WithTimeLayout(time.CSTLayout),
		WithFileP(config.Server.SQLLogFile),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.GORMLogger Error: %v", err)
	}
}

func Close() {
	_ = Logger.Sync()
	_ = GORMLogger.Sync()
}
