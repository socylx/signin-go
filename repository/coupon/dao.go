package coupon

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func Detail(ctx core.StdContext, couponID uint32) (data *Coupon, err error) {
	db := mysql.DB.WithContext(ctx)
	data = &Coupon{}
	err = db.Table("coupon").Where("coupon.id = ?", couponID).First(data).Error
	return
}
