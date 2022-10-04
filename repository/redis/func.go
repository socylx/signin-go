package redis

import (
	"fmt"
	"signin-go/global/time"
)

/*
获取续卡目标的RedisKey
*/
func GetRenewTargeValueRedisKey(studioID, staffUserID uint32) string {
	return fmt.Sprintf("HistoricalRenewTargeValue_%v_%v", studioID, staffUserID)
}

/*
获取跟进过的学员IDs的RedisKey
*/
func GetConsultantFollowUserIDsRedisKey(startTime, endTime time.Time, studioID, staffUserID uint32) string {
	return fmt.Sprintf(
		"consultantFollowUser_%s_%s_%v_%v",
		time.CSTLayoutString(startTime, time.YYYYMMDD),
		time.CSTLayoutString(endTime, time.YYYYMMDD),
		studioID, staffUserID,
	)
}

/*
获取某一门店/某个顾问的某段时间内续卡率的RedisKey
*/
func GetConsultantRenewRateRedisKey(startTime, endTime time.Time, studioID, staffUserID uint32) string {
	return fmt.Sprintf(
		"ConsultantRenewRate_%s_%s_%v_%v",
		time.CSTLayoutString(startTime, time.YYYYMMDD),
		time.CSTLayoutString(endTime, time.YYYYMMDD),
		studioID, staffUserID,
	)
}

/*
获取某一门店/某个顾问的某段时间内续卡金额的RedisKey
*/
func GetConsultantRenewAmountRedisKey(startTime, endTime time.Time, studioID, staffUserID uint32) string {
	return fmt.Sprintf(
		"ConsultantRenewAmount_%s_%s_%v_%v",
		time.CSTLayoutString(startTime, time.YYYYMMDD),
		time.CSTLayoutString(endTime, time.YYYYMMDD),
		studioID, staffUserID,
	)
}
