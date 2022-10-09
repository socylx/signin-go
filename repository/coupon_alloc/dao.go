package coupon_alloc

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

type Filter struct {
	UserID uint32
	Status []CouponAllocStatus
}

func GetCouponAllocs(ctx core.StdContext, filter *Filter) (data []*CouponAlloc, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("coupon_alloc").
		Joins("JOIN coupon on coupon_alloc.coupon_id = coupon.id").
		Where("coupon_alloc.is_del = 0")

	if len(filter.Status) > 0 {
		query = query.Where("coupon_alloc.status IN ?", filter.Status)
	}
	if filter.UserID > 0 {
		query = query.Where("coupon_alloc.user_id = ?", filter.UserID)
	}
	err = query.Find(&data).Error
	return
}
