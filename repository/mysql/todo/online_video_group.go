package todo

import (
	"time"
)

// OnlineVideoGroup [...]
type OnlineVideoGroup struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID uint32    `gorm:"column:activity_id" json:"activity_id"` // 对应的课程id
	Title      string    `gorm:"column:title" json:"title"`             // 录播组标题
	Decs       string    `gorm:"column:decs" json:"decs"`               // 录播课程描述
	SubTitle   string    `gorm:"column:sub_title" json:"sub_title"`     // 副标题
	Poster     string    `gorm:"column:poster" json:"poster"`           // 封面
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideoGroup) TableName() string {
	return "online_video_group"
}
