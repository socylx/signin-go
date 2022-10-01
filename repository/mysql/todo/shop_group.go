package todo

import (
	"time"
)

// ShopGroup [...]
type ShopGroup struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Desc        string    `gorm:"column:desc" json:"desc"`             // 活动描述
	Min         uint32    `gorm:"column:min" json:"min"`               // 最小成团人数
	Max         uint32    `gorm:"column:max" json:"max"`               // 参团最多人数
	GetLimit    uint32    `gorm:"column:get_limit" json:"get_limit"`   // 限制参与活动次数
	Duration    uint32    `gorm:"column:duration" json:"duration"`     // 分钟，拼团持续时间
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"` // 拼团开始时间
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`     // 拼团结束时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`
	OptUserID   uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`     // 最后编辑人id
	Status      uint32    `gorm:"column:status" json:"status"`               // 1-正常, 2-停用
	RuleDesc    string    `gorm:"column:rule_desc" json:"rule_desc"`         // 参与规则描述
	PosterBgImg string    `gorm:"column:poster_bg_img" json:"poster_bg_img"` // 分享海报的背景图
}

// TableName get sql table name.获取数据库表名
func (m *ShopGroup) TableName() string {
	return "shop_group"
}
