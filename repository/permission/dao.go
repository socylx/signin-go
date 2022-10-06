package permission

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

/*
获取某位员工有权限的门店IDs
*/
func PermissionApplyStudioIDs(ctx core.StdContext, roleID uint32) (applyStudioIDs []uint32, err error) {
	db := mysql.DB.WithContext(ctx)

	applyStudioIDs = []uint32{}
	err = db.Table("permission").
		Select("permission.apply_studio_id").
		Where("permission.is_del = 0 AND permission.role_id = ?", roleID).
		Find(&applyStudioIDs).Error

	return
}
