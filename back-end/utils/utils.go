package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
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

func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func Copy_file(file, src, dest string) {
	// copy certain file from src folder to dest folder
	fmt.Println(file, src, dest)
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

func FileCheckWithC(c *gin.Context, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		response.ErrFileNotValidResponse(c)
	}
}

func Ordinal(num int) string {

	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}

	// math.Abs() is to convert negative number to positive
	floatNum := math.Abs(float64(num))
	positiveNum := int(floatNum)

	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return "th"
	}

	return ordinalDictionary[positiveNum]

}

func Ordinalize(num int) string {
	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}
	// math.Abs() is to convert negative number to positive
	floatNum := math.Abs(float64(num))
	positiveNum := int(floatNum)

	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return strconv.Itoa(num) + "th"
	}

	return strconv.Itoa(num) + ordinalDictionary[positiveNum]
}

func MakeAnswertar(path string) bool {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("tar", "cvf", "answer.tar", "answer")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		color.Yellow(err.Error() + stdout.String() + stderr.String())
		return false
	} else {
		return true
	}
}

func CheckModule() {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("make")
	cmd.Dir = "./autograder/"
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
}

func writejson(test models.GraderTest, path string, name string) error {
	target_path := path + name + ".json"
	file, err := os.OpenFile(target_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	switch name {
	case "answer":
		if test.Answers.TestAutograder == nil {
			return errors.New("the answer can not be found")
		}
		data, _ := json.Marshal(test.Answers)

		_, err = fmt.Fprint(file, string(data))
		return err
	case "solution":
		if test.Solutions.TestAutograder == nil {
			return errors.New("the solution can not be found")
		}
		data, _ := json.Marshal(test.Solutions)

		_, err = fmt.Fprint(file, string(data))
		return err
	}
	return nil
}

func WriteGraderTest(test models.GraderTest, path string) error {
	err := writejson(test, path, "answer")
	if err != nil {
		return err
	}
	err = writejson(test, path, "solution")
	if err != nil {
		return err
	}
	return nil
}
