package main

import (
	"testing"
)

// const (
// 	auth_url string = "/auth/"
// 	user_url string = "/user/"
// )

func TestInfo(t *testing.T) {
	// server := initialize.SetupServer()
	// req, _ := http.NewRequest("GET", auth_url+"info", nil)
	// w := httptest.NewRecorder()
	// server.ServeHTTP(w, req)

	// expectedStatus := http.StatusOK
	// // expectedContent := "hello world"

	// assert.Equal(t, expectedStatus, w.Code)

	// var data response.Response
	// json.Unmarshal(w.Body.Bytes(), &data)
	// assert.NotEmpty(t, data)
	// assert.NotEmpty(t, data.Data)
	// // assert.IsType(t, models.Autolab_Info_Front{}, data.Data)
	// assert.Contains(t, data.Data, "clientId")
	// assert.Contains(t, w.Body.String(), expectedContent)
}

// func TestCreateTokenInstructor(t *testing.T) {
// 	teardownTest, server := SetupTest(t)
// 	defer teardownTest(t)

// 	token, err := CreateInstructorToken()
// 	fmt.Println("==============")
// 	fmt.Println(token)
// 	fmt.Println("==============")
// 	assert.Nil(t, err)

// 	req, _ := http.NewRequest("GET", user_url+"courses", nil)
// 	req.Header.Add("Authorization", "Bearer "+token)
// 	w := httptest.NewRecorder()
// 	server.ServeHTTP(w, req)

// 	expectedStatus := http.StatusOK

// 	assert.Equal(t, expectedStatus, w.Code)
// 	data := &[]models.User_Courses{}
// 	// var response response.Response
// 	response := response.Response{Data: data}
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.NotEmpty(t, response.Data)

// 	assert.Equal(t, "instructor", (*data)[0].Auth_level)
// }

// func TestCreateTokenTA(t *testing.T) {
// 	token, err := CreateTAToken()
// 	assert.Nil(t, err)

// 	server := initialize.SetupServer()
// 	req, _ := http.NewRequest("GET", user_url+"courses", nil)
// 	req.Header.Add("Authorization", "Bearer "+token)
// 	w := httptest.NewRecorder()
// 	server.ServeHTTP(w, req)

// 	expectedStatus := http.StatusOK

// 	assert.Equal(t, expectedStatus, w.Code)
// 	data := &[]models.User_Courses{}
// 	// var response response.Response
// 	response := response.Response{Data: data}
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.NotEmpty(t, response.Data)
// 	assert.Equal(t, "course_assistant", (*data)[0].Auth_level)
// }

// func TestCreateTokenStudent(t *testing.T) {
// 	token, err := CreateStudentToken()
// 	fmt.Println("==============")
// 	fmt.Println("student")
// 	fmt.Println(token)
// 	fmt.Println("==============")
// 	assert.Nil(t, err)

// 	server := initialize.SetupServer()
// 	req, _ := http.NewRequest("GET", user_url+"courses", nil)
// 	req.Header.Add("Authorization", "Bearer "+token)
// 	w := httptest.NewRecorder()
// 	server.ServeHTTP(w, req)

// 	expectedStatus := http.StatusOK

// 	assert.Equal(t, expectedStatus, w.Code)
// 	data := &[]models.User_Courses{}
// 	// var response response.Response
// 	response := response.Response{Data: data}
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.NotEmpty(t, response.Data)
// 	assert.Equal(t, "student", (*data)[0].Auth_level)
// }
