package todo

import (
	"time"
)

// SubscribeItem 订阅的具体项目，例如订阅A老师，B舞种，则本表里就有两条数据
type SubscribeItem struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	SubscribeID uint32    `gorm:"column:subscribe_id" json:"subscribe_id"` // 所属订阅id
	Type        uint32    `gorm:"column:type" json:"type"`                 // 类型，1老师，2舞种，3门店，4课程类型，5课程难度，6上课方式
	OptUserID   uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`   // 操作人id
	ForID       uint32    `gorm:"column:for_id" json:"for_id"`             // 实际对应的值id, 例如type=1是，这个值就代表teacher_id
	ForID2      uint32    `gorm:"column:for_id_2" json:"for_id_2"`         // 备用，实际对应的值id, 例如type=1是，这个值就代表teacher_id
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *SubscribeItem) TableName() string {
	return "subscribe_item"
}
