package course

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

const (
	Grader_Path = "tmp/grader"
)

type Grader_Creat struct {
	Name   string       `form:"name" json:"name" binding:"required"`
	Blanks []dao.Blanks `form:"blanks" json:"blanks" binding:"required"`
}

type Grader_Create_Validate struct {
	Name       string       `form:"name" json:"name" binding:"required"`
	BaseCourse string       `form:"base_course" json:"base_course" binding:"required"`
	Blanks     []dao.Blanks `form:"blanks" json:"blanks"`
}

type Grader_Upload struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

type Grader_Update struct {
	Blanks []dao.Blanks `form:"blanks" json:"blanks" binding:"required"`
}

type Grader_Store struct {
	Name       string                `form:"name" json:"name" binding:"required"`
	BaseCourse string                `form:"base_course" json:"base_course" binding:"required"`
	File       *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

type Grader_Valid struct {
	Valid bool `json:"valid"`
}

type Test_Grader struct {
	Answer   string `json:"answer"`
	Solution string `json:"solution"`
}

func GetCourseGrader(c *gin.Context) (string, string) {
	course := GetCourse(c)
	grader := c.Param("grader_name")
	if status := dao.ValidateGraderName(grader, course); status {
		response.ErrGraderNotValidResponse(c, course, grader)
	}
	return course, grader
}

func GetBaseCourseGrader(c *gin.Context) (string, string) {
	_, base := GetCourseBaseCourse(c)
	grader := c.Param("grader_name")
	if status := dao.ValidateGraderName(grader, base); status {
		response.ErrGraderNotValidResponse(c, base, grader)
	}
	return base, grader
}

func StoreFile(c *gin.Context, grader Grader_Store) {
	course_path := filepath.Join(Grader_Path, grader.BaseCourse)
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

func GetBasicGraderLenDict() map[string]int {
	grader_dict := make(map[string]int)
	for _, grader := range global.Settings.Basic_Grader {
		grader_dict[grader] = 1
	}
	return grader_dict
}
