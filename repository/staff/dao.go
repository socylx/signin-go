package staff

import (
	"signin-go/common/types"
	"signin-go/global/mysql"
	"signin-go/internal/core"
	"signin-go/repository/system_page"
)

/*
员工对应的数据

	员工ID
	角色ID
	-
	-
	-
*/
type StaffRolePage struct {
	ID                uint32
	RoleID            uint32
	RolePageID        uint32
	SystemPageID      uint32
	SystemPagePageKey string
}

// 员工的角色和队友的SystemPagePageKey
func StaffRolePageData(ctx core.StdContext, userID uint32) (data *StaffRolePage, err error) {
	db := mysql.DB.WithContext(ctx)

	data = &StaffRolePage{}
	err = db.Table("staff").
		Select("staff.id,staff.role_id,role_page.id as role_page_id,system_page.id as system_page_id,system_page.page_key as system_page_page_key").
		Joins("join role_page on staff.role_id = role_page.role_id").
		Joins("join system_page on role_page.system_page_id = system_page.id").
		Where("staff.is_del = 0 AND staff.user_id = ? AND role_page.is_del = 0 AND system_page.is_del = 0", userID).
		Order("staff.user_id ASC").
		Limit(1).
		Find(&data).Error
	return
}

// 获取某个门店的店长和课程顾问对应的ID
func StudioConsultantOnlyID(ctx core.StdContext, studioID uint32) (data []*types.StudioConsultantOnlyID, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table("staff").
		Select("staff.id, staff.user_id as staff_user_id").
		Joins("JOIN role_page on staff.role_id = role_page.role_id").
		Joins("JOIN permission on staff.role_id = permission.role_id").
		Where("staff.is_del = 0").
		Where("role_page.is_del = 0 AND role_page.system_page_id IN ?", [2]uint32{system_page.Studio.ID, system_page.Consultant.ID}).
		Where("permission.is_del = 0 AND permission.apply_studio_id = ?", studioID).
		Find(&data).Error

	return
}
