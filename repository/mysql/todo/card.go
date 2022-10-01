package todo

import (
	"time"
)

// Card [...]
type Card struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	BrandID       int       `gorm:"column:brand_id" json:"brand_id"` // 所属品牌id
	Type          int       `gorm:"column:type" json:"type"`         // 卡类型
	Level         string    `gorm:"column:level" json:"level"`       // 级别
	Desc          string    `gorm:"column:desc" json:"desc"`
	Priority      int       `gorm:"column:priority" json:"priority"`               // 权重
	ApplyStudioID int       `gorm:"column:apply_studio_id" json:"apply_studio_id"` // 适用场馆id
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`         // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`         // 更新时间
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`                   // 是否被删除
	CreateUserID  int       `gorm:"column:create_user_id" json:"create_user_id"`   // 创建者
}

// TableName get sql table name.获取数据库表名
func (m *Card) TableName() string {
	return "card"
}
