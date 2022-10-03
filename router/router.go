package router

import (
	v1 "signin-go/api/v1"
	"signin-go/internal/alert"
	"signin-go/internal/core"
	"signin-go/internal/metrics"
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
		core.WithAlertNotify(alert.NotifyHandler()),
		core.WithRecordMetrics(metrics.RecordHandler()),
		core.WithEnableCors(),
	)

	routerGroup := mux.Group("/v1")
	v1.Router(routerGroup)

	return mux
}
