package initialize

import "github.com/gin-gonic/gin"

func SetupServer() *gin.Engine {
	InitConfig()

	server := Routers()

	InitLogger()
	InitMysqlDB()

	return server
}
