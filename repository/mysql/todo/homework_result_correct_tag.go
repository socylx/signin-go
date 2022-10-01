package todo

import (
	"time"
)

// HomeworkResultCorrectTag [...]
type HomeworkResultCorrectTag struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ScoreMin   uint32    `gorm:"column:score_min" json:"score_min"`
	ScoreMax   uint32    `gorm:"column:score_max" json:"score_max"`
	Name       string    `gorm:"column:name" json:"name"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *HomeworkResultCorrectTag) TableName() string {
	return "homework_result_correct_tag"
}
