package error

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(ErrorResponse); ok {
		message := gin.H{
			"error": gin.H{
				err.Scope: err.Message,
			},
		}
		c.JSON(err.Status, message)
		data, _ := json.Marshal(message)
		c.Error(errors.New(string(data)))
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

var (
	Authentication = "Authentication"
	Unauthorized   = AuthRespnse("Token error")
)

type ErrorResponse struct {
	Status  int
	Scope   string
	Message string
}

func EnvResponse(title string) ErrorResponse {
	return InternalErrorResponse(Authentication, title+" is not existed in settings-dev.yaml, please add it and restart the server")
}

func AuthRespnse(message string) ErrorResponse {
	return ErrorResponse{Status: http.StatusUnauthorized, Scope: Authentication, Message: message}
}

func InternalErrorResponse(scope, message string) ErrorResponse {
	return ErrorResponse{Status: http.StatusInternalServerError, Scope: scope, Message: message}
}
