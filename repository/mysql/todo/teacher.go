package todo

import (
	"time"
)

// Teacher [...]
type Teacher struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID           int       `gorm:"column:user_id" json:"user_id"`                       // 所属用户user_id 一对一
	BrandID          int       `gorm:"column:brand_id" json:"brand_id"`                     // 所属品牌团体id
	Desc             string    `gorm:"column:desc" json:"desc"`                             // 教师简介
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`               // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`               // 更新时间
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`                         // 1代表删除
	IsProxy          bool      `gorm:"column:is_proxy" json:"is_proxy"`                     // 是否代课老师
	Avatar           string    `gorm:"column:avatar" json:"avatar"`                         // 老师的宣传头像
	PayeeName        string    `gorm:"column:payee_name" json:"payee_name"`                 // 老师真实姓名，用于工资发放
	PayeeAccount     string    `gorm:"column:payee_account" json:"payee_account"`           // 收款账号
	AccountBelongsTo string    `gorm:"column:account_belongs_to" json:"account_belongs_to"` // 账号所属，支付宝/微信/开户行
	JobType          uint32    `gorm:"column:job_type" json:"job_type"`                     // 1：全职，2：在职，3：兼职
	SalaryType       uint32    `gorm:"column:salary_type" json:"salary_type"`               // 1：月结，2：现金现结
	ArrangeCourse    bool      `gorm:"column:arrange_course" json:"arrange_course"`         // 是否给此老师排课
	Remark           string    `gorm:"column:remark" json:"remark"`                         // 备注
	ManagerUserID    uint32    `gorm:"column:manager_user_id" json:"manager_user_id"`       // 负责人id
}

// TableName get sql table name.获取数据库表名
func (m *Teacher) TableName() string {
	return "teacher"
}
