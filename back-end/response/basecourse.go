package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrBasecourseNotValidResponse(c *gin.Context) {
	err := Error{Type: Course, Message: "This base course is already in use and cannot be edited."}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrBasecourseNotExistsResponse(c *gin.Context) {
	err := Error{Type: Course, Message: "This base course not exists."}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}
