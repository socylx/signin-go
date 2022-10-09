package studio

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

func GetStudioIDs(ctx core.StdContext) (data []uint32, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("studio").
		Select("studio.id").
		Where("studio.is_del = 0").Find(&data).Error
	return
}
