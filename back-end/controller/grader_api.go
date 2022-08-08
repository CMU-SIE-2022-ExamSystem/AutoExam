package controller

import (
	"bytes"
	"os/exec"
	"syscall"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/fatih/color"
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
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
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
// @Summary read all validated graders configuration with basic grader
// @Schemes
// @Description read all validated graders configuration with basic grader
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=[]string} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/list [get]
func ReadGraderList_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)
	grader, err := dao.ReadAllValidGraders(base_course)
	if err != nil {
		response.ErrMySQLReadAllResponse(c, Grader_Model)
	}

	var list []string
	for _, grade := range grader {
		list = append(list, grade.Name)
	}

	// add basic grader
	list = append(global.Settings.Basic_Grader, list...)

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
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
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
		Modules:      body.Modules,
	}
	grader, err := dao.Insert_grader(instance)
	if err != nil {
		response.ErrMySQLCreateResponse(c, Grader_Model)
	}

	response.CreatedResponse(c, grader)
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
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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
// @Description update a grader configuration and cannot update any grader that is already used in some questions
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param data body course.Grader_Update true "body data"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.GraderUpdateNotSafeError} "not update safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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

	grader, err := dao.Update_blanks_grader(instance)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}

	response.SuccessResponse(c, grader)
}

// UpdateGrader_Handler godoc
// @Summary update a grader configuration forcely
// @Schemes
// @Description update a grader configuration forcely, this would not validate whether this grader is used in some questions or not
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param data body course.Grader_Update true "body data"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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
// @Description delete a grader configuration and cannot delete any grader that is already used in some questions
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Success 204 "no content"
// @Failure 400 {object} response.BadRequestResponse{error=response.GraderDeleteNotSafeError} "not delete safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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
// @Failure 400 {object} response.BadRequestResponse{error=response.GraderNoUploadError} "no uploaded file or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}/valid  [put]
func ValidGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, grader_name := course.GetBaseCourseGrader(c)

	var body course.Grader_Valid
	validate.ValidateJson(c, &body)

	if body.Valid && !dao.GetUploadStatus(grader_name, base_course) {
		response.ErrGraderNoUploadResponse(c, grader_name)
	}

	grader, err := dao.UpdateGraderValid(grader_name, base_course, body.Valid)
	if err != nil {
		response.ErrMySQLUpdateResponse(c, Grader_Model)
	}
	response.SuccessResponse(c, grader)
}

// UploadGrader_Handler godoc
// @Summary upload a grader file with .py extension
// @Schemes
// @Description upload a grader file with .py extension and cannot upload the file for any grader that is already used in some questions
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.GraderUpdateNotSafeError} "not update safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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
// @Description upload a grader file with .py extension forcely, this would not validate whether this grader is used in some questions or not
// @Tags grader
// @Accept mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Param file formData file true "the python file"
// @Success 200 {object} response.Response{data=dao.Grader_API} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid of grader or course"
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

// function for upload python file and store to file system
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

// Testgrader_Handler godoc
// @Summary test a grader
// @Schemes
// @Description test a grader file to make it work in our system
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Success 200 "success"
// @Failure 500 "not valid"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/autograder/{grader_name}/test [get]
func Testgrader_Handler(c *gin.Context) {
	base_course, question_type := course.GetBaseCourseGrader(c)
	color.Yellow(base_course)
	dao.SearchAndStore_grader(c, question_type, base_course, "./autograder/exec/autograders/")
	dao.SearchAndStore_module(c, question_type, base_course, "./autograder/exec/")
	utils.CheckModule()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("./driver.sh", question_type)
	cmd.Dir = "./autograder/exec/"
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	// os.Remove("./autograder/exec/autograders/" + question_type + ".py")
	// os.Remove("./autograder/exec/requirements.txt")
	if err != nil {
		dao.UpdateGraderValid(question_type, base_course, false)
		response.ErrorInternaWithData(c, err.Error(), stdout.String()+stderr.String())
	} else {
		dao.UpdateGraderValid(question_type, base_course, true)
		response.SuccessResponse(c, stdout.String())
	}
}
