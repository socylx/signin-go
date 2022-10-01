package todo

import (
	"time"
)

// MembershipLog 对会员卡的修改记录日志
type MembershipLog struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	MembershipID int       `gorm:"column:membership_id" json:"membership_id"`   // 对应修改的会员卡id
	Deadline     time.Time `gorm:"column:deadline" json:"deadline"`             // 有效期
	Remains      float32   `gorm:"column:remains" json:"remains"`               // 余额
	EntityCardID string    `gorm:"column:entity_card_id" json:"entity_card_id"` // 实体卡号
	Remark       string    `gorm:"column:remark" json:"remark"`                 // 修改备注，修改人填写
	OptUserID    int       `gorm:"column:opt_user_id" json:"opt_user_id"`       // 修改人id
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	SysRemark    string    `gorm:"column:sys_remark" json:"sys_remark"` // 系统自动生成的备注，例如 余额有10变为20
	OptType      uint32    `gorm:"column:opt_type" json:"opt_type"`     // 操作类型，1常规日志，2退款退卡，3开错删除，4余额转为其他卡
}

// TableName get sql table name.获取数据库表名
func (m *MembershipLog) TableName() string {
	return "membership_log"
}
