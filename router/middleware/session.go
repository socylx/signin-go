package middleware

import (
	"encoding/base64"
	"gsteps-go/global/redis"
	"gsteps-go/global/utils"
	"gsteps-go/internal/core"
	"gsteps-go/internal/proposal"
	"gsteps-go/repository/permission"
	"gsteps-go/repository/staff"
	"strings"
)

type sesssion struct {
	AccessTokken string `json:"access_token"` //令牌
	OpenID       string `json:"open_id"`      //open
	UserID       uint32 `json:"user_id"`      //用户id
	UnionID      string `json:"union_id"`     //唯一id
}

/*
向Context写入当前请求人的 ID/页面权限/门店权限 数据
*/
func SetSessionUserInfo(c core.Context) {
	sessionUserInfo := proposal.SessionUserInfo{
		UserID:           0,
		SystemPage:       map[string]bool{},
		StudioPermission: map[uint32]bool{},
	}
	defer c.SetSessionUserInfo(sessionUserInfo)

	cookie, cookidErr := c.GetCookie("session_id")
	if cookidErr != nil {
		return
	}

	cookies := strings.Split(cookie, "|")
	if len(cookies) != 6 {
		return
	}

	redisKey := strings.Split(cookies[4], ":")[1]
	byteRedisKey, decodeErr := base64.StdEncoding.DecodeString(redisKey)
	if decodeErr != nil {
		return
	}

	redisDataStr, redisErr := redis.Redis.Get(c.RequestContext(), string(byteRedisKey)).Result()
	if redisErr != nil {
		return
	}

	redisData := &sesssion{}
	unmarshalErr := utils.Json.Unmarshal([]byte(redisDataStr), redisData)
	if unmarshalErr != nil {
		return
	}
	if redisData.UserID <= 0 {
		return
	}
	sessionUserInfo.UserID = redisData.UserID

	userStaffRolePageData, err := staff.StaffRolePageData(c.RequestContext(), redisData.UserID)
	if err != nil {
		return
	}

	if userStaffRolePageData.SystemPagePageKey != "" {
		sessionUserInfo.SystemPage[userStaffRolePageData.SystemPagePageKey] = true
	}

	permissionApplyStudioIDs, err := permission.PermissionApplyStudioIDs(c.RequestContext(), userStaffRolePageData.RoleID)
	if err != nil || len(permissionApplyStudioIDs) <= 0 {
		return
	}

	sessionUserInfo.SystemPage[permission.Admin] = true
	for _, permissionApplyStudioID := range permissionApplyStudioIDs {
		sessionUserInfo.StudioPermission[permissionApplyStudioID] = true
	}
}
