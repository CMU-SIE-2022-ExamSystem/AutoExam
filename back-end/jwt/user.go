package jwt

import (
	"fmt"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func UserRefreshHandler(c *gin.Context) {
	user_email := GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	refresh := user.Refresh_token

	auth := global.Settings.Autolabinfo

	http_body := models.Http_Body_Refresh{
		Grant_type:    "refresh_token",
		Refresh_token: refresh,
		Scope:         auth.Scope,
		Client_id:     auth.Client_id,
		Client_secret: auth.Client_secret,
	}

	autolab_resp, flag := autolab.AutolabAuthHandler(c, "/oauth/token", http_body)

	if flag {
		user.Access_token = autolab_resp.Access_token
		user.Refresh_token = autolab_resp.Refresh_token
		user.Create_at = utils.GetNowTime()
		user.Expires_in = autolab_resp.Expires_in
		color.Yellow(user.Access_token)
		color.Yellow(user.Refresh_token)
		global.DB.Save(&user)
	}
}

//todo: these two check functions are testing user auth_level
func Check_authlevel_API(c *gin.Context) {
	course_name := c.Param("course_name")
	user_email := GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := autolab.AutolabGetHandler(c, token, "/courses")

	if strings.Contains(string(body), course_name) {
		autolab_resp := utils.User_courses_trans(string(body))
		autolab_map := utils.Map_user_authlevel(autolab_resp)
		fmt.Println(autolab_map)
		if autolab_map[course_name] == "student" {
			response.SuccessResponse(c, "student")
		} else if autolab_map[course_name] == "course_assistant" {
			response.SuccessResponse(c, "course_assistant")
		} else if autolab_map[course_name] == "instructor" {
			response.SuccessResponse(c, "instructor")
		}
	}
}

func Check_authlevel_DB(c *gin.Context) {
	course_name := c.Param("course_name")
	user_email := GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)

	auth_level := dao.Check_authlevel(user.ID, course_name)
	response.SuccessResponse(c, auth_level)
}

func Get_authlevel_DB(c *gin.Context) (auth_level string) {
	course_name := c.Param("course_name")
	user_email := GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)

	auth_level = dao.Check_authlevel(user.ID, course_name)
	if auth_level == "" {
		response.ErrorInternalResponse(c, response.Error{Type: "Database", Message: "There is no this user in database, please try again."})
	}
	return
}

func Check_authlevel_Student(c *gin.Context) {
	auth := Get_authlevel_DB(c)

	if auth != "student" {
		response.ErrUnauthResponse(c, "The user is not a student in this course")
	}

}

func Check_authlevel_Instructor(c *gin.Context) {
	auth := Get_authlevel_DB(c)

	if auth != "instructor" {
		response.ErrUnauthResponse(c, "The user is not an instructor in this course")
		c.Abort()
	}
}

func Check_authlevel_Assistant(c *gin.Context) {
	auth := Get_authlevel_DB(c)

	if auth != "course_assistant" {
		response.ErrUnauthResponse(c, "The user is not an assistant in this course")
	}
}

func Check_authlevel_Assistant_and_Instructor(c *gin.Context) {
	auth := Get_authlevel_DB(c)

	if auth == "student" {
		response.ErrUnauthResponse(c, "The user is not an assistant or an instructor in this course")
	}

}
