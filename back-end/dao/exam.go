package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetQuestions() []models.Header {
	collection := global.Mongo.Database("exam").Collection("question_bank")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var questions []models.Header
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var question models.Header
		err := cur.Decode(&question)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(question)
		questions = append(questions, question)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return questions
}
