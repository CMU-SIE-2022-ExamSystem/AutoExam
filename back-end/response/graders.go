package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Grader = "Grader"
)

func ErrGraderReadFileResponse(c *gin.Context, err any) {
	err = ErrorJson{Type: Grader, Message: err}
	ErrorResponseWithStatus(c, err, http.StatusInternalServerError)
}

func ErrGraderTestResponse(c *gin.Context, err any) {
	err = ErrorJson{Type: Grader, Message: err}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrGraderNotValidResponse(c *gin.Context, course, grader string) {
	var temp GraderNotValidError
	err := Error{Type: Grader, Message: ReplaceMessageCourseGraderName(&temp, course, grader)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrGraderDeleteNotSafeResponse(c *gin.Context, grader_name string) {
	err := Error{Type: Grader, Message: ReplaceMessageGraderName(&GraderDeleteNotSafeError{}, grader_name)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrGraderUpdateNotSafeResponse(c *gin.Context, grader_name string) {
	err := Error{Type: Grader, Message: ReplaceMessageGraderName(&GraderUpdateNotSafeError{}, grader_name)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrGraderNoUploadResponse(c *gin.Context, grader_name string) {
	err := Error{Type: Grader, Message: ReplaceMessageGraderName(&GraderNoUploadError{}, grader_name)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ReplaceMessageGraderName(str interface{}, grader_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "grader_name", grader_name)
}

func ReplaceMessageCourseGraderName(str interface{}, course_name, grader_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "course_name", course_name)
	return strings.ReplaceAll(message, "grader_name", grader_name)
}

type GraderNotValidError struct {
	Type    string `json:"type" example:"Grader"`
	Message string `json:"message" example:"There is no this grader 'grader_name' in such base course 'course_name'"`
}

type GraderUpdateNotSafeError struct {
	Type    string `json:"type" example:"Grader"`
	Message string `json:"message" example:"This grader name 'grader_name' is already valid. It would be dangerous to upload a new grader"`
}

type GraderDeleteNotSafeError struct {
	Type    string `json:"type" example:"Grader"`
	Message string `json:"message" example:"This grader name 'grader_name' is used in some questions. Therefore, it cannot be deleted!"`
}

type GraderNoUploadError struct {
	Type    string `json:"type" example:"Grader"`
	Message string `json:"message" example:"There is no file uploaded to this grader 'grader_name'"`
}
