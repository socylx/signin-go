package signin

import "fmt"

func GetAllSigninSpendRedisKey(userID uint32) string {
	return fmt.Sprintf("allSigninSpend_%v", userID)
}
