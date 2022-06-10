package utils

import (
	"encoding/json"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/fatih/color"
)

func Autolab_err_trans(str string) models.Autolab_err_Response {
	var response models.Autolab_err_Response
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func Autolab_resp_trans(str string) models.Autolab_Response {
	var response models.Autolab_Response
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func User_info_trans(str string) models.User_Info {
	var response models.User_Info
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func User_courses_trans(str string) []models.User_Courses {
	var response []models.User_Courses
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}
