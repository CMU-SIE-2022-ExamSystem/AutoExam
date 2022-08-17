package dao

import (
	"context"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Ass_Collection_Name string = "assessment_bank"
)

func CreateExam(exam AutoExam_Assessments) (result *mongo.InsertOneResult, err error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := bson.Marshal(exam)
	if err != nil {
		panic(err)
	}

	result, err = collection.InsertOne(ctx, data)
	return
}

func GetAllExams(course string) ([]models.Assessments, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)

	filter := bson.D{{Key: "course", Value: course}}

	cursor, err := collection.Find(context.TODO(), filter)

	var assessments []models.Assessments

	for cursor.Next(context.TODO()) {
		var assessment AutoExam_Assessments
		cursor.Decode(&assessment)
		assessments = append(assessments, assessment.ToAssessments())
	}

	return assessments, err
}

func ReadExam(course, assessment_name string) (AutoExam_Assessments, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)

	filter := bson.D{{Key: "course", Value: course}, {Key: "id", Value: assessment_name}}
	var assessment AutoExam_Assessments
	err := collection.FindOne(context.TODO(), filter).Decode(&assessment)
	return assessment, err
}

func UpdateExam(course, assessment_name string, exam AutoExam_Assessments) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "course", Value: course}, {Key: "id", Value: assessment_name}}
	data, err := bson.Marshal(exam)
	if err != nil {
		return err
	}

	_, err = collection.ReplaceOne(ctx, filter, data)
	return err
}

func DeleteExam(course, assessment_name string) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)

	filter := bson.D{{Key: "course", Value: course}, {Key: "id", Value: assessment_name}}
	_, err := collection.DeleteOne(context.TODO(), filter)

	return err
}

// return true for there is no assessment in mongo
func ValidateAssessmentByCourse(course string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Ass_Collection_Name)
	filter := bson.D{{Key: "course", Value: course}}

	var assessment AutoExam_Assessments
	err := collection.FindOne(context.TODO(), filter).Decode(&assessment)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, err
}
