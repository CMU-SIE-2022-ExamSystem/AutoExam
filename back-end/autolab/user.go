package autolab

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Userinfo_Handler(c *gin.Context, autolab_resp models.Autolab_Response) {
	body := controller.Autolab_User_Handler(c, autolab_resp.Access_token, "/user")
	// fmt.Println(string(body))

	userinfo_resp := utils.User_info_trans(string(body))

	user, flag := controller.FindUserInfo(userinfo_resp.Email)

	if flag {
		color.Yellow("User is already in our DB!")
		user.Access_token = autolab_resp.Access_token
		user.Refresh_token = autolab_resp.Refresh_token
		user.Create_at = utils.GetNowTime()
		user.Expires_in = autolab_resp.Expires_in

		global.DB.Save(&user)
		jwt_token := controller.CreateToken(c, user.ID, user.Email)
		set_cookie(c, jwt_token)

		user_info := models.User_Info_Front{Token: jwt_token, First_name: user.First_name, Last_name: user.Last_name}
		response.SuccessResponse(c, user_info)
	} else {
		color.Yellow("User is not in our DB!")
		new_user := models.User{
			Email:         userinfo_resp.Email,
			First_name:    userinfo_resp.First_name,
			Last_name:     userinfo_resp.Last_name,
			Access_token:  autolab_resp.Access_token,
			Refresh_token: autolab_resp.Refresh_token,
			Create_at:     utils.GetNowTime(),
			Expires_in:    autolab_resp.Expires_in,
		}

		global.DB.Create(&new_user)
		jwt_token := controller.CreateToken(c, new_user.ID, new_user.Email)
		set_cookie(c, jwt_token)

		user_info := models.User_Info_Front{Token: jwt_token, First_name: new_user.First_name, Last_name: new_user.Last_name}
		response.SuccessResponse(c, user_info)
	}
}

func set_cookie(c *gin.Context, cookie string) {
	// c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("gin_cookie", "Bearer "+cookie, int(controller.Expire_time), "/", "", false, false)
}

func Usercourses_Handler(c *gin.Context) {
	user_email := controller.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := controller.Autolab_User_Handler(c, token, "/courses")
	fmt.Println(string(body))

	autolab_resp := utils.User_courses_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}
