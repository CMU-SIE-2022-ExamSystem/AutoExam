package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrImageNotExistsResponse(c *gin.Context) {
	err := Error{Type: Course, Message: "no image found according to the base course and image id in the dadabase!"}
	ErrorResponseWithStatus(c, err, http.StatusNotFound)
}
