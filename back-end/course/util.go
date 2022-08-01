package course

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func GetCourse(c *gin.Context) string {
	course := c.Param("course_name")
	validate_course(c, course)

	return course
}

func GetCourseAssessment(c *gin.Context) (string, string) {
	course := GetCourse(c)
	assessment := c.Param("assessment_name")

	// assessment not in mongodb
	if !Validate_autoexam_assessment(c, course, assessment) {
		response.ErrAssessmentNotValidResponse(c, course, assessment)
	}
	return course, assessment
}

func GetFilteredAssessments(c *gin.Context) []models.Assessments {
	course := GetCourse(c)
	assessments := get_assessments(c, course)
	var autolab_index []int
	auth := jwt.Get_authlevel_DB(c)
	for index, assessment := range assessments {
		if assessment.Autolab && !assessment.AutoExam {
			autolab_index = append(autolab_index, index)
		} else if auth == "student" && (!assessment.Autolab || assessment.Draft) { // student cannot see draft assessment
			autolab_index = append(autolab_index, index)
		}
	}

	for i := range autolab_index {
		index := autolab_index[len(autolab_index)-1-i]
		assessments = removeIndex(assessments, index)
	}
	return assessments
}

func removeIndex(s []models.Assessments, index int) []models.Assessments {
	return append(s[:index], s[index+1:]...)
}

func validate_course(c *gin.Context, course string) {
	user := jwt.GetEmail(c)
	courses := dao.Get_all_courses(user.ID)
	if !slices.Contains(courses, course) {
		response.ErrCourseNotValidResponse(c, course)
	}
}

func validate_assessment(c *gin.Context, course, assessment string) bool {
	filtered_resp := get_assessments(c, course)
	for _, resp := range filtered_resp {
		if assessment == resp.Name {
			return true
		}
	}
	return false
}

func Validate_autoexam_assessment(c *gin.Context, course, assessment string) bool {
	assessments, err := dao.GetAllExams(course)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when reading all assessments from mongodb")
	}
	for _, temp := range assessments {
		if assessment == temp.Name {
			return true
		}
	}
	return false
}

func Validate_autolab_assessment(c *gin.Context, course, assessment string) bool {
	assessments := get_autolab_assessments(c, course)
	return check_in_autolab_assessments(assessments, assessment) != -1
}

func Validate_assessment_name(c *gin.Context, course, assessment string) {
	if validate_assessment(c, course, assessment) {
		response.ErrAssessmentNameNotValidResponse(c, http.StatusConflict, "The name of assessment '"+assessment+"' is already in this course '"+course+"'")
	}

	firstLetter := assessment[0:1]
	_, err := strconv.Atoi(firstLetter)
	if err == nil {
		response.ErrAssessmentNameNotValidResponse(c, http.StatusBadRequest, "The name of assessment '"+assessment+"' can't have leading numeral")
	}
}

func get_assessments(c *gin.Context, course string) []models.Assessments {
	// get autolab's assessments
	autolab_assessments := get_autolab_assessments(c, course)

	// get mongodb's assessments
	autoexam_assessments, err := dao.GetAllExams(course)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when reading all assessments from mongodb")
	}

	// modify assessments' start_at , end_at and due_at
	length := len(autoexam_assessments)
	for i := 0; i < length; i++ {
		autoexam := autoexam_assessments[i]
		if index := check_in_autolab_assessments(autolab_assessments, autoexam.Name); index != -1 {
			autolab_assessments[index].Start_at = autoexam.Start_at
			autolab_assessments[index].Due_at = autoexam.End_at
			autolab_assessments[index].End_at = autoexam.End_at
			autolab_assessments[index].AutoExam = true
			autolab_assessments[index].Draft = autoexam.Draft
		} else {
			autolab_assessments = append(autolab_assessments, autoexam)
		}
	}
	return autolab_assessments
}

func check_in_autolab_assessments(autolab []models.Assessments, name string) int {
	for i, assessment := range autolab {
		if name == assessment.Name {
			return i
		}
	}
	return -1
}

func get_autolab_assessments(c *gin.Context, course string) []models.Assessments {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course+"/assessments")
	// fmt.Println(string(body))

	autolab_resp := utils.Course_assessments_trans(string(body))
	filtered_resp := utils.ExamNameFilter(autolab_resp)
	var temp []models.Assessments
	for _, resp := range filtered_resp {
		temp = append(temp, resp.ToAssessments())
	}
	return temp
}

func CourseUserData(c *gin.Context) ([]models.Course_User_Data, error) {
	var users []models.Course_User_Data
	jwt.Check_authlevel_Instructor(c)

	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := GetCourse(c)

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/course_user_data")
	// fmt.Println(string(body))

	if strings.Contains(string(body), "error") {
		return users, errors.New(string(body))
	} else {
		users := utils.Course_user_trans(string(body))
		return users, nil
	}
}
