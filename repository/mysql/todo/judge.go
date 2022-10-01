package todo

import (
	"time"
)

// Judge [...]
type Judge struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	Key               string    `gorm:"column:key" json:"key"`                       // 字符串标识
	Title             string    `gorm:"column:title" json:"title"`                   // 标题，龙德店开业活动
	GUIDeImg          string    `gorm:"column:guide_img" json:"guide_img"`           // 操作指导图，指定怎么上传图片，朋友圈截图示例
	GUIDeContent      string    `gorm:"column:guide_content" json:"guide_content"`   // 操作指定描述
	AwardContent      string    `gorm:"column:award_content" json:"award_content"`   // 奖励描述
	SysMark           string    `gorm:"column:sys_mark" json:"sys_mark"`             // 系统备注
	Type              uint32    `gorm:"column:type" json:"type"`                     // 类型，1转发公众号文章，2海报
	ModifyUserID      uint32    `gorm:"column:modify_user_id" json:"modify_user_id"` // 最后一次修改的人
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	StartTime         time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime           time.Time `gorm:"column:end_time" json:"end_time"`
	CheckAwardContent string    `gorm:"column:check_award_content" json:"check_award_content"` // 查看奖励说明
	CheckAwardURL     string    `gorm:"column:check_award_url" json:"check_award_url"`         // 查看奖励图文
}

// TableName get sql table name.获取数据库表名
func (m *Judge) TableName() string {
	return "judge"
}
