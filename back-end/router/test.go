package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/test"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("test")
	{
		AuthRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})

		AuthRouter.GET("/users", middlewares.JWTAuth(), test.GetUsers)
		// AuthRouter.GET("/users", test.GetUsers)
		AuthRouter.GET("/login", test.Login)
		AuthRouter.GET("/courses", autolab.Usercourses_Handler)

	}
}
