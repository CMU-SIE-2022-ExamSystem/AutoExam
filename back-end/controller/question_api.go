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

// QuestionTag_Handler godoc
// @Summary get the possible list of all question tags
// @Schemes
// @Description get the possible list of all question tags
// @Tags question
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=dao.Tags_Return} "desc"
// @Security ApiKeyAuth
// @Router /questions/tags [get]
func QuestionTag_Handler(c *gin.Context) {
	tags, err := dao.GetTags()
	if err != nil {
		response.ErrDBResponse(c, "There is an error when reading question tags from mongodb")
	}
	data := dao.Tags_Return{Tags: tags}
	response.SuccessResponse(c, data)
}

// CreateQuestion_Handler godoc
// @Summary create an new question configuration
// @Schemes
// @Description create an new question configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param data body dao.AutoExam_Assessments_Create true "body data"
// @Success 201 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /questions/ [post]
// @Deprecated
func CreateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	var body dao.AutoExam_Assessments_Create
	validate.Validate(c, &body)
	course_name := course.GetCourse(c)
	course.Validate_assessment_name(c, course_name, body.Name)
	assessment := body.ToAutoExamAssessments(course_name)
	_, err := dao.CreateExam(assessment)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when storing a new assessment to mongodb")
	}
	response.SuccessResponse(c, assessment)
}

// ReadQuestion_Handler godoc
// @Summary read an question configuration
// @Schemes
// @Description read an question configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		question_id			path	string	true	"Questions Id"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /questions/{question_id} [get]
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
// @Summary update an question configuration
// @Schemes
// @Description update an question configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		question_id			path	string	true	"Questions Id"
// @Param data body dao.AutoExam_Assessments_Update true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /questions/{question_id} [put]
// @Deprecated
func UpdateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)

	var body dao.AutoExam_Assessments_Update
	validate.Validate(c, &body)
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
// @Summary delete an question configuration
// @Schemes
// @Description delete an question configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		question_id			path	string	true	"Questions Id"
// @Success 204
// @Security ApiKeyAuth
// @Router /questions/{question_id} [delete]
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
