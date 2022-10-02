package router

import (
	v1 "signin-go/api/v1"
	"signin-go/internal/core"
)

type Response struct {
	Code    int         `json:"code"` // 业务码
	Message string      `json:"msg"`  // 描述信息
	Res     interface{} `json:"res"`
}

func HTTPServer() core.Mux {
	mux := core.New()

	v1.Init(mux)

	return mux
}
