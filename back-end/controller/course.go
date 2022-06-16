package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

// Usercourses_Handler godoc
// @Summary get course assessments
// @Schemes
// @Description get course assessments list
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments [get]
func Assessments_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")

	body := autolab.AutolabUserHandler(c, token, "/courses/"+course_name+"/assessments")
	// fmt.Println(string(body))

	autolab_resp := utils.Course_assessments_trans(string(body))
	filtered_resp := utils.ExamNameFilter(autolab_resp)

	response.SuccessResponse(c, filtered_resp)
}
