package course

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func GetBaseCourseQuestion(c *gin.Context) (string, string) {
	_, base := GetCourseBaseCourse(c)
	question := c.Param("question_id")
	fmt.Println(dao.ValidateQuestionById(base, question))
	if status, _ := dao.ValidateQuestionById(base, question); status {
		response.ErrQuestionNotValidResponse(c, base, question)
	}
	return base, question
}
