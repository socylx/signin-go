package users

import (
	"gsteps-go/global/time"
	"gsteps-go/internal/core"
	"gsteps-go/internal/errors"
	"gsteps-go/repository/coupon"
	"gsteps-go/repository/coupon_alloc"
	"gsteps-go/repository/fission_map"
	"gsteps-go/repository/follow"
	"gsteps-go/repository/judge_user"
	"gsteps-go/repository/membership"
	"gsteps-go/repository/order"
	"gsteps-go/repository/order_item"
	pageAccessRepo "gsteps-go/repository/page_access"
	"gsteps-go/repository/page_event"
	signinRepo "gsteps-go/repository/signin"
	"gsteps-go/repository/user_before_member"
	"gsteps-go/repository/users"

	pageAccessServ "gsteps-go/service/page_access"
	"gsteps-go/service/show_video"
	signinServ "gsteps-go/service/signin"
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
		couponAllocData  *users.CouponAllocData = &users.CouponAllocData{
			CouponAllocs:            []*coupon_alloc.CouponAllocData{},
			LastNewUserCouponSignin: &signinRepo.SigninData{},
		}
		signins        []*signinRepo.SigninData
		fissionMap     []*fission_map.FissionMapData
		judgeUserData  []*judge_user.JudgeUserData
		pageAccessData *users.PageAccessData = &users.PageAccessData{}
		pageEventData  *users.PageEventData  = &users.PageEventData{
			AccessLocation: &users.AccessLocation{},
		}
		showVideoCount int64
		allSigninSpend float64
		orders         []*users.Order
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
		wg.Add(10)
		go func() {
			defer wg.Done()
			memberships, _ = membership.GetMembershipDatas(ctx, &membership.MembershipFilter{
				UserID: uint32(user.ID),
			})
		}()
		go func() {
			defer wg.Done()
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
						Remains:    couponAlloc.Remain,
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
			var signinData *signinRepo.SigninData = &signinRepo.SigninData{}
			if couponAllocID > 0 {
				signinData, _ = signinRepo.GetSigninDataByCouponAllocID(ctx, couponAllocID)
			}
			couponAllocData = &users.CouponAllocData{
				CouponAllocs:            couponAllocDatas,
				LastNewUserCouponSignin: signinData,
			}
		}()
		go func() {
			defer wg.Done()
			todayDate := time.TodayDate()
			signins, _ = signinRepo.GetSigninDatas(
				ctx, &signinRepo.Filter{
					UserID:              uint32(user.ID),
					ActivityStartTimeGE: todayDate.AddDate(0, 0, -90),
					ActivityStartTimeLT: todayDate.AddDate(0, 0, 7),
				})
		}()
		go func() {
			defer wg.Done()
			fissionMap, _ = fission_map.GetFissionMapData(
				ctx, &fission_map.Filter{
					ShareUserID: uint32(user.ID),
					Type:        fission_map.Member,
					StatusIn:    fission_map.STATUS_AFTER_ACCEPT,
				},
			)
		}()
		go func() {
			defer wg.Done()
			judgeUserData, _ = judge_user.GetJudgeUserDatas(ctx, uint32(user.ID))
		}()
		go func() {
			defer wg.Done()
			var belongsStudioID uint32
			if user.BelongsStudioID > 0 {
				belongsStudioID = user.BelongsStudioID
			} else if userBeforeMember.BelongStudioID > 0 {
				belongsStudioID = userBeforeMember.BelongStudioID
			}
			pageAccessData, _ = pageAccessServ.GetPageAccessData(ctx, uint32(user.ID), belongsStudioID)
		}()
		go func() {
			defer wg.Done()
			lastPageEvent, _ := page_event.GetLastPageEvent(ctx, &page_event.Filter{
				UserID:   uint32(user.ID),
				Type:     pageAccessRepo.MiniProgram,
				EventKey: page_event.Location,
			})
			pageEventData = &users.PageEventData{
				AccessLocation: &users.AccessLocation{
					Longitude: lastPageEvent.Data1,
					Latitude:  lastPageEvent.Data2,
				},
			}
		}()
		go func() {
			defer wg.Done()
			showVideoCount = show_video.GetShowVideoCount(ctx, uint32(user.ID))
		}()
		go func() {
			defer wg.Done()
			allSigninSpend = signinServ.GetAllSigninSpend(ctx, uint32(user.ID))
		}()
		go func() {
			defer wg.Done()
			orderDatas, _ := order.GetOrderDatas(ctx, uint32(user.ID))
			for _, orderData := range orderDatas {
				orderItems, _ := order_item.GetOrderIemDatas(ctx, orderData.ID)
				order := &users.Order{}
				order.ID = orderData.ID
				order.Status = orderData.Status
				order.OrderItems = orderItems
				orders = append(orders, order)
			}
		}()
	}
	wg.Wait()

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
		Signins:         signins,
		FissionMap:      fissionMap,
		JudgeUserData:   judgeUserData,
		PageAccessData:  pageAccessData,
		PageEventData:   pageEventData,
		ShowVideoCount:  showVideoCount,
		AllSigninSpend:  allSigninSpend,
		Orders:          orders,
	}
	return
}
