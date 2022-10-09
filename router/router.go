package router

import (
	v1 "gsteps-go/api/v1"
	v2 "gsteps-go/api/v2"
	"gsteps-go/internal/core"
	"gsteps-go/router/middleware"
)

type Response struct {
	Code    int         `json:"code"` // 业务码
	Message string      `json:"msg"`  // 描述信息
	Res     interface{} `json:"res"`
}

func HTTPServer() core.Mux {
	mux := core.New(
		core.WithDisablePProf(),
		core.WithDisableSwagger(),
		core.WithDisablePrometheus(),
		// core.WithAlertNotify(alert.NotifyHandler()),
		// core.WithRecordMetrics(metrics.RecordHandler()),
		core.WithEnableCors(),
	)

	v1RouterGroup := mux.Group("/v1", middleware.SetSessionUserInfo)
	v1.Router(v1RouterGroup)

	v2RouterGroup := mux.Group("/v2", middleware.SetSessionUserInfo)
	v2.Router(v2RouterGroup)

	return mux
}
