package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

func CheckModule(path string) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("make")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
}

func RemoveGraderTestFile(path string) {
	os.Remove(path + "requirements.txt")
	os.Remove(path + "answer.json")
	os.Remove(path + "solution.json")

	dir, _ := ioutil.ReadDir(path + "autograders")
	for _, d := range dir {
		if strings.Contains(d.Name(), ".py") {
			if d.Name() != "__init__.py" {
				os.Remove(path + "autograders/" + d.Name())
			}
		}
	}
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
