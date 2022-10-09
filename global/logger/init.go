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
	Logger, err = NewLogger(
		WithDisableConsole(),
		WithTimeLayout(time.CSTLayout),
		WithJSONEncoder(),
		WithLevelLog(config.Server.ServerLogDir),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.Logger Error: %v", err)
	}

	GORMLogger, err = NewLogger(
		WithDisableConsole(),
		WithTimeLayout(time.CSTLayout),
		WithFileP(config.Server.SQLLogFile),
		WithNORMALEncoder(),
	)
	if err != nil {
		log.Fatalf("global.logger.Init.GORMLogger Error: %v", err)
	}
}

func Close() {
	_ = Logger.Sync()
	_ = GORMLogger.Sync()
}
