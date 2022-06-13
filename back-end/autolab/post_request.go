package autolab

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/gin-gonic/gin"
)

func autolab_oauth_url(endpoint string) string {
	autolab_url := global.Settings.Autolabinfo.Protocol + "://" + global.Settings.Autolabinfo.Ip
	autolab_api_url := autolab_url + endpoint
	return autolab_api_url
}

func Autolab_Auth_Handler(c *gin.Context, endpoint string, http_body interface{}) (models.Autolab_Response, bool) {
	resp_body, _ := json.Marshal(http_body)
	resp, err := http.Post(autolab_oauth_url(endpoint), "application/json", bytes.NewBuffer(resp_body))

	if err != nil {
		Autolab_Error_Hander(c, resp, err)
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
