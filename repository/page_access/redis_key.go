package page_access

import (
	"fmt"
	"signin-go/global/time"
)

func GetPageAccessCountRedisKey(t time.Time, userID uint32) string {
	return fmt.Sprintf("pageAccessCount_%s_%v", time.CSTLayoutString(t, time.YYYYMMDD), userID)
}

func GetAccessBuyCardCountRedisKey(t time.Time, userID uint32) string {
	return fmt.Sprintf("accessBuyCardCount_%s_%v", time.CSTLayoutString(t, time.YYYYMMDD), userID)
}

func GetCurrentStudioAccessActivityCountRedisKey(t time.Time, userID uint32) string {
	return fmt.Sprintf("currentStudioAccessActivityCount_%s_%v", time.CSTLayoutString(t, time.YYYYMMDD), userID)
}

func GetLastPageAccessTimeRedisKey(userID uint32) string {
	return fmt.Sprintf("lastPageAccessTime_%v", userID)
}
