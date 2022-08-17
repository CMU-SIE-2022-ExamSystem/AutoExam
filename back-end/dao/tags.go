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
	Tag_Collection_Name string = "tag"
)

type AutoExam_Tags struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `yaml:"name" json:"name" bson:"name" form:"name" binding:"required"`
	Course   string             `yaml:"base_course" json:"course" bson:"base_course" form:"base_course" binding:"required"`
}

type AutoExam_Tags_Create struct {
	Name   string `yaml:"name" json:"name" bson:"name" form:"name" binding:"required"`
	Course string `yaml:"base_course" json:"course" bson:"base_course" form:"base_course" binding:"required"`
}

type Tags struct {
	Id     string `bson:"id" json:"id"`
	Name   string `yaml:"name" json:"name" bson:"name" form:"name" binding:"required"`
	Course string `yaml:"base_course" json:"course" bson:"base_course" form:"base_course" binding:"required"`
}

type Tags_API struct {
	Name string `yaml:"name" json:"name" bson:"name" form:"name" binding:"required"`
}

func (autoexam *AutoExam_Tags) ToTags() Tags {
	tags := Tags{
		Id:     autoexam.ObjectID.Hex(),
		Name:   autoexam.Name,
		Course: autoexam.Course,
	}
	return tags
}

func ReadAllTags(base_course string) ([]Tags, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}}

	cursor, err := collection.Find(context.TODO(), filter)

	var tags []Tags

	for cursor.Next(context.TODO()) {
		var tag AutoExam_Tags
		cursor.Decode(&tag)
		tags = append(tags, tag.ToTags())
	}
	return tags, err
}

func CreateTag(tag AutoExam_Tags_Create) (result *mongo.InsertOneResult, err error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := bson.Marshal(tag)
	if err != nil {
		return nil, err
	}

	result, err = collection.InsertOne(ctx, data)
	return
}

func ReadTag(base_course, name string) (Tags, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)

	filter := bson.D{{Key: "base_course", Value: base_course}, {Key: "name", Value: name}}
	var tags AutoExam_Tags
	err := collection.FindOne(context.TODO(), filter).Decode(&tags)
	return tags.ToTags(), err
}

func UpdateTag(id string, tag AutoExam_Tags_Create) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}

	data, err := bson.Marshal(tag)
	if err != nil {
		return err
	}

	_, err = collection.ReplaceOne(ctx, filter, data)
	return err
}

func ReadTagById(id string) (Tags, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)
	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}
	var tags AutoExam_Tags
	err := collection.FindOne(context.TODO(), filter).Decode(&tags)
	return tags.ToTags(), err
}

func DeleteTagById(id string) error {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)

	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}}
	_, err := collection.DeleteOne(context.TODO(), filter)

	return err
}

// return true for no such object in mongo
func ValidateTag(base_course, name string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)
	filter := bson.D{{Key: "base_course", Value: base_course}, {Key: "name", Value: name}}

	var tags AutoExam_Tags
	err := collection.FindOne(context.TODO(), filter).Decode(&tags)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, err
}

// return true for no such object in mongo
func ValidateTagById(base_course, id string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Tag_Collection_Name)
	objectid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectid}, {Key: "base_course", Value: base_course}}

	var tags AutoExam_Tags
	err := collection.FindOne(context.TODO(), filter).Decode(&tags)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, err
}

// return true for safe delete
func ValidateTagUsedById(id string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "question_tag", Value: id}}
	var questions AutoExam_Questions
	err := collection.FindOne(context.TODO(), filter).Decode(&questions)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func ReadTagName(id string) string {
	tags, _ := ReadTagById(id)
	return tags.Name
}
