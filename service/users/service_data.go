package users

import (
	"signin-go/internal/core"
	"signin-go/internal/errors"
	"signin-go/repository/coupon"
	"signin-go/repository/coupon_alloc"
	"signin-go/repository/follow"
	"signin-go/repository/membership"
	"signin-go/repository/signin"
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
		couponAllocData  *users.CouponAllocData
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
		wg.Add(2)
		go func() {
			memberships, _ = membership.GetMembershipDatas(ctx, &membership.MembershipFilter{
				UserID: uint32(user.ID),
			})
			wg.Done()
		}()

		go func() {
			couponAllocs, _ := coupon_alloc.GetCouponAllocs(
				ctx,
				&coupon_alloc.Filter{
					Status: []coupon_alloc.CouponAllocStatus{coupon_alloc.Init, coupon_alloc.Used},
					UserID: uint32(user.ID),
				},
			)
			var couponAllocID uint32
			couponAllocDatas := make([]*coupon_alloc.CouponAllocData, 0, len(couponAllocs))
			for _, couponAlloc := range couponAllocs {
				coupon, _ := coupon.Detail(ctx, couponAlloc.CouponID)
				couponAllocDatas = append(
					couponAllocDatas, &coupon_alloc.CouponAllocData{
						ID:         couponAlloc.ID,
						CreateTime: couponAlloc.CreateTime,
						Deadline:   couponAlloc.Deadline,
						Remain:     couponAlloc.Remain,
						GetType:    couponAlloc.GetType,
						Coupon: &coupon_alloc.Coupon{
							ID:         coupon.ID,
							Type:       coupon.Type,
							AmountType: coupon.Type,
							IsNewUser:  coupon.IsNewUser,
							IsDel:      coupon.IsDel,
						},
					},
				)
				if coupon.IsNewUser {
					couponAllocID = couponAlloc.ID
				}
			}
			var signinData *signin.SigninData
			if couponAllocID > 0 {
				signinData, _ = signin.GetSigninDataByCouponAllocID(ctx, couponAllocID)
			}
			couponAllocData = &users.CouponAllocData{
				CouponAllocs:            couponAllocDatas,
				LastNewUserCouponSignin: signinData,
			}
			wg.Done()
		}()
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
		Memberships:     memberships,
		CouponAllocData: couponAllocData,
	}
	return
}
