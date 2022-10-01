package todo

import (
	"time"
)

// Signin [...]
type Signin struct {
	ID                        uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID                    uint32    `gorm:"column:user_id" json:"user_id"`                   // 约课人id
	MembershipID              uint32    `gorm:"column:membership_id" json:"membership_id"`       // 预约使用的卡id
	ActivityID                uint32    `gorm:"column:activity_id" json:"activity_id"`           // 课程活动id
	OrderID                   uint32    `gorm:"column:order_id" json:"order_id"`                 // 现金预约课程时，订单id
	Spend                     float32   `gorm:"column:spend" json:"spend"`                       // 约课所消耗的卡次
	Discount                  float32   `gorm:"column:discount" json:"discount"`                 // 折扣
	SysRemark                 string    `gorm:"column:sys_remark" json:"sys_remark"`             // 折扣描述
	Score                     uint32    `gorm:"column:score" json:"score"`                       // 评分
	Comment                   string    `gorm:"column:comment" json:"comment"`                   // 评价
	CreateTime                time.Time `gorm:"column:create_time" json:"create_time"`           // 创建时间
	UpdateTime                time.Time `gorm:"column:update_time" json:"update_time"`           // 修改时间
	IsDel                     bool      `gorm:"column:is_del" json:"is_del"`                     // 1为删除
	CurrentRemains            float32   `gorm:"column:current_remains" json:"current_remains"`   // 当前记录对应剩余次数
	Type                      int       `gorm:"column:type" json:"type"`                         // 操作类型，1签到，2取消签到，3预约
	OptUserID                 int       `gorm:"column:opt_user_id" json:"opt_user_id"`           // 执行这次操作的用户id
	IsBack                    bool      `gorm:"column:is_back" json:"is_back"`                   // 这条签到是否被取消，0没有取消，1 已取消
	IsNotice                  bool      `gorm:"column:is_notice" json:"is_notice"`               // 是否发送了上课提醒
	MembershipSpend           float32   `gorm:"column:membership_spend" json:"membership_spend"` // 卡消耗次数
	EndTime                   time.Time `gorm:"column:end_time" json:"end_time"`                 // 过期时间
	ActivityGroupUserRecordID uint32    `gorm:"column:activity_group_user_record_id" json:"activity_group_user_record_id"`
}

// TableName get sql table name.获取数据库表名
func (m *Signin) TableName() string {
	return "signin"
}
