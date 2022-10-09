package coupon

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

func Detail(ctx core.StdContext, couponID uint32) (data *Coupon, err error) {
	db := mysql.DB.WithContext(ctx)
	data = &Coupon{}
	err = db.Table("coupon").Where("coupon.id = ?", couponID).First(data).Error
	return
}
