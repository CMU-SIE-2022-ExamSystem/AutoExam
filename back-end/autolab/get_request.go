package autolab

import (
	"io/ioutil"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

func Userinfo_Handler(c *gin.Context, token string, refresh string) {
	autolab := global.Settings.Autolabinfo

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "http://"+autolab.Ip+"/api/v1/user", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		response.ErrUnauthResponse(c, "There may be something wrong with Autolab's web server, please try again later.")
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
	autolab := global.Settings.Autolabinfo

	user_token := utils.GetToken(c)
	user := models.User{ID: user_token.ID}
	global.DB.Find(&user)
	token := user.Access_token

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "http://"+autolab.Ip+"/api/v1//courses", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		response.ErrUnauthResponse(c, "There may be something wrong with Autolab's web server, please try again later.")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	autolab_resp := utils.User_courses_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}
