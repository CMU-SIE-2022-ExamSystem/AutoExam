package initialize

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/docs"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
)

func InitSwagger() {
	docs.SwaggerInfo.Title = "Swagger Exam API"
	docs.SwaggerInfo.Description = "This is a backend server of auto exam"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = global.Settings.Ip + ":8080"
	// docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
