package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})

		TestRouter.GET("/users", jwt.JWTAuth(), controller.GetUsers)
		TestRouter.GET("/login", controller.Login)
		TestRouter.GET("/refresh", jwt.JWTAuth(), jwt.UserRefreshHandler)
		TestRouter.GET("/cookie", controller.CookieTest)
		TestRouter.GET("/:course_name/check", jwt.JWTAuth(), jwt.Check_authlevel_API)
		TestRouter.GET("/:course_name/checkDB", jwt.JWTAuth(), jwt.Check_authlevel_DB)
		TestRouter.POST("/:course_name/assessments/:assessment_name/submit", jwt.JWTAuth(), controller.Take_exam_Test)
		// TestRouter.GET("/f/:course/:assessment/:user_id/", controller.FolderTest)
		TestRouter.POST("/:course_name/:assessment_name/config", controller.Examconfig_Handler)
		TestRouter.POST("/quesion_update", controller.Question_Handler)
		TestRouter.GET("/exam", jwt.JWTAuth(), controller.Test_exam)
		TestRouter.GET("autograder/:question_type", controller.Autograder_Test)
	}
}
