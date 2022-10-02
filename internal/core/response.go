package core

type Response struct {
	Code    int         `json:"code"` // 业务码
	Message string      `json:"msg"`  // 描述信息
	Res     interface{} `json:"res"`
}
