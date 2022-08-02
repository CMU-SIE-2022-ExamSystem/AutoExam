package course

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func CreateAnswerFolder(path string) bool {
	if _, err := os.Stat(path + "answer/"); err == nil {
		color.Yellow("answer folder already exists!")
	} else if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path+"answer/", 0777)
		if err != nil {
			return false
		}
		color.Yellow("create answer folder for course!")
	}
	return true
}

func PrepareSolution(student dao.Assessment_Student, path string) error {
	target := path + "answer/solution.json"
	os.Remove(target)
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if student.Solutions == nil {
		return errors.New("the solution can not be found")
	}
	data, _ := json.Marshal(student.Solutions)

	_, err = fmt.Fprint(file, string(data))
	return err
}

func PrepareConfig(student dao.Assessment_Student, path string) error {
	target := path + "answer/config.yaml"
	os.Remove(target)
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if student.Problems == nil {
		return errors.New("the config can not be found")
	}
	data, _ := yaml.Marshal(student.Problems)

	_, err = fmt.Fprint(file, string(data))
	return err
}

func PrepareAnswer(student dao.Assessment_Student, path string) error {
	target := path + "answer/answer.json"
	os.Remove(target)
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if student.Answers == nil {
		return errors.New("the answer can not be found")
	}
	data, _ := json.Marshal(student.Answers)

	_, err = fmt.Fprint(file, string(data))
	return err
}

func CheckSubmission(c *gin.Context) bool {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name, assessment_name := GetCourseAssessment(c)
	assessment, _ := dao.ReadExam(course_name, assessment_name)
	max := assessment.General.MaxSubmissions
	if max == -1 {
		return true
	}

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submissions")

	autolab_resp := utils.Assessments_submissions_trans(string(body))
	if len(autolab_resp) != 0 {
		return autolab_resp[0].Version < max
	} else {
		return true
	}
}
