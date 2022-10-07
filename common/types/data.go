package types

/*
课程顾问

	员工对应的UserID
	续卡目标值
	上周续卡率
	近三十天续卡率
	已完成的续卡金额
	已跟进的UserIDs
*/
type ConsultantRenewData struct {
	UserID            uint32   `json:"user_id"`
	TargetValue       uint64   `json:"target_value"`
	LastweekRenewRate int      `json:"lastweek_renew_rate"`
	Nearly30RenewRate int      `json:"nearly_30_renew_rate"`
	RenewAmount       float64  `json:"renew_amount"`
	FollowUserIDs     []uint64 `json:"follow_user_ids"`
}

/*
店长或顾问的ID

	StaffID
	StaffUserID
*/
type StudioConsultantOnlyID struct {
	ID     uint32
	UserID uint32
}
