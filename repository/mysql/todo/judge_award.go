package todo

import (
	"time"
)

// JudgeAward [...]
type JudgeAward struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeID    uint32    `gorm:"column:judge_id" json:"judge_id"`
	Type       uint32    `gorm:"column:type" json:"type"` // 奖励类型，1课次，2现金，100自定义
	Num        uint32    `gorm:"column:num" json:"num"`   // 最小额度 ,int ,现金是分为单位
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	AwardNum   uint32    `gorm:"column:award_num" json:"award_num"` // 奖励的数量
}

// TableName get sql table name.获取数据库表名
func (m *JudgeAward) TableName() string {
	return "judge_award"
}
