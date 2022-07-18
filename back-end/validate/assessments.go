package validate

import (
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func AssessmentUpdateValidation(sl validator.StructLevel) {

	assessment := sl.Current().Interface().(dao.AutoExam_Assessments_Update)

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
	tags, err := dao.GetTags()
	if err != nil {
		panic(err)
	}

	// validate settings
	if len(assessment.Settings) > 0 {
		for i, setting := range assessment.Settings {
			if setting.Tag == "" {
				sl.ReportError(setting.Tag, "tag", "Tag", "noTag", utils.Ordinalize(i+1))
			}
			if !slices.Contains(tags, setting.Tag) {
				sl.ReportError(setting.Tag, "tag", "Tag", "notValidTag", utils.Ordinalize(i+1))
			}
			if setting.Max_score < 1 {
				sl.ReportError(setting.Max_score, "max_score", "Max_score", "maxscore", "1, "+utils.Ordinalize(i+1))
			}
			// TODO validate score, id and number of questions relationship
			if setting.Score != nil {
			}
		}
	}

}
