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
// @Param		tag_id				query	string	false	"Tag Id"
// @Param		hidden				query	bool	false	"Show Hidden Question"
// @Success 200 {object} response.Response{data=[]dao.Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadAllError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions [get]
func ReadAllQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)
	tag_id := course.GetQueryTagId(c)
	hidden := course.GetQueryHidden(c)
	questions, err := dao.ReadAllQuestionsByTag(base_course, tag_id, hidden)
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
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBCreateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/ [post]
func CreateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, base_course := course.GetCourseBaseCourse(c)

	var body dao.Questions_Create_Validate
	body.BaseCourse = base_course
	validate.ValidateJson(c, &body)

	question, err := dao.CreateQuestion(body.ToAutoExamQuestions())
	if err != nil {
		response.ErrMongoDBCreateResponse(c, Que_Model)
	}

	response.CreatedResponse(c, question)
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
// @Success 200 {object} response.Response{data=dao.Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.QuestionNotValidError} "not valid of question or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [get]
func ReadQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, question_id := course.GetBaseCourseQuestion(c)
	question, err := dao.ReadQuestionById(question_id)
	if err != nil {
		response.ErrMongoDBReadResponse(c, Que_Model)
	}
	response.SuccessResponse(c, question)
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
// @Param data body dao.Questions_Create true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.QuestionModifyNotSafeError} "not update safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.QuestionNotValidError} "not valid of question or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [put]
func UpdateQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, question_id := course.GetBaseCourseQuestion(c)

	if status, err := dao.ValidateQuestionUsedById(question_id); !status {
		response.ErrQuestionModifyNotSafeResponse(c, question_id)
	} else if err != nil {
		response.ErrMongoDBReadResponse(c, Student_Model)
	}

	var body dao.Questions_Create_Validate
	body.BaseCourse = base_course
	validate.ValidateJson(c, &body)

	err := dao.UpdateQuestions(question_id, body.ToAutoExamQuestions())
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Que_Model)
	}

	question, _ := dao.ReadQuestionById(question_id)

	response.SuccessResponse(c, question)
}

// UpdateForceQuestion_Handler godoc
// @Summary update a question configuration without check used or not
// @Schemes
// @Description update a question configuration without check used or not
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		question_id			path	string	true	"Questions Id"
// @Param data body dao.Questions_Create true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.QuestionModifyNotSafeError} "not update safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.QuestionNotValidError} "not valid of question or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id}/force [put]
func UpdateForceQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	base_course, question_id := course.GetBaseCourseQuestion(c)

	var body dao.Questions_Create_Validate
	body.BaseCourse = base_course
	validate.ValidateJson(c, &body)

	err := dao.UpdateQuestions(question_id, body.ToAutoExamQuestions())
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Que_Model)
	}

	question, _ := dao.ReadQuestionById(question_id)

	response.SuccessResponse(c, question)
}

// DeleteQuestion_Handler godoc
// @Summary hard or soft delete a question configuration
// @Schemes
// @Description hard or soft delete a question configuration
// @Tags question
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		question_id			path	string	true	"Questions Id"
// @Param		hard				query	bool	false	"Hard Delete"
// @Success 204 "no content"
// @Failure 400 {object} response.BadRequestResponse{error=response.QuestionModifyNotSafeError} "not delete safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.QuestionNotValidError} "not valid of question or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBDeleteError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/questions/{question_id} [delete]
func DeleteQuestion_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, question_id := course.GetBaseCourseQuestion(c)
	hard := course.GetQueryHard(c)

	if hard {
		if status, err := dao.ValidateQuestionUsedById(question_id); !status {
			response.ErrQuestionDeleteNotSafeResponse(c, question_id)
		} else if err != nil {
			response.ErrMongoDBReadResponse(c, Student_Model)
		}
	}

	err := dao.DeleteQuestionById(question_id, hard)
	if err != nil {
		response.ErrMongoDBDeleteResponse(c, Que_Model)
	}

	response.NonContentResponse(c)
}

// TODO check  check update or force update
// TODO question & tag put delete CORS
