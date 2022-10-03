package staff

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

type StaffRolePage struct {
	ID                uint32
	RoleID            uint32
	RolePageID        uint32
	SystemPageID      uint32
	SystemPagePageKey string
}

func StaffRolePageData(ctx core.Context, userID uint32) (data *StaffRolePage, err error) {
	db := mysql.DB.WithContext(ctx.RequestContext())

	data = &StaffRolePage{}
	err = db.Table(
		tableName(),
	).Select(
		"staff.id,staff.role_id,role_page.id as role_page_id,system_page.id as system_page_id,system_page.page_key as system_page_page_key",
	).Joins(
		"join role_page on staff.role_id = role_page.role_id",
	).Joins(
		"join system_page on role_page.system_page_id = system_page.id",
	).Where(
		"staff.is_del = 0 AND staff.user_id = ? AND role_page.is_del = 0 AND system_page.is_del = 0", userID,
	).Order(
		"staff.user_id ASC",
	).Limit(1).Find(&data).Error
	return
}
