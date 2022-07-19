package dao

import (
	"context"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	Que_Collection_Name string = "question_bank"
)

func GetQuestion(id int) ([]string, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{}}
	results, err := collection.Distinct(context.TODO(), "questionTag", filter)
	var tags []string = make([]string, len(results))
	for i, result := range results {
		tags[i] = result.(string)
	}
	return tags, err
}

// func GetQuestionByTag(tag_id string) ([]Question_Header, error) {

// }
