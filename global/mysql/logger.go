package mysql

import (
	"context"
	"errors"
	"fmt"
	glogger "gsteps-go/global/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	loggerConfig
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	logger                              *zap.Logger
}

type loggerConfig struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	LogLevel                  logger.LogLevel
}

func newLogger(c loggerConfig) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if c.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &gormLogger{
		loggerConfig: c,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		logger:       glogger.GORMLogger,
	}
}

func (log *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l := *log
	l.LogLevel = level
	return &l
}

func (log *gormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if log.LogLevel >= logger.Info {
		log.logger.Sugar().Infof(log.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, args...)...)
		// log.logger.Info(
		// 	log.infoStr+msg+fmt.Sprintln(append([]interface{}{utils.FileWithLineNum()}, data...)),
		// 	zap.Any("TraceID", "SNJ179293HNA51792LS"),
		// )
	}
}

func (log *gormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if log.LogLevel >= logger.Warn {
		log.logger.Sugar().Warnf(log.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, args...)...)
	}
}

func (log *gormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if log.LogLevel >= logger.Error {
		log.logger.Sugar().Errorf(log.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, args...)...)
	}
}

// Trace print sql message
func (log *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if log.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && log.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !log.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			log.logger.Sugar().Errorf(log.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.logger.Sugar().Errorf(log.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > log.SlowThreshold && log.SlowThreshold != 0 && log.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", log.SlowThreshold)
		if rows == -1 {
			log.logger.Sugar().Warnf(log.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.logger.Sugar().Warnf(log.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case log.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			log.logger.Sugar().Infof(log.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.logger.Sugar().Infof(log.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
