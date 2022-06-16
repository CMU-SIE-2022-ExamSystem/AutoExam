package initialize

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// logger and recovery policy
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	//development cors
	Router.Use(cors.Default())

	ApiGroup := Router.Group("/")
	router.AuthRouter(ApiGroup)
	router.UserRouter(ApiGroup)
	router.SwaggerRouter(ApiGroup)
	router.TestRouter(ApiGroup)
	router.CourseRouter(ApiGroup)
	return Router
}
