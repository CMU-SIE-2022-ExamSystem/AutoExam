package course

import (
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
)

var (
	Course_path = "./source/courses/"
)

func Build_Course(course string) string {
	path := filepath.Join(Course_path, course)
	utils.CreateFolder(path)
	return path
}
