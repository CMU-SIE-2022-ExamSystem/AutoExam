package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrTagNotValidResponse(c *gin.Context, course, tag string) {
	err := Error{Type: Course, Message: "There is no this tag id '" + tag + "' in such course '" + course + "'"}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrTagNotSafeResponse(c *gin.Context, tag_name string) {
	err := Error{Type: Course, Message: "This tag name '" + tag_name + "' is used in some questions. Therefore, it cannot be deleted!"}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}
