package autolab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func user_info_trans(str string) user_Info {
	var response user_Info
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("json transfer error>>> ", err)
	}
	return response
}

func user_courses_trans(str string) []user_Courses {
	var response []user_Courses
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("json transfer error>>> ", err)
	}
	return response
}

func Userinfo_Handler(c *gin.Context, token string) {
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

	autolab_resp := user_info_trans(string(body))
	//todo store user information into database

	response.SuccessResponse(c, autolab_resp)
}

func Usercourses_Handler(c *gin.Context) {
	autolab := global.Settings.Autolabinfo

	//todo: get token from database
	token := c.Query("token")

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

	autolab_resp := user_courses_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}
