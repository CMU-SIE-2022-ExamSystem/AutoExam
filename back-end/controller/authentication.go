package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

// Authinfo_Handler godoc
// @Summary get auth info
// @Schemes
// @Description get autolab authentication info
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Autolab_Info_Front} "desc"
// @Router /auth/info [get]
func Authinfo_Handler(c *gin.Context) {
	auth := global.Settings.Autolabinfo
	user_front := models.Autolab_Info_Front{Client_id: auth.Client_id, Scope: auth.Scope}
	response.SuccessResponseWithType(c, user_front, 1)
}

// Authtoken_Handler godoc
// @Summary get auth token
// @Schemes
// @Description get autolab authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.Auth_Code true "body data"
// @Success 200 {object} response.Response{data=models.User_Info_Front} "desc"
// @Router /auth/token [post]
func Authtoken_Handler(c *gin.Context) {
	auth := global.Settings.Autolabinfo

	body := models.Auth_Code{}
	c.BindJSON(&body)

	http_body := models.Http_Body{
		Grant_type:    "authorization_code",
		Code:          body.Code,
		Redirect_uri:  auth.Redirect_uri,
		Client_id:     auth.Client_id,
		Client_secret: auth.Client_secret,
	}

	autolab_resp, flag := autolab.AutolabAuthHandler(c, "/oauth/token", http_body)

	if flag {
		Userinfo_Handler(c, autolab_resp)
	}
}
