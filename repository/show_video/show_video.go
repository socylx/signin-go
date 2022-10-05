package show_video

import (
	"time"
)

// ShowVideo 展示视频，可能是结课视频，也可能单独的视频
type ShowVideo struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityID   uint32    `gorm:"column:activity_id" json:"activity_id"`       // 课程id
	Title        string    `gorm:"column:title" json:"title"`                   // 标题
	QqVideoID    string    `gorm:"column:qq_video_id" json:"qq_video_id"`       // 腾讯视频的id
	VideoURL     string    `gorm:"column:video_url" json:"video_url"`           // 视频url
	Poster       string    `gorm:"column:poster" json:"poster"`                 // 封面
	OptUserID    uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`       // 上传人id
	TeacherID    uint32    `gorm:"column:teacher_id" json:"teacher_id"`         // 老师id
	StudioID     uint32    `gorm:"column:studio_id" json:"studio_id"`           // 门店id
	CourseKindID uint32    `gorm:"column:course_kind_id" json:"course_kind_id"` // 舞种id
	Word         string    `gorm:"column:word" json:"word"`                     // 关键字预存，为了方便搜索
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	ShowVideo    bool      `gorm:"column:show_video" json:"show_video"` // 是否显示视频，1已经转码完
	ShootTime    time.Time `gorm:"column:shoot_time" json:"shoot_time"` // 视频拍摄时间
	ShowMode     uint32    `gorm:"column:show_mode" json:"show_mode"`   // 显示方式，1小程序列表展示，2列表不展示（只展示详情页）
}
