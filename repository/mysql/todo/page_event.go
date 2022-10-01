package todo

import (
	"time"
)

// PageEvent [...]
type PageEvent struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`             // 事件发生的页面
	EventKey   string    `gorm:"column:event_key" json:"event_key"` // 事件key
	Params     string    `gorm:"column:params" json:"params"`       // 页面的参数，？后的值
	Data1      string    `gorm:"column:data_1" json:"data_1"`
	Data2      string    `gorm:"column:data_2" json:"data_2"`
	Data3      string    `gorm:"column:data_3" json:"data_3"`
	Type       string    `gorm:"column:type" json:"type"` // mini : 小程序, admin: 后台, fuli: 福利小程序
	SessionID  string    `gorm:"column:session_id" json:"session_id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	TempUserID string    `gorm:"column:temp_user_id" json:"temp_user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *PageEvent) TableName() string {
	return "page_event"
}
