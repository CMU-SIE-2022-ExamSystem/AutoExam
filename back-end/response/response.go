package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status"`
	Type   int         `json:"type"`
	Error  any         `json:"error"`
	Data   interface{} `json:"data"`
}

type SwaggerResponse struct {
	Type  int         `json:"type"`
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ErrorJson struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
}

var (
	Authentication = "Authentication"
	FileSystem     = "FileSystem"
	Course         = "Course"
	Database       = "Database"
	Validation     = "Validation"
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

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"type":  0,
		"error": struct{}{},
		"data":  data,
	})
}

func NonContentResponse(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(http.StatusNoContent)
}

func SuccessResponseWithType(c *gin.Context, data interface{}, t int) {
	c.JSON(http.StatusOK, gin.H{
		"type":  t,
		"error": struct{}{},
		"data":  data,
	})
}

func ErrorResponseWithStatus(c *gin.Context, err any, status int) {
	c.JSON(status, gin.H{
		"type":  0,
		"error": err,
		"data":  struct{}{},
	})
	panic(err)
}

func ErrorInternalResponse(c *gin.Context, err Error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"type":  0,
		"error": err,
		"data":  struct{}{},
	})
	panic(err)
}

func ErrorInternaWithType(c *gin.Context, err Error, t int) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"type":  t,
		"error": err,
		"data":  struct{}{},
	})
	panic(err)
}

func ErrorInternaWithData(c *gin.Context, err any, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"type":  0,
		"error": err,
		"data":  data,
	})
	panic(err)
}

func ErrUnauthResponse(c *gin.Context, message string) {
	err := Error{Type: Authentication, Message: message}
	ErrorResponseWithStatus(c, err, http.StatusUnauthorized)
}

func ErrForbiddenResponse(c *gin.Context, message string) {
	err := Error{Type: Authentication, Message: message}
	ErrorResponseWithStatus(c, err, http.StatusForbidden)
}

func ErrFileNotValidResponse(c *gin.Context) {
	err := Error{Type: FileSystem, Message: "Target file does not exist or it is empty."}
	ErrorResponseWithStatus(c, err, http.StatusInternalServerError)
}

func ErrFileStoreResponse(c *gin.Context) {
	err := Error{Type: FileSystem, Message: "There is an error happened when storing a file."}
	ErrorResponseWithStatus(c, err, http.StatusInternalServerError)
}

func ErrCourseNotValidResponse(c *gin.Context, course string) {
	err := Error{Type: Course, Message: "This user is not registered in such course '" + course + "'"}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

func ErrValidateResponse(c *gin.Context, message interface{}) {
	err := ErrorJson{Type: Validation, Message: message}
	ErrorResponseWithStatus(c, err, http.StatusBadRequest)
}

type BadRequestResponse struct {
	Status int         `json:"status" example:"400"`
	Type   int         `json:"type" example:"0"`
	Error  any         `json:"error"`
	Data   interface{} `json:"data"`
}

type NotValidResponse struct {
	Status int         `json:"status" example:"404"`
	Type   int         `json:"type" example:"0"`
	Error  any         `json:"error"`
	Data   interface{} `json:"data"`
}

type ForbiddenResponse struct {
	Status int         `json:"status" example:"403"`
	Type   int         `json:"type" example:"0"`
	Error  any         `json:"error"`
	Data   interface{} `json:"data"`
}

type ForbiddenError struct {
	Type    string `json:"type" example:"Authentication"`
	Message string `json:"message" example:"The user is not 'auth_level' in this course"`
}
