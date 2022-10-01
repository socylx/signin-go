package todo

import (
	"time"
)

// CourseLevel [...]
type CourseLevel struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type       string    `gorm:"column:type" json:"type"`               // 班级类型
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否被删除
	IsBundling bool      `gorm:"column:is_bundling" json:"is_bundling"` // 否绑定销售，例如小班课，集训课需要一次性预约所以的课程
	Notice     string    `gorm:"column:notice" json:"notice"`           // 注意事项
}

// TableName get sql table name.获取数据库表名
func (m *CourseLevel) TableName() string {
	return "course_level"
}
