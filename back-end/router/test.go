package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})
		// TestRouter.GET("/info", autolab.Userinfo_Handler)
		TestRouter.GET("/courses", autolab.Usercourses_Handler)
	}
}
