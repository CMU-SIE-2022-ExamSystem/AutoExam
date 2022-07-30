package dao

import (
	"context"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Student_Collection_Name = "student_bank"
)

type Assessment_Student struct {
	Email      string              `json:"email" bson:"email"`
	Assessment string              `json:"assessment" bson:"assessment"`
	Course     string              `json:"course" bson:"course"`
	Questions  []string            `json:"questions" bson:"questions"`
	Problems   []Student_Problems  `json:"problems" bson:"problems"`
	Answers    []Student_Questions `json:"answers" bson:"answers"`
	Solutions  []Student_Questions `json:"solutions" bson:"solutions"`
}

type Student_Questions struct {
	Key   string                  `json:"key" bson:"key"`
	Value []Student_Sub_Questions `json:"value" bson:"value"`
}

type Student_Sub_Questions struct {
	Key   string                      `json:"key" bson:"key"`
	Value []Student_Sub_Sub_Questions `json:"value" bson:"value"`
}
type Student_Sub_Sub_Questions struct {
	Key   string   `json:"key" bson:"key"`
	Value []string `json:"value" bson:"value"`
}

type Student_Problems struct {
	Name     string  `yaml:"name" json:"name" bson:"name"`
	Grader   string  `yaml:"grader" json:"grader" bson:"grader"` // grader name
	MaxScore float64 `yaml:"max_score" json:"max_score" bson:"max_score"`
}

// @Description questions model info
type Questions_Student struct {
	Id                string                       `bson:"id" json:"-"`                                    // question id
	Title             string                       `json:"title" bson:"title"`                             // question title
	Description       string                       `json:"description" bson:"description"`                 // question content details
	Tag               string                       `json:"question_tag" bson:"question_tag"`               // tag of the question, would return tag name
	SubQuestions      []Sub_Question_Blank_Student `json:"sub_questions" bson:"sub_questions"`             // detail of sub_questions
	SubQuestionNumber int                          `json:"sub_question_number" bson:"sub_question_number"` // number of sub_questions
}

type Sub_Question_Blank_Student struct {
	Description string   `json:"description" bson:"description"` // sub question's content
	Choices     []Choice `json:"choices" bson:"choices"`         // required for "choices" type sub question
	Blanks      []Blanks `json:"blanks" bson:"blanks"`           // detail of blanks of sub question, based on grader
}

type Answers_Upload struct {
	Answers []Student_Questions `json:"answers"`
}

type Answers_Upload_Validate struct {
	Student Assessment_Student
	Answers []Student_Questions `json:"answers" binding:"required"`
}

// create student
func CreateOrUpdateStudent(student Assessment_Student) (Assessment_Student, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Student_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "course", Value: student.Course}, {Key: "email", Value: student.Email}, {Key: "assessment", Value: student.Assessment}}
	opts := options.Replace().SetUpsert(true)
	data, err := bson.Marshal(student)
	if err != nil {
		return Assessment_Student{}, err
	}
	_, err = collection.ReplaceOne(ctx, filter, data, opts)
	if err != nil {
		return Assessment_Student{}, err
	}
	return ReadStudent(student.Course, student.Assessment, student.Email)
}

// read student by course, assessment, id
func ReadStudent(course, assessment, email string) (Assessment_Student, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Student_Collection_Name)

	filter := bson.D{{Key: "course", Value: course}, {Key: "email", Value: email}, {Key: "assessment", Value: assessment}}
	var student Assessment_Student
	err := collection.FindOne(context.TODO(), filter).Decode(&student)
	return student, err
}

func (student *Assessment_Student) ToRealQuestions() []Questions_Student {
	var questions_student []Questions_Student
	assessment, _ := ReadExam(student.Course, student.Assessment)
	for i, id := range student.Questions {
		autoexam, _ := ReadQuestionById(id)
		question := autoexam.ToQuestionsStudent()
		question.Title = assessment.Settings[i].Title
		questions_student = append(questions_student, question)
	}
	return questions_student
}

func (student *Assessment_Student) ToAnwerStruct() {
	solutions := student.Solutions
	for i := range solutions {
		// sub layers
		for j := range solutions[i].Value {
			// sub sub layers
			for z := range solutions[i].Value[j].Value {
				var value []string
				value = append(value, "")
				solutions[i].Value[j].Value[z].Value = value
			}
		}
	}
}
