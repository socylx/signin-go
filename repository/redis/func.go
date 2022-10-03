package redis

import (
	"fmt"
	"signin-go/global/time"
)

func GetRenewTargeValueRedisKey(studioID, staffUserID uint32) string {
	return fmt.Sprintf("HistoricalRenewTargeValue_%v_%v", studioID, staffUserID)
}

func GetConsultantFollowUserIDsRedisKey(startTime, endTime time.Time, studioID, staffUserID uint32) string {
	return fmt.Sprintf(
		"consultantFollowUser_%s_%s_%v_%v",
		time.CSTLayoutString(startTime, time.YYYYMMDD),
		time.CSTLayoutString(endTime, time.YYYYMMDD),
		studioID, staffUserID,
	)
}
