package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/test"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})

		TestRouter.GET("/users", middlewares.JWTAuth(), test.GetUsers)
		// AuthRouter.GET("/users", test.GetUsers)
		TestRouter.GET("/login", test.Login)
		TestRouter.GET("/courses", middlewares.JWTAuth(), autolab.Usercourses_Handler)

	}
}
