package todo

import (
	"time"
)

// ShopIndex 商城首页数据
type ShopIndex struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"` // 首页备注
	Data       string    `gorm:"column:data" json:"data"`
	DataURL    string    `gorm:"column:data_url" json:"data_url"`       // json链接，数据会传到阿里云oss,以json格式存起来
	IsShow     bool      `gorm:"column:is_show" json:"is_show"`         // 是否显示到首页，1显示，不显示。同一时间，只有一个会显得到首页
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 商品类别，1商品，2课程
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopIndex) TableName() string {
	return "shop_index"
}
