package todo

import (
	"time"
)

// TeacherFree [...]
type TeacherFree struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherID  uint32    `gorm:"column:teacher_id" json:"teacher_id"`
	Value      uint32    `gorm:"column:value" json:"value"`           // 值
	ValueType  uint32    `gorm:"column:value_type" json:"value_type"` // 1、门店, 2、周几
	Start      uint32    `gorm:"column:start" json:"start"`
	End        uint32    `gorm:"column:end" json:"end"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherFree) TableName() string {
	return "teacher_free"
}
