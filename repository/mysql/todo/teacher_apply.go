package todo

import (
	"time"
)

// TeacherApply [...]
type TeacherApply struct {
	ID                   uint32    `gorm:"primaryKey;column:id" json:"id"`
	Nickname             string    `gorm:"column:nickname" json:"nickname"`                 // 艺名
	Name                 string    `gorm:"column:name" json:"name"`                         // 真实姓名
	Phone                string    `gorm:"column:phone" json:"phone"`                       // 电话
	Sex                  uint32    `gorm:"column:sex" json:"sex"`                           // 性别
	Birthday             time.Time `gorm:"column:birthday" json:"birthday"`                 // 出生日期
	DanceStartTime       time.Time `gorm:"column:dance_start_time" json:"dance_start_time"` // 开始跳舞的时间
	TeachStartTime       time.Time `gorm:"column:teach_start_time" json:"teach_start_time"` // 开始授课的时间
	Content              string    `gorm:"column:content" json:"content"`                   // 舞蹈生涯/荣誉
	Dance                string    `gorm:"column:dance" json:"dance"`                       // 申请教授的舞种ID字符串, 如: 1,2,3,8
	Studio               string    `gorm:"column:studio" json:"studio"`                     // 意向门店ID字符串, 如: 1,2,3,8
	ProvinceID           uint32    `gorm:"column:province_id" json:"province_id"`           // 省id
	CityID               uint32    `gorm:"column:city_id" json:"city_id"`                   // 市id
	CountyID             uint32    `gorm:"column:county_id" json:"county_id"`               // 区id
	Address              string    `gorm:"column:address" json:"address"`                   // 住址
	GraduateSchool       string    `gorm:"column:graduate_school" json:"graduate_school"`   // 毕业院校
	Remark               string    `gorm:"column:remark" json:"remark"`                     // 备注
	Status               uint32    `gorm:"column:status" json:"status"`                     // 状态
	UserID               uint32    `gorm:"column:user_id" json:"user_id"`                   // 申请人
	ShareUserID          uint32    `gorm:"column:share_user_id" json:"share_user_id"`       // 邀请人
	ResultMsg            string    `gorm:"column:result_msg" json:"result_msg"`             // 审核结果描述
	Score                uint32    `gorm:"column:score" json:"score"`                       // 得分平均值
	CreateTime           time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime           time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel                bool      `gorm:"column:is_del" json:"is_del"`
	TeacherManagerRemark string    `gorm:"column:teacher_manager_remark" json:"teacher_manager_remark"` // 师资备注
	Wechat               string    `gorm:"column:wechat" json:"wechat"`                                 // 微信号
	TeamBackground       string    `gorm:"column:team_background" json:"team_background"`               // 团队背景
}

// TableName get sql table name.获取数据库表名
func (m *TeacherApply) TableName() string {
	return "teacher_apply"
}
