package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/courses", jwt.JWTAuth(), controller.Usercourses_Handler)
	}
}
