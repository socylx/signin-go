package todo

import (
	"time"
)

// Course [...]
type Course struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	BrandID       uint32    `gorm:"column:brand_id" json:"brand_id"`
	CreateUserID  uint32    `gorm:"column:create_user_id" json:"create_user_id"`   // 创建人用户id
	CourseKindID  uint32    `gorm:"column:course_kind_id" json:"course_kind_id"`   // 课程类型id
	CourseLevelID uint32    `gorm:"column:course_level_id" json:"course_level_id"` // 课程级别id
	CourseGradeID uint32    `gorm:"column:course_grade_id" json:"course_grade_id"` // 课程等级id
	CourseWay     uint32    `gorm:"column:course_way" json:"course_way"`           // 授课方式，1线下，2线上
	Duration      uint32    `gorm:"column:duration" json:"duration"`               // 课程时长
	CardIDs       string    `gorm:"column:card_ids" json:"card_ids"`               // 允许使用的卡的类型id
	Spend         float32   `gorm:"column:spend" json:"spend"`                     // 大于0时，可以现金支付
	Desc          string    `gorm:"column:desc" json:"desc"`                       // 课程描述
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
	Name          string    `gorm:"column:name" json:"name"`                   // 课程名称
	IsBundling    bool      `gorm:"column:is_bundling" json:"is_bundling"`     // 课程是否捆绑销售，0不是，1是；1代表整个课程组需要一次性预约（例如集训课）
	VideoURL      string    `gorm:"column:video_url" json:"video_url"`         // 老师的开课视频
	VideoPoster   string    `gorm:"column:video_poster" json:"video_poster"`   // 视频封面
	VideoHeight   uint32    `gorm:"column:video_height" json:"video_height"`   // 预告视频高
	VideoWidth    uint32    `gorm:"column:video_width" json:"video_width"`     // 预告视频width
	VideoUpDate   time.Time `gorm:"column:video_up_date" json:"video_up_date"` // 视频上传时间
	ShowVideo     bool      `gorm:"column:show_video" json:"show_video"`       // 是否显示这个视频，视频需要编码，编码后才将这个值设置为1
	TeacherMark   string    `gorm:"column:teacher_mark" json:"teacher_mark"`   // 老师对这课程组的建议
}

// TableName get sql table name.获取数据库表名
func (m *Course) TableName() string {
	return "course"
}
