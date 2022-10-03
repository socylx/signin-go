package middleware

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
)

func CheckPermission(c core.Context, key string) {
	sessionUserInfo := c.SessionUserInfo()
	if !sessionUserInfo.SystemPage[key] {
		c.AbortWithError(core.Error(
			code.PermissionError,
			code.Text(code.PermissionError)),
		)
		return
	}
}

func CheckPermissions(c core.Context, keys ...string) {
	sessionUserInfo := c.SessionUserInfo()

	var permission bool
	for _, key := range keys {
		if sessionUserInfo.SystemPage[key] {
			permission = true
			break
		}
	}
	if !permission {
		c.AbortWithError(core.Error(
			code.PermissionError,
			code.Text(code.PermissionError)),
		)
		return
	}
}
