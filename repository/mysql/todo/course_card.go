package todo

import (
	"time"
)

// CourseCard 课程可使用的卡，以及支付数额表
type CourseCard struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	CardID     uint32    `gorm:"column:card_id" json:"card_id"`     // 可以使用的卡id,对应card表id
	CourseID   int       `gorm:"column:course_id" json:"course_id"` // 对应的课程id，对应 course 表的id
	Type       uint32    `gorm:"column:type" json:"type"`           // 1-课程，2-系列课
	Spend      float32   `gorm:"column:spend" json:"spend"`         // 本次课需支付数额，默认1
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *CourseCard) TableName() string {
	return "course_card"
}
