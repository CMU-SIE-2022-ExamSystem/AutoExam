package controller

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/gin-gonic/gin"
)

const (
	Grader_Model = "grader"
)

var Basic_Grader = []string{"multiple_blank", "multiple_choice", "single_blank", "single_choice"}

// ReadAllGrader_Handler godoc
// @Summary read all graders configuration except basic grader
// @Schemes
// @Description read all graders configuration except basic grader
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=[]dao.Grader_API} "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders [get]
func ReadAllGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)
	grader, err := dao.ReadAllGraders(base_course)
	if err != nil {
		response.ErrMySQLReadAllResponse(c, Grader_Model)
	}
	response.SuccessResponse(c, grader)
}

// ReadAllGrader_Handler godoc
// @Summary read all graders configuration with basic grader
// @Schemes
// @Description read all graders configuration with basic grader
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=[]string} "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/list [get]
func ReadGraderList_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)
	grader, err := dao.ReadAllGraders(base_course)
	if err != nil {
		response.ErrMySQLReadAllResponse(c, Grader_Model)
	}

	// TODO add basic grader
	// TODO should not show not uploaded or not valided
	var list []string
	for _, grade := range grader {
		list = append(list, grade.Name)
	}
	list = append(Basic_Grader, list...)

	response.SuccessResponse(c, list)
}

// CreateGrader_Handler godoc
// @Summary create a new grader configuration without grader file
// @Schemes
// @Description create a new grader configuration without grader file
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body course.Grader_Creat true "body data"
// @Success 201 {object} response.Response{data=dao.Grader_API} "created"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/ [post]
func CreateGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	// validate
	var body course.Grader_Create_Validate
	_, base_course := course.GetCourseBaseCourse(c)
	body.BaseCourse = base_course
	validate.ValidateJson(c, &body)

	instance := dao.PythonFile{
		QuestionType: body.Name,
		BaseCourse:   body.BaseCourse,
		Blanks:       body.Blanks,
	}
	grader, err := dao.Insert_grader(instance)
	if err != nil {
		response.ErrMySQLCreateResponse(c, Grader_Model)
	}

	response.SuccessResponse(c, grader)
}

// ReadGrader_Handler godoc
// @Summary read a grader configuration
// @Schemes
// @Description read a grader configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [get]
func ReadGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)
	grader, err := dao.ReadGrader(grader_name, base_course)
	if err != nil {
		response.ErrMongoDBReadResponse(c, Grader_Model)
	}
	response.SuccessResponse(c, grader)
}

// UpdateGrader_Handler godoc
// @Summary update a grader configuration
// @Schemes
// @Description update a grader configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param data body course.Grader_Update true "body data"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.UpdateNotSafeError} "not update safe"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [put]
func UpdateGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	var body course.Grader_Update
	validate.ValidateJson(c, &body)

	// check used in question or not
	if status, err := dao.ValidateGraderUsed(grader_name, base_course); err != nil {
		response.ErrMySQLDeleteResponse(c, Tag_Model)
	} else if !status {
		response.ErrGraderUpdateNotSafeResponse(c, grader_name)
	}

	instance := dao.PythonFile{
		QuestionType: grader_name,
		BaseCourse:   base_course,
		Blanks:       body.Blanks,
	}

	fmt.Println(instance)

	grader, err := dao.Update_blanks_grader(instance)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	response.SuccessResponse(c, grader)
}

// UpdateGrader_Handler godoc
// @Summary update a grader configuration forcely
// @Schemes
// @Description update a grader configuration forcely, this would not validate whether this grader is used or not
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param data body course.Grader_Update true "body data"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.GraderResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/force [put]
func UpdateForceGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	var body course.Grader_Update
	validate.ValidateJson(c, &body)

	instance := dao.PythonFile{
		QuestionType: grader_name,
		BaseCourse:   base_course,
		Blanks:       body.Blanks,
	}

	grader, err := dao.Update_blanks_grader(instance)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	response.SuccessResponse(c, grader)
}

// DeleteGrader_Handler godoc
// @Summary delete a grader configuration
// @Schemes
// @Description delete a grader configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Success 204 "no content"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.DeleteNotSafeError} "not delete safe"
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}  [delete]
func DeleteGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	// check used in question or not
	if status, err := dao.ValidateGraderUsed(grader_name, base_course); err != nil {
		response.ErrMySQLDeleteResponse(c, Tag_Model)
	} else if !status {
		response.ErrGraderDeleteNotSafeResponse(c, grader_name)
	}

	err := dao.Delete_grader(grader_name, base_course)
	if err != nil {
		response.ErrMySQLDeleteResponse(c, Grader_Model)
	}
	response.NonContentResponse(c)
}

// ValidGrader_Handler godoc
// @Summary edit a grader valid configuration
// @Schemes
// @Description edit a grader valid configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param data body course.Grader_Valid true "body data"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/valid  [put]
func ValidGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, grader_name := course.GetCourseGrader(c)

	var body course.Grader_Valid
	validate.ValidateJson(c, &body)
	grader, err := dao.UpdateGraderValid(grader_name, course_name, body.Valid)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}
	response.SuccessResponse(c, grader)
}

// UploadGrader_Handler godoc
// @Summary upload a grader file with .py extension
// @Schemes
// @Description upload a grader file with .py extension
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.UpdateNotSafeError} "not update safe"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/upload [put]
func UploadGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	var body course.Grader_Upload
	validate.ValidateForm(c, &body)

	// check used in question or not
	if status, err := dao.ValidateGraderUsed(grader_name, base_course); err != nil {
		response.ErrMySQLDeleteResponse(c, Tag_Model)
	} else if !status {
		response.ErrGraderUpdateNotSafeResponse(c, grader_name)
	}

	upload_and_store_grader(c, base_course, grader_name, body)
}

// UploadForceGrader_Handler godoc
// @Summary upload a grader file with .py extension forcely
// @Schemes
// @Description upload a grader file with .py extension forcely, this would not validate whether this grader is used or not
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 404 {object} response.GraderResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/upload/force [put]
func UploadForceGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	var body course.Grader_Upload
	validate.ValidateForm(c, &body)

	upload_and_store_grader(c, base_course, grader_name, body)
}

func upload_and_store_grader(c *gin.Context, base_course, grader_name string, body course.Grader_Upload) {
	instance := dao.PythonFile{
		QuestionType: grader_name,
		BaseCourse:   base_course,
		PythonGrader: course.FileToByte(c, body.File),
	}

	grader, err := dao.Update_python_grader(instance)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	file := course.Grader_Store{
		Name:       grader_name,
		BaseCourse: base_course,
		File:       body.File,
	}

	// store to file system
	course.StoreFile(c, file)

	response.SuccessResponse(c, grader)
}
