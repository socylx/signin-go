package todo

import (
	"time"
)

// VideoMarkMake [...]
type VideoMarkMake struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	ForID        uint32    `gorm:"column:for_id" json:"for_id"`
	ForType      uint32    `gorm:"column:for_type" json:"for_type"`
	Status       uint32    `gorm:"column:status" json:"status"`
	MakeUserID   uint32    `gorm:"column:make_user_id" json:"make_user_id"`
	ReviewUserID uint32    `gorm:"column:review_user_id" json:"review_user_id"`
	OptUserID    uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	SubmitTime   time.Time `gorm:"column:submit_time" json:"submit_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`           // 1代表删除
}

// TableName get sql table name.获取数据库表名
func (m *VideoMarkMake) TableName() string {
	return "video_mark_make"
}
