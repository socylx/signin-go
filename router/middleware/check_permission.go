package middleware

import (
	"signin-go/internal/code"
	"signin-go/internal/core"
)

func CheckPermission(key string) core.HandlerFunc {
	return func(c core.Context) {
		sessionUserInfo := c.SessionUserInfo()
		if !sessionUserInfo.SystemPage[key] {
			c.AbortWithError(core.Error(
				code.PermissionError,
				code.Text(code.PermissionError)),
			)
			return
		}
	}
}

func CheckPermissions(keys ...string) core.HandlerFunc {
	return func(c core.Context) {
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
}
