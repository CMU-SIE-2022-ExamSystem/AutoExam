package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Tag = "Tag"
)

func ErrTagNotValidResponse(c *gin.Context, course, tag string) {
	err := Error{Type: Tag, Message: ReplaceMessageCourseQuestionName(&TagNotValidError{}, course, tag)}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrTagDeleteNotSafeResponse(c *gin.Context, tag string) {
	err := Error{Type: Tag, Message: ReplaceMessageTagName(&TagDeleteNotSafeError{}, tag)}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}
func ReplaceMessageTagName(str interface{}, tag string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "tag_id", tag)
}

func ReplaceMessageCourseTagnName(str interface{}, course_name, tag_id string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	message := strings.ReplaceAll(field.Tag.Get("example"), "course_name", course_name)
	return strings.ReplaceAll(message, "tag_id", tag_id)
}

type TagNotValidError struct {
	Type    string `json:"type" example:"Tag"`
	Message string `json:"message" example:"There is no this tag id 'tag_id' in such base course 'course_name'"`
}

type TagDeleteNotSafeError struct {
	Type    string `json:"type" example:"Tag"`
	Message string `json:"message" example:"This tag id 'tag_id' is already used in some questions. Therefore, it cannot be deleted!"`
}
