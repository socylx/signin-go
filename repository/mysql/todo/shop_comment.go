package todo

import (
	"time"
)

// ShopComment 商品评论
type ShopComment struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"` // 评论人id
	SpuID      uint32    `gorm:"column:spu_id" json:"spu_id"`   // 评论的商品spu id
	SkuID      uint32    `gorm:"column:sku_id" json:"sku_id"`   // 评论的商品sku id
	Index      uint32    `gorm:"column:index" json:"index"`     // 评论显示的排序，值越小越靠前
	Title      string    `gorm:"column:title" json:"title"`
	IsStar     bool      `gorm:"column:is_star" json:"is_star"` // 是否是星级好评
	Content    string    `gorm:"column:content" json:"content"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopComment) TableName() string {
	return "shop_comment"
}
