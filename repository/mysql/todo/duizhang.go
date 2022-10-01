package todo

import (
	"time"
)

// Duizhang [...]
type Duizhang struct {
	ID                    uint32    `gorm:"primaryKey;column:id" json:"id"`
	BankSerialNumber      string    `gorm:"column:bank_serial_number" json:"bank_serial_number"`           // 银行流水号
	GatheringPurposesCode string    `gorm:"column:gathering_purposes_code" json:"gathering_purposes_code"` // 收款用途code
	GatheringPurposesName string    `gorm:"column:gathering_purposes_name" json:"gathering_purposes_name"` // 收款用途name
	Remark                string    `gorm:"column:remark" json:"remark"`                                   // 备注
	CreateTime            time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                 bool      `gorm:"column:is_del" json:"is_del"`
	Status                uint32    `gorm:"column:status" json:"status"`   // 状态
	IsPush                bool      `gorm:"column:is_push" json:"is_push"` // 是否推送到金蝶
	OptUserID             uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	CompleteTime          time.Time `gorm:"column:complete_time" json:"complete_time"`
	PaymentBankID         uint32    `gorm:"column:payment_bank_id" json:"payment_bank_id"`
	SysRemark             string    `gorm:"column:sys_remark" json:"sys_remark"`     // 系统日志
	KingdeeCode           string    `gorm:"column:kingdee_code" json:"kingdee_code"` // 金蝶收款单编码/易快报单据编码
	Amtcdr                string    `gorm:"column:amtcdr" json:"amtcdr"`             // 借贷标记
}

// TableName get sql table name.获取数据库表名
func (m *Duizhang) TableName() string {
	return "duizhang"
}
