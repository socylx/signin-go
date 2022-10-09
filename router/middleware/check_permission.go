package middleware

import (
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
)

/*
检查是否有某个角色的权限
*/
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

/*
检查是否有任意一个角色的权限
*/
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
