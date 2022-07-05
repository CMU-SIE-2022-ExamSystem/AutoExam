package utils

import (
	"fmt"
	"os"
	"time"
)

func GetNowTime() int64 {
	now := time.Now().Unix()
	return now
}

func GetNowFormatTodayTime() string {

	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			panic(err)
		}
	}
}

func FileCheck(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}
}
