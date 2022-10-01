package todo

import (
	"time"
)

// TeacherSalaryActivity 老师的费用结算记录，每节课都有一条记录
type TeacherSalaryActivity struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID uint32    `gorm:"column:activity_id" json:"activity_id"` // 课程id
	Type       uint32    `gorm:"column:type" json:"type"`               // 1月结，2现金现结
	Salary     float32   `gorm:"column:salary" json:"salary"`           // 支付金额
	Status     uint32    `gorm:"column:status" json:"status"`           // 1初始状态，2财务已确认
	Mark       string    `gorm:"column:mark" json:"mark"`               // 结算备注
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人
	SalaryType uint32    `gorm:"column:salary_type" json:"salary_type"` // 账单类型, 1: 课程账单, 2: 迟到扣除，3：交通费
	TeacherID  uint32    `gorm:"column:teacher_id" json:"teacher_id"`   // 老师ID
}

// TableName get sql table name.获取数据库表名
func (m *TeacherSalaryActivity) TableName() string {
	return "teacher_salary_activity"
}
