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

// ReadAllGrader_Handler godoc
// @Summary read all graders configuration
// @Schemes
// @Description read all graders configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=dao.Tags_Return} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders [get]
func ReadAllGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name := c.Param("course_name")
	fmt.Println(course_name)
}

// CreateGrader_Handler godoc
// @Summary create a new grader configuration
// @Schemes
// @Description create a new grader configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body dao.AutoExam_Assessments_Create true "body data"
// @Success 201 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/ [post]
// @Deprecated
func CreateGrader_Handler(c *gin.Context) {
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

// ReadGrader_Handler godoc
// @Summary read a grader configuration
// @Schemes
// @Description read a grader configuration
// @Tags grader
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		grader_name			path	string	true	"Grader Name"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [get]
// @Deprecated
func ReadGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	assessment, err := dao.ReadExam(course_name, assessment_name)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when reading an assessment from mongodb")
	}
	response.SuccessResponse(c, assessment)
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
// @Param data body dao.AutoExam_Assessments_Update true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name} [put]
// @Deprecated
func UpdateGrader_Handler(c *gin.Context) {
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
// @Security ApiKeyAuth
// @Router /courses/{course_name}/graders/{grader_name}  [delete]
// @Deprecated
func DeleteGrader_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	err := dao.DeleteExam(course_name, assessment_name)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when deleting an assessment to mongodb")
	}
	response.NonContentResponse(c)
}
