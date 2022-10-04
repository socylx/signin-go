package users

import (
	"signin-go/internal/core"
	"signin-go/internal/errors"
	"signin-go/repository/follow"
	"signin-go/repository/membership"
	"signin-go/repository/user_before_member"
	"signin-go/repository/users"
	"sync"
)

type DataID struct {
	UserID             uint32 // 用户ID
	UserBeforeMemberID uint32 // 线索ID
}

func Data(ctx core.StdContext, dataID *DataID) (data *users.Data, err error) {
	var (
		user             *users.Users
		userBeforeMember *user_before_member.UserBeforeMember
		follows          []*follow.Follow
		memberships      []*membership.MembershipData
	)

	if dataID.UserID > 0 {
		user, _ = users.Detail(ctx, dataID.UserID)
		userBeforeMember, _ = user_before_member.Detail(ctx, 0, uint32(user.ID))
		if userBeforeMember.ID > 0 {
			follows, _ = follow.GetFollows(ctx, userBeforeMember.ID)
		}

	} else if dataID.UserBeforeMemberID > 0 {
		userBeforeMember, _ = user_before_member.Detail(ctx, dataID.UserBeforeMemberID, 0)
		user, _ = users.Detail(ctx, userBeforeMember.UserID)
	}
	if user.ID <= 0 && userBeforeMember.ID <= 0 {
		return nil, errors.New("无用户/线索")
	}

	var wg sync.WaitGroup

	if user.ID > 0 {
		wg.Add(1)
		go func(ctx core.StdContext) {
			memberships, _ = membership.GetMembershipDatas(ctx, &membership.MembershipFilter{
				UserID: uint32(user.ID),
			})
			// userData.Store("coupon_alloc_data", s.couponAllocData(userID))
			wg.Done()
		}(ctx)
	}

	// run.Async(ctx, func() {
	// 	userData.Store("signins", s.signins(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("fission_map", s.fissionMaps(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("judge_user_data", s.judgeUserData(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("page_access_data", s.pageAccessData(userID, bs.ID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("page_event_data", s.pageEventData(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("show_video_count", s.showVideoCount(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("all_signin_spend", s.allSigninSpend(userID))
	// 	wg.Done()
	// })
	// run.Async(ctx, func() {
	// 	userData.Store("orders", s.orders(userID))
	// 	wg.Done()
	// })

	// wg.Wait()
	// data = map[string]interface{}{}
	// userData.Range(func(k, v interface{}) bool {
	// 	if k != "user_id" && k != "user_before_member_id" {
	// 		data[k.(string)] = v
	// 	}
	// 	return true
	// })

	data = &users.Data{
		UserBeforeMember: &users.UserBeforeMember{
			ID:             userBeforeMember.ID,
			TransferTime:   userBeforeMember.TransferTime,
			UserID:         userBeforeMember.UserID,
			SourceID:       userBeforeMember.SourceID,
			ManagerUserID:  userBeforeMember.ManagerUserID,
			BelongStudioID: userBeforeMember.BelongStudioID,
			Follows:        follows,
		},
		User: &users.User{
			ID:              uint32(user.ID),
			BelongsStudioID: user.BelongsStudioID,
			ManagerUserID:   user.ManagerUserID,
		},
		Memberships: memberships,
	}
	return
}
