package utils

import (
	"fmt"
	"time"
)

func GetNowFormatTodayTime() string {

	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func GetNowTime() int64 {
	now := time.Now().Unix()
	return now
}
