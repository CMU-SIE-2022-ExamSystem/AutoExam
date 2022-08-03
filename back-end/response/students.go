package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrStudentNotValidResponse(c *gin.Context, assessment, student string) {
	var temp StudentNotValidError
	err := Error{Type: Assessment, Message: ReplaceMessageAssessmentStudentName(&temp, assessment, student)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ReplaceMessageStudentName(str interface{}, student_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "student_name", student_name)
}

func ReplaceMessageAssessmentStudentName(str interface{}, assessment_name, student_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "assessment_name", assessment_name)
	return strings.ReplaceAll(message, "student_name", student_name)
}

func ReplaceMessageCourseAssessmentStudentName(str interface{}, course_name, assessment_name, student_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "course_name", course_name)
	message = strings.ReplaceAll(message, "assessment_name", assessment_name)
	return strings.ReplaceAll(message, "student_name", student_name)
}

type StudentNotValidError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"There is some wrong for this assessment 'assessment_name' of user 'student_name'. Please confirm your instructors!"`
}
