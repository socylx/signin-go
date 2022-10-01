package todo

import (
	"time"
)

// ActivityGroup [...]
type ActivityGroup struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	Title           string    `gorm:"column:title" json:"title"`                       // 系列课名称
	CoverImg        string    `gorm:"column:cover_img" json:"cover_img"`               // 封面图片
	DescVideo       string    `gorm:"column:desc_video" json:"desc_video"`             // 视频介绍
	Abstract        string    `gorm:"column:abstract" json:"abstract"`                 // 摘要
	CourseKindID    uint32    `gorm:"column:course_kind_id" json:"course_kind_id"`     // 舞种
	CourseGradeID   uint32    `gorm:"column:course_grade_id" json:"course_grade_id"`   // 难度
	Desc            string    `gorm:"column:desc" json:"desc"`                         // 课程介绍
	Status          uint32    `gorm:"column:status" json:"status"`                     // 状态
	BuyInformation  string    `gorm:"column:buy_information" json:"buy_information"`   // 购买须知
	CertificateName string    `gorm:"column:certificate_name" json:"certificate_name"` // 证书名称
	PriceAmount     float32   `gorm:"column:price_amount" json:"price_amount"`         // 现金价格
	OriginalPrice   float32   `gorm:"column:original_price" json:"original_price"`     // 原价
	EffectiveDays   uint32    `gorm:"column:effective_days" json:"effective_days"`     // 有效天数
	PublishStatus   bool      `gorm:"column:publish_status" json:"publish_status"`     // 是否发布
	PublishTime     time.Time `gorm:"column:publish_time" json:"publish_time"`         // 发布时间
	CreateUserID    uint32    `gorm:"column:create_user_id" json:"create_user_id"`     // 创建人ID
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityGroup) TableName() string {
	return "activity_group"
}
