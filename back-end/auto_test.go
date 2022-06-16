package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/stretchr/testify/assert"
)

const (
	url string = "/auth/"
)

func TestInfo(t *testing.T) {
	server := initialize.SetupServer()
	req, _ := http.NewRequest("GET", url+"info", nil) // 建立一個請求
	w := httptest.NewRecorder()                       // 建立一個ResponseRecorder其實作http.ResponseWriter，用來記錄response狀態
	server.ServeHTTP(w, req)                          // gin.Engine.ServerHttp實作http.Handler介面，用來處理HTTP請求及回應。

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
