package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func QuestionRouter(Router *gin.RouterGroup) {
	QuestionRouter := Router.Group("questions")
	{
		QuestionRouter.GET("/tags", jwt.JWTAuth(), controller.QuestionTag_Handler)

		// questions CRUD
		QuestionRouter.POST("/", jwt.JWTAuth(), controller.CreateQuestion_Handler)
		QuestionRouter.GET("/:question_id/", jwt.JWTAuth(), controller.ReadQuestion_Handler)
		QuestionRouter.PUT("/:question_id/", jwt.JWTAuth(), controller.UpdateQuestion_Handler)
		QuestionRouter.DELETE("/:question_id/", jwt.JWTAuth(), controller.DeleteQuestion_Handler)
	}
}
