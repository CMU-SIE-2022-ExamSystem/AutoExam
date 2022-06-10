package autolab

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

var (
	autolab_url     string = global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url string = autolab_url + "/api/v1"
)

func Userinfo_Handler(c *gin.Context, token string, refresh string) {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_api_url+"/user", nil)
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
	user_token := utils.GetToken(c)
	user := models.User{ID: user_token.ID}
	global.DB.Find(&user)
	token := user.Access_token

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_api_url+"/courses", nil)
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

func refresh_Handler(c *gin.Context) {
	user_token := utils.GetToken(c)
	user := models.User{ID: user_token.ID}
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

	autolab_resp := Autolab_Auth_Handler(c, http_body)

	user.Access_token = autolab_resp.Access_token
	user.Refresh_token = autolab_resp.Refresh_token
	global.DB.Save(&user)
}

func Autolab_Auth_Handler(c *gin.Context, http_body interface{}) models.Autolab_Response {
	resp_body, _ := json.Marshal(http_body)
	resp, err := http.Post(autolab_url+"/oauth/token", "application/json", bytes.NewBuffer(resp_body))

	if err != nil {
		Autolab_Error_Hander(c, resp, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	if strings.Contains(string(body), "error") {
		err_response := utils.Autolab_err_trans(string(body))
		response.ErrUnauthResponse(c, err_response.Error_description)
		c.Abort()
	}

	autolab_resp := utils.Autolab_resp_trans(string(body))
	return autolab_resp
}

func Autolab_Error_Hander(c *gin.Context, resp *http.Response, err error) {
	if err != nil {
		response.ErrUnauthResponse(c, "There may be something wrong with Autolab's web server, please try again later.")
	}

	status := resp.StatusCode
	if status >= http.StatusOK && status <= 299 {
		return
	} else {
		if status == http.StatusUnauthorized {
		} else {
			response.ErrorInternalResponse(c, response.Error{Type: "Autolab", Message: "Unknown error"})
		}
	}
}
