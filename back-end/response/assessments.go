package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Assessment = "Assessment"
)

func ErrAssessmentNotValidResponse(c *gin.Context, course, assessment string) {
	var temp AssessmentNotValidError
	err := Error{Type: Assessment, Message: ReplaceMessageCourseAssessmentName(&temp, course, assessment)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrAssessmentModifyNotSafeResponse(c *gin.Context, assessment_name string) {
	err := Error{Type: Assessment, Message: ReplaceMessageAssessmentName(&AssessmentModifyNotSafeError{}, assessment_name)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrAssessmentNotInAutolabResponse(c *gin.Context, course, assessment string) {
	err := Error{Type: Assessment, Message: ReplaceMessageCourseAssessmentName(&AssessmentNotInAutolabError{}, course, assessment)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrAssessmentNameNotValidResponse(c *gin.Context, status int, message string) {
	err := Error{Type: Assessment, Message: message}
	ErrorResponseWithStatus(c, err, status)
}

func ErrAssessmentNoSettingsResponse(c *gin.Context, assessment string) {
	err := Error{Type: Assessment, Message: ReplaceMessageAssessmentName(&AssessmentNoSettingsError{}, assessment)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrAssessmentBeforeStartAtResponse(c *gin.Context, assessment string) {
	err := Error{Type: Assessment, Message: ReplaceMessageAssessmentName(&AssessmentBeforeStartAtError{}, assessment)}
	ErrorResponseWithStatus(c, err, http.StatusForbidden)
}

func ErrAssessmentInternaldResponse(c *gin.Context, message string) {
	err := Error{Type: Assessment, Message: message}
	ErrorInternalResponse(c, err)
}

func ErrAssessmentGenerateResponse(c *gin.Context, course, assessment string) {
	var temp AssessmentNotValidError
	err := Error{Type: Assessment, Message: ReplaceMessageCourseAssessmentName(&temp, course, assessment)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ReplaceMessageAssessmentName(str interface{}, assessment_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "assessment_name", assessment_name)
}

func ReplaceMessageAssessmentNameErrorMessage(str interface{}, assessment_name, error_message string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "assessment_name", assessment_name)
	return strings.ReplaceAll(message, "error_message", error_message)
}

func ReplaceMessageCourseAssessmentName(str interface{}, course_name, assessment_name string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "course_name", course_name)
	return strings.ReplaceAll(message, "assessment_name", assessment_name)
}

type AssessmentNotValidError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"There is no this assessment 'assessment_name' in course 'course_name'"`
}

type AssessmentModifyNotSafeError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"This assessment name 'assessment_name' is already uploaded to autolab. It would be dangerous to modify this assessment"`
}

type AssessmentNotInAutolabError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"There is no this assessment 'assessment_name' in such course 'course_name' on autolab, please download this assessment and uploaded the tar file to the specific course 'course_name' on autolab"`
}

type AssessmentNoSettingsError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"This assessment name 'assessment_name' does not have any settings for the configuration of the exam or quiz"`
}

type AssessmentBeforeStartAtError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"It's too early to take this assessment name 'assessment_name' now. Please try again later."`
}

type AssessmentGenerateError struct {
	Type    string `json:"type" example:"Assessment"`
	Message string `json:"message" example:"There are some errors when generating assessment name 'assessment_name' with error message 'error_message'"`
}
