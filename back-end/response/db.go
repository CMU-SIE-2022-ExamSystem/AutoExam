package response

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	MongoDB = "MongoDB"
	MySQL   = "MySQL"
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
	message := ReplaceMessageModelName(&MongoDBReadAllError{}, model)
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBReadResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MongoDBReadError{}, model)
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBCreateResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MongoDBCreateError{}, model)
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBUpdateResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MongoDBUpdateError{}, model)
	ErrMongoDBResponse(c, message)
}

func ErrMongoDBDeleteResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MongoDBDeleteError{}, model)
	ErrMongoDBResponse(c, message)
}

func ErrMySQLResponse(c *gin.Context, message string) {
	err := Error{Type: MySQL, Message: message}
	ErrorInternaWithType(c, err, -1)
}

func ErrMySQLReadAllResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MySQLReadAllError{}, model)
	ErrMySQLResponse(c, message)
}

func ErrMySQLReadResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MySQLReadError{}, model)
	ErrMySQLResponse(c, message)
}

func ErrMySQLCreateResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MySQLCreateError{}, model)
	ErrMySQLResponse(c, message)
}

func ErrMySQLUpdateResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MySQLUpdateError{}, model)
	ErrMySQLResponse(c, message)
}

func ErrMySQLDeleteResponse(c *gin.Context, model string) {
	message := ReplaceMessageModelName(&MySQLDeleteError{}, model)
	ErrMySQLResponse(c, message)
}

func ReplaceMessageModelName(str interface{}, model string) string {
	field, _ := reflect.TypeOf(str).Elem().FieldByName("Message")
	return strings.ReplaceAll(field.Tag.Get("example"), "model_name", model)
}

type DBesponse struct {
	Status int         `json:"status" example:"500"`
	Type   int         `json:"type" example:"-1"`
	Error  any         `json:"error"`
	Data   interface{} `json:"data"`
}

type MongoDBReadAllError struct {
	Type    string `json:"type" example:"MongoDB"`
	Message string `json:"message" example:"There is an error when reading all model_name from mongodb"`
}

type MongoDBReadError struct {
	Type    string `json:"type" example:"MongoDB"`
	Message string `json:"message" example:"There is an error when reading one model_name from mongodb"`
}

type MongoDBCreateError struct {
	Type    string `json:"type" example:"MongoDB"`
	Message string `json:"message" example:"There is an error when writing a new model_name to mongodb"`
}

type MongoDBUpdateError struct {
	Type    string `json:"type" example:"MongoDB"`
	Message string `json:"message" example:"There is an error when update one model_name to mongodb"`
}

type MongoDBDeleteError struct {
	Type    string `json:"type" example:"MongoDB"`
	Message string `json:"message" example:"There is an error when delete one model_name from mongodb"`
}

type MySQLReadAllError struct {
	Type    string `json:"type" example:"MySQL"`
	Message string `json:"message" example:"There is an error when reading all model_name from MySQL"`
}

type MySQLReadError struct {
	Type    string `json:"type" example:"MySQL"`
	Message string `json:"message" example:"There is an error when reading one model_name from MySQL"`
}

type MySQLCreateError struct {
	Type    string `json:"type" example:"MySQL"`
	Message string `json:"message" example:"There is an error when writing a new model_name to MySQL"`
}

type MySQLUpdateError struct {
	Type    string `json:"type" example:"MySQL"`
	Message string `json:"message" example:"There is an error when update one model_name to MySQL"`
}

type MySQLDeleteError struct {
	Type    string `json:"type" example:"MySQL"`
	Message string `json:"message" example:"There is an error when delete one model_name from MySQL"`
}
