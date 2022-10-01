package todo

import (
	"time"
)

// CouponAllocLog 用户券的使用记录
type CouponAllocLog struct {
	ID            uint32    `gorm:"primaryKey;column:id" json:"id"`
	CouponAllocID uint32    `gorm:"column:coupon_alloc_id" json:"coupon_alloc_id"` // 券实例id
	Type          uint32    `gorm:"column:type" json:"type"`                       // 券用于什么场景，1约课，2订单
	ForID         uint32    `gorm:"column:for_id" json:"for_id"`                   // 券支付对象id, 例如签到id signin_id
	Spend         float32   `gorm:"column:spend" json:"spend"`                     // 消耗多少
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	SysMark       string    `gorm:"column:sys_mark" json:"sys_mark"`       // 系统自动生成备注
	Remark        string    `gorm:"column:remark" json:"remark"`           // 管理填写的备注
	OptUserID     uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 操作人id
}

// TableName get sql table name.获取数据库表名
func (m *CouponAllocLog) TableName() string {
	return "coupon_alloc_log"
}
