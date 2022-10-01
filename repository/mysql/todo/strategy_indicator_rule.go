package todo

import (
	"time"
)

// StrategyIndicatorRule [...]
type StrategyIndicatorRule struct {
	ID                  uint32    `gorm:"primaryKey;column:id" json:"id"`
	StrategyIndicatorID uint32    `gorm:"column:strategy_indicator_id" json:"strategy_indicator_id"`
	Name                string    `gorm:"column:name" json:"name"`
	Min                 string    `gorm:"column:min" json:"min"`
	Max                 string    `gorm:"column:max" json:"max"`
	CreateTime          time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime          time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel               bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *StrategyIndicatorRule) TableName() string {
	return "strategy_indicator_rule"
}
