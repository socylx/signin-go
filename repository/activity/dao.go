package activity

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

/*
获取某一课程的详情
*/
func Detail(ctx core.StdContext, activityID uint32) (activity *Activity, err error) {
	db := mysql.DB.WithContext(ctx)

	activity = &Activity{}
	err = db.Table("activity").
		Where("activity.is_del = 0 AND activity.id = ?", activityID).First(activity).Error
	return
}

type Filter struct {
	IncludeIDs []uint32
	StudioID   uint32
}

/*
获取课程IDs
*/
func GetActivityIDs(ctx core.StdContext, filter *Filter) (data []uint32, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("activity").
		Select("activity.id").
		Joins("JOIN classroom ON activity.classroom_id = classroom.id").
		Where("activity.is_del = 0 AND activity.id IN ? AND classroom.is_del = 0 AND classroom.studio_id = ?", filter.IncludeIDs, filter.StudioID).
		Find(&data).Error
	return
}
