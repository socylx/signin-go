package signin

import (
	"signin-go/global/mysql"
	"signin-go/global/time"
	"signin-go/internal/core"
)

func GetSigninDataByCouponAllocID(ctx core.StdContext, couponAllocID uint32) (data *SigninData, err error) {
	db := mysql.DB.WithContext(ctx)
	data = &SigninData{}
	err = db.Table("signin").
		Select("signin.id, signin.user_id as signin_user_id,activity.id as activity_id,activity.start_time as activity_start_time,course.id as course_id, course.course_level_id as course_level_id").
		Joins("JOIN coupon_alloc_log on signin.id = coupon_alloc_log.for_id").
		Joins("JOIN activity on signin.activity_id = activity.id").
		Joins("JOIN course on activity.course_id = course.id").
		Where("signin.is_del = 0 AND signin.is_back = 0 AND signin.type IN (1, 3) AND coupon_alloc_log.is_del = 0 AND coupon_alloc_log.type = 1 AND coupon_alloc_log.coupon_alloc_id = ?", couponAllocID).
		Order("signin.id").First(data).Error
	return
}

type Filter struct {
	UserID              uint32
	ActivityStartTimeGE time.Time
	ActivityStartTimeLT time.Time
	// Type                []SigninType
}

func GetSigninDatas(ctx core.StdContext, filter *Filter) (data []*SigninData, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("signin").
		Select("signin.id, signin.user_id as signin_user_id,activity.id as activity_id,activity.start_time as activity_start_time,course.id as course_id, course.course_level_id as course_level_id").
		Joins("JOIN activity on signin.activity_id = activity.id").
		Joins("JOIN course on activity.course_id = course.id").
		Where("signin.is_del = 0 AND signin.is_back = 0 AND signin.type IN (1,3) AND activity.is_del = 0").
		Where("")

	if filter.UserID > 0 {
		query = query.Where("signin.user_id = ?", filter.UserID)
	}
	if filter.ActivityStartTimeGE != time.TimeZeroTime {
		query = query.Where("activity.start_time >= ?", filter.ActivityStartTimeGE)
	}
	if filter.ActivityStartTimeLT != time.TimeZeroTime {
		query = query.Where("activity.start_time < ?", filter.ActivityStartTimeLT)
	}
	err = query.Find(&data).Error
	return
}
