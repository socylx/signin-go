package todo

import (
	"time"
)

// Questionnaire [...]
type Questionnaire struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	Title        string    `gorm:"column:title" json:"title"`
	DeadlineTime time.Time `gorm:"column:deadline_time" json:"deadline_time"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	CreateUserID uint32    `gorm:"column:create_user_id" json:"create_user_id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	Type         uint32    `gorm:"column:type" json:"type"` // 1-普通问卷, 2-全职老师打分问卷
}

// TableName get sql table name.获取数据库表名
func (m *Questionnaire) TableName() string {
	return "questionnaire"
}
