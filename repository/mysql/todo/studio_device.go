package todo

import (
	"time"
)

// StudioDevice [...]
type StudioDevice struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"id"`
	StudioID       uint32    `gorm:"column:studio_id" json:"studio_id"`
	DeviceType     uint32    `gorm:"column:device_type" json:"device_type"`           // 设备类型, 1: 收款终端设备
	DeviceTypeCode string    `gorm:"column:device_type_code" json:"device_type_code"` // 设备字符串类型码
	SerialNum      string    `gorm:"column:serial_num" json:"serial_num"`             // 终端序列号
	DeviceID       string    `gorm:"column:device_id" json:"device_id"`               // 终端设备号
	Status         uint32    `gorm:"column:status" json:"status"`                     // 是否在使用
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel          bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *StudioDevice) TableName() string {
	return "studio_device"
}
