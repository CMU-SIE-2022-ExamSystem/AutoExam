package authentication

import (
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gin-gonic/gin"
)

type Info struct {
	Client_id string `mapstructure:"client_id"`
	Scope     string `mapstructure:"scope"`
}

func read() (info Info) {
	auth := global.Settings.Autolabinfo
	info = Info{Client_id: auth.Client_id, Scope: auth.Scope}
	return
}

// AuthInfo godoc
// @Summary get auth info
// @Schemes
// @Description get autolab authentication info
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} Info "desc"
// @Router /auth/info [get]
func AuthInfo(c *gin.Context) {
	auth := read()
	c.JSON(http.StatusOK, auth)
}
