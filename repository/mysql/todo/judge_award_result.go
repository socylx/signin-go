package todo

import (
	"time"
)

// JudgeAwardResult [...]
type JudgeAwardResult struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeUserID  uint32    `gorm:"column:judge_user_id" json:"judge_user_id"`
	JudgeAwardID uint32    `gorm:"column:judge_award_id" json:"judge_award_id"`
	Status       uint32    `gorm:"column:status" json:"status"`
	Mark         string    `gorm:"column:mark" json:"mark"`
	InstanceID   string    `gorm:"column:instance_id" json:"instance_id"` // 奖励产生的券ID/会员卡ID/现金奖励订单号
	AwardTime    time.Time `gorm:"column:award_time" json:"award_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	Type         uint32    `gorm:"column:type" json:"type"`
	AwardValue   uint32    `gorm:"column:award_value" json:"award_value"`
}

// TableName get sql table name.获取数据库表名
func (m *JudgeAwardResult) TableName() string {
	return "judge_award_result"
}
