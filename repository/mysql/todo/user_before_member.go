package todo

import (
	"time"
)

// UserBeforeMember 售前用户表，可以由管理员创建，当用户自己注册系统后，本表会和user表关联
type UserBeforeMember struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID          uint32    `gorm:"column:user_id" json:"user_id"`                     // 对应系统注册用户id
	SourceID        uint32    `gorm:"column:source_id" json:"source_id"`                 // 来源id
	Name            string    `gorm:"column:name" json:"name"`                           // 名字
	Phone           string    `gorm:"column:phone" json:"phone"`                         // 手机号，当用户注册后，会用来自动对应
	VirtualPhone    string    `gorm:"column:virtual_phone" json:"virtual_phone"`         // 虚拟手机号
	Sex             uint32    `gorm:"column:sex" json:"sex"`                             // 性别，1男，2女
	Age             uint32    `gorm:"column:age" json:"age"`                             // 年龄
	BelongStudioID  uint32    `gorm:"column:belong_studio_id" json:"belong_studio_id"`   // 所属门店id
	LixiaoUserID    string    `gorm:"column:lixiao_user_id" json:"lixiao_user_id"`       // 线索对应的励销里的用户id
	Wechat          string    `gorm:"column:wechat" json:"wechat"`                       // 微信号
	Mark            string    `gorm:"column:mark" json:"mark"`                           // 备注
	Mark2           string    `gorm:"column:mark2" json:"mark2"`                         // 备注2
	FollowStatusID  uint32    `gorm:"column:follow_status_id" json:"follow_status_id"`   // 售前用户状态
	CreateUserID    uint32    `gorm:"column:create_user_id" json:"create_user_id"`       // 客户创建人id
	ManagerUserID   uint32    `gorm:"column:manager_user_id" json:"manager_user_id"`     // 负责人id
	LixiaoManagerID string    `gorm:"column:lixiao_manager_id" json:"lixiao_manager_id"` // 励销里的所属人id
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	CustomerType    uint32    `gorm:"column:customer_type" json:"customer_type"` // 客户类型，1成人客户，2少儿客户
	CustomerCity    uint32    `gorm:"column:customer_city" json:"customer_city"` // 客户区域，1北京，2非北京
	TransferTime    time.Time `gorm:"column:transfer_time" json:"transfer_time"` // 转移到门店的时间
}

// TableName get sql table name.获取数据库表名
func (m *UserBeforeMember) TableName() string {
	return "user_before_member"
}
