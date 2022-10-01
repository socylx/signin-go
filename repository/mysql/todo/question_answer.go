package todo

import (
	"time"
)

// QuestionAnswer [...]
type QuestionAnswer struct {
	ID                    uint32    `gorm:"primaryKey;column:id" json:"id"`
	QuestionnaireAnswerID uint32    `gorm:"column:questionnaire_answer_id" json:"questionnaire_answer_id"`
	QuestionID            uint32    `gorm:"column:question_id" json:"question_id"`
	QuestionOptionID      uint32    `gorm:"column:question_option_id" json:"question_option_id"`
	AnswerContent         string    `gorm:"column:answer_content" json:"answer_content"`
	CreateTime            time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                 bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *QuestionAnswer) TableName() string {
	return "question_answer"
}
