package todo

import (
	"time"
)

// Homework [...]
type Homework struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID   uint32    `gorm:"column:activity_id" json:"activity_id"`
	Content      string    `gorm:"column:content" json:"content"`
	Music        string    `gorm:"column:music" json:"music"`
	TeacherID    uint32    `gorm:"column:teacher_id" json:"teacher_id"`
	CreateUserID uint32    `gorm:"column:create_user_id" json:"create_user_id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Homework) TableName() string {
	return "homework"
}
