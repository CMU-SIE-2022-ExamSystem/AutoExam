package validate

import (
	"mime/multipart"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/go-playground/validator/v10"
)

func GraderCreateValidation(sl validator.StructLevel) {
	grader := sl.Current().Interface().(course.Grader_Create_Validate)

	if !dao.ValidateGrader(grader.Name, grader.BaseCourse) {
		sl.ReportError(grader.Name, "name", "Name", "name", grader.BaseCourse)
	}

	blanksType(sl, grader.Blanks)
}

func GraderUpdateValidation(sl validator.StructLevel) {
	grader := sl.Current().Interface().(course.Grader_Update)

	blanksType(sl, grader.Blanks)
}

func GraderUploadValidation(sl validator.StructLevel) {
	grader := sl.Current().Interface().(course.Grader_Upload)

	if fileRequired(grader.File) {
		// validate file extension
		if !fileExtValidate(grader.File.Filename) {
			sl.ReportError(grader.File, "file", "File", "extension", ".py")
		}
	}
}

func fileRequired(file *multipart.FileHeader) bool {
	return file != nil
}

func fileExtValidate(name string) bool {
	return filepath.Ext(name) == ".py"
}

func blanksType(sl validator.StructLevel, blanks []dao.Blanks) {
	if len(blanks) == 0 {
		sl.ReportError(blanks, "blanks", "Blanks", "required", "")
	}

	for i, blank := range blanks {
		if blank.Type == "" {
			sl.ReportError(blank.Type, "type", "Type", "requiredType", utils.Ordinalize(i+1))
			return
		}
		if blank.Type != "string" && blank.Type != "code" {
			sl.ReportError(blank.Type, "type", "Type", "oneofType", "string code,"+utils.Ordinalize(i+1))
		}

	}
}
