package todo

import (
	"time"
)

// TeacherApplyConfirmContent [...]
type TeacherApplyConfirmContent struct {
	ID                    uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherApplyConfirmID uint32    `gorm:"column:teacher_apply_confirm_id" json:"teacher_apply_confirm_id"`
	ForID                 uint32    `gorm:"column:for_id" json:"for_id"`
	ForType               int       `gorm:"column:for_type" json:"for_type"` // 1-舞种,2-标签,3-问卷
	CreateTime            time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                 bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherApplyConfirmContent) TableName() string {
	return "teacher_apply_confirm_content"
}
