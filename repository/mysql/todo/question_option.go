package todo

import (
	"time"
)

// QuestionOption [...]
type QuestionOption struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint32    `gorm:"column:question_id" json:"question_id"`
	Content    string    `gorm:"column:content" json:"content"`
	Index      uint32    `gorm:"column:index" json:"index"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	Value      uint32    `gorm:"column:value" json:"value"` // 分值
}

// TableName get sql table name.获取数据库表名
func (m *QuestionOption) TableName() string {
	return "question_option"
}
