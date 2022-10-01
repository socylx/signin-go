package todo

import (
	"time"
)

// LiveHeartstack 直播课用户留存心跳记录
type LiveHeartstack struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserIDList string    `gorm:"column:user_id_list" json:"user_id_list"` // 在线用户id数组，逗号分割
	ActivityID uint32    `gorm:"column:activity_id" json:"activity_id"`   // 直播课程id
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 这条记录修改时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 心跳的创建时间
}

// TableName get sql table name.获取数据库表名
func (m *LiveHeartstack) TableName() string {
	return "live_heartstack"
}
