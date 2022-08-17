package utils

import (
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

func ExamNameFilter(body []models.Autolab_Assessments) []models.Autolab_Assessments {
	var result []models.Autolab_Assessments
	for i := 0; i < len(body); i++ {
		if strings.Contains(body[i].Category_name, dao.Assessment_Catergories[0]) {
			tmp := body[i]
			result = append(result, tmp)
		} else if strings.Contains(body[i].Category_name, dao.Assessment_Catergories[1]) {
			tmp := body[i]
			result = append(result, tmp)
		}
	}
	return result
}
