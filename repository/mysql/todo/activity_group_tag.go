package todo

import (
	"time"
)

// ActivityGroupTag [...]
type ActivityGroupTag struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityGroupID uint32    `gorm:"column:activity_group_id" json:"activity_group_id"`
	TagDetailID     uint32    `gorm:"column:tag_detail_id" json:"tag_detail_id"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityGroupTag) TableName() string {
	return "activity_group_tag"
}
