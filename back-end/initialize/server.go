package initialize

import "github.com/gin-gonic/gin"

func SetupServer() *gin.Engine {
	InitConfig()

	server := Routers()

	InitLogger()
	InitMysqlDB()
	InitMongoDB()

	return server
}
