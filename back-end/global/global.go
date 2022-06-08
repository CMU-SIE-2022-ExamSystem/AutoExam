package global

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/config"
	"go.uber.org/zap"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
)
