package authentication

import (
	"net/http"

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
	return
}

// Auth godoc
// @Summary get auth info
// @Schemes
// @Description get autolab authentication info
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} Authentication "desc"
// @Router /auth [get]
func Auth(c *gin.Context) {
	auth := read()
	c.JSON(http.StatusOK, auth)
}
