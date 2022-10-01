package todo

import (
	"time"
)

// QuestionnaireAnswer [...]
type QuestionnaireAnswer struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID          uint32    `gorm:"column:user_id" json:"user_id"`
	QuestionnaireID uint32    `gorm:"column:questionnaire_id" json:"questionnaire_id"`
	Contact         string    `gorm:"column:contact" json:"contact"`
	ContactName     string    `gorm:"column:contact_name" json:"contact_name"`
	AskUserID       uint32    `gorm:"column:ask_user_id" json:"ask_user_id"`
	Latitude        string    `gorm:"column:latitude" json:"latitude"`   // 位置纬度
	Longitude       string    `gorm:"column:longitude" json:"longitude"` // 位置经度
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *QuestionnaireAnswer) TableName() string {
	return "questionnaire_answer"
}
