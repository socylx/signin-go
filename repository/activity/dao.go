package activity

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func Detail(ctx core.StdContext, activityID uint32) (activity *Activity, err error) {
	db := mysql.DB.WithContext(ctx)

	activity = &Activity{}
	err = db.Table(
		tableName(),
	).Where(
		"activity.is_del = 0 AND activity.id = ?", activityID,
	).First(activity).Error
	return
}
