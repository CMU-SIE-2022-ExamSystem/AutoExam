package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/config"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/error"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gin-gonic/gin"
)

func read() (auth config.AutolabConfig) {
	auth = global.Settings.Autolabinfo
	if auth.Ip == "" {
		panic(error.EnvResponse("ip"))
	} else if auth.Client_id == "" {
		panic(error.EnvResponse("client_id"))
	} else if auth.Client_secret == "" {
		panic(error.EnvResponse("client_secret"))
	} else if auth.Redirect_uri == "" {
		panic(error.EnvResponse("redirect_uri"))
	} else if auth.Scope == "" {
		panic(error.EnvResponse("scope"))
	}
	return auth
}

func Autolab_Err(str string) autolab_err_Response {
	var response autolab_err_Response
	err := json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("json transfer error>>> ", err)
	}
	return response
}

func Autolab_Response(str string) autolab_Response {
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
	auth := read()
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
	auth := read()
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
	fmt.Println(string(body))

	if strings.Contains(string(body), "error") {
		err_response := Autolab_Err(string(body))
		c.JSON(http.StatusOK, err_response)
		return
	}

	response := Autolab_Response(string(body))

	c.JSON(http.StatusOK, response)
}
