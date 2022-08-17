package course

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func GetBaseCourseQuestion(c *gin.Context) (string, string) {
	_, base := GetCourseBaseCourse(c)
	question := c.Param("question_id")

	if status, _ := dao.ValidateQuestionById(base, question); status {
		// fmt.Println(status)
		response.ErrQuestionNotValidResponse(c, base, question)
	}
	return base, question
}

func GetQueryTagId(c *gin.Context) string {
	_, base := GetCourseBaseCourse(c)
	tag_id := c.Query("tag_id")
	if tag_id == "" {
		return tag_id
	}
	Validate_tag(c, base, tag_id)
	return tag_id
}
func GetQueryHard(c *gin.Context) bool {
	hard := c.Query("hard")
	if hard == "" || hard == "false" {
		return false
	}
	return true
}

func GetQueryHidden(c *gin.Context) bool {
	hidden := c.Query("hidden")
	if hidden == "" || hidden == "false" {
		return false
	}
	return true
}
