package activity

import (
	"time"
)

// Activity [...]
type Activity struct {
	ID                  uint32    `gorm:"primaryKey;column:id" json:"id"`
	CourseID            uint32    `gorm:"column:course_id" json:"course_id"`                 // 课程id
	CustomTitle         string    `gorm:"column:custom_title" json:"custom_title"`           // 课程自定义标题
	AdminUserID         uint32    `gorm:"column:admin_user_id" json:"admin_user_id"`         // 负责人user_id, 可能不止是老师
	Index               int       `gorm:"column:index" json:"index"`                         // 课程周期的第几次课程
	IsProxy             bool      `gorm:"column:is_proxy" json:"is_proxy"`                   // 是否是代理人
	StartTime           time.Time `gorm:"column:start_time" json:"start_time"`               // 课程开始时间
	ClassroomID         uint32    `gorm:"column:classroom_id" json:"classroom_id"`           // 教室id
	MaxMember           uint32    `gorm:"column:max_member" json:"max_member"`               // 最多上课人数
	MinMember           uint32    `gorm:"column:min_member" json:"min_member"`               // 最少上课人数
	CreateTime          time.Time `gorm:"column:create_time" json:"create_time"`             // 创建时间
	UpdateTime          time.Time `gorm:"column:update_time" json:"update_time"`             // 更新时间
	IsDel               bool      `gorm:"column:is_del" json:"is_del"`                       // 1为删除
	TeacherID           uint32    `gorm:"column:teacher_id" json:"teacher_id"`               // 教课老师id
	PublishTime         time.Time `gorm:"column:publish_time" json:"publish_time"`           // 课程发布时间，对外课预约
	IsIsolation         bool      `gorm:"column:is_isolation" json:"is_isolation"`           // 本堂课是否为独立课
	AudioURL            string    `gorm:"column:audio_url" json:"audio_url"`                 // 课堂音频url
	LivePushURL         string    `gorm:"column:live_push_url" json:"live_push_url"`         // 直播推流源
	LivePlayURL         string    `gorm:"column:live_play_url" json:"live_play_url"`         // 直播播放源
	LivePlayURL2        string    `gorm:"column:live_play_url_2" json:"live_play_url_2"`     // 直播播放源
	MarkType            uint32    `gorm:"column:mark_type" json:"mark_type"`                 // 1结课，2拍摄
	SongName            string    `gorm:"column:song_name" json:"song_name"`                 // 歌曲名字
	TeacherComeTime     time.Time `gorm:"column:teacher_come_time" json:"teacher_come_time"` // 老师到店时间
	PreviewVideoURL     string    `gorm:"column:preview_video_url" json:"preview_video_url"` // 本次课的预告视频
	PreviewVideoPoster  string    `gorm:"column:preview_video_poster" json:"preview_video_poster"`
	PreviewVideoHeight  uint32    `gorm:"column:preview_video_height" json:"preview_video_height"`
	PreviewVideoWidth   uint32    `gorm:"column:preview_video_width" json:"preview_video_width"`
	PreviewShowVideo    bool      `gorm:"column:preview_show_video" json:"preview_show_video"` // 是否显示
	PreviewVideoUpDate  time.Time `gorm:"column:preview_video_up_date" json:"preview_video_up_date"`
	EffectiveDays       uint32    `gorm:"column:effective_days" json:"effective_days"` // 有效时长
	DependMembership    bool      `gorm:"column:depend_membership" json:"depend_membership"`
	EffectiveDesc       string    `gorm:"column:effective_desc" json:"effective_desc"`     // 课程有效期说明
	TranslationDesc     string    `gorm:"column:translation_desc" json:"translation_desc"` // 课堂翻译
	ShowTranslationDesc bool      `gorm:"column:show_translation_desc" json:"show_translation_desc"`
}
