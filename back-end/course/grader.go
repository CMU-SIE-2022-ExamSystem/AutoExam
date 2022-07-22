package course

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

const (
	Grader_Path = "tmp/grader"
)

type Grader_Creat struct {
	Name string                `form:"name" json:"name" binding:"required"`
	File *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

type Grader_Create_Validate struct {
	Name   string                `form:"name" json:"name" binding:"required"`
	Course string                `form:"course" json:"course" binding:"required"`
	File   *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

type Grader_Update struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

type Grader_Valid struct {
	Valid bool `json:"valid"`
}

func GetCourseGrader(c *gin.Context) (string, string) {
	course := GetCourse(c)
	grader := c.Param("grader_name")
	if status := dao.ValidateGrader(grader, course); status {
		response.ErrGraderNotValidResponse(c, course, grader)
	}
	return course, grader
}

func StoreFile(c *gin.Context, grader Grader_Create_Validate) {
	course_path := filepath.Join(Grader_Path, grader.Course)
	utils.CreateFolder(course_path)
	if err := c.SaveUploadedFile(grader.File, filepath.Join(course_path, grader.Name+".py")); err != nil {
		response.ErrFileStoreResponse(c)
	}
}

func FileToByte(c *gin.Context, file *multipart.FileHeader) []byte {
	fileContent, err := file.Open()
	if err != nil {
		response.ErrGraderReadFileResponse(c, err)
	}
	byteContent, err := ioutil.ReadAll(fileContent)
	if err != nil {
		response.ErrGraderReadFileResponse(c, err)
	}
	return byteContent
}
