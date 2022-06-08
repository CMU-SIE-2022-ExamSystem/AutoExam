package main

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"go.uber.org/zap"
)

func main() {

	initialize.InitConfig()

	Router := initialize.Routers()

	initialize.InitLogger()

	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("error function", zap.String("error", "start error!"))
	}

}
