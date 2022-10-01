package todo

import (
	"time"
)

// MembershipSnapshot [...]
type MembershipSnapshot struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`
	MID           uint32    `gorm:"column:m_id" json:"m_id"` // membership 表的id
	MCardID       int       `gorm:"column:m_card_id" json:"m_card_id"`
	MUserID       int       `gorm:"column:m_user_id" json:"m_user_id"`
	MCreateTime   time.Time `gorm:"column:m_create_time" json:"m_create_time"`
	MRemains      float32   `gorm:"column:m_remains" json:"m_remains"`
	MEntityCardID string    `gorm:"column:m_entity_card_id" json:"m_entity_card_id"`
	MPayment      int       `gorm:"column:m_payment" json:"m_payment"`
	MRemark       string    `gorm:"column:m_remark" json:"m_remark"`
	UName         string    `gorm:"column:u_name" json:"u_name"`
	UPhone        string    `gorm:"column:u_phone" json:"u_phone"`
	USource       string    `gorm:"column:u_source" json:"u_source"`
	UNickname     string    `gorm:"column:u_nickname" json:"u_nickname"`
	PPrice        int       `gorm:"column:p_price" json:"p_price"`
	PExpire       int       `gorm:"column:p_expire" json:"p_expire"`
	PRemark       string    `gorm:"column:p_remark" json:"p_remark"`
	PLimit        int       `gorm:"column:p_limit" json:"p_limit"`
	POnlyStaff    int       `gorm:"column:p_only_staff" json:"p_only_staff"`
	SName         string    `gorm:"column:s_name" json:"s_name"`
	SnapshotTime  time.Time `gorm:"column:snapshot_time" json:"snapshot_time"`
	MDeadline     time.Time `gorm:"column:m_deadline" json:"m_deadline"`
	IsDel         bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *MembershipSnapshot) TableName() string {
	return "membership_snapshot"
}
