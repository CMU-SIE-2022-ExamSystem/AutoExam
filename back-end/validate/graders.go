package validate

import (
	"mime/multipart"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/go-playground/validator/v10"
)

func GraderCreateValidation(sl validator.StructLevel) {
	grader := sl.Current().Interface().(course.Grader_Create_Validate)

	// validate grader name
	if status := dao.ValidateGrader(grader.Name, grader.Course); !status {
		sl.ReportError(grader.Name, "name", "Name", "name", grader.Course)
	}

	if fileRequired(grader.File) {
		// validate file extension
		if !fileExtValidate(grader.File.Filename) {
			sl.ReportError(grader.File, "file", "File", "extension", ".py")
		}
	}
}

func GraderUpdateValidation(sl validator.StructLevel) {
	grader := sl.Current().Interface().(course.Grader_Update)

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
