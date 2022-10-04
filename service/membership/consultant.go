package membership

import (
	"signin-go/internal/core"
	"signin-go/repository/card"
	"signin-go/repository/membership"
	"time"
)

/*
获取顾问某段时间内的续卡金额
*/
func GetConsultantRenewAmount(ctx core.StdContext, startTime, endTime time.Time, studioID, staffUserID uint32) (amount uint64) {
	memberships, _ := membership.GetMembershipDatas(ctx, &membership.MembershipFilter{
		CreateTimeGE: startTime,
		CreateTimeLT: endTime,
		CardID:       card.AdultCard,
	})

	userIDs := []uint32{}
	for _, m := range memberships {
		userIDs = append(userIDs, m.UserID)
	}

	beforeMemberships, _ := membership.GetMembershipDatas(ctx, &membership.MembershipFilter{
		CreateTimeLT:   startTime,
		CardID:         card.AdultCard,
		IncludeUserIDs: userIDs,
	})

	beforeUserMembership := map[uint32]bool{}
	for _, beforeMembership := range beforeMemberships {
		beforeUserMembership[beforeMembership.UserID] = true
	}

	for _, m := range memberships {
		if m.BelongsStudioID != studioID || (staffUserID > 0 && staffUserID != m.SalesUserID) {
			beforeUserMembership[m.UserID] = true
			continue
		}
		if !beforeUserMembership[m.UserID] {
			beforeUserMembership[m.UserID] = true
			continue
		}
		amount += uint64(m.Amount * 100)
	}
	return
}
