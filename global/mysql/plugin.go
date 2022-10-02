package mysql

import (
	"log"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/internal/trace"

	"gorm.io/gorm"
	"gorm.io/gorm/utils"
)

const (
	callBackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"
)

type TracePlugin struct{}

func (op *TracePlugin) Name() string {
	return "tracePlugin"
}

func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &TracePlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	return
}

func after(db *gorm.DB) {
	_ctx := db.Statement.Context

	log.Println("after._ctx: ", _ctx)
	ctx, ok := _ctx.(core.StdContext)
	if !ok {
		return
	}
	log.Println("after.ctx: ", ctx)

	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}
	log.Println("after._ts: ", _ts)

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}
	log.Println("after.ts: ", ts)

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	log.Println("after.sql: ", sql)

	sqlInfo := new(trace.SQL)
	sqlInfo.Timestamp = time.CSTLayoutString()
	sqlInfo.SQL = sql
	sqlInfo.Stack = utils.FileWithLineNum()
	sqlInfo.Rows = db.Statement.RowsAffected
	sqlInfo.CostSeconds = time.Since(ts).Seconds()

	if ctx.Trace != nil {
		ctx.Trace.AppendSQL(sqlInfo)
	}
	return
}
