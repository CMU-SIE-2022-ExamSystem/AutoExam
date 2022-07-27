package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrCourseNoBaseCourseResponse(c *gin.Context, course string) {
	message := ReplaceMessageCourseName(&CourseNoBaseCourseError{}, course)
	err := Error{Type: Course, Message: message}
	ErrorResponseWithStatus(c, err, http.StatusForbidden)
}

func ReplaceMessageCourseName(str interface{}, course string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "course_name", course)
}

type CourseNoBaseCourseError struct {
	Type    string `json:"type" example:"Course"`
	Message string `json:"message" example:"The course 'course_name' does not belong to any base course! Should use other api to set up the relationship between course and base course"`
}

type CourseNotValidError struct {
	Type    string `json:"type" example:"Course"`
	Message string `json:"message" example:"This user is not registered in such course 'course_name'"`
}
