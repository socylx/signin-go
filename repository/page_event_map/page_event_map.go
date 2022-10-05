package page_event_map

import (
	"time"
)

// PageEventMap [...]
type PageEventMap struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	EventKey   string    `gorm:"column:event_key" json:"event_key"`
	EventTitle string    `gorm:"column:event_title" json:"event_title"`
	Mark       string    `gorm:"column:mark" json:"mark"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}
