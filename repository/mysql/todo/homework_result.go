package todo

import (
	"time"
)

// HomeworkResult [...]
type HomeworkResult struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	HomeworkID uint32    `gorm:"column:homework_id" json:"homework_id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`
	Permission uint32    `gorm:"column:permission" json:"permission"`
	Show       bool      `gorm:"column:show" json:"show"`
	Status     uint32    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	JoinID     uint32    `gorm:"column:join_id" json:"join_id"`
}

// TableName get sql table name.获取数据库表名
func (m *HomeworkResult) TableName() string {
	return "homework_result"
}
