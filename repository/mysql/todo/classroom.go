package todo

import (
	"time"
)

// Classroom [...]
type Classroom struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	StudioID   int       `gorm:"column:studio_id" json:"studio_id"`     // 所属工作室id
	Name       string    `gorm:"column:name" json:"name"`               // 教室名称
	Capacity   int       `gorm:"column:capacity" json:"capacity"`       // 可容纳人数
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否被删除
}

// TableName get sql table name.获取数据库表名
func (m *Classroom) TableName() string {
	return "classroom"
}
