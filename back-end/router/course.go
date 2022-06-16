package router

import (
	"github.com/gin-gonic/gin"
)

func CourseRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("course")
	{
		// CourseRouter.GET("/info", controller.Authinfo_Handler)
		// CourseRouter.POST("/token", controller.Authtoken_Handler)
	}
}
