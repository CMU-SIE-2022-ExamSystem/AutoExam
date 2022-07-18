package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v3"
)

func GradeGen(course, assessment, student_id, path string, answer []byte) (string, error) {
	client := global.Mongo
	//write the solution.json
	errMessage, err := solutionHelper(client, course, assessment, student_id, path)
	if err != nil {
		return errMessage, err
	}
	//turn struct into yaml file again
	errMessage, err = yamlHelper(client, course, assessment, path)
	if err != nil {
		return errMessage, err
	}
	//save student's answer.json
	studentHelper(path, answer)
	err = nil
	return errMessage, err
}

func studentHelper(path string, answer []byte) {
	ioutil.WriteFile(path+"answer.json", answer, 0644)
}

func yamlHelper(client *mongo.Client, str1, str2, path string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	examCollection := client.Database("auto_exam").Collection("assessment_bank")
	var exam AutoExam_Assessments
	filter := bson.D{{Key: "course", Value: str1}, {Key: "id", Value: str2}}
	err := examCollection.FindOne(ctx, filter).Decode(&exam)
	if err == mongo.ErrNoDocuments {
		return "no such exam config!", err
	} else if err != nil {
		return "fatal error!", err
	}
	data, _ := yaml.Marshal(exam)
	ioutil.WriteFile(path+"config.yaml", data, 0644)
	return "", nil
}

//generate solution.json
//str1 is course id; str2 is exam id; str3 is student id
func solutionHelper(client *mongo.Client, str1, str2, str3, path string) (string, error) {
	// get the the context object，it can be used to set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//student bank collection
	stuCollection := client.Database("auto_exam").Collection("student_bank")
	var student Student
	filter := bson.D{{Key: "course", Value: str1}, {Key: "examId", Value: str2}, {Key: "studentId", Value: str3}}
	//judge if we have stored this student; if not we should generate config for him
	err := stuCollection.FindOne(ctx, filter).Decode(&student)
	if err == mongo.ErrNoDocuments {
		return "no such student config!", err
	} else if err != nil {
		return "fatal error", err
	}
	///question_bank_new collection
	questCollection := client.Database("auto_exam").Collection("question_bank_new")

	//call the helper func
	helper(&student, questCollection, path)
	return "", nil
}

func helper(student *Student, questCollection *mongo.Collection, path string) {
	// get the the context object，it can be used to set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//level1
	var result Header
	var filter bson.D
	level1 := map[string]interface{}{}
	for i, id := range student.Settings {
		str1 := fmt.Sprintf("Q%s", strconv.Itoa(i+1))
		//level2
		filter = bson.D{{Key: "headerId", Value: id}}
		questCollection.FindOne(ctx, filter).Decode(&result)
		level2 := map[string]interface{}{}
		questions := result.Questions
		for questionId, question := range questions {
			questionType := question.Type
			answer := question.Answer
			str2 := fmt.Sprintf("%s_sub%s", str1, strconv.Itoa(questionId+1))
			if questionType != "multiple_blank" {
				//level3
				level3 := map[string]string{}
				str3 := str2
				level3[str3] = answer[0]
				level2[str2] = level3
			} else {
				//level3
				level3 := map[string]string{}
				for answerId, value := range answer {
					str3 := fmt.Sprintf("%s_sub%s", str2, strconv.Itoa(answerId+1))
					level3[str3] = value
				}
				level2[str2] = level3
			}
		}
		level1[str1] = level2
	}
	solution, _ := json.Marshal(level1)
	ioutil.WriteFile(path+"solution_new.json", solution, 0644)
}
