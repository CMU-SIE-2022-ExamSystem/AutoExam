package initialize

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// logger and recovery policy
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	//development cors
	// Router.Use(cors.Default())
	Router.Use(CORSMiddleware())

	ApiGroup := Router.Group("/")
	router.AuthRouter(ApiGroup)
	router.BaseCourseRouter(ApiGroup)
	router.UserRouter(ApiGroup)
	router.SwaggerRouter(ApiGroup)
	router.CourseRouter(ApiGroup)
	router.ImageRouter(ApiGroup)
	return Router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
