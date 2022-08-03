package dao

import (
	"context"
	"strconv"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Que_Collection_Name string = "question_bank"
)

func ReadAllQuestions(base_course string, hidden bool) ([]Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	var filter primitive.D
	if hidden {
		filter = bson.D{{Key: "base_course", Value: base_course}}
	} else {
		filter = bson.D{{Key: "base_course", Value: base_course}, {Key: "hidden", Value: false}}
	}

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

func ReadOrgQuestionById(id string) (AutoExam_Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}

	var question AutoExam_Questions
	err := collection.FindOne(context.TODO(), filter).Decode(&question)
	return question, err
}

func ReadAllQuestionsByTag(base_course, tag_id string, hidden bool) ([]Questions, error) {
	if tag_id == "" {
		return ReadAllQuestions(base_course, hidden)
	}

	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	var filter primitive.D
	if hidden {
		filter = bson.D{{Key: "question_tag", Value: tag_id}}
	} else {
		filter = bson.D{{Key: "question_tag", Value: tag_id}, {Key: "hidden", Value: false}}
	}

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

// create or update Question
func CreateOrUpdateQuestions(question AutoExam_Questions) (AutoExam_Questions, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: question.ObjectID}}
	opts := options.Replace().SetUpsert(true)
	data, err := bson.Marshal(question)
	if err != nil {
		return AutoExam_Questions{}, err
	}
	_, err = collection.ReplaceOne(ctx, filter, data, opts)
	if err != nil {
		return AutoExam_Questions{}, err
	}
	return ReadOrgQuestionById(question.ToQuestions().Id)
}

func DeleteQuestionById(id string, hard bool) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}
	var err error
	if hard {
		_, err = collection.DeleteOne(context.TODO(), filter)
	} else {
		var question AutoExam_Questions
		question, _ = ReadOrgQuestionById(id)
		question.Hidden = true
		_, err = CreateOrUpdateQuestions(question)
	}
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

// return true for safe delete or update
func ValidateQuestionUsedById(id string) (bool, error) {

	client := global.Mongo
	// get the collection instance
	collection := client.Database("auto_exam").Collection(Student_Collection_Name)
	filter := bson.D{{Key: "questions", Value: id}}

	var student Assessment_Student
	err := collection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
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

// only show not hidden
func GetAllQuestionIDBySubQuestionNumber(base_course, tag_id string, sub_question_number int) ([]string, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}, {Key: "question_tag", Value: tag_id}, {Key: "sub_question_number", Value: sub_question_number}, {Key: "hidden", Value: false}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []string{}, err
	}

	var id []string

	for cursor.Next(context.TODO()) {
		var autoexam AutoExam_Questions
		err := cursor.Decode(&autoexam)
		if err != nil {
			return []string{}, err
		}
		id = append(id, autoexam.ToQuestions().Id)
	}
	return id, err
}
