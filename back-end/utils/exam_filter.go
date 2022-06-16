package utils

import (
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

func ExamNameFilter(body []models.Assessments) []models.Assessments {
	var result []models.Assessments
	for i := 0; i < len(body); i++ {
		if strings.Contains(body[i].Category_name, "Exam") {
			tmp := body[i]
			result = append(result, tmp)
		} else if strings.Contains(body[i].Category_name, "Quiz") {
			tmp := body[i]
			result = append(result, tmp)
		}
	}
	return result
}
