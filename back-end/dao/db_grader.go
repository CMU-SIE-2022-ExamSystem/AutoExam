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
	"golang.org/x/exp/slices"
)

type PythonFile struct {
	ID           uint     `json:"-" gorm:"primaryKey"`
	QuestionType string   `gorm:"type:varchar(255)"`
	PythonGrader []byte   `gorm:"type:longblob"`
	BaseCourse   string   `grom:"type:varchar(255)"`
	Valid        bool     `gorm:"type:bool"`
	Blanks       []Blanks `gorm:"constraint:OnUpdate:CASCADE,,OnDelete:CASCADE,foreignKey:PythonFileID"`
	Uploaded     bool     `gorm:"type:bool"`
}

// @Description grader model info
type Grader_API struct {
	Name     string   `json:"name" binding:"required"`
	Blanks   []Blanks `json:"blanks"`                   // describing the structure of blanks in this grader
	Valid    bool     `json:"valid" binding:"required"` // whether the grader is valid by /courses/{course_name}/graders/{grader_name}/valid api
	Uploaded bool     `json:"uploaded"`                 // whether the python file is uploaded
}

// @Description describing the type of a certain blank
type Blanks struct {
	ID           uint   `json:"-" gorm:"primaryKey"`
	PythonFileID uint   `json:"-"`
	Type         string `json:"type" binding:"required,oneof=string code"`
}

func insert(instance PythonFile) error {
	if err := global.DB.Create(&instance).Error; err != nil {
		return err
	}
	return nil
}

func Insert_grader(instance PythonFile) (Grader_API, error) {
	err := insert(instance)
	if err != nil {
		return Grader_API{}, err
	}
	return ReadGrader(instance.QuestionType, instance.BaseCourse)
}

func Update_blanks_grader(instance PythonFile) (Grader_API, error) {
	old_instance, _ := search_grader(instance.QuestionType, instance.BaseCourse)
	// delete old_instance's blanks
	for _, blank := range old_instance.Blanks {
		global.DB.Delete(&blank)
	}

	// insert new blanks into old_instance
	global.DB.Model(&old_instance).Association("Blanks").Replace(instance.Blanks)

	return ReadGrader(instance.QuestionType, instance.BaseCourse)
}

func Update_python_grader(instance PythonFile) (Grader_API, error) {
	old_instance, _ := search_grader(instance.QuestionType, instance.BaseCourse)
	old_instance.PythonGrader = instance.PythonGrader
	old_instance.Uploaded = true

	global.DB.Save(&old_instance)

	return ReadGrader(instance.QuestionType, instance.BaseCourse)
}

func InsertOrUpddbate_grader(question_type string, byte_array []byte, course string) (Grader_API, error) {
	var instance PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type, BaseCourse: course}).Find(&instance)
	if rows.RowsAffected < 1 {
		// no grader file of this question type; need insert
		new_instance := PythonFile{
			QuestionType: question_type,
			PythonGrader: byte_array,
			BaseCourse:   course,
			Valid:        false,
		}
		err := insert(new_instance)
		return new_instance.ToGraderAPI(), err
	} else {
		// the grader file of this type is already exsited, need to update
		if len(byte_array) != 0 {
			instance := PythonFile{
				PythonGrader: byte_array,
				Uploaded:     true,
			}
			if err := global.DB.Model(new(PythonFile)).Where("question_type=? AND base_course=?", question_type, course).Updates(instance).Error; err != nil {
				return Grader_API{}, err
			}
		} else {
			if err := global.DB.Model(new(PythonFile)).Where("question_type=? AND base_course=?", question_type, course).Update("python_grader", byte_array).Error; err != nil {
				return Grader_API{}, err
			}
		}

		return ReadGrader(question_type, course)
	}
}

func search_grader(question_type, course string) (PythonFile, error) {
	var instance PythonFile
	result := global.DB.Preload("Blanks").Where(&PythonFile{QuestionType: question_type, BaseCourse: course}).Find(&instance)
	return instance, result.Error
}

func search_all_grader(course string) ([]PythonFile, error) {
	var instances []PythonFile
	result := global.DB.Preload("Blanks").Where(&PythonFile{BaseCourse: course}).Find(&instances)
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

// read all validated graders
func ReadAllValidGraders(course string) ([]Grader_API, error) {
	instances, err := search_all_grader(course)
	if err != nil {
		return nil, err
	}

	var api []Grader_API
	for _, instance := range instances {
		if instance.Valid {
			api = append(api, instance.ToGraderAPI())
		}
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
	instance, _ := search_grader(question_type, course)
	// association delete
	result := global.DB.Select("Blanks").Delete(&instance)
	return result.Error
}

// return true for no grader in MySQL
func ValidateGrader(question_type, course string) bool {
	if slices.Contains(global.Settings.Basic_Grader, question_type) {
		return false
	}

	var instance PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type, BaseCourse: course}).Find(&instance)
	return rows.RowsAffected < 1
}

// return true for safe delete
func ValidateGraderUsed(question_type, course string) (bool, error) {
	client := global.Mongo
	//get the collection instance
	collection := client.Database("auto_exam").Collection(Que_Collection_Name)
	filter := bson.D{{Key: "questions.questionType", Value: question_type}, {Key: "course", Value: course}}
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

func UpdateGraderValid(question_type, course string, valid bool) (Grader_API, error) {
	if err := global.DB.Model(new(PythonFile)).Where("question_type=? AND base_course=?", question_type, course).Update("valid", valid).Error; err != nil {
		return Grader_API{}, err
	}
	return ReadGrader(question_type, course)
}

func GetCode(question_type, course string) string {
	grader, _ := search_grader(question_type, course)
	return grader.Code()
}

func GetUploadStatus(question_type, course string) bool {
	grader, _ := search_grader(question_type, course)
	return grader.Uploaded
}

func (grader *PythonFile) Code() string {
	return string(grader.PythonGrader[:])
}

func (grader *PythonFile) ToGraderAPI() Grader_API {
	api := Grader_API{
		Name:     grader.QuestionType,
		Valid:    grader.Valid,
		Blanks:   grader.Blanks,
		Uploaded: grader.Uploaded,
	}
	return api
}
