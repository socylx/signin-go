package todo

import (
	"time"
)

// MiniIndex 小程序的首页推荐 banner 数据
type MiniIndex struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`           // 点击后的跳转链接
	Img        string    `gorm:"column:img" json:"img"`           // 展示的头图
	Title      string    `gorm:"column:title" json:"title"`       // 标图
	Desc       string    `gorm:"column:desc" json:"desc"`         // 描述
	Type       int       `gorm:"column:type" json:"type"`         // 首页banner的类型，1公众号文章，2本小程序路由，3其他小程序
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // banner 属于哪个模块，1首页，2课程列表, 3课堂视频，4福利小程序首页
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`     // 1为删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Index      int       `gorm:"column:index" json:"index"` // 权重，数值越大，越排靠前
}

// TableName get sql table name.获取数据库表名
func (m *MiniIndex) TableName() string {
	return "mini_index"
}
