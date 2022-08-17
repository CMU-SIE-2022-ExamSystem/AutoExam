package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRouter(Router *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/"
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
