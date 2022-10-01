package todo

import (
	"time"
)

// ExpenditureExpectedIncome 预计收入设置表
type ExpenditureExpectedIncome struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`   // 开始时间
	EndTime    time.Time `gorm:"column:end_time" json:"end_time"`       // 结束时间
	StudioID   uint32    `gorm:"column:studio_id" json:"studio_id"`     // 对应门店id
	Value      float32   `gorm:"column:value" json:"value"`             // 设置的每天收入
	Status     uint32    `gorm:"column:status" json:"status"`           // 1正常，2弃用
	Mark       string    `gorm:"column:mark" json:"mark"`               // 描述
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // user_id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ExpenditureExpectedIncome) TableName() string {
	return "expenditure_expected_income"
}
