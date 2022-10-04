package strategy

import (
	"crypto/md5"
	"fmt"
	"signin-go/global/time"
)

const strategyKeyBaseString = "30ec877eaf21e960b504398cc7f48efc"

func GenerateStrategyKey() string {
	data := []byte(fmt.Sprintf("%s %s", strategyKeyBaseString, time.CSTLayoutString(time.Now(), time.CSTLayout)))
	sumStr := fmt.Sprintf("%x", md5.Sum(data))
	return sumStr
}
