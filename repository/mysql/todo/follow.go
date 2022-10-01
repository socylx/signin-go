package todo

import (
	"time"
)

// Follow 用户跟进记录表
type Follow struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID         uint32    `gorm:"column:user_id" json:"user_id"`                   // 用户id
	UserSnapshotID uint32    `gorm:"column:user_snapshot_id" json:"user_snapshot_id"` // 用户快照Id
	ForType        uint32    `gorm:"column:for_type" json:"for_type"`                 // 1是注册用户跟进，2线索跟进
	OptUserID      uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`           // 添加记录的管理员
	Description    string    `gorm:"column:description" json:"description"`           // 跟进记录描述
	FollowStatusID uint32    `gorm:"column:follow_status_id" json:"follow_status_id"` // 状态id
	Mark           string    `gorm:"column:mark" json:"mark"`                         // 备注
	ImgURL         string    `gorm:"column:img_url" json:"img_url"`                   // 图片附件
	FileURL        string    `gorm:"column:file_url" json:"file_url"`                 // 任意文件附件
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
	NextFollowDate time.Time `gorm:"column:next_follow_date" json:"next_follow_date"` // 下次跟进时间
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	FollowType     uint32    `gorm:"column:follow_type" json:"follow_type"`     // 跟进类型 1=售前 2=售后（续费）3=售后（会员服务）
	FollowWeight   uint32    `gorm:"column:follow_weight" json:"follow_weight"` // 跟进权重 100=低 200=中 300=高 400=超级高 0=默认排最低
	FollowTag      uint32    `gorm:"column:follow_tag" json:"follow_tag"`       // 1: 客服跟进标记
}

// TableName get sql table name.获取数据库表名
func (m *Follow) TableName() string {
	return "follow"
}
