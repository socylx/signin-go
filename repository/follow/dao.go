package follow

import (
	"signin-go/global/mysql"
	"signin-go/global/time"
	"signin-go/internal/core"
)

/*
获取某个门店某一时间内续卡跟进过的学员UserIDs
*/
func GetFollowUserIDs(ctx core.StdContext, startTime, endTime time.Time, studioID, staffUserID uint32) (followUserIDs []uint32, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table(tableName()).
		Select("follow.user_id").
		Joins("JOIN staff on follow.opt_user_id = staff.user_id").
		Joins("JOIN role_page on staff.role_id = role_page.role_id").
		Joins("JOIN permission on role_page.role_id = permission.role_id").
		Where("follow.create_time >= ? AND follow.create_time < ? AND follow.is_del = 0 AND follow.for_type = ? AND follow.follow_type = ?", startTime, endTime, User, Renewal).
		Where("role_page.system_page_id = 5 AND role_page.is_del = 0").
		Where("permission.is_del = 0 AND permission.apply_studio_id = ?", studioID)
	if staffUserID > 0 {
		query = query.Where("staff.user_id = ?", staffUserID)
	}
	err = query.Find(&followUserIDs).Error

	return
}
