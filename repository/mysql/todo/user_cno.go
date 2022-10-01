package todo

import (
	"time"
)

// UserCno [...]
type UserCno struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Cno        string    `gorm:"column:cno" json:"cno"`                 // 座席号
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`         // user_id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *UserCno) TableName() string {
	return "user_cno"
}
