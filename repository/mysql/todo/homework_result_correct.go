package todo

import (
	"time"
)

// HomeworkResultCorrect [...]
type HomeworkResultCorrect struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	HomeworkResultID uint32    `gorm:"column:homework_result_id" json:"homework_result_id"`
	Content          string    `gorm:"column:content" json:"content"`
	Score            uint32    `gorm:"column:score" json:"score"`
	CorrectUserID    uint32    `gorm:"column:correct_user_id" json:"correct_user_id"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *HomeworkResultCorrect) TableName() string {
	return "homework_result_correct"
}
