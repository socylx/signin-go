package todo

import (
	"time"
)

// OnlineVideo [...]
type OnlineVideo struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type              uint32    `gorm:"column:type" json:"type"`                       // 1元素片段，2录播视频
	Title             string    `gorm:"column:title" json:"title"`                     // 视频标题
	Decs              string    `gorm:"column:decs" json:"decs"`                       // 视频描述
	URL               string    `gorm:"column:url" json:"url"`                         // 视频链接
	Poster            string    `gorm:"column:poster" json:"poster"`                   // 封面
	Duration          uint32    `gorm:"column:duration" json:"duration"`               // 视频时长
	TeacherID         uint32    `gorm:"column:teacher_id" json:"teacher_id"`           // 视频对应的老师
	CourseGradeID     uint32    `gorm:"column:course_grade_id" json:"course_grade_id"` // 视频难度等级
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID         uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 上传人id
	PublishTime       time.Time `gorm:"column:publish_time" json:"publish_time"`
	PublishUserID     uint32    `gorm:"column:publish_user_id" json:"publish_user_id"`
	MarkPublishTime   time.Time `gorm:"column:mark_publish_time" json:"mark_publish_time"`
	MarkPublishUserID uint32    `gorm:"column:mark_publish_user_id" json:"mark_publish_user_id"`
	Status            uint32    `gorm:"column:status" json:"status"`
}

// TableName get sql table name.获取数据库表名
func (m *OnlineVideo) TableName() string {
	return "online_video"
}
