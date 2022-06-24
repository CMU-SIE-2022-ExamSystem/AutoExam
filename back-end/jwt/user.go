package jwt

import (
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
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

func Check_authlevel(c *gin.Context) {
	course_name := c.Param("course_name")
	user_email := GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := autolab.AutolabGetHandler(c, token, "/courses")

	if strings.Contains(string(body), course_name) {
		autolab_resp := utils.User_courses_trans(string(body))
		autolab_map := utils.Map_user_authlevel(autolab_resp)
		if autolab_map[course_name] == "student" {
			response.SuccessResponse(c, "You do not have permission to access here.")
		} else {
			response.SuccessResponse(c, "You have permission to access here:)")
		}
	} else {
		response.SuccessResponse(c, "You do not have permission to access here.")
	}
	// test_map := make(map[string]string)
	// test_map["18613"] = "instructor"
	// test_map["18741"] = "instructor"
	// test_map["18749"] = "student"
	// test_map["19673"] = "course_assistant"
	// test_map["18989"] = "course_assistant"
	// test_map["39699"] = "student"
	// test_map["17637"] = "student"
	// fmt.Println(utils.Map_DBcheck(test_map))
}
