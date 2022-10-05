package signin

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetSigninDataByCouponAllocID(ctx core.StdContext, couponAllocID uint32) (data *SigninData, err error) {
	db := mysql.DB.WithContext(ctx)
	data = &SigninData{}
	err = db.Table("signin").
		Select("signin.id, signin.user_id as signin_user_id, activity.start_time as activity_start_time, course.course_level_id as course_level_id").
		Joins("JOIN coupon_alloc_log on signin.id = coupon_alloc_log.for_id").
		Joins("JOIN activity on signin.activity_id = activity.id").
		Joins("JOIN course on activity.course_id = course.id").
		Where("signin.is_del = 0 AND signin.is_back = 0 AND signin.type IN (1, 3) AND coupon_alloc_log.is_del = 0 AND coupon_alloc_log.type = 1 AND coupon_alloc_log.coupon_alloc_id = ?", couponAllocID).
		Order("signin.id").First(data).Error
	return
}
