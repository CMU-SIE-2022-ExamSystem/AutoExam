package validate

import (
	"fmt"
	"strconv"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/go-playground/validator/v10"
)

func QuestionsCreateValidation(sl validator.StructLevel) {
	question := sl.Current().Interface().(dao.Questions_Create_Validate)

	// validate tag
	if status, _ := dao.ValidateTagById(question.BaseCourse, question.Tag); status {
		sl.ReportError(question.Tag, "question_tag", "Tag", "notValidTag", question.BaseCourse)
	}

	grader_dict := dao.GetBasicGraderDict()
	fmt.Println(grader_dict)
	// validate sub questions
	for i, sub_question := range question.SubQuestions {
		grader_name := sub_question.Grader

		// validate required grader
		if grader_name == "" {
			sl.ReportError(sub_question.Grader, "grader", "Grader", "requiredGrader", utils.Ordinalize(i+1))
			break
		}
		// validate grader in the system or not
		if !dao.ValidateGrader(sub_question.Grader, question.BaseCourse) {
			sl.ReportError(sub_question.Grader, "grader", "Grader", "notValidGrader", question.BaseCourse+","+utils.Ordinalize(i+1))
			break
		}

		// validate solutions number equal to grader
		_, ok := grader_dict[grader_name]
		if !ok {
			grader, _ := dao.ReadGrader(grader_name, question.BaseCourse)
			grader_dict[grader_name] = grader
		}
		if grader_dict[grader_name].BlanksNumber != len(sub_question.Solutions) {
			sl.ReportError(sub_question.Solutions, "solutions", "Solutions", "lenAnswer", grader_name+","+strconv.Itoa(grader_dict[grader_name].BlanksNumber)+","+utils.Ordinalize(i+1))
			break
		}

		question.SubQuestions[i].Blanks = grader_dict[grader_name].Blanks

		// check whether there is a choice sub blank
		is_choice := false
		for _, blank := range grader_dict[grader_name].Blanks {
			if blank.IsChoice {
				is_choice = true
			}
		}
		if is_choice {
			if len(sub_question.Choices) == 0 {
				sl.ReportError(sub_question.Choices, "choices", "Choices", "requiredChoice", utils.Ordinalize(i+1))
				break
			}
			if len(sub_question.Choices) != len(grader_dict[grader_name].Blanks) {
				sl.ReportError(sub_question.Choices, "choices", "Choices", "lenChoice", utils.Ordinalize(i+1)+","+strconv.Itoa(len(grader_dict[grader_name].Blanks)))
				break
			}
			for j, blank := range grader_dict[grader_name].Blanks {
				if blank.IsChoice {
					if len(sub_question.Choices[i]) == 0 {
						sl.ReportError(sub_question.Choices, "choices", "Choices", "notValidChoiceZero", utils.Ordinalize(i+1)+","+utils.Ordinalize(j+1))
					}
				} else {
					if len(sub_question.Choices[i]) != 0 {
						sl.ReportError(sub_question.Choices, "choices", "Choices", "notValidChoiceNotZero", utils.Ordinalize(i+1)+","+utils.Ordinalize(j+1))
					}
				}
			}
		} else {
			if len(sub_question.Choices) != 0 {
				sl.ReportError(sub_question.Choices, "choices", "Choices", "notRequiredChoice", utils.Ordinalize(i+1))
				break
			}
		}
	}

}
