package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/error"
	"github.com/gin-gonic/gin"
)

func autolab_err_trans(str string) autolab_err_Response {
	var response autolab_err_Response
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("json transfer error>>> ", err)
	}
	return response
}

func autolab_resp_trans(str string) autolab_Response {
	var response autolab_Response
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("json transfer error>>> ", err)
	}
	return response
}

// Authinfo_Handler godoc
// @Summary get auth info
// @Schemes
// @Description get autolab authentication info
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} Authentication "desc"
// @Router /auth/info [get]
func Authinfo_Handler(c *gin.Context) {
	auth := autolab.Read_Autolab_Env()
	c.JSON(http.StatusOK, gin.H{
		"Client_id": auth.Client_id,
		"Scope":     auth.Scope,
	})
}

// Authtoken_Handler godoc
// @Summary get auth token
// @Schemes
// @Description get autolab authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} Authentication "desc"
// @Router /auth [post]
func Authtoken_Handler(c *gin.Context) {
	auth := autolab.Read_Autolab_Env()
	code := c.Query("code")

	http_body := http_Body{
		Grant_type:    "authorization_code",
		Code:          code,
		Redirect_uri:  auth.Redirect_uri,
		Client_id:     auth.Client_id,
		Client_secret: auth.Client_secret,
	}

	resp_body, _ := json.Marshal(http_body)
	resp, err := http.Post("http://"+auth.Ip+"/oauth/token", "application/json", bytes.NewBuffer(resp_body))

	if err != nil {
		error.ErrorHandler(c, error.AutolabResponse())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	if strings.Contains(string(body), "error") {
		err_response := autolab_err_trans(string(body))
		c.JSON(http.StatusOK, err_response)
		return
	}

	response := autolab_resp_trans(string(body))
	//todo store user information into database

	c.JSON(http.StatusOK, response)
}
