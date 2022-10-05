package page_access

import (
	"time"
)

// PageAccess [...]
type PageAccess struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`       // 页面路径
	Title      string    `gorm:"column:title" json:"title"`   // 页面标题
	Params     string    `gorm:"column:params" json:"params"` // 页面的参数，？后的值
	Data1      string    `gorm:"column:data_1" json:"data_1"`
	Data2      string    `gorm:"column:data_2" json:"data_2"`
	Data3      string    `gorm:"column:data_3" json:"data_3"`
	Type       string    `gorm:"column:type" json:"type"`             // mini : 小程序, admin: 后台, fuli: 福利小程序
	SessionID  string    `gorm:"column:session_id" json:"session_id"` // 会话session
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	TempUserID string    `gorm:"column:temp_user_id" json:"temp_user_id"`
	SourceID   uint32    `gorm:"column:source_id" json:"source_id"` // 来源
	SceneID    string    `gorm:"column:scene_id" json:"scene_id"`   // 小程序场景值
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}
