package todo

import (
	"time"
)

// QuestionnaireForType [...]
type QuestionnaireForType struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	QuestionnaireID uint32    `gorm:"column:questionnaire_id" json:"questionnaire_id"`
	Type            uint32    `gorm:"column:type" json:"type"` // 面向人券,  0不限制，1有效会员，2失效会员，3新人，4非会员   5 老师 6 管理员
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *QuestionnaireForType) TableName() string {
	return "questionnaire_for_type"
}
