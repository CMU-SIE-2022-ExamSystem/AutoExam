package autolab

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

func autolab_Oauth_Url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + endpoint
	return autolab_api_url
}

func fileUploadRequest(uri string, paramName string, path string, token string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	return req, err
}

func AutolabAuthHandler(c *gin.Context, endpoint string, http_body interface{}) (models.Autolab_Response, bool) {
	resp_body, _ := json.Marshal(http_body)
	resp, err := http.Post(autolab_Oauth_Url(endpoint), "application/json", bytes.NewBuffer(resp_body))

	if err != nil {
		AutolabErrorHander(c, resp, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var autolab_resp models.Autolab_Response

	if strings.Contains(string(body), "error") {
		err_response := utils.Autolab_err_trans(string(body))
		response.ErrUnauthResponse(c, err_response.Error_description)
		return autolab_resp, false
	}

	autolab_resp = utils.Autolab_resp_trans(string(body))

	return autolab_resp, true
}

func AutolabSubmitHandler(c *gin.Context, token string, endpoint string, path string) []byte {
	client := &http.Client{}
	request, internal_err := fileUploadRequest(autolab_Api_Url(endpoint), "submission[file]", path, token)

	var resp *http.Response
	if internal_err != nil {
		response.ErrFileResponse(c)
	} else {
		var err error
		resp, err = client.Do(request)
		if err != nil {
			AutolabErrorHander(c, resp, err)
		}
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	return body
}
