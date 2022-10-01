package todo

import (
	"time"
)

// TeacherSalary 老师的费用，jazz常规课多少钱一节
type TeacherSalary struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	TeacherID        uint32    `gorm:"column:teacher_id" json:"teacher_id"`           // 教师id
	CourseLevelID    uint32    `gorm:"column:course_level_id" json:"course_level_id"` // 课程类型，常规课、小班课等
	DanceTypeID      uint32    `gorm:"column:dance_type_id" json:"dance_type_id"`     // 舞种
	BrandID          uint32    `gorm:"column:brand_id" json:"brand_id"`               // 品牌id
	StudioID         uint32    `gorm:"column:studio_id" json:"studio_id"`             // 门店id
	MinPlayerNum     uint32    `gorm:"column:min_player_num" json:"min_player_num"`   // 上课人数最小值，例如达到多少人工资不一样
	MaxPlayerNum     uint32    `gorm:"column:max_player_num" json:"max_player_num"`   // 上课人数最大值
	Salary           float32   `gorm:"column:salary" json:"salary"`                   // 费用
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID        uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`             // 操作人
	StudentType      uint32    `gorm:"column:student_type" json:"student_type"`           // 学员类型
	ActivityDuration uint32    `gorm:"column:activity_duration" json:"activity_duration"` // 课程时长
	MinTeacherNum    uint32    `gorm:"column:min_teacher_num" json:"min_teacher_num"`     // 上课的老师数
	MaxTeacherNum    uint32    `gorm:"column:max_teacher_num" json:"max_teacher_num"`     // 最大上课老师数
}

// TableName get sql table name.获取数据库表名
func (m *TeacherSalary) TableName() string {
	return "teacher_salary"
}
