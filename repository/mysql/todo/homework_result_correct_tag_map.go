package todo

import (
	"time"
)

// HomeworkResultCorrectTagMap [...]
type HomeworkResultCorrectTagMap struct {
	ID                         uint32    `gorm:"primaryKey;column:id" json:"id"`
	HomeworkResultCorrectID    uint32    `gorm:"column:homework_result_correct_id" json:"homework_result_correct_id"`
	HomeworkResultCorrectTagID uint32    `gorm:"column:homework_result_correct_tag_id" json:"homework_result_correct_tag_id"`
	CreateTime                 time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime                 time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *HomeworkResultCorrectTagMap) TableName() string {
	return "homework_result_correct_tag_map"
}
