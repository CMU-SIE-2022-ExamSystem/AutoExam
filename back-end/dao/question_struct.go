package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

///the version for storing, with answer
type Choice struct {
	ChoiceId string `json:"choice_id" bson:"choice_id"`
	Content  string `json:"content" bson:"content"`
}

type Sub_Question_Blank struct {
	Grader      string     `json:"grader" bson:"grader"`           // sub question's grader
	Description string     `json:"description" bson:"description"` // sub question's content
	Choices     []Choice   `json:"choices" bson:"choices"`         // required for "choices" type sub question
	Solutions   [][]string `json:"solutions" bson:"solutions"`     // solutions of the sub question, the design for 2D slices is that the first dimension would be capable of multiple blanks while the second dimension would be used when if multiple solutions are all correct\n example: [["A", "B"], ["C"]]
	Blanks      []Blanks   `json:"blanks" bson:"blanks"`           // detail of blanks of sub question, based on grader
}

type AutoExam_Questions struct {
	ObjectID          primitive.ObjectID   `bson:"_id" json:"_id"`
	Title             string               `json:"title" bson:"title"`
	Description       string               `json:"description" bson:"description"`
	BaseCourse        string               `yaml:"base_course" json:"base_course" bson:"base_course" form:"base_course" binding:"required"`
	Tag               string               `json:"question_tag" bson:"question_tag"`
	SubQuestions      []Sub_Question_Blank `json:"sub_questions" bson:"sub_questions"`
	SubQuestionNumber int                  `json:"sub_question_number" bson:"sub_question_number"`
}

type AutoExam_Questions_Create struct {
	Title             string               `json:"title" bson:"title"`
	Description       string               `json:"description" bson:"description"`
	BaseCourse        string               `yaml:"base_course" json:"base_course" bson:"base_course" form:"base_course" binding:"required"`
	Tag               string               `json:"question_tag" bson:"question_tag"`
	SubQuestions      []Sub_Question_Blank `json:"sub_questions" bson:"sub_questions"`
	SubQuestionNumber int                  `json:"sub_question_number" bson:"sub_question_number"`
}

type Tags_Return struct {
	Tags []string `yaml:"tags" json:"tags"`
}

// @Description questions model info
type Questions struct {
	Id                string               `bson:"id" json:"id"`                                   // question id
	Title             string               `json:"title" bson:"title"`                             // question title
	Description       string               `json:"description" bson:"description"`                 // question content details
	Tag               string               `json:"question_tag" bson:"question_tag"`               // tag of the question, would return tag name
	SubQuestions      []Sub_Question_Blank `json:"sub_questions" bson:"sub_questions"`             // detail of sub_questions
	SubQuestionNumber int                  `json:"sub_question_number" bson:"sub_question_number"` // number of sub_questions
}

// @Description questions model info
type Questions_Create struct {
	Title        string         `json:"title" bson:"title"`                 // question title
	Description  string         `json:"description" bson:"description"`     // question content details
	Tag          string         `json:"question_tag" bson:"question_tag"`   // tag of the question, only accept the tag id
	SubQuestions []Sub_Question `json:"sub_questions" bson:"sub_questions"` // detail of sub_questions
}

type Sub_Question struct {
	Grader      string     `json:"grader" bson:"grader"`           // sub question's grader
	Description string     `json:"description" bson:"description"` // sub question's content
	Choices     []Choice   `json:"choices" bson:"choices"`         // required for "choices" type sub question
	Solutions   [][]string `json:"solutions" bson:"solutions"`     // solutions of the sub question, the design for 2D slices is that the first dimension would be capable of multiple blanks while the second dimension would be used when if multiple solutions are all correct\n example: [["A", "B"], ["C"]]
}

type Questions_Create_Validate struct {
	BaseCourse   string               `yaml:"base_course" json:"base_course" bson:"base_course" form:"base_course" binding:"required"`
	Title        string               `json:"title" bson:"title"`               // question title
	Description  string               `json:"description" bson:"description"`   // question content details
	Tag          string               `json:"question_tag" bson:"question_tag"` // tag of the question, only accept the tag id
	SubQuestions []Sub_Question_Blank `json:"sub_questions" bson:"sub_questions"`
}

func (autoexam *AutoExam_Questions) ToQuestions() Questions {
	questions := Questions{
		Id:                autoexam.ObjectID.Hex(),
		Title:             autoexam.Title,
		Description:       autoexam.Description,
		Tag:               ReadTagName(autoexam.Tag),
		SubQuestions:      autoexam.SubQuestions,
		SubQuestionNumber: autoexam.SubQuestionNumber,
	}
	return questions
}

func (question *Questions_Create_Validate) ToAutoExamQuestions() AutoExam_Questions_Create {
	instance := AutoExam_Questions_Create{
		Title:             question.Title,
		Description:       question.Description,
		Tag:               question.Tag,
		SubQuestions:      question.SubQuestions,
		BaseCourse:        question.BaseCourse,
		SubQuestionNumber: len(question.SubQuestions),
	}
	return instance
}

func (autoexam *Questions) ToQuestionsStudent(score float64, sub_scores []float64) Questions_Student {
	var sub_question []Sub_Question_Blank_Student
	for i, quest := range autoexam.SubQuestions {
		sub_question = append(sub_question, quest.ToSubQuestionsBlankStudent(sub_scores[i]))
	}

	questions := Questions_Student{
		Id:                autoexam.Id,
		Title:             autoexam.Title,
		Description:       autoexam.Description,
		Tag:               autoexam.Tag,
		SubQuestions:      sub_question,
		SubQuestionNumber: autoexam.SubQuestionNumber,
		Score:             score,
	}
	return questions
}

func (autoexam *Sub_Question_Blank) ToSubQuestionsBlankStudent(score float64) Sub_Question_Blank_Student {
	student := Sub_Question_Blank_Student{
		Description: autoexam.Description,
		Choices:     autoexam.Choices,
		Blanks:      autoexam.Blanks,
		Score:       score,
	}
	return student
}
