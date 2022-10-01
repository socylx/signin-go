package todo

import (
	"time"
)

// MiniFormID 存储小程序生成的 formId, 用来推送
type MiniFormID struct {
	ID      int       `gorm:"primaryKey;column:id" json:"id"`
	FormID  string    `gorm:"column:form_id" json:"form_id"`   // 微信小程序的 formId
	UserID  int       `gorm:"column:user_id" json:"user_id"`   // 属于的用户id
	EndTime time.Time `gorm:"column:end_time" json:"end_time"` // formId 过期时间
}

// TableName get sql table name.获取数据库表名
func (m *MiniFormID) TableName() string {
	return "mini_form_id"
}
