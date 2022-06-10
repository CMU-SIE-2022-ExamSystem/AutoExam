package global

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
	DB       *gorm.DB
)
