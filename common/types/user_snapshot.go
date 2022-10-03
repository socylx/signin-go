package types

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
