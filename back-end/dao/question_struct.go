package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

///the version for storing, with answer
type Choice struct {
	ChoiceId string `json:"choice_id" bson:"choiceId"`
	Content  string `json:"content" bson:"content"`
}

// sub question with answer
type Sub_Question struct {
	Type        string     `json:"question_type" bson:"questionType"`
	QuestionId  int        `json:"question_id" bson:"questionId"`
	Description string     `json:"description" bson:"description"`
	Choices     []Choice   `json:"choices" bson:"choices"`
	Answers     [][]string `json:"answers" bson:"answers"` // answer is changed into an array of strings
}

// question header with answer
type AutoExam_Questions struct {
	ObjectID    primitive.ObjectID `bson:"_id" json:"_id"`
	Course      string             `yaml:"base_course" json:"base_course" bson:"base_course" form:"base_course" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Tag         string             `json:"question_tag" bson:"questionTag"`
	Questions   []Sub_Question     `json:"questions" bson:"questions"`
}

type AutoExam_Questions_Create struct {
	Course      string         `yaml:"course" json:"course" bson:"course" form:"course" binding:"required"`
	Description string         `json:"description" bson:"description"`
	Tag         string         `json:"question_tag" bson:"questionTag"`
	Questions   []Sub_Question `json:"questions" bson:"questions"`
}

type Tags_Return struct {
	Tags []string `yaml:"tags" json:"tags"`
}

func (autoexam *AutoExam_Questions) ToQuestions() Questions {
	questions := Questions{
		Id:          autoexam.ObjectID.Hex(),
		Description: autoexam.Description,
		Tag:         autoexam.Tag,
		Questions:   autoexam.Questions,
	}
	return questions
}

type Questions struct {
	Id          string         `bson:"id" json:"id"`
	Description string         `json:"description" bson:"description"`
	Tag         string         `json:"question_tag" bson:"questionTag"`
	Questions   []Sub_Question `json:"questions" bson:"questions"`
}

type Questions_Create struct {
	Description string         `json:"description" bson:"description"`
	Tag         string         `json:"question_tag" bson:"questionTag"`
	Questions   []Sub_Question `json:"questions" bson:"questions"`
}

type Questions_Create_Validate struct {
	Course      string         `yaml:"base_course" json:"base_course" bson:"base_course" form:"base_course" binding:"required"`
	Description string         `json:"description" bson:"description"`
	Tag         string         `json:"question_tag" bson:"questionTag" binding:"required"`
	Questions   []Sub_Question `json:"questions" bson:"questions" binding:"required"`
}
