package fission_map

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

type Filter struct {
	ShareUserID uint32
	Type        FissionType
	StatusIn    []FissionStatus
}

func GetFissionMapData(ctx core.StdContext, filter *Filter) (data []*FissionMapData, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table("fission_map").Select("fission_map.id").Where("fission_map.is_del = 0")

	if filter.ShareUserID > 0 {
		query = query.Where("fission_map.share_user_id = ?", filter.ShareUserID)
	}
	if filter.Type > 0 {
		query = query.Where("fission_map.type = ?", filter.Type)
	}
	if len(filter.StatusIn) > 0 {
		query = query.Where("fission_map.status IN ?", filter.StatusIn)
	}
	err = query.Find(&data).Error
	return
}
