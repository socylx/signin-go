package show_video

import "fmt"

func GetShowVideoCountRedisKey(userID uint32) string {
	return fmt.Sprintf("showVideoCount_%v", userID)
}
