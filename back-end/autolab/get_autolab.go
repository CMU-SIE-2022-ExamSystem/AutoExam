package autolab

import (
	"io/ioutil"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gin-gonic/gin"
)

func autolab_Api_Url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + "/api/v1" + endpoint
	return autolab_api_url
}

func AutolabGetHandler(c *gin.Context, token string, endpoint string) []byte {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, autolab_Api_Url(endpoint), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(request)

	if err != nil {
		AutolabErrorHander(c, resp, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
