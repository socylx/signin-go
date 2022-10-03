package proposal

import (
	"signin-go/global/utils"
)

// SessionUserInfo 当前用户会话信息
type SessionUserInfo struct {
	UserID           uint32          `json:"user_id"`     // 用户ID
	SystemPage       map[string]bool `json:"system_page"` // system_page.page_key map bool
	StudioPermission map[uint32]bool `json:"permission"`  // Permission
}

// Marshal 序列化到JSON
func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = utils.Json.Marshal(user)
	return
}
