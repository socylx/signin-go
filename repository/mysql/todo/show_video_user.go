package todo

import (
	"time"
)

// ShowVideoUser 结课视频如果没有对应的课程id, 则这个表里设置参与者（老师和学生）
type ShowVideoUser struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShowVideoID uint32    `gorm:"column:show_video_id" json:"show_video_id"` // 对应的结课视频id
	Type        uint32    `gorm:"column:type" json:"type"`                   // 参与人类型,1学员，2老师
	PlayerID    uint32    `gorm:"column:player_id" json:"player_id"`         // 参与人id，type=1是user_id, type=2是teacher_id
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShowVideoUser) TableName() string {
	return "show_video_user"
}
