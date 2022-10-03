package permission

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func PermissionApplyStudioIDs(ctx core.Context, userID uint32) (applyStudioIDs []uint32, err error) {
	db := mysql.DB.WithContext(ctx.RequestContext())

	applyStudioIDs = []uint32{}
	err = db.Table(tableName()).
		Select("permission.apply_studio_id").
		Where("permission.is_del = 0 AND permission.role_id = ?", userID).
		Find(&applyStudioIDs).Error

	return
}
