package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.GET("/info", controller.Authinfo_Handler)
		AuthRouter.POST("/token", controller.Authtoken_Handler)
	}
}
