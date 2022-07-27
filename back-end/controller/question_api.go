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
	Que_Model = "question"
)

// ReadAllQuestion_Handler godoc
// @Summary read all questions configuration
// @Schemes
// @Description read all questions configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=[]dao.Questions} "success"
// @Failure 403 {object} response.ForbiddenResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadAllError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions [get]
func ReadAllQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)

	questions, err := dao.ReadAllQuestions(base_course)
	if err != nil {
		response.ErrMongoDBReadAllResponse(c, Que_Model)
	}
	response.SuccessResponse(c, questions)
}

// CreateQuestion_Handler godoc
// @Summary create a new question configuration
// @Schemes
// @Description create a new question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body dao.Questions_Create true "body data"
// @Success 201 {object} response.Response{data=dao.Questions} "created"
// @Failure 403 {object} response.ForbiddenResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBCreateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/ [post]
func CreateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, base_course := course.GetCourseBaseCourse(c)

	var body dao.Questions_Create_Validate
	body.Course = base_course
	validate.ValidateJson(c, &body)

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
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 404 {object} response.NotValidResponse{} "not valid"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [get]
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
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 404 {object} response.NotValidResponse{error=response.GraderNotValidError} "not valid"
// @Failure 400 {object} response.GraderResponse{error=response.UpdateNotSafeError} "not update safe"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [put]
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
// @Success 204 "no content"
// @Failure 400 {object} response.GraderResponse{error=response.DeleteNotSafeError} "not delete safe"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBDeleteError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [delete]
func DeleteQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	err := dao.DeleteExam(course_name, assessment_name)
	if err != nil {
		response.ErrDBResponse(c, "There is an error when deleting an assessment to mongodb")
	}
	response.NonContentResponse(c)
}
