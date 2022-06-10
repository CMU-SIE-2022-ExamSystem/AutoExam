package main

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"go.uber.org/zap"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {

	initialize.InitConfig()

	Router := initialize.Routers()

	initialize.InitLogger()
	initialize.InitMysqlDB()

	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))

	if err != nil {
		zap.L().Info("error function", zap.String("error", "start error!"))
		fmt.Println(err)
	}

}
