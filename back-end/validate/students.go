package validate

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/go-playground/validator/v10"
)

func AnswersUpdateValidation(sl validator.StructLevel) {
	answers := sl.Current().Interface().(dao.Answers_Upload_Validate)
	answer := answers.Answers
	solutions := answers.Student.Solutions
	if len(answer) == 0 {
		return
	}
	if len(solutions) != len(answer) {
		sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "question")
	}

	// 1st layers
	for i := range solutions {
		if _, found := answer[i]; !found {
			sl.ReportError(answer, "data", "Data", "noKey", i)
		}
		if len(solutions[i]) != len(answer[i]) {
			sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "sub_question")
		}
		for j := range solutions[i] {

			if _, found := answer[i][j]; !found {
				sl.ReportError(answers.Answers, "data", "Data", "noKey", j)
			}
			if len(solutions[i][j]) != len(answer[i][j]) {
				sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "subsub__question")
			}
			for z := range solutions[i][j] {
				if _, found := answer[i][j][z]; !found {
					sl.ReportError(answers.Answers, "data", "Data", "noKey", z)
				}
			}
		}
	}
}
