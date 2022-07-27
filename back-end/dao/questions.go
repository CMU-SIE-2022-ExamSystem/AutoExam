package dao

import (
	"context"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Que_Collection_Name string = "question_bank"
)

func ReadAllQuestions(base_course string) ([]Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}}

	cursor, err := collection.Find(context.TODO(), filter)

	var questions []Questions

	for cursor.Next(context.TODO()) {
		var autoexam AutoExam_Questions
		cursor.Decode(&autoexam)
		questions = append(questions, autoexam.ToQuestions())
	}
	return questions, err
}

func CreateQuestion(question Questions) (result *mongo.InsertOneResult, err error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := bson.Marshal(question)
	if err != nil {
		return nil, err
	}

	result, err = collection.InsertOne(ctx, data)
	return
}

func ReadQuestionById(id string) (Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}

	var tags AutoExam_Questions
	err := collection.FindOne(context.TODO(), filter).Decode(&tags)
	return tags.ToQuestions(), err
}

func UpdateQuestions(id string, question Questions_Create_Validate) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}

	data, err := bson.Marshal(question)
	if err != nil {
		return err
	}

	_, err = collection.ReplaceOne(ctx, filter, data)
	return err
}

func DeleteQuestionById(id string) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}
	_, err := collection.DeleteOne(context.TODO(), filter)

	return err
}
