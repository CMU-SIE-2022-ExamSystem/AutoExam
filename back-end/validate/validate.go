package validate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "datetime":
		return "This field should be in the format of " + fe.Param()
	case "ltfield":
		return "This field should be less than the field of " + fe.Param()
	case "oneof":
		return "Should be one of [" + strings.ReplaceAll(fe.Param(), " ", ", ") + "]"
	// assessment's tag
	case "submission":
		return "Should be only 1 when category_name is 'Exam'"
	}
	fmt.Println(fe.Namespace())
	fmt.Println(fe.Field())
	fmt.Println(fe.StructNamespace())
	fmt.Println(fe.StructField())
	fmt.Println(fe.Tag())
	fmt.Println(fe.ActualTag())
	fmt.Println(fe.Kind())
	fmt.Println(fe.Type())
	fmt.Println(fe.Value())
	fmt.Println(fe.Param())
	fmt.Println()
	return "Unknown error"
}

func transErrorMsg(c *gin.Context, err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	} else {
		fmt.Println("===========")
		fmt.Println(reflect.TypeOf(err))
		fmt.Println("===========")
		response.ErrValidateResponse(c, err)
	}
	return nil
}

func Validate(c *gin.Context, obj any) {
	if err := c.ShouldBindJSON(&obj); err != nil {
		if msg := transErrorMsg(c, err); msg != nil {
			response.ErrValidateResponse(c, msg)
		}
	}
}
