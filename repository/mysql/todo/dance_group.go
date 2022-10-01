package todo

import (
	"time"
)

// DanceGroup [...]
type DanceGroup struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`               // 团体名称
	Desc       string    `gorm:"column:desc" json:"desc"`               // 团体描述
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 1代表删除
}

// TableName get sql table name.获取数据库表名
func (m *DanceGroup) TableName() string {
	return "dance_group"
}
