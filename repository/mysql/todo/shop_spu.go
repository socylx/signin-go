package todo

import (
	"time"
)

// ShopSpu 商城，商品spu
type ShopSpu struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	PreCode        string    `gorm:"column:pre_code" json:"pre_code"`         // spu的前置值 1000
	Title          string    `gorm:"column:title" json:"title"`               // 标题
	SubTitle       string    `gorm:"column:sub_title" json:"sub_title"`       // 副标题，存放促销信息等
	Status         uint32    `gorm:"column:status" json:"status"`             // 1未发布，2发布，3下架
	Desc           string    `gorm:"column:desc" json:"desc"`                 // 商品描述
	Keyword        string    `gorm:"column:keyword" json:"keyword"`           // 商品关键字，用于搜索等
	Limit          uint32    `gorm:"column:limit" json:"limit"`               // 每人限购数量，0是不限制
	Category1ID    uint32    `gorm:"column:category1_id" json:"category1_id"` // 第一级分类
	Category2ID    uint32    `gorm:"column:category2_id" json:"category2_id"` // 第二级分类
	Category3ID    uint32    `gorm:"column:category3_id" json:"category3_id"` // 第三级分类
	OptUserID      uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`   // 修改人id
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	ShopBrandID    uint32    `gorm:"column:shop_brand_id" json:"shop_brand_id"`       // 商品品牌ID
	ShortName      string    `gorm:"column:short_name" json:"short_name"`             // 商品短名称
	SellingPoint   string    `gorm:"column:selling_point" json:"selling_point"`       // 卖点
	NeedExpressFee bool      `gorm:"column:need_express_fee" json:"need_express_fee"` // 是否需要运费，1需要，0免运费
}

// TableName get sql table name.获取数据库表名
func (m *ShopSpu) TableName() string {
	return "shop_spu"
}
