package todo

import (
	"time"
)

// JudgeUser [...]
type JudgeUser struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeID           uint32    `gorm:"column:judge_id" json:"judge_id"`
	UserID            uint32    `gorm:"column:user_id" json:"user_id"`
	Mark              string    `gorm:"column:mark" json:"mark"`
	Status            uint32    `gorm:"column:status" json:"status"` // 1待审核，2审核通过，3不通过
	Msg               string    `gorm:"column:msg" json:"msg"`       // 审核结果描述
	AwardMark         string    `gorm:"column:award_mark" json:"award_mark"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	CurrentNickname   string    `gorm:"column:current_nickname" json:"current_nickname"`
	CurrentHeadimgurl string    `gorm:"column:current_headimgurl" json:"current_headimgurl"`
	AwardType         uint32    `gorm:"column:award_type" json:"award_type"`     // 1 券 2 会员卡 3 现金
	AwardValue        uint32    `gorm:"column:award_value" json:"award_value"`   //  课次\现金分
	AwardID           string    `gorm:"column:award_id" json:"award_id"`         // 奖励产生的券ID/会员卡ID/现金奖励订单号
	AwardStatus       uint32    `gorm:"column:award_status" json:"award_status"` // 奖励状态
	AwardTime         time.Time `gorm:"column:award_time" json:"award_time"`     // 发放奖励成功时间
}

// TableName get sql table name.获取数据库表名
func (m *JudgeUser) TableName() string {
	return "judge_user"
}
