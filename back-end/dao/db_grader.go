package dao

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PythonFile struct {
	QuestionType string `gorm:"type:varchar(255)"`
	PythonGrader []byte `gorm:"type:longblob"`
	Course       string `grom:"type:varchar(255)"`
	Valid        bool   `gorm:"type:bool"`
}

type Grader_API struct {
	Name   string `json:"name" binding:"required"`
	Course string `json:"course" binding:"required"`
	Valid  bool   `json:"valid" binding:"required"`
}

func insert(instance PythonFile) error {
	if err := global.DB.Create(&instance).Error; err != nil {
		fmt.Println("insert failed")
		fmt.Println(err)
		return err
	}
	fmt.Println("insert successfully")
	return nil
}

func InsertOrUpddbate_grader(question_type string, byte_array []byte, course string) (Grader_API, error) {
	var instance PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type, Course: course}).Find(&instance)
	if rows.RowsAffected < 1 {
		// no grader file of this question type; need insert
		new_instance := PythonFile{
			QuestionType: question_type,
			PythonGrader: byte_array,
			Course:       course,
			Valid:        false,
		}
		err := insert(new_instance)
		return new_instance.ToGraderAPI(), err
	} else {
		// the grader file of this type is already exsited, need to update
		if err := global.DB.Model(new(PythonFile)).Where("question_type=? AND course=?", question_type, course).Update("python_grader", byte_array).Error; err != nil {
			fmt.Println("update failed!")
			fmt.Println(err)
			return Grader_API{}, nil
		}
		fmt.Println("update succeed!")
		return ReadGrader(question_type, course)
	}
}

func search_grader(question_type, course string) (PythonFile, error) {
	var instance PythonFile
	result := global.DB.Where(&PythonFile{QuestionType: question_type, Course: course}).Find(&instance)
	return instance, result.Error
}

func search_all_grader(course string) ([]PythonFile, error) {
	var instances []PythonFile
	result := global.DB.Where(&PythonFile{Course: course}).Find(&instances)
	return instances, result.Error
}

func ReadAllGraders(course string) ([]Grader_API, error) {
	instances, err := search_all_grader(course)
	if err != nil {
		return nil, err
	}

	var api []Grader_API
	for _, instance := range instances {
		api = append(api, instance.ToGraderAPI())
	}
	return api, nil
}

func ReadGrader(question_type, course string) (Grader_API, error) {
	instance, err := search_grader(question_type, course)
	if err != nil {
		return instance.ToGraderAPI(), err
	}
	return instance.ToGraderAPI(), nil
}

func SearchAndStore_grader(c *gin.Context, question_type string, file_path string) {
	var new_data PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type}).Find(&new_data)
	if rows.RowsAffected < 1 {
		response.ErrDBResponse(c, "The corresponding grader of this question type can not be found.")
		return
	}
	byte_array := new_data.PythonGrader
	stored_file_name := fmt.Sprintf("%s.py", question_type)
	write_file(stored_file_name, byte_array, file_path)
}

func write_file(file_name string, byte_rray []byte, file_path string) {
	pathAndName := fmt.Sprintf("%s/%s", file_path, file_name)
	err := ioutil.WriteFile(pathAndName, byte_rray, 0666)
	if err != nil {
		color.Yellow("write fail!")
		return
	}
	color.Yellow("write success")
}

func Delete_grader(question_type, course string) error {
	result := global.DB.Where(&PythonFile{QuestionType: question_type, Course: course}).Delete(&PythonFile{})
	return result.Error
}

// return true for no grader in MySQL
func ValidateGrader(question_type, course string) bool {
	var instance PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type, Course: course}).Find(&instance)
	return rows.RowsAffected < 1
}

// return true for safe delete
func ValidateGraderUsed(question_type, course string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "questions.questionType", Value: question_type}, {Key: "course", Value: course}}
	var questions Question_Header
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

func UpdateGraderValid(question_type, course string, valid bool) (Grader_API, error) {
	if err := global.DB.Model(new(PythonFile)).Where("question_type=? AND course=?", question_type, course).Update("valid", valid).Error; err != nil {
		fmt.Println("update failed!")
		fmt.Println(err)
		return Grader_API{}, nil
	}
	return ReadGrader(question_type, course)
}

func GetCode(question_type, course string) string {
	grader, _ := search_grader(question_type, course)
	return grader.Code()
}

func (grader *PythonFile) Code() string {
	return string(grader.PythonGrader[:])
}

func (grader *PythonFile) ToGraderAPI() Grader_API {
	api := Grader_API{
		Name:   grader.QuestionType,
		Course: grader.Course,
		Valid:  grader.Valid,
	}
	return api
}
