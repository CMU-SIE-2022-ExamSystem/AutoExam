// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
)

func main() {

	initialize.SetupServer()

	// register task
	global.Redis.Register("generate", generate)

	// start workers (non-blocking call)
	global.Redis.StartWorker()
	fmt.Println("worker start")

	// loop forever to wait client request
	select {}
}

func generate(course_name, assessment_name, email string) bool {
	fmt.Println("generate task start")
	assessment := get_assessments(course_name, assessment_name)
	token := refresh_token(email)
	users, _ := coures_user_data(course_name, token)

	var err error
	for _, user := range users {
		student := assessment.GenerateAssessmentStudent(user.Email, course_name, assessment_name)
		_, err = dao.CreateOrUpdateStudent(student)
		if err != nil {
			assessment.Generated = -1
			assessment.GeneratedError = "There is an error happen when generating " + student.Email + "'s exam with error message: " + err.Error()
		}
	}
	if err == nil {
		assessment.Generated = 1
		assessment.GeneratedError = ""
	}
	assessment.Statistic = dao.Statistic{}
	dao.UpdateExam(course_name, assessment_name, assessment)
	fmt.Println("generate task finish")
	return err == nil
}

func coures_user_data(course_name, token string) ([]models.Course_User_Data, error) {
	var users []models.Course_User_Data
	body := autolab_get(token, "/courses/"+course_name+"/course_user_data")

	if strings.Contains(string(body), "error") {
		return users, errors.New(string(body))
	} else {
		users := utils.Course_user_trans(string(body))
		return users, nil
	}
}

func get_assessments(course_name, assessment_name string) dao.AutoExam_Assessments {
	// read certain assessment
	assessment, _ := dao.ReadExam(course_name, assessment_name)
	return assessment
}

func autolab_api_url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + "/api/v1" + endpoint
	return autolab_api_url
}

func autolab_get(token string, endpoint string) []byte {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_api_url(endpoint), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, _ := client.Do(request)
	fmt.Println(resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func refresh_token(email string) string {
	user := models.User{}
	global.DB.Where("email = ?", email).Find(&user)
	refresh := user.Refresh_token

	auth := global.Settings.Autolabinfo

	http_body := models.Http_Body_Refresh{
		Grant_type:    "refresh_token",
		Refresh_token: refresh,
		Scope:         auth.Scope,
		Client_id:     auth.Client_id,
		Client_secret: auth.Client_secret,
	}

	autolab_resp, flag := autolab_auth("/oauth/token", http_body)

	if flag {
		user.Access_token = autolab_resp.Access_token
		user.Refresh_token = autolab_resp.Refresh_token
		user.Create_at = utils.GetNowTime()
		user.Expires_in = autolab_resp.Expires_in
		color.Yellow(user.Access_token)
		color.Yellow(user.Refresh_token)
		global.DB.Save(&user)
	} else {
		log.Println("error happens when refresh token")
	}
	return autolab_resp.Access_token
}

func autolab_Oauth_Url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + endpoint
	return autolab_api_url
}

func autolab_auth(endpoint string, http_body interface{}) (models.Autolab_Response, bool) {
	resp_body, err := json.Marshal(http_body)
	if err != nil {
		log.Println(err.Error())
	}

	resp, err := http.Post(autolab_Oauth_Url(endpoint), "application/json", bytes.NewBuffer(resp_body))
	if err != nil {
		log.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	var autolab_resp models.Autolab_Response
	if strings.Contains(string(body), "error") {
		err_response := utils.Autolab_err_trans(string(body))
		log.Printf(err_response.Error_description)
	}

	autolab_resp = utils.Autolab_resp_trans(string(body))

	return autolab_resp, true
}
