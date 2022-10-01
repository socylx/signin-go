package todo

import (
	"time"
)

// Pv [...]
type Pv struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`
	ClientIP   string    `gorm:"column:client_ip" json:"client_ip"` // 访问者ip
	UnionID    string    `gorm:"column:union_id" json:"union_id"`   // 访问者唯一union_id
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel      int16     `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Pv) TableName() string {
	return "pv"
}
