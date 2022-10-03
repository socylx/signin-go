package types

type ConsultantRenewData struct {
	UserID            uint32   `json:"user_id"`
	TargetValue       uint64   `json:"target_value"`
	LastweekRenewRate int      `json:"lastweek_renew_rate"`
	Nearly30RenewRate int      `json:"nearly_30_renew_rate"`
	RenewAmount       float64  `json:"renew_amount"`
	FollowUserIDs     []uint64 `json:"follow_user_ids"`
}

type StudioConsultantOnlyID struct {
	ID     uint32
	UserID uint32
}
