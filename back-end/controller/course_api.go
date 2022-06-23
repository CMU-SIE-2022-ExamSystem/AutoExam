package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
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
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/exam [get]
func Exam_Handler(c *gin.Context) {
	response.SuccessResponse(c, dao.GetQuestions())
	// response.SuccessResponse(c, "1")
}

// Usercourses_Handler godoc
// @Summary get course assessments
// @Schemes
// @Description get course assessments list
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Assessments} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments [get]
func Assessments_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments")
	// fmt.Println(string(body))

	autolab_resp := utils.Course_assessments_trans(string(body))
	filtered_resp := utils.ExamNameFilter(autolab_resp)

	response.SuccessResponse(c, filtered_resp)
}

// Usercourses_Handler godoc
// @Summary get assessment submissions
// @Schemes
// @Description get assessment submissions list
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Submissions} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/submissions [get]
func Submissions_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submissions")
	// fmt.Println(string(body))

	autolab_resp := utils.Assessments_submissions_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}

func Usersubmit_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	body := autolab.AutolabSubmitHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submit", "./tmp/answer.tar")
	// fmt.Println(string(body))

	autolab_resp := utils.User_submit_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}
