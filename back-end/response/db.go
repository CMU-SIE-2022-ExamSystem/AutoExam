package response

import (
	"github.com/gin-gonic/gin"
)

var (
	MongoDB = "MongoDB"
)

func ErrDBResponse(c *gin.Context, message string) {
	err := Error{Type: Database, Message: message}
	ErrorInternaWithType(c, err, -1)
}

func ErrMongoDBResponse(c *gin.Context, message string) {
	err := Error{Type: MongoDB, Message: message}
	ErrorInternaWithType(c, err, -1)
}

func ErrMongoDBReadAllResponse(c *gin.Context, model string) {
	message := "There is an error when reading all " + model + "s from mongodb"
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBReadResponse(c *gin.Context, model string) {
	message := "There is an error when reading an old " + model + " from mongodb"
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBCreateResponse(c *gin.Context, model string) {
	message := "There is an error when writing a new " + model + " to mongodb"
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBUpdateResponse(c *gin.Context, model string) {
	message := "There is an error when update an old " + model + " to mongodb"
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBDeleteResponse(c *gin.Context, model string) {
	message := "There is an error when delete an old " + model + " from mongodb"
	ErrMongoDBResponse(c, message)
}
