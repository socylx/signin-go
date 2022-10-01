package time

import (
	"log"
	"time"
)

func Init() {
	log.Println("global.time.Init Start...")
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		log.Fatalf("global.time.Init Error: %v", err)
	}

	// 默认设置为中国时区
	time.Local = cst
}
