package todo

import (
	"time"
)

// CallRecord [...]
type CallRecord struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	UniqueID       string    `gorm:"column:unique_id" json:"unique_id"` // 通话记录唯一标识
	UserID         uint32    `gorm:"column:user_id" json:"user_id"`
	UserBeforeID   uint32    `gorm:"column:user_before_id" json:"user_before_id"`
	Hotline        string    `gorm:"column:hotline" json:"hotline"`                 // 热线号码
	CustomerNumber string    `gorm:"column:customer_number" json:"customer_number"` // 客户号码
	Cno            string    `gorm:"column:cno" json:"cno"`                         // 座席号，要求只能是 4-6 位数字
	ClientName     string    `gorm:"column:client_name" json:"client_name"`         // 坐席姓名
	ClientNumber   string    `gorm:"column:client_number" json:"client_number"`     // 坐席号码
	CallType       string    `gorm:"column:call_type" json:"call_type"`             // 呼叫类型
	StartTime      time.Time `gorm:"column:start_time" json:"start_time"`           // 开始时间
	EndTime        time.Time `gorm:"column:end_time" json:"end_time"`               // 结束时间
	BridgeTime     time.Time `gorm:"column:bridge_time" json:"bridge_time"`         // 接通时间
	BridgeDuration uint32    `gorm:"column:bridge_duration" json:"bridge_duration"` // 接通时长
	TotalDuration  uint32    `gorm:"column:total_duration" json:"total_duration"`   // 总时长
	Status         string    `gorm:"column:status" json:"status"`                   // 接听状态
	EndReason      string    `gorm:"column:end_reason" json:"end_reason"`           // 挂机方
	SipCause       string    `gorm:"column:sip_cause" json:"sip_cause"`             // 呼叫结果
	UserField      string    `gorm:"column:user_field" json:"user_field"`           // 自定义字段
	Mark           int       `gorm:"column:mark" json:"mark"`                       // 标记
	Tags           string    `gorm:"column:tags" json:"tags"`                       // 标签
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`         // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`         // 更新时间
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
	RecordFile     string    `gorm:"column:record_file" json:"record_file"`           // 录音文件
	UserSnapshotID uint32    `gorm:"column:user_snapshot_id" json:"user_snapshot_id"` // 拨打电话时对应的用户剩余卡次节点ID
	RecordType     uint32    `gorm:"column:record_type" json:"record_type"`           // 1:线索电话跟进， 2：会员跟进
}

// TableName get sql table name.获取数据库表名
func (m *CallRecord) TableName() string {
	return "call_record"
}
