package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	BaseCourse = "BaseCourse"
)

func ErrBasecourseNotValidResponse(c *gin.Context) {
	err := Error{Type: BaseCourse, Message: "This base course is already in use and cannot be edited."}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrBasecourseNotExistsResponse(c *gin.Context) {
	err := Error{Type: BaseCourse, Message: "This base course not exists."}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrBasecourseRelationRecreatedResponse(c *gin.Context) {
	err := Error{Type: BaseCourse, Message: "This base course relationship is already in use and cannot be re-created."}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrBasecourseRelationNotExistsResponse(c *gin.Context) {
	err := Error{Type: BaseCourse, Message: "This base course relationship not exists."}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}

func ErrBasecourseRelationNotValidResponse(c *gin.Context) {
	err := Error{Type: BaseCourse, Message: "This base course is already in use and cannot be edited."}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

type BasecourseNotValidError struct {
	Type    string `json:"type" example:"BaseCourse"`
	Message string `json:"message" example:"This base course is already in use and cannot be edited."`
}

type BasecourseNotExistsError struct {
	Type    string `json:"type" example:"BaseCourse"`
	Message string `json:"message" example:"This base course not exists."`
}

type BasecourseRelationRecreatedError struct {
	Type    string `json:"type" example:"BaseCourse"`
	Message string `json:"message" example:"This base course relationship is already in use and cannot be re-created."`
}

type BasecourseRelationNotExistsError struct {
	Type    string `json:"type" example:"BaseCourse"`
	Message string `json:"message" example:"This base course relationship not exists."`
}
type BasecourseRelationNotValidError struct {
	Type    string `json:"type" example:"BaseCourse"`
	Message string `json:"message" example:"his base course is already in use and cannot be edited."`
}
