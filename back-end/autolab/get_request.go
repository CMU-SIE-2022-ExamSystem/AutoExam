package autolab

import (
	"io/ioutil"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func autolab_api_url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + "/api/v1" + endpoint
	return autolab_api_url
}

func Userinfo_Handler(c *gin.Context, token string, refresh string) {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_api_url("/user"), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		Autolab_Error_Hander(c, resp, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	autolab_resp := utils.User_info_trans(string(body))

	user := models.User{
		Email:         autolab_resp.Email,
		First_name:    autolab_resp.First_name,
		Last_name:     autolab_resp.Last_name,
		Access_token:  token,
		Refresh_token: refresh,
	}

	global.DB.Create(&user)
	jwt_token := utils.CreateToken(c, user.ID, token)
	response.SuccessResponse(c, gin.H{
		"token":     jwt_token,
		"firstName": autolab_resp.First_name,
		"lastName":  autolab_resp.Last_name,
	})
}

func Usercourses_Handler(c *gin.Context) {
	user_token := utils.GetEmail(c)
	user := models.User{ID: user_token.ID}
	global.DB.Find(&user)
	token := user.Access_token

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_api_url("/courses"), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		Autolab_Error_Hander(c, resp, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	autolab_resp := utils.User_courses_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}

func Refresh_Handler(c *gin.Context) {
	user_email := utils.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	// token := user.Access_token
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
		color.Yellow(user.Access_token)
		color.Yellow(user.Refresh_token)
		global.DB.Save(&user)
	}
}
