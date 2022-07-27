package validate

import (
	"fmt"
	"reflect"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/go-playground/validator/v10"
)

func QuestionsCreateValidation(sl validator.StructLevel) {
	question := sl.Current().Interface().(dao.Questions_Create_Validate)

	// TODO validate question_type & number & answers
	for i, sub_question := range question.Questions {
		SubQuestionValidate(sub_question, i, sl)
	}

	// validate tag
	if status, _ := dao.ValidateTagById(question.Course, question.Tag); status {
		sl.ReportError(question.Tag, "question_tag", "Tag", "notValidTag", question.Course)
	}
	fmt.Println(question)
}

func SubQuestionValidate(question dao.Sub_Question, i int, sl validator.StructLevel) {
	// validate single and choices
	// if strings.Contains(question.Type, "single") || strings.Contains(question.Type, "choice") {
	// 	if len(question.Blanks) != 1 {
	// 		sl.ReportError(question.Blanks, "blanks", "Blanks", "singleBlank", question.Type+","+utils.Ordinalize(i+1))
	// 	}
	// }

	fmt.Println(reflect.TypeOf(question.Answers))
	// validate answers type
	if reflect.TypeOf(question.Answers).Elem().Kind() != reflect.Slice {
		sl.ReportError(question.Answers, "answers", "Answers", "notValidAnswers", utils.Ordinalize(i+1))
	}
	// for _, blank := range question.Blanks {
	// 	if !(blank.Type == "string" || blank.Type == "code") {
	// 		sl.ReportError(question.Blanks, "type", "Type", "oneof", "string code")
	// 	}
	// }
}
