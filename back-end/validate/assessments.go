package validate

import (
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/go-playground/validator/v10"
)

func AssessmentUpdateValidation(sl validator.StructLevel) {

	assessment := sl.Current().Interface().(dao.AutoExam_Assessments_Update)

	// if len(user.FirstName) == 0 && len(user.LastName) == 0 {
	// 	sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
	// 	sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	// }
	startTime, _ := time.Parse(models.TimeFormat, assessment.General.Start_at)
	endTime, _ := time.Parse(models.TimeFormat, assessment.General.End_at)
	gradingTime, _ := time.Parse(models.TimeFormat, assessment.General.Grading_deadline)

	if !startTime.Before(endTime) {
		sl.ReportError(assessment.General.Start_at, "start_at", "StartAt", "ltfield", "end_at")
	}
	if !(endTime.Before(gradingTime) || endTime.Equal(gradingTime)) {
		sl.ReportError(assessment.General.End_at, "end_at", "EndAt", "ltfield", "grading_deadline")
	}

	if assessment.General.Category_name == "Exam" {
		if assessment.General.MaxSubmissions != 1 {
			sl.ReportError(assessment.General.MaxSubmissions, "max_submissions", "MaxSubmissions", "submission", "1")
		}
	}

	// TODO maybe should validate Settings and Problems
}
