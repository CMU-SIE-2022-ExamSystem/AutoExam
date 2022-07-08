package course

import (
	"strconv"

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

func GetCourseAssessment(c *gin.Context, validated bool) (string, string) {
	course := GetCourse(c)
	assessment := c.Param("assessment_name")
	status := validate_assessment(c, course, assessment)
	if validated {
		if !status {
			response.ErrAssessmentNotValidResponse(c, course, assessment)
		}
	} else {
		// TODO should validate name is not in current assessments list
		// validate assessment no numeric start
		validate_assessment_name(c, assessment)
	}

	return course, assessment
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

func validate_assessment_name(c *gin.Context, assessment string) {
	firstLetter := assessment[0:1]
	_, err := strconv.Atoi(firstLetter)
	if err == nil {
		response.ErrAssessmentNameNotValidResponse(c, "The name of assessment '"+assessment+"' can't have leading numeral")
	}
}

func get_assessments(c *gin.Context, course string) []models.Assessments {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course+"/assessments")
	// fmt.Println(string(body))

	autolab_resp := utils.Course_assessments_trans(string(body))
	filtered_resp := utils.ExamNameFilter(autolab_resp)
	return filtered_resp
}
