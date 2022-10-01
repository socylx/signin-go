package todo

import (
	"time"
)

// ActivityVideo 上课后的课堂视频
type ActivityVideo struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	VideoURL          string    `gorm:"column:video_url" json:"video_url"`       // 视频链接
	VideoPoster       string    `gorm:"column:video_poster" json:"video_poster"` // 视频封面
	ActivityID        uint32    `gorm:"column:activity_id" json:"activity_id"`   // 对应的课程id
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID         uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 上传人id
	Decs              string    `gorm:"column:decs" json:"decs"`               // 视频描述
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	VideoHeight       uint32    `gorm:"column:video_height" json:"video_height"` // 视频 height
	VideoWidth        uint32    `gorm:"column:video_width" json:"video_width"`   // 视频width
	ShowVideo         bool      `gorm:"column:show_video" json:"show_video"`     // 是否显示这个视频，视频需要编码，编码后才将这个值设置为1
	PublishTime       time.Time `gorm:"column:publish_time" json:"publish_time"`
	PublishUserID     uint32    `gorm:"column:publish_user_id" json:"publish_user_id"`
	MarkPublishTime   time.Time `gorm:"column:mark_publish_time" json:"mark_publish_time"`
	MarkPublishUserID uint32    `gorm:"column:mark_publish_user_id" json:"mark_publish_user_id"`
	Status            uint32    `gorm:"column:status" json:"status"`
	Title             string    `gorm:"column:title" json:"title"`
	Type              uint32    `gorm:"column:type" json:"type"` // 1-正常视频, 2-试看视频
}

// TableName get sql table name.获取数据库表名
func (m *ActivityVideo) TableName() string {
	return "activity_video"
}
