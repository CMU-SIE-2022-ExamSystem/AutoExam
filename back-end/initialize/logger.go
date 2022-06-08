package initialize

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"go.uber.org/zap"
)

func InitLogger() {

	cfg := zap.NewDevelopmentConfig()

	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()), //
		"stdout",
	}
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
}
