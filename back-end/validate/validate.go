package validate

import (
	"encoding/json"
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

// customized error message with Tag
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

	// assessment
	case "name":
		return "This name is already used in this base course '" + fe.Param() + "'"
	case "submission":
		return "Should be only 1 when category_name is 'Exam'"
	case "noTag":
		return "This field is required in the " + fe.Param() + " settings"
	case "notValidTag":
		s := strings.Split(fe.Param(), ",")
		return "This field is not one of the valid tags id in the base course '" + s[0] + "' in the " + s[1] + " settings"
	case "gteSetting":
		s := strings.Split(fe.Param(), ",")
		return "This field is required or should be greate than or equal than " + s[0] + " in the " + s[1] + " settings"
	case "mongo":
		return "There is an internal mongodb error when validate this field"
	case "notValidSub":
		s := strings.Split(fe.Param(), ",")
		return "Should be one of [" + strings.ReplaceAll(s[0], " ", ", ") + "]" + " in the " + s[1] + " settings"
	case "notValidIdTag":
		s := strings.Split(fe.Param(), ",")
		return "The tag of question id '" + s[0] + "' is not equal to the setting' tag in the " + s[1] + " settings"
	case "notValidIdNumber":
		s := strings.Split(fe.Param(), ",")
		return "The sub question number of question id '" + s[0] + "' is not equal to the setting' sub_question_number in the " + s[1] + " settings"
	case "notValidScoreLength":
		s := strings.Split(fe.Param(), ",")
		return "The length of scores should be equal to the setting' sub_question_number in the " + s[1] + " settings"
	case "notValidScore":
		return "The sum of scores should be equal to max_score in the " + fe.Param() + " settings"
	case "numberNotDivisible":
		return "The max_score is not divisible by sub_question_number in the " + fe.Param() + " settings. Please specifiy the field of scores"
	case "answerNotValid":
		return "The structure of data is wrong. Please use /courses/{course_name}/assessments/{assessment_name}/answers/struct to check the structure"
	case "answerWrongName":
		return "The key '" + fe.Param() + "' is wrong. Please use /courses/{course_name}/assessments/{assessment_name}/answers/struct to check the structure"

	// grader
	case "requiredType":
		return "This field is required in the " + fe.Param() + " blanks"
	case "oneofType":
		s := strings.Split(fe.Param(), ",")
		return "This field should be one of [" + strings.ReplaceAll(s[0], " ", ", ") + "] in the " + s[1] + " blanks"
	case "extension":
		return "This file's extension should be '" + fe.Param() + "'"

	// question
	case "notValidGrader":
		s := strings.Split(fe.Param(), ",")
		return "This field is not one of the valid grader in the base course '" + s[0] + "' in the " + s[1] + " sub_questions"
	case "requiredGrader":
		return "This field is required in the " + fe.Param() + " sub_questions"
	case "singleBlank":
		s := strings.Split(fe.Param(), ",")
		return "The length of array is required to be 1 when the question_type is '" + s[0] + "' in the " + s[1] + " sub_questions"
	case "lenAnswer":
		s := strings.Split(fe.Param(), ",")
		return "The length of the solutions does not equal to the blank length " + s[1] + " of grader '" + s[0] + "' in the " + s[2] + " sub_questions"
	case "notValidAnswer":
		return "The type of this field should two-dimensional array of string in the " + fe.Param() + " sub_questions"
	case "requiredChoice":
		return "This field is required when the grader name contains 'choice' in the " + fe.Param() + " sub_questions"
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

func TransErrorMsg(c *gin.Context, err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	} else if ute, ok := err.(*json.UnmarshalTypeError); ok {
		var out []ErrorMsg
		fmt.Println(err)
		out = append(out, ErrorMsg{ute.Field, "This field's type is not correct"})
		return out
	} else {
		fmt.Println("===========")
		fmt.Println(reflect.TypeOf(err))
		fmt.Println("===========")
		response.ErrValidateResponse(c, err)
	}
	return nil
}

func ValidateJson(c *gin.Context, obj any) {
	if err := c.ShouldBindJSON(&obj); err != nil {
		if msg := TransErrorMsg(c, err); msg != nil {
			response.ErrValidateResponse(c, msg)
		}
	}
}

func ValidateForm(c *gin.Context, obj any) {
	if err := c.ShouldBind(obj); err != nil {
		if msg := TransErrorMsg(c, err); msg != nil {
			response.ErrValidateResponse(c, msg)
		}
	}
}
