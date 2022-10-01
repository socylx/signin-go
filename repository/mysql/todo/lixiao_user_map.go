package todo

import (
	"time"
)

// LixiaoUserMap 励销系统客户，线索 与我们系统用户的对应表
type LixiaoUserMap struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`         // 对应我们系统里的user_id
	CustomerID string    `gorm:"column:customer_id" json:"customer_id"` // 励销系统里的客户id
	LeadsID    string    `gorm:"column:leads_id" json:"leads_id"`       // 励销系统里的线索id
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *LixiaoUserMap) TableName() string {
	return "lixiao_user_map"
}
