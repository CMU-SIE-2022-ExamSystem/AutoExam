package initialize

import (
	"reflect"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})
		// v.RegisterCustomTypeFunc(validate.ValidateTime, models.AutoTime{})
		v.RegisterStructValidation(validate.AssessmentUpdateValidation, dao.AutoExam_Assessments_Update{})
		v.RegisterStructValidation(validate.TagsNameValidation, dao.AutoExam_Tags_Create{})
	}
}
