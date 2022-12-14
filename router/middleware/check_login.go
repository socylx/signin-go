package middleware

import (
	"gsteps-go/internal/code"
	"gsteps-go/internal/core"
)

/*
检查是否登录
*/
func CheckLogin() core.HandlerFunc {
	return func(c core.Context) {
		sessionUserInfo := c.SessionUserInfo()
		if sessionUserInfo.UserID <= 0 {
			c.AbortWithError(core.Error(
				code.AuthorizationError,
				code.Text(code.AuthorizationError)),
			)
			return
		}
	}
}
