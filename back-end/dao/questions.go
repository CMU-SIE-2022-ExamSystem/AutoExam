package dao

import (
	"context"
	"strconv"
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

func ReadAllQuestionsByTag(base_course, tag_id string) ([]Questions, error) {
	if tag_id == "" {
		return ReadAllQuestions(base_course)
	}

	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "question_tag", Value: tag_id}}

	cursor, err := collection.Find(context.TODO(), filter)

	var questions []Questions

	for cursor.Next(context.TODO()) {
		var autoexam AutoExam_Questions
		cursor.Decode(&autoexam)
		questions = append(questions, autoexam.ToQuestions())
	}
	return questions, err
}

func CreateQuestion(question AutoExam_Questions_Create) (Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := bson.Marshal(question)
	if err != nil {
		return Questions{}, err
	}

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return Questions{}, err
	}
	return ReadQuestionById(result.InsertedID.(primitive.ObjectID).Hex())
}

func UpdateQuestions(id string, question AutoExam_Questions_Create) error {
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

// return true for no such object in mongo
func ValidateQuestionById(course, id string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}, {Key: "base_course", Value: course}}

	var question AutoExam_Questions
	err := collection.FindOne(context.TODO(), filter).Decode(&question)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, err
}

// TODO finish it after assessment is done
// return true for safe delete
func ValidateQuestionUsedById(id string) (bool, error) {
	// client := global.Mongo
	//get the collection instance
	// collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	// filter := bson.D{{Key: "questionTag", Value: id}}
	// var questions AutoExam_Questions
	// err := collection.FindOne(context.TODO(), filter).Decode(&questions)
	// if err != nil {
	// 	// ErrNoDocuments means that the filter did not match any documents in the collection
	// 	if err == mongo.ErrNoDocuments {
	// 		return true, nil
	// 	}
	// 	return false, err
	// }
	return false, nil
}

func GetAllSubQuestionNumber(base_course, tag_id string) ([]int, []string) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}, {Key: "question_tag", Value: tag_id}}

	results, _ := collection.Distinct(context.TODO(), "sub_question_number", filter)

	var numbers []int
	var numbers_text []string
	for _, result := range results {
		numbers = append(numbers, int(result.(int32)))
		numbers_text = append(numbers_text, strconv.Itoa(int(result.(int32))))
	}
	return numbers, numbers_text
}

func GetAllQuestionIDBySubQuestionNumber(base_course, tag_id string, sub_question_number int) []string {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}, {Key: "question_tag", Value: tag_id}, {Key: "sub_question_number", Value: sub_question_number}}

	cursor, _ := collection.Find(context.TODO(), filter)

	var id []string

	for cursor.Next(context.TODO()) {
		var autoexam AutoExam_Questions
		cursor.Decode(&autoexam)
		id = append(id, autoexam.ToQuestions().Id)
	}
	return id
}
