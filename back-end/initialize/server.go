package initialize

import (
	"crypto/tls"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	InitConfig()
	tls_check()
	server := Routers()

	InitLogger()
	InitMysqlDB()
	InitMongoDB()

	return server
}

func tls_check() {
	autolab := global.Settings.Autolabinfo
	if autolab.Protocol == "https" && autolab.Skip_Secure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}
