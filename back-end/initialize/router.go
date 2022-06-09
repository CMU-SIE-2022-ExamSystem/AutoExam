package initialize

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	Router.Use(middlewares.GinLogger())
	// Router.Use(gin.CustomRecovery(error.ErrorHandler))
	ApiGroup := Router.Group("/")
	router.AuthRouter(ApiGroup)
	router.SwaggerRouter(ApiGroup)
	router.TestRouter(ApiGroup)
	return Router
}
