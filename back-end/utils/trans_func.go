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

func Course_assessments_trans(str string) []models.Autolab_Assessments {
	var response []models.Autolab_Assessments
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func Assessments_submissions_trans(str string) []models.Submissions {
	var response []models.Submissions
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}

	if len(response) != 0 {
		for i, submission := range response {
			sum := 0.0
			for index := range submission.Scores {
				sum += submission.Scores[index]
			}
			response[i].TotalScore = sum
		}
	}

	return response
}

func User_submit_trans(str string) models.Submit {
	var response models.Submit
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func Course_user_trans(str string) []models.Course_User_Data {
	var response []models.Course_User_Data
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func Course_user_err_trans(str string) models.Course_User_err {
	var response models.Course_User_err
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
	}
	return response
}

func Assessments_submissionscheck_trans(str string) ([]models.Submissions, bool) {
	var response []models.Submissions
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		color.Yellow("json transfer error>>> " + err.Error())
		return response, false
	}
	return response, true
}
