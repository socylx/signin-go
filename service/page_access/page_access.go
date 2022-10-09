package page_access

import (
	"gsteps-go/global/redis"
	"gsteps-go/global/time"
	"gsteps-go/internal/core"
	"gsteps-go/repository/activity"
	"gsteps-go/repository/page_access"
	"gsteps-go/repository/users"
	"log"

	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

/*
获取用户近三十天访问小程序的数据
*/
func GetPageAccessData(ctx core.StdContext, userID uint32, belongStudioID uint32) (data *users.PageAccessData, err error) {
	var (
		pageAccessCount, currentStudioAccessActivityCount, accessBuyCardCount uint64
		lastPageAccessTime                                                    time.Time

		wg   sync.WaitGroup
		days int

		todayDate = time.TodayDate()
		startTime = todayDate.AddDate(0, 0, -30)
		endTime   = todayDate.AddDate(0, 0, 1)
	)
	data = &users.PageAccessData{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		lastPageAccessTimeRedisKey := page_access.GetLastPageAccessTimeRedisKey(userID)
		lastPageAccessTimeRedisData, _ := redis.Redis.Get(ctx, lastPageAccessTimeRedisKey).Result()
		lastPageAccessTime, err = time.ParseCSTInLocation(lastPageAccessTimeRedisData, time.CSTLayout)
		if err == nil {
			return
		}
		lastPageAccessTime, _ = page_access.GetLastPageAccessTime(ctx, &page_access.Filter{UserID: userID})
	}()

	for {
		queryStartTime := startTime.AddDate(0, 0, days)
		queryEndTime := queryStartTime.AddDate(0, 0, 1)

		wg.Add(1)
		go func(sTime, eTime time.Time) {
			defer func() {
				wg.Done()
			}()
			pageAccessCountRedisKey := page_access.GetPageAccessCountRedisKey(sTime, userID)
			accessBuyCardCountRedisKey := page_access.GetAccessBuyCardCountRedisKey(sTime, userID)
			currentStudioAccessActivityCountRedisKey := page_access.GetCurrentStudioAccessActivityCountRedisKey(sTime, userID)

			pageAccessCountRedisData, err1 := redis.Redis.Get(ctx, pageAccessCountRedisKey).Result()
			accessBuyCardCountRedisData, err2 := redis.Redis.Get(ctx, accessBuyCardCountRedisKey).Result()
			currentStudioAccessActivityCountRedisData, err3 := redis.Redis.Get(ctx, currentStudioAccessActivityCountRedisKey).Result()
			if err1 == nil && err2 == nil && err3 == nil {
				pac, pacerr := strconv.ParseUint(pageAccessCountRedisData, 10, 64)
				abcc, abccerr := strconv.ParseUint(accessBuyCardCountRedisData, 10, 64)
				csaac, csaacerr := strconv.ParseUint(currentStudioAccessActivityCountRedisData, 10, 64)
				if pacerr == nil && abccerr == nil && csaacerr == nil {
					atomic.AddUint64(&pageAccessCount, pac)
					atomic.AddUint64(&accessBuyCardCount, abcc)
					atomic.AddUint64(&currentStudioAccessActivityCount, csaac)
					return
				}
			}

			pageAccess, err := page_access.GetPageAccess(ctx, &page_access.Filter{
				UserID:       userID,
				CreateTimeGE: sTime,
				CreateTimeLT: eTime,
				Type:         page_access.MiniProgram,
			})
			if err != nil {
				log.Println("pageAccessData: ", err)
				return
			}

			sessionIDM := map[string]bool{}
			buyCardSessionID := map[string]bool{}

			activityIDs := []uint32{}
			activityAccessSessions := map[uint32][]string{}

			for _, pa := range pageAccess {
				sessionIDM[pa.SessionID] = true
				if pa.URL == page_access.CourseDetailIndex && strings.HasPrefix(pa.Params, "id=") {
					activityID, err := strconv.Atoi(strings.Split(pa.Params, "=")[1])
					if err != nil {
						continue
					}
					activityIDs = append(activityIDs, uint32(activityID))
					activityAccessSessions[uint32(activityID)] = append(activityAccessSessions[uint32(activityID)], pa.SessionID)
				}
				if pa.URL == page_access.BuyCardIndex {
					buyCardSessionID[pa.SessionID] = true
				}
			}
			var duration time.Duration
			if queryStartTime.Before(todayDate) {
				duration = 1 * 7 * 24 * time.Hour
			} else {
				duration = 20 * time.Minute
			}
			redis.Redis.Set(ctx, pageAccessCountRedisKey, uint64(len(sessionIDM)), duration)
			redis.Redis.Set(ctx, accessBuyCardCountRedisKey, uint64(len(buyCardSessionID)), duration)
			atomic.AddUint64(&pageAccessCount, uint64(len(sessionIDM)))
			atomic.AddUint64(&accessBuyCardCount, uint64(len(buyCardSessionID)))

			if belongStudioID > 0 && len(activityIDs) > 0 {
				activitySessionID := map[string]bool{}
				currentStudioActivityIDs, _ := activity.GetActivityIDs(ctx, &activity.Filter{
					IncludeIDs: activityIDs,
					StudioID:   belongStudioID,
				})
				for _, currentStudioActivityID := range currentStudioActivityIDs {
					for _, activityAccessSession := range activityAccessSessions[currentStudioActivityID] {
						activitySessionID[activityAccessSession] = true
					}
				}
				redis.Redis.Set(ctx, currentStudioAccessActivityCountRedisKey, uint64(len(activitySessionID)), duration)
				atomic.AddUint64(&currentStudioAccessActivityCount, uint64(len(activitySessionID)))
			}
		}(queryStartTime, queryEndTime)
		if queryEndTime.Equal(endTime) || queryEndTime.After(endTime) {
			break
		}
		days += 1
	}
	wg.Wait()
	data.PageAccessCount = pageAccessCount
	data.LastPageAccessTime = lastPageAccessTime
	data.CurrentStudioAccessActivityCount = currentStudioAccessActivityCount
	data.AccessBuyCardCount = accessBuyCardCount
	// data = &users.PageAccessData{
	// 	PageAccessCount:                  pageAccessCount,
	// 	LastPageAccessTime:               lastPageAccessTime,
	// 	CurrentStudioAccessActivityCount: currentStudioAccessActivityCount,
	// 	AccessBuyCardCount:               accessBuyCardCount,
	// }
	return
}
