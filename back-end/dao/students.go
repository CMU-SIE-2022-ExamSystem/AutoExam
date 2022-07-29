package dao

import (
	"context"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	Student_Collection_Name = "student_bank"
)

type Assessment_Student struct {
	Id         int                `json:"id" bson:"id"`
	Assessment string             `json:"assessment" bson:"assessment"`
	Course     string             `json:"course" bson:"course"`
	Questions  []string           `json:"questions" bson:"questions"`
	Problems   []Student_Problems `json:"problems" bson:"problems"`
	Answers    interface{}        `json:"answers" bson:"answers"`
	Solutions  interface{}        `json:"solutions" bson:"solutions"`
}

type Student_Problems struct {
	Name     string  `yaml:"name" json:"name" bson:"name"`
	Grader   string  `yaml:"grader" json:"grader" bson:"grader"` // grader name
	MaxScore float64 `yaml:"max_score" json:"max_score" bson:"max_score"`
}

// create student
func CreateStudent(student Assessment_Student) (Assessment_Student, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Student_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := bson.Marshal(student)
	if err != nil {
		return Assessment_Student{}, err
	}

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return Assessment_Student{}, err
	}
	return ReadStudent(student.Course, student.Assessment, student.Id)
}

// read student by course, assessment, id
func ReadStudent(course, assessment string, id int) (Assessment_Student, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Student_Collection_Name)

	filter := bson.D{{Key: "course", Value: course}, {Key: "id", Value: id}, {Key: "assessment", Value: assessment}}
	var student Assessment_Student
	err := collection.FindOne(context.TODO(), filter).Decode(&student)
	return student, err
}
