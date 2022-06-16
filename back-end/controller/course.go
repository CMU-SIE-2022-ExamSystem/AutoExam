package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

// Exam_Handler godoc
// @Summary get exam question
// @Schemes
// @Description get exam question
// @Tags exam
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "desc"
// @Param        course_name   		path      string  true  "Course Name"
// @Param        assessment_name   	path      string  true  "Assessment Name"
// @Router /courses/{course_name}/assessments/{assessment_name}/exam [get]
func Exam_Handler(c *gin.Context) {
	response.SuccessResponse(c, dao.GetQuestions())
	// response.SuccessResponse(c, "1")
}
