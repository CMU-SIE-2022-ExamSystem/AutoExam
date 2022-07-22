package validate

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/go-playground/validator/v10"
)

func TagsNameValidation(sl validator.StructLevel) {
	tags := sl.Current().Interface().(dao.AutoExam_Tags_Create)
	status, err := dao.ValidateTag(tags.Course, tags.Name)
	if err != nil {
		sl.ReportError(tags.Name, "name", "Name", "mongo", "")
	}
	if !status {
		sl.ReportError(tags.Name, "name", "Name", "name", tags.Course)
	}
}
