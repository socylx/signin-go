package todo

import (
	"time"
)

// LiveComment 课程留言表
type LiveComment struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID           uint32    `gorm:"column:user_id" json:"user_id"`                     // 评论人id
	ActivityID       uint32    `gorm:"column:activity_id" json:"activity_id"`             // 评论的课程id
	MarkType         uint32    `gorm:"column:mark_type" json:"mark_type"`                 // 1正常留言，2恶毒留言，3违规
	Title            string    `gorm:"column:title" json:"title"`                         // 评论的标题
	Content          string    `gorm:"column:content" json:"content"`                     // 内容
	TransformContent string    `gorm:"column:transform_content" json:"transform_content"` // 转换后的内容，原内容需要进行审核
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	OptUserID        uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
}

// TableName get sql table name.获取数据库表名
func (m *LiveComment) TableName() string {
	return "live_comment"
}
