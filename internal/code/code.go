package code

import (
	_ "embed"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	CacheSetError      = 10106
	CacheGetError      = 10107
	CacheDelError      = 10108
	CacheNotExist      = 10109
	ResubmitError      = 10110
	HashIdsEncodeError = 10111
	HashIdsDecodeError = 10112
	PermissionError    = 10113
	RedisConnectError  = 10114
	MySQLConnectError  = 10115
	WriteConfigError   = 10116
	SendEmailError     = 10117
	MySQLExecError     = 10118
	GoVersionError     = 10119
	SocketConnectError = 10120
	SocketSendError    = 10121

	UsersCreateError             = 20101
	UsersListError               = 20102
	UsersDeleteError             = 20103
	UsersUpdateError             = 20104
	UsersModifyPersonalInfoError = 20105
	UsersDetailError             = 20106
	UsersDataError               = 20107

	StrategyCreateError    = 20201
	StrategyQueryError     = 20202
	StrategyDeleteError    = 20203
	StrategyUpdateError    = 20204
	SetStrategyStatusError = 20205

	StrategyIndicatorQueryError = 20301

	StudioQueryError = 20401

	UserBeforeMemberQueryError = 20501
)

func Text(code int) string {
	return text[code]
}
