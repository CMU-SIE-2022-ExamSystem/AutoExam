package validate

import (
	"math"
	"strings"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

const float64EqualityThreshold = 1e-9

func AssessmentUpdateValidation(sl validator.StructLevel) {

	assessment := sl.Current().Interface().(dao.AutoExam_Assessments_Update_Validate)

	// validate general

	// validate time
	startTime, _ := time.Parse(models.TimeFormat, assessment.General.Start_at)
	endTime, _ := time.Parse(models.TimeFormat, assessment.General.End_at)
	gradingTime, _ := time.Parse(models.TimeFormat, assessment.General.Grading_deadline)
	if !startTime.Before(endTime) {
		sl.ReportError(assessment.General.Start_at, "start_at", "StartAt", "ltfield", "end_at")
	}
	if !(endTime.Before(gradingTime) || endTime.Equal(gradingTime)) {
		sl.ReportError(assessment.General.End_at, "end_at", "EndAt", "ltfield", "grading_deadline")
	}

	// validate max_submissions
	if assessment.General.Category_name == "Exam" {
		if assessment.General.MaxSubmissions != 1 {
			sl.ReportError(assessment.General.MaxSubmissions, "max_submissions", "MaxSubmissions", "submission", "1")
		}
	}

	// validate settings
	base_course := assessment.BaseCourse

	if len(assessment.Settings) > 0 {
		for i, setting := range assessment.Settings {
			// validate tag
			if setting.Tag == "" {
				sl.ReportError(setting.Tag, "tag", "Tag", "noTag", utils.Ordinalize(i+1))
				break
			}
			if status, _ := dao.ValidateTagById(base_course, setting.Tag); status {
				sl.ReportError(setting.Tag, "tag", "Tag", "notValidTag", base_course+","+utils.Ordinalize(i+1))
				break
			}

			// validate sub question number
			if setting.SubQuestionNumber == 0 {
				sl.ReportError(setting.SubQuestionNumber, "sub_question_number", "SubQuestionNumber", "gteSetting", "1,"+utils.Ordinalize(i+1))
				break
			}
			distinct_sub_number, distinct_sub_number_string := dao.GetAllSubQuestionNumber(base_course, setting.Tag)
			if !slices.Contains(distinct_sub_number, setting.SubQuestionNumber) {
				sl.ReportError(setting.SubQuestionNumber, "sub_question_number", "SubQuestionNumber", "notValidSub", strings.Join(distinct_sub_number_string, " ")+", "+utils.Ordinalize(i+1))
			}

			// validate max_score
			if setting.Max_score < 1 {
				sl.ReportError(setting.Max_score, "max_score", "Max_score", "gteSetting", "1,"+utils.Ordinalize(i+1))
			}

			// validate id
			if len(setting.Id) != 0 {
				for _, id := range setting.Id {
					quetsion, _ := dao.ReadAutoExamQuestionById(id)
					// validate id's tag
					if quetsion.Tag != setting.Tag {
						sl.ReportError(setting.Id, "id", "Id", "notValidIdTag", id+","+utils.Ordinalize(i+1))
					}
					// validate id's sub question number
					if quetsion.SubQuestionNumber != setting.SubQuestionNumber {
						sl.ReportError(setting.Id, "id", "Id", "notValidIdNumber", id+","+utils.Ordinalize(i+1))
					}
				}
			}

			// validate scores
			if len(setting.Scores) != 0 {
				// validate scores's length
				if len(setting.Scores) != setting.SubQuestionNumber {
					sl.ReportError(setting.Scores, "scores", "Scores", "notValidScoreLength", utils.Ordinalize(i+1))
				}
				// validate sum of scores
				sum := 0.0
				for i := range setting.Scores {
					sum += setting.Scores[i]
				}
				if math.Abs(sum-setting.Max_score) > float64EqualityThreshold {
					sl.ReportError(setting.Scores, "scores", "Scores", "notValidScore", utils.Ordinalize(i+1))
				}
			} else {
				// validate scores divisible
				var scores []float64
				score := math.Round((setting.Max_score/float64(setting.SubQuestionNumber))*100) / 100
				if math.Abs(score*float64(setting.SubQuestionNumber)-setting.Max_score) > float64EqualityThreshold {
					sl.ReportError(setting.Scores, "scores", "Scores", "numberNotDivisible", utils.Ordinalize(i+1))
					break
				}
				for i := 0; i < setting.SubQuestionNumber; i++ {
					scores = append(scores, score)
				}
				assessment.Settings[i].Scores = scores
			}
		}
	}

}
