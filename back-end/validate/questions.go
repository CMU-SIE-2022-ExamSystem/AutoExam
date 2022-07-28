package validate

import (
	"strconv"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
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

	grader_dict := course.GetBasicGraderLenDict()

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
			grader_dict[grader_name] = grader.BlanksNumber
		}
		if grader_dict[grader_name] != len(sub_question.Solutions) {
			sl.ReportError(sub_question.Solutions, "solutions", "Solutions", "lenAnswer", grader_name+","+strconv.Itoa(grader_dict[grader_name])+","+utils.Ordinalize(i+1))
		}

		// validate choices required for "choices" grader
		if strings.Contains(grader_name, "choice") {
			if len(sub_question.Choices) == 0 {
				sl.ReportError(sub_question.Choices, "choices", "Choices", "requiredChoice", utils.Ordinalize(i+1))
			}
		}

	}

}
