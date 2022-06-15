package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/test"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})

		TestRouter.GET("/users", jwt.JWTAuth(), test.GetUsers)
		TestRouter.GET("/login", test.Login)
		TestRouter.GET("/refresh", jwt.JWTAuth(), jwt.UserRefreshHandler)
		TestRouter.GET("/cookie", test.CookieTest)
	}
}
