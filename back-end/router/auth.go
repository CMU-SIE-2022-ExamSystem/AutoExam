package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/authentication"
	"github.com/gin-gonic/gin"
)

func AuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.GET("/info", authentication.AuthInfo)
	}
}
