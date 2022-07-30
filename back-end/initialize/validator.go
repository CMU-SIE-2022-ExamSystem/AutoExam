package initialize

import (
	"reflect"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
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
		v.RegisterStructValidation(validate.AssessmentUpdateValidation, dao.AutoExam_Assessments_Update_Validate{})
		v.RegisterStructValidation(validate.TagsNameValidation, dao.AutoExam_Tags_Create{})
		v.RegisterStructValidation(validate.GraderCreateValidation, course.Grader_Create_Validate{})
		v.RegisterStructValidation(validate.GraderUpdateValidation, course.Grader_Update{})
		v.RegisterStructValidation(validate.GraderUploadValidation, course.Grader_Upload{})
		v.RegisterStructValidation(validate.QuestionsCreateValidation, dao.Questions_Create_Validate{})
		v.RegisterStructValidation(validate.AnswersUpdateValidation, dao.Answers_Upload_Validate{})
	}
}
