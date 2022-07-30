package validate

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/go-playground/validator/v10"
)

func AnswersUpdateValidation(sl validator.StructLevel) {
	answers := sl.Current().Interface().(dao.Answers_Upload_Validate)
	data := answers.Answers
	solutions := answers.Student.Solutions
	if len(answers.Answers) == 0 {
		return
	}
	if len(solutions) != len(answers.Answers) {
		sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "")
	}
	// 1st layers
	for i := range solutions {
		answer := data[i]
		solution := solutions[i]
		if solution.Key != answer.Key {
			sl.ReportError(answers.Answers, "data", "Data", "answerWrongName", answer.Key)
		}
		if len(solution.Value) != len(answer.Value) {
			sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "")
		}
		// sub layers
		for j := range solution.Value {
			sub_answer, sub_solution := answer.Value[j], solution.Value[j]
			if sub_solution.Key != sub_answer.Key {
				sl.ReportError(answers.Answers, "data", "Data", "answerWrongName", sub_answer.Key)
			}
			if len(sub_solution.Value) != len(sub_answer.Value) {
				sl.ReportError(answers.Answers, "data", "Data", "answerNotValid", "")
			}

			// sub sub layers
			for z := range sub_solution.Value {
				sub_sub_answer, sub_sub_solution := sub_answer.Value[z], sub_solution.Value[z]
				if sub_sub_solution.Key != sub_sub_answer.Key {
					sl.ReportError(answers.Answers, "data", "Data", "answerWrongName", sub_sub_answer.Key)
				}
			}
		}
	}
}
