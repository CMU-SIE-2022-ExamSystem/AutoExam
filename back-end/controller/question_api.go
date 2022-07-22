package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

//TODO: need some works
func Question_Handler(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	color.Yellow(string(data))
}

// ReadAllQuestion_Handler godoc
// @Summary read all questions configuration
// @Schemes
// @Description read all questions configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=dao.Tags_Return} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions [get]
func ReadAllQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name := c.Param("course_name")
	fmt.Println(course_name)
}

// CreateQuestion_Handler godoc
// @Summary create a new question configuration
// @Schemes
// @Description create a new question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body dao.Question_Header_Create true "body data"
// @Success 201 {object} response.Response{data=dao.Question_Header} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/ [post]
// @Deprecated
func CreateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	var body dao.Question_Header_Create
	validate.ValidateJson(c, &body)
	// course_name := course.GetCourse(c)
	// course.Validate_assessment_name(c, course_name, body.Name)
	// assessment := body.ToAutoExamAssessments(course_name)
	// _, err := dao.CreateExam(assessment)
	// if err != nil {
	// 	response.ErrDBResponse(c, "There is an error when storing a new assessment to mongodb")
	// }
	// response.SuccessResponse(c, assessment)
}

// ReadQuestion_Handler godoc
// @Summary read a question configuration
// @Schemes
// @Description read a question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		question_id			path	string	true	"Questions Id"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [get]
// @Deprecated
func ReadQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	assessment, err := dao.ReadExam(course_name, assessment_name)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when reading an assessment from mongodb")
	}
	response.SuccessResponse(c, assessment)
}

// UpdateQuestion_Handler godoc
// @Summary update a question configuration
// @Schemes
// @Description update a question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		question_id			path	string	true	"Questions Id"
// @Param data body dao.AutoExam_Assessments_Update true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [put]
// @Deprecated
func UpdateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)

	var body dao.AutoExam_Assessments_Update
	validate.ValidateJson(c, &body)
	assessment := body.ToAutoExamAssessments(course_name)

	// check whether new data's name is same as the original assessment's name
	if !(assessment.Id == assessment_name) {
		if course.Validate_autoexam_assessment(c, course_name, assessment.Id) {
			response.ErrAssessmentNotValidResponse(c, course_name, assessment.Id)
		}
	}

	err := dao.UpdateExam(course_name, assessment_name, assessment)

	if err != nil {
		fmt.Println("==========")
		fmt.Println(err)
		fmt.Println("==========")
		response.ErrDBResponse(c, "There is an error when updating an assessment to mongodb")
	}

	response.SuccessResponse(c, assessment)
}

// DeleteQuestion_Handler godoc
// @Summary delete a question configuration
// @Schemes
// @Description delete a question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		question_id			path	string	true	"Questions Id"
// @Success 204
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [delete]
// @Deprecated
func DeleteQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	err := dao.DeleteExam(course_name, assessment_name)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when deleting an assessment to mongodb")
	}
	response.NonContentResponse(c)
}
