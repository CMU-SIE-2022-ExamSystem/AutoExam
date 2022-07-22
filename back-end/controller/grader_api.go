package controller

import (
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
	course_name := c.Param("course_name")
	grader, err := dao.ReadAllGraders(course_name)
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
	course_name := c.Param("course_name")
	grader, err := dao.ReadAllGraders(course_name)
	if err != nil {
		response.ErrMySQLReadAllResponse(c, Grader_Model)
	}

	var list []string
	for _, grade := range grader {
		list = append(list, grade.Name)
	}
	list = append(list, Basic_Grader...)

	response.SuccessResponse(c, list)
}

// CreateGrader_Handler godoc
// @Summary create a new grader configuration
// @Schemes
// @Description create a new grader configuration
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data formData course.Grader_Creat true "body data"
// @Param file formData file true "the python file"
// @Success 201 {object} response.Response{data=dao.Grader_API} "created"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/ [post]
func CreateGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	// validate
	var body course.Grader_Create_Validate
	body.Course = course.GetCourse(c)
	validate.ValidateForm(c, &body)

	// read code to []byte{}
	grader, err := dao.InsertOrUpddbate_grader(body.Name, course.FileToByte(c, body.File), body.Course)
	if err != nil {
		response.ErrMySQLCreateResponse(c, Grader_Model)
	}

	// store to file system
	course.StoreFile(c, body)

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
// @Success 200 {object} response.Response{data=dao.Grader_API} "desc"
// @Failure 404 {object} response.NotValudResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [get]
func ReadGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, grader_name := course.GetCourseGrader(c)
	grader, err := dao.ReadGrader(grader_name, course_name)
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
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "desc"
// @Failure 404 {object} response.NotValudResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.UpdateNotSafeError} "not update safe"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [put]
func UpdateGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, grader_name := course.GetCourseGrader(c)

	var body course.Grader_Update
	validate.ValidateForm(c, &body)

	// check used in question or not
	if status, err := dao.ValidateGraderUsed(grader_name, course_name); err != nil {
		response.ErrMySQLDeleteResponse(c, Tag_Model)
	} else if !status {
		response.ErrGraderUpdateNotSafeResponse(c, grader_name)
	}

	grader, err := dao.InsertOrUpddbate_grader(grader_name, course.FileToByte(c, body.File), course_name)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	file := course.Grader_Create_Validate{
		Name:   grader_name,
		Course: course_name,
		File:   body.File,
	}

	// store to file system
	course.StoreFile(c, file)

	response.SuccessResponse(c, grader)
}

// UpdateGrader_Handler godoc
// @Summary update a grader configuration forcely
// @Schemes
// @Description update a grader configuration forcely, this would not validate whether this grader is used or not
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "desc"
// @Failure 404 {object} response.GraderResponse{error=response.GraderNotValidError} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/force [put]
func UpdateForceGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, grader_name := course.GetCourseGrader(c)

	var body course.Grader_Update
	validate.ValidateForm(c, &body)

	grader, err := dao.InsertOrUpddbate_grader(grader_name, course.FileToByte(c, body.File), course_name)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	file := course.Grader_Create_Validate{
		Name:   grader_name,
		Course: course_name,
		File:   body.File,
	}

	// store to file system
	course.StoreFile(c, file)

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
// @Success 204
// @Failure 404 {object} response.NotValudResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.DeleteNotSafeError} "not delete safe"
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}  [delete]
func DeleteGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, grader_name := course.GetCourseGrader(c)

	// check used in question or not
	if status, err := dao.ValidateGraderUsed(grader_name, course_name); err != nil {
		response.ErrMySQLDeleteResponse(c, Tag_Model)
	} else if !status {
		response.ErrGraderDeleteNotSafeResponse(c, grader_name)
	}

	err := dao.Delete_grader(grader_name, course_name)
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
// @Success 200 {object} response.Response{data=dao.Grader_API} "desc"
// @Failure 404 {object} response.NotValudResponse{error=response.GraderNotValidError} "not valid"
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
