package todo

// TeacherDancetype [...]
type TeacherDancetype struct {
	ID          uint32 `gorm:"primaryKey;column:id" json:"id"`
	TeacherID   int    `gorm:"column:teacher_id" json:"teacher_id"`       // 教师id
	DanceTypeID int    `gorm:"column:dance_type_id" json:"dance_type_id"` // 舞种
	IsDel       bool   `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *TeacherDancetype) TableName() string {
	return "teacher_dancetype"
}
