package dao

import (
	"context"
	"encoding/json"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TakeExam(str1, str2, str3 string) []byte {
	client := global.Mongo
	// get the the context object，it can be used to set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//get the assessment_bank collection
	collection := client.Database("auto_exam").Collection("assessment_bank")

	//read exam config and use it to generate student questions
	var exam AutoExam_Assessments
	getExam(collection, &exam, str1, str2)
	//get the student_bank collection
	stuCollection := client.Database("auto_exam").Collection("student_bank")
	var student Student
	filter := bson.D{{Key: "course", Value: str1}, {Key: "examId", Value: str2}, {Key: "studentId", Value: str3}}
	//judge if we have stored this student; if not we should generate config for him
	err := stuCollection.FindOne(ctx, filter).Decode(&student)
	if err == mongo.ErrNoDocuments {
		//we should care for this poor student, it will take some time...
		// fmt.Println("student not exists, but gotcha!")
		genStudent(stuCollection, &exam, &student, str1, str2, str3)
	} else if err != nil {
		//the err we never want to see
		panic(err)
	}

	//generate questions json using student config
	//get the question_bank_new collection
	question_collection := client.Database("auto_exam").Collection("question_bank_new")
	container := getQuestions(question_collection, &student)
	return container
}

//get questions using student config
func getQuestions(collection *mongo.Collection, student *Student) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//declare some arguments
	var container Container1
	var questions []Header1
	var question Header1
	var filter bson.D
	settings := student.Settings

	for _, value := range settings {
		filter = bson.D{{Key: "headerId", Value: value}}
		cursor := collection.FindOne(ctx, filter)
		cursor.Decode(&question)
		questions = append(questions, question)
	}
	container.Data = questions
	jsonData, _ := json.Marshal(container)
	return jsonData
}

//the helper func for reading the assessment config from the mongodb
func getExam(collection *mongo.Collection, exam *AutoExam_Assessments, id1, id2 string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "course", Value: id1}, {Key: "id", Value: id2}}
	cursor := collection.FindOne(ctx, filter)
	cursor.Decode(exam)
}

//generate the student test config based on assessment and store it
func genStudent(collection *mongo.Collection, exam *AutoExam_Assessments, student *Student, id1, id2, id3 string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	list := exam.Settings
	settings := make(map[int]int)

	for key := range list {
		settings[key] = key + 1
	}
	student.CourseId = id1
	student.ExamId = id2
	student.StudentId = id3
	student.Settings = settings
	data, err := bson.Marshal(student)
	if err != nil {
		// fmt.Println("serialization fail")
		panic(err)
	}
	collection.InsertOne(ctx, data)
}
