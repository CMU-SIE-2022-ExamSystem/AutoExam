package dao

import (
	"fmt"
	"io/ioutil"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

type PythonFile struct {
	QuestionType string `gorm:"type:varchar(255)"`
	PythonGrader []byte `gorm:"type:longblob"`
}

func insert(instance PythonFile) {
	if err := global.DB.Create(&instance).Error; err != nil {
		fmt.Println("insert failed")
		fmt.Println(err)
		return
	}
	fmt.Println("insert successfully")
}

func InsertOrUpddbate_grader(question_type string, byte_array []byte) {
	var instance PythonFile
	rows := global.DB.Where(&PythonFile{QuestionType: question_type}).Find(&instance)
	if rows.RowsAffected < 1 {
		// no grader file of this question type; need insert
		new_instance := PythonFile{
			question_type,
			byte_array,
		}
		insert(new_instance)
	} else {
		// the grader file of this type is already exsited, need to update
		if err := global.DB.Model(new(PythonFile)).Where("question_type=?", question_type).Update("python_grader", byte_array).Error; err != nil {
			fmt.Println("update failed!")
			fmt.Println(err)
			return
		}
		fmt.Println("update succeed!")
	}
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

func Delete_grader(question_type string) {
	global.DB.Where(&PythonFile{QuestionType: question_type}).Delete(&PythonFile{})
}
