package strategy

import (
	"signin-go/common/types"
	redisGlob "signin-go/global/redis"
	"signin-go/global/time"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/coupon"
	redisRepo "signin-go/repository/redis"
	"signin-go/repository/strategy"

	"signin-go/service/follow"
	"signin-go/service/users"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type recommendRequest struct {
	StudioID         uint32  `form:"studio_id" binding:"required"`
	Type             uint32  `form:"type" binding:"required"`
	Count            uint64  `form:"count" binding:"required"`
	StaffUserID      uint32  `form:"staff_user_id"`
	MinRemain        float32 `form:"min_remain"`
	MaxRemain        float32 `form:"max_remain"`
	MinActivityCount int32   `form:"min_activity_count"`
	MaxActivityCount int32   `form:"max_activity_count"`
}

func recommend(c core.Context) {
	request := new(recommendRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.StudioID <= 0 {
		c.AbortWithError(core.Error(code.ParamBindError, "无门店ID参数"))
		return
	}

	valid := strategy.StrategyTypeCheck[request.Type]
	if !valid {
		c.AbortWithError(core.Error(code.ParamBindError, "策略类型不支持"))
		return
	}

	redisKey := redisRepo.GetStrategyRecommendIDsRedisKey(request.StudioID, request.Type)
	redisDatas, err := redisGlob.Redis.ZRevRange(c.RequestContext(), redisKey, 0, -1).Result()
	if err != nil {
		c.AbortWithError(core.Error(
			code.CacheGetError,
			code.Text(code.CacheGetError)).WithError(err),
		)
		return
	}
	request.Count = 1000

	var (
		today             = time.TodayDate()
		thisWeekStartDate = today.AddDate(0, 0, -int(today.Weekday()-1))
		thisWeekEndDate   = thisWeekStartDate.AddDate(0, 0, 7)

		strategyTypeXuka bool = request.Type == strategy.Xuka

		followUserStatus types.FollowUserStatus

		count               uint64
		userIDs             = []uint64{}
		userBeforeMemberIDs = []uint64{}
		wg                  sync.WaitGroup

		appendUserIDsMutex             sync.Mutex
		appendUserBeforeMemberIDsMutex sync.Mutex

		appendUserIDs = func(userID uint64) {
			appendUserIDsMutex.Lock()
			defer func() {
				appendUserIDsMutex.Unlock()
			}()
			userIDs = append(userIDs, userID)
		}
		appendUserBeforeMemberIDs = func(userBeforeMemberID uint64) {
			appendUserBeforeMemberIDsMutex.Lock()
			defer func() {
				appendUserBeforeMemberIDsMutex.Unlock()
			}()
			userBeforeMemberIDs = append(userBeforeMemberIDs, userBeforeMemberID)
		}

		requestRemainCond        bool = request.MinRemain >= 0 && request.MaxRemain >= 0 && request.MinRemain < request.MaxRemain
		requestActivityCountCond bool = request.MinActivityCount >= 0 && request.MaxActivityCount >= 0 && request.MinActivityCount < request.MaxActivityCount
	)
	if strategyTypeXuka {
		followUserStatus, _ = follow.GetConsultantFollowUserStatus(c.RequestContext(), thisWeekStartDate, thisWeekEndDate, request.StudioID, 0)
	}

	for _, redisData := range redisDatas {
		if count >= request.Count {
			break
		}

		IDs := strings.Split(redisData, "_")
		if len(IDs) != 2 {
			continue
		}

		userID, err1 := strconv.ParseUint(IDs[0], 10, 64)
		userBeforeMemberID, err2 := strconv.ParseUint(IDs[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		wg.Add(1)
		go func(uID, uBeforeID uint64) {
			defer wg.Done()

			var (
				remainCond   bool = true
				activityCond bool = true
				managerCond  bool = true
			)
			if requestRemainCond || requestActivityCountCond || request.StaffUserID > 0 {
				userData, err := users.Data(c.RequestContext(), &users.DataID{
					UserID:             uint32(uID),
					UserBeforeMemberID: uint32(uBeforeID),
				})
				if err != nil {
					return
				}

				if requestRemainCond {
					var remains float32
					for _, membership := range userData.Memberships {
						remains += membership.Remains
					}
					todayDate := time.TodayDate()
					for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
						deadline := couponAlloc.Deadline
						if deadline != time.TimeZeroTime || deadline.Before(todayDate) {
							continue
						}
						if couponAlloc.Coupon.Type == coupon.CashType && couponAlloc.Coupon.AmountType == coupon.NumberAmountType {
							remains += couponAlloc.Remains
						}
					}
					remainCond = request.MinRemain <= remains && remains <= request.MaxRemain
				}

				if requestActivityCountCond {
					var activityCount int32
					todayDate := time.TodayDate()
					startDate := todayDate.AddDate(0, 0, -30)
					endDate := todayDate.AddDate(0, 0, 1)
					for _, signin := range userData.Signins {
						activityStartTime := signin.ActivityStartTime
						if activityStartTime.After(startDate) && activityStartTime.Before(endDate) {
							activityCount += 1
						}
					}
					activityCond = request.MinActivityCount <= activityCount && activityCount <= request.MaxActivityCount
				}

				if request.StaffUserID > 0 {
					if strategyTypeXuka {
						managerCond = userData.User.ManagerUserID == request.StaffUserID
					} else {
						managerCond = userData.UserBeforeMember.ManagerUserID == request.StaffUserID
					}
				}
			}

			if uID > 0 {
				if strategyTypeXuka && followUserStatus[uID] {
					return
				}
				if remainCond && activityCond && managerCond {
					appendUserIDs(uID)
				}
				atomic.AddUint64(&count, 1)
			} else if uBeforeID > 0 {
				if remainCond && activityCond && managerCond {
					appendUserBeforeMemberIDs(uBeforeID)
				}
				atomic.AddUint64(&count, 1)
			}
		}(userID, userBeforeMemberID)
	}
	wg.Wait()
	response := &struct {
		UserIDs             []uint64 `json:"user_ids"`
		UserBeforeMemberIDs []uint64 `json:"user_before_member_ids"`
	}{
		UserIDs:             userIDs,
		UserBeforeMemberIDs: userBeforeMemberIDs,
	}
	c.Payload(response)
}
