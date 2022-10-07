package types

/*
接口 /user_snapshot/accesstorenew 的 Response

	门店ID
	门店员工UserID
	续卡金额目标值
	上周续卡率
	近三十天续卡率
	已完成的续卡金额
	已跟进的学员UserIDs
	若`门店员工UserID`为0,则`ConsultantTargets`对应门店下员工的续卡目标的数据
*/
type AccessToRenewResponse struct {
	StudioID          uint32                 `json:"studio_id"`
	StaffUserID       uint32                 `json:"staff_user_id"`
	TargetValue       uint64                 `json:"target_value"`
	LastweekRenewRate int                    `json:"lastweek_renew_rate"`
	Nearly30RenewRate int                    `json:"nearly_30_renew_rate"`
	RenewAmount       float64                `json:"renew_amount"`
	FollowUserIDs     []uint64               `json:"follow_user_ids"`
	ConsultantTargets []*ConsultantRenewData `json:"consultant_targets"`
}
