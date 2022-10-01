package todo

import (
	"time"
)

// UserBeforeMemberTag 售前用户表，用户特征数据
type UserBeforeMemberTag struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserBeforeID uint32    `gorm:"column:user_before_id" json:"user_before_id"` // 售前客户id
	Type         uint32    `gorm:"column:type" json:"type"`                     // 特征类型，1舞种
	ItemID       uint32    `gorm:"column:item_id" json:"item_id"`               // 特征对应的其他表id，例如type=1(舞种)，本字段字对应舞种表的id
	Title        string    `gorm:"column:title" json:"title"`                   // 特征的值，冗余存储
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *UserBeforeMemberTag) TableName() string {
	return "user_before_member_tag"
}
