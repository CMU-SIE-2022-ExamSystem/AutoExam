package autolab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/error"
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

func Userinfo_Handler(c *gin.Context) {
	autolab := Read_Autolab_Env()

	//todo: get token from mysql
	token := c.Query("token")

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "http://"+autolab.Ip+"/api/v1/user", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		error.ErrorHandler(c, error.AutolabResponse())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	response := user_info_trans(string(body))

	c.JSON(http.StatusOK, response)
}

func Usercourses_Handler(c *gin.Context) {
	autolab := Read_Autolab_Env()

	//todo: get token from mysql
	token := c.Query("token")

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "http://"+autolab.Ip+"/api/v1//courses", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		error.ErrorHandler(c, error.AutolabResponse())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	response := user_courses_trans(string(body))

	c.JSON(http.StatusOK, response)
}
