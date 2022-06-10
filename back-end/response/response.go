package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int
	Type   int
	Error  Error
	Data   interface{}
}

type SwaggerResponse struct {
	Type  int
	Error Error
	Data  interface{}
}

type Error struct {
	Type    string
	Message string
}

var (
	Authentication = "Authentication"
)

func NormalResponse(c *gin.Context, response Response) {
	c.JSON(response.Status, gin.H{
		"type":  response.Type,
		"error": response.Error,
		"data":  response.Data,
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	SuccessResponseWithType(c, data, 0)
}

func SuccessResponseWithType(c *gin.Context, data interface{}, t int) {
	c.JSON(http.StatusOK, gin.H{
		"type":  t,
		"error": struct{}{},
		"data":  data,
	})
}

func ErrorResponseWithStatus(c *gin.Context, err Error, status int) {
	c.JSON(status, gin.H{
		"type":  0,
		"error": err,
		"data":  struct{}{},
	})
}

func ErrorInternalResponse(c *gin.Context, err Error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"type":  0,
		"error": err,
		"data":  nil,
	})
}

func ErrUnauthResponse(c *gin.Context, message string) {
	err := Error{Type: Authentication, Message: message}
	ErrorResponseWithStatus(c, err, http.StatusUnauthorized)
}
