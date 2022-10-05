package judge_user_resources

import (
	"time"
)

// JudgeUserResources [...]
type JudgeUserResources struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeUserID uint32    `gorm:"column:judge_user_id" json:"judge_user_id"`
	URL         string    `gorm:"column:url" json:"url"`
	Type        uint32    `gorm:"column:type" json:"type"` // 1图片，2视频，3音频
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
}
