package controller

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func FindUserInfo(email string) (*models.User, bool) {
	var user models.User
	rows := global.DB.Where(&models.User{Email: email}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func Userrefresh_Handler(c *gin.Context) {
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

	autolab_resp, flag := Autolab_Auth_Handler(c, "/oauth/token", http_body)

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
