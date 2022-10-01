package todo

import (
	"time"
)

// Relation 账号关系表
type Relation struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`        // 主键 id
	ParentID   int       `gorm:"column:parent_id" json:"parent_id"`     // 母账号id
	ChildID    int       `gorm:"column:child_id" json:"child_id"`       // 子账号id
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	Remark     string    `gorm:"column:remark" json:"remark"`           // 备注
}

// TableName get sql table name.获取数据库表名
func (m *Relation) TableName() string {
	return "relation"
}
