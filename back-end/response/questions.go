package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Question = "Question"
)

func ErrQuestionNotValidResponse(c *gin.Context, course, question string) {
	var temp QuestionNotValidError
	err := Error{Type: Question, Message: ReplaceMessageCourseQuestionName(&temp, course, question)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrQuestionModifyNotSafeResponse(c *gin.Context, question_id string) {
	err := Error{Type: Question, Message: ReplaceMessageQuestionName(&QuestionModifyNotSafeError{}, question_id)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrQuestionDeleteNotSafeResponse(c *gin.Context, question_id string) {
	err := Error{Type: Question, Message: ReplaceMessageQuestionName(&QuestionDeleteNotSafeError{}, question_id)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrQuestionNoUploadResponse(c *gin.Context, question_id string) {
	err := Error{Type: Question, Message: ReplaceMessageQuestionName(&QuestionNoUploadError{}, question_id)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ReplaceMessageQuestionName(str interface{}, question_id string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "question_id", question_id)
}

func ReplaceMessageCourseQuestionName(str interface{}, course_name, question_id string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "course_name", course_name)
	return strings.ReplaceAll(message, "question_id", question_id)
}

type QuestionNotValidError struct {
	Type    string `json:"type" example:"Question"`
	Message string `json:"message" example:"There is no this question id 'question_id' in such base course 'course_name'"`
}

type QuestionModifyNotSafeError struct {
	Type    string `json:"type" example:"Question"`
	Message string `json:"message" example:"This question id 'question_id' is already used in some exams. It would be dangerous to modify this question"`
}

type QuestionDeleteNotSafeError struct {
	Type    string `json:"type" example:"Question"`
	Message string `json:"message" example:"This question id 'question_id' is already used in some exams. Therefore, it cannot be deleted!"`
}

type QuestionNoUploadError struct {
	Type    string `json:"type" example:"Question"`
	Message string `json:"message" example:"There is no file uploaded to this question id 'question_id'"`
}
