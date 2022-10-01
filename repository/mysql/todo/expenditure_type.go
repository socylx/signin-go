package todo

import (
	"time"
)

// ExpenditureType 支出类型表
type ExpenditureType struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`               // 支出名称，工资，房租，装修 等等
	Type       uint32    `gorm:"column:type" json:"type"`               // 类别，备用
	Desc       string    `gorm:"column:desc" json:"desc"`               // 描述
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // user_id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ExpenditureType) TableName() string {
	return "expenditure_type"
}
