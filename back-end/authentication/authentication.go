package authentication

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

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
	auth := global.Settings.Autolabinfo
	response.SuccessResponseWithType(c, gin.H{
		"clientId": auth.Client_id,
		"scope":    auth.Scope,
	}, 1)
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
	auth := global.Settings.Autolabinfo
	code := c.Query("code")

	http_body := models.Http_Body{
		Grant_type:    "authorization_code",
		Code:          code,
		Redirect_uri:  auth.Redirect_uri,
		Client_id:     auth.Client_id,
		Client_secret: auth.Client_secret,
	}

	// resp_body, _ := json.Marshal(http_body)
	// resp, err := http.Post("http://"+auth.Ip+"/oauth/token", "application/json", bytes.NewBuffer(resp_body))

	// if err != nil {
	// 	response.ErrUnauthResponse(c, "There may be something wrong with Autolab's web server, please try again later.")
	// }

	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// // fmt.Println(string(body))

	// if strings.Contains(string(body), "error") {
	// 	err_response := utils.Autolab_err_trans(string(body))
	// 	response.ErrUnauthResponse(c, err_response.Error_description)
	// 	return
	// }

	autolab_resp, flag := autolab.Autolab_Auth_Handler(c, "/oauth/token", http_body)

	if flag {
		autolab.Userinfo_Handler(c, autolab_resp.Access_token, autolab_resp.Refresh_token)
	}
}
