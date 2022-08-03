package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

const (
	auth_url string = "/auth/"
	user_url string = "/user/"
)

func CreateToken(Id uint, email string) (string, error) {
	j := jwt.NewJWT()
	claims := jwt.CustomClaims{
		ID:    uint(Id),
		Email: email,
		StandardClaims: jwt_go.StandardClaims{
			NotBefore: utils.GetNowTime(),
			ExpiresAt: utils.GetNowTime() + jwt.Expire_time,
			Issuer:    "test",
		},
	}
	token, err := j.CreateToken(claims)
	return token, err
}

func CreateInstructorToken() (string, error) {
	return CreateToken(1, "sie2022@andrew.cmu.edu")
}

func CreateStudentToken() (string, error) {
	return CreateToken(2, "student@andrew.cmu.edu")
}

func CreateTAToken() (string, error) {
	return CreateToken(3, "ta@andrew.cmu.edu")
}

func TestInfo(t *testing.T) {
	server := initialize.SetupServer()
	req, _ := http.NewRequest("GET", auth_url+"info", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	expectedStatus := http.StatusOK
	// expectedContent := "hello world"

	assert.Equal(t, expectedStatus, w.Code)

	var data response.Response
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.NotEmpty(t, data)
	assert.NotEmpty(t, data.Data)
	// assert.IsType(t, models.Autolab_Info_Front{}, data.Data)
	assert.Contains(t, data.Data, "clientId")
	// assert.Contains(t, w.Body.String(), expectedContent)
}

func TestCreateTokenInstructor(t *testing.T) {
	token, err := CreateInstructorToken()
	fmt.Println("==============")
	fmt.Println(token)
	fmt.Println("==============")
	assert.Nil(t, err)

	server := initialize.SetupServer()
	req, _ := http.NewRequest("GET", user_url+"courses", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	expectedStatus := http.StatusOK

	assert.Equal(t, expectedStatus, w.Code)
	data := &[]models.User_Courses{}
	// var response response.Response
	response := response.Response{Data: data}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotEmpty(t, response.Data)

	assert.Equal(t, "instructor", (*data)[0].Auth_level)
}

func TestCreateTokenTA(t *testing.T) {
	token, err := CreateTAToken()
	assert.Nil(t, err)

	server := initialize.SetupServer()
	req, _ := http.NewRequest("GET", user_url+"courses", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	expectedStatus := http.StatusOK

	assert.Equal(t, expectedStatus, w.Code)
	data := &[]models.User_Courses{}
	// var response response.Response
	response := response.Response{Data: data}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotEmpty(t, response.Data)
	assert.Equal(t, "course_assistant", (*data)[0].Auth_level)
}

func TestCreateTokenStudent(t *testing.T) {
	token, err := CreateStudentToken()
	fmt.Println("==============")
	fmt.Println("student")
	fmt.Println(token)
	fmt.Println("==============")
	assert.Nil(t, err)

	server := initialize.SetupServer()
	req, _ := http.NewRequest("GET", user_url+"courses", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	expectedStatus := http.StatusOK

	assert.Equal(t, expectedStatus, w.Code)
	data := &[]models.User_Courses{}
	// var response response.Response
	response := response.Response{Data: data}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotEmpty(t, response.Data)
	assert.Equal(t, "student", (*data)[0].Auth_level)
}
