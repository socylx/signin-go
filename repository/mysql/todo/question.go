package todo

import (
	"time"
)

// Question [...]
type Question struct {
	ID                     uint32    `gorm:"primaryKey;column:id" json:"id"`
	QuestionnaireID        uint32    `gorm:"column:questionnaire_id" json:"questionnaire_id"`
	ParentQuestionOptionID uint32    `gorm:"column:parent_question_option_id" json:"parent_question_option_id"`
	Content                string    `gorm:"column:content" json:"content"`
	Type                   uint32    `gorm:"column:type" json:"type"`       // 1:单选	2:多选	3:问答
	IsMust                 bool      `gorm:"column:is_must" json:"is_must"` // 是否必答
	Index                  uint32    `gorm:"column:index" json:"index"`
	Remark                 string    `gorm:"column:remark" json:"remark"`
	Weight                 uint32    `gorm:"column:weight" json:"weight"` // 权重，评断分数
	CreateTime             time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime             time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                  bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Question) TableName() string {
	return "question"
}
