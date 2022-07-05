package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
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

func Copy_file(file, src, dest string) {
	// copy certain file from src folder to dest folder

	src = filepath.Join(src, file)
	dest = filepath.Join(dest, file)
	if _, err := os.Stat(src); os.IsNotExist(err) {
		if err != nil {
			panic(err)
		}
	}
	cp.Copy(src, dest)
}

func FileCheck(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}
}
